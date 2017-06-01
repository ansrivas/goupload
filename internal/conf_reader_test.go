package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetConfig(t *testing.T) {
	assert := assert.New(t)

	conf, err := NewConfig("./testdir/config.yaml")
	assert.Nil(err, "File present, error should be nil")
	config := conf.GetConfig()
	expected := "world"
	assert.Equal(expected, config.GetString("field.hello"), "expected to read world from the config")

	_, err = NewConfig(".")
	assert.NotNil(err, "File NOT present, error should be notnil")

}
