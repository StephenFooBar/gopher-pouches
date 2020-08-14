package config

import (
	"testing"

	"github.com/StephenFooBar/gopher-pouches/command"
	"github.com/stretchr/testify/assert"
)

func TestGetConfigShouldShowErrorMessageWhenYamlFileIsMissing(t *testing.T) {
	expected := command.YamlFileMissing
	actual := Get()
	assert.Equal(t, false, actual.Success)
	assert.Equal(t, expected, actual.Message)
}
