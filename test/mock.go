package test

import "github.com/StephenFooBar/gopher-pouches/datastore/common"

var (
	EmptyFeed          = common.Feed{}
	MissingURLMockFeed = common.Feed{"Mock Feed", ""}
	InvalidURLMockFeed = common.Feed{"Mock Feed", "not url"}

	NonRssMockFeed        = common.Feed{"Mock Feed", MockServerUrl + "/"}
	NotFoundURLMockFeed   = common.Feed{"Mock Feed", MockServerUrl + "/feeds/notexist"}
	MissingRssTagFeed     = common.Feed{"Mock Feed", MockServerUrl + "/feeds/missing-rss-tag.xml"}
	MissingChannelTagFeed = common.Feed{"Mock Feed", MockServerUrl + "/feeds/missing-channel-tag.xml"}

	ValidRedisConnection = "host=:6379,database=2"
	MockPort             = "41914"
	MockServerUrl        = "http://localhost:" + MockPort
)
