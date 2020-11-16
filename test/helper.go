package test

import (
	"testing"

	"github.com/StephenFooBar/gopher-pouches/command"
	"github.com/StephenFooBar/gopher-pouches/config"
	"github.com/StephenFooBar/gopher-pouches/datastore"
	"github.com/stretchr/testify/assert"
)

func AssertFailure(t *testing.T, expected string, actual command.Response) {
	assert.False(t, actual.Success)
	assert.Equal(t, expected, actual.Message)
}

//create a func to call redis.initializeDB.
// the func name might be: InitializeRedis
func InitializeRedis(conf config.Config) {
	r := datastore.GetInstance(conf.Connection)
	r.InitializeDb()
	// needs to have config info
	// based on the config info, create redis instance
	// call initialize DB method in redis
}
