package test

import "github.com/StephenFooBar/gopher-pouches/datastore/common"

var (
	EmptyFeed          = common.Feed{}
	MissingURLMockFeed = common.Feed{"Mock Feed", ""}
	InvalidURLMockFeed = common.Feed{"Mock Feed", "not url"}

	NonRssMockFeed      = common.Feed{"Mock Feed", "https://www.google.com"}
	NotFoundURLMockFeed = common.Feed{"Mock Feed", "https://www.google.com/notexist"}
	MissingRssTagFeed   = common.Feed{"Mock Feed", "http://localhost:" + MockPort + "/feeds/missing-rss-tag.xml"}

	ValidRedisConnection = "host=:6379,database=2"
	MockPort             = "41914"
)
