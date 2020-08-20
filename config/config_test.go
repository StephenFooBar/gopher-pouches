package config

import (
	"testing"

	"github.com/StephenFooBar/gopher-pouches/command"
	"github.com/stretchr/testify/assert"
)

func TestGetConfigShouldShowErrorMessageWhenYamlFileIsMissing(t *testing.T) {
	expected := command.ConfigFileMissing
	actual := Get("not-existing-file")
	assert.Equal(t, false, actual.Success)
	assert.Equal(t, expected, actual.Message)
}

func TestGetConfigShouldShowErrorMessageWhenConfigFileIsInvalid(t *testing.T) {
	expected := command.InvalidConfig
	actual := Get("test-invalid-config.yml")
	assert.Equal(t, false, actual.Success)
	assert.Equal(t, expected, actual.Message)
}

func TestGetConfigShouldShowErrorMessageWhenDatabaseTypeIsMissingInConfig(t *testing.T) {
	expected := command.ConfigEntryMissing
	actual := Get("test-datastore-type-missing.yml")
	assert.Equal(t, false, actual.Success)
	assert.Equal(t, expected, actual.Message)
}
