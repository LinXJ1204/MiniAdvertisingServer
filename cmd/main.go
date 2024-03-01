package main

import (
	"context"
	"fmt"

	common "../common"
	db "../db"
)

func main() {
	common.NewAd()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	mongoDb := db.NewMongoDatabase(ctx, "localhost:27017")

	/* test := &common.Ad{StartAt: "2024-02-25T03:00:00.000Z", EndAt: "2024-12-31T16:00:00.000Z", Title: "AD 65"}
	mongoDb.NewAd(test) */
	r := mongoDb.Search(&common.SearchCondition{})
	for _, v := range r {
		fmt.Println(v.Title, v.EndAt)
	}
}
