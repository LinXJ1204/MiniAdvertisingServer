package db

import (
	"context"
	"fmt"
	"time"

	"../common"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDbAdsType struct {
	Title             string             `json:"title"`
	StartAt           primitive.DateTime `bson:"startAt"`
	EndAt             primitive.DateTime `bson:"endAt"`
	*common.Condition `json:"condition"`
}

func NewMongoDatabase(ctx context.Context, dbUrl string) AdsDatabase {
	d := &MongoAdsDatabase{}
	d.newdb(ctx, dbUrl)
	return d
}

type MongoAdsDatabase struct {
	ActiceCollection       *mongo.Collection
	InActiceCollection     *mongo.Collection
	ExpireActiceCollection *mongo.Collection
	latestExpire           string
	latestStart            string
}

func (d *MongoAdsDatabase) newdb(ctx context.Context, dbUrl string) error {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://"+dbUrl))
	if err != nil {
		return err
	}

	d.ActiceCollection = client.Database("AdsDb").Collection("ActiceCollection")
	d.InActiceCollection = client.Database("AdsDb").Collection("InActiceCollection")
	d.InActiceCollection = client.Database("AdsDb").Collection("ExpireActiceCollection")

	return nil
}

func (d *MongoAdsDatabase) NewAd(ad *common.Ad) error {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	_, err := d.ActiceCollection.InsertOne(ctx, d.AdCommonTypeToMongoDbType(ad))
	//Todo: check latestExpire
	return err
}

func (d *MongoAdsDatabase) Search(c *common.SearchCondition) []*common.Respond {
	var Ads []MongoDbAdsType
	var resAds []*common.Respond
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	//filter active ads
	filter := bson.D{
		/* {Key: "startAt", Value: bson.D{
			{Key: "$lt", Value: primitive.NewDateTimeFromTime(time.Now())},
		}}, {
			Key: "endAt",
			Value: bson.D{
				{Key: "$gte", Value: primitive.NewDateTimeFromTime(time.Now())},
			},
		}, */
	}

	cursor, err := d.ActiceCollection.Find(ctx, filter)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	cursor.All(ctx, &Ads)

	for _, v := range Ads {
		resAds = append(resAds, d.MongoDbTypeToRespond(v))
	}

	return resAds
}

func (d *MongoAdsDatabase) AutoUpdate() {

}

func (d *MongoAdsDatabase) AdCommonTypeToMongoDbType(ad *common.Ad) *MongoDbAdsType {
	res := &MongoDbAdsType{}
	start, err := time.Parse(time.RFC3339Nano, ad.StartAt)
	if err != nil {
		return nil
	}
	end, err := time.Parse(time.RFC3339Nano, ad.EndAt)
	if err != nil {
		return nil
	}

	res.StartAt = primitive.NewDateTimeFromTime(start)
	res.EndAt = primitive.NewDateTimeFromTime(end)
	res.Title = ad.Title
	res.Condition = ad.Condition

	return res
}

func (d *MongoAdsDatabase) MongoDbTypeToRespond(a MongoDbAdsType) *common.Respond {
	res := &common.Respond{}
	res.Title = a.Title
	res.EndAt = a.EndAt.Time().Format(time.RFC3339Nano)
	return res
}
