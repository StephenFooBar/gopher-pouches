package feed

import (
	"github.com/StephenFooBar/gopher-pouches/datastore/common"
)

var (
	emptyFeed            = common.Feed{}
	mockFeed             = common.Feed{"Mock Feed", "https://www.google.com"}
	validRedisConnection = "host=:6379,database=2"
)
