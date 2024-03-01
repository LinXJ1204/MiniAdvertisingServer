package db

import (
	"../common"
)

type AdsDatabase interface {
	NewAd(*common.Ad) error
	Search(*common.SearchCondition) []*common.Respond
}
