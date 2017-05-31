package logger

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(Logger.Logger.Level.String(), "info", "Default log level should be info")
}
