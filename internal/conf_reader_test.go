package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetConfig(t *testing.T) {
	assert := assert.New(t)

	conf, err := GetConfig()
	assert.NotNil(err, "Error should be not nil")
	assert.Equal(err.Error(), ErrConfigUninitalized.Error())

	err = InitConfig("./testdir/config.yaml")
	assert.Nil(err, "File present, error should be nil")
	conf, err = GetConfig()
	expected := "world"
	assert.Equal(expected, conf.GetString("field.hello"), "expected to read world from the config")

	err = InitConfig(".")
	assert.NotNil(err, "File NOT present, error should be notnil")

}
