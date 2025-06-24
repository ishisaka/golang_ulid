package func_sample

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestCreateULID(t *testing.T) {
	ti := time.Unix(1000000, 0)
	ulid := CreateULID(ti)
	assert.NotNil(t, ulid)
	assert.NotEmpty(t, ulid.String())
	assert.Equal(t,
		ulid.String(),
		"0000XSNJG0MQJHBF4QX1EFD6Y3")
}
