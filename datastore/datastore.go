package datastore

import (
	"github.com/StephenFooBar/gopher-pouches/config"
	"github.com/StephenFooBar/gopher-pouches/datastore/common"
	"github.com/StephenFooBar/gopher-pouches/datastore/redis"
)

func GetDatastore(conf config.Config) common.Datastore {
	switch conf.Datastore {
	case "redis":
		return redis.GetInstance(conf.Connection)
	default:
		return nil
	}
}
