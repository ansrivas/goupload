package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Assets(t *testing.T) {
	assert := assert.New(t)

	templates := SetAssetsPath("./testdir")

	assert.NotNil(templates.Login, "Template should be properly initialized")
	assert.NotNil(templates.Public, "Template should be properly initialized")
	assert.NotNil(templates.Protected, "Template should be properly initialized")
}
