package uuid3

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetUUID(t *testing.T) {
	uuidGen := NewUuid3Generator()

	uuidTest, err := uuidGen.GenerateID()

	assert.NoError(t, err)
	assert.NotEqual(t, uuidTest, "")

}
