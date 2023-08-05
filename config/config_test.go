package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfigValues(t *testing.T) {
	err := Init("test_config_file", "./")
	assert.Nil(t, err)
	assert.Equal(t, "1000", GetConfigValue("section1.value"))
	assert.Equal(t, "some-name", GetConfigValue("section2.name"))
}

func TestInvalidFile(t *testing.T) {
	err := Init("non-existing-file", "./")
	assert.NotNil(t, err)
}
