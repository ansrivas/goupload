package logger

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_(t *testing.T) {
	assert := assert.New(t)

	log := SetupLogger("testapp")
	assert.Equal(log.Logger.Level.String(), "info", "Default log level should be info")
}
