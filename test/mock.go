package test

import "github.com/StephenFooBar/gopher-pouches/datastore/common"

var (
	EmptyFeed            = common.Feed{}
	NonRssMockFeed       = common.Feed{"Mock Feed", "https://www.google.com"}
	RssMockFeed          = common.Feed{"Mock Rss Feed", ""}
	ValidRedisConnection = "host=:6379,database=2"
)
