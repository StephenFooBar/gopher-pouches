package config

import (
	"testing"

	"github.com/StephenFooBar/gopher-pouches/command"
	"github.com/stretchr/testify/assert"
)

func TestGetConfigShouldShowErrorMessageWhenYamlFileIsMissing(t *testing.T) {
	expected := command.ConfigFileMissing
	_, actual := Get("not-existing-file")
	assert.Equal(t, false, actual.Success)
	assert.Equal(t, expected, actual.Message)
}

func TestGetConfigShouldShowErrorMessageWhenConfigFileIsInvalid(t *testing.T) {
	expected := command.InvalidConfig
	_, actual := Get(Root + "/test/test-invalid-config.yml")
	//t.Log(Root)
	assert.Equal(t, false, actual.Success)
	assert.Equal(t, expected, actual.Message)
}

func TestGetConfigShouldShowErrorMessageWhenDatabaseTypeIsMissingInConfig(t *testing.T) {
	expected := command.ConfigEntryMissing
	_, actual := Get(Root + "/test/test-datastore-type-missing.yml")
	assert.Equal(t, false, actual.Success)
	assert.Equal(t, expected, actual.Message)
}

func TestGetConfigShouldShowErrorMessageWhenConnectionInfoIsMissingInConfig(t *testing.T) {
	expected := command.ConfigEntryMissing
	_, actual := Get(Root + "/test/test-connection-missing.yml")
	assert.Equal(t, false, actual.Success)
	assert.Equal(t, expected, actual.Message)
}

func TestGetConfigShouldReturnSuccessfullyWhenConfigFileIsValid(t *testing.T) {
	//expected := command.Success
	//iactual
}
