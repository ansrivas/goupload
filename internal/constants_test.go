package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Assets(t *testing.T) {
	assert := assert.New(t)

	SetAssetsPath("./testdir")

	assert.NotNil(Templates.Login, "Template should be properly initialized")
	assert.NotNil(Templates.Public, "Template should be properly initialized")
	assert.NotNil(Templates.Protected, "Template should be properly initialized")
}
