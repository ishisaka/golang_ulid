package func_sample

import (
	"fmt"
	"github.com/oklog/ulid"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestCreateULID(t *testing.T) {
	ti := time.Unix(1000000, 0)
	id := CreateULID(ti)
	assert.NotNil(t, id)
	assert.NotEmpty(t, id.String())
	assert.Equal(t,
		id.String(),
		"0000XSNJG0MQJHBF4QX1EFD6Y3")
}

func TestParseULID(t *testing.T) {
	s := "0000XSNJG0MQJHBF4QX1EFD6Y3"
	id, err := ParseULID(s)
	assert.NoError(t, err)
	assert.NotNil(t, id)
	assert.IsType(t, ulid.ULID{}, id)
	assert.Equal(t, id.String(), s)
	ti := time.Unix(1000000, 0)
	assert.Equal(t, ulid.Time(id.Time()), ti)
	s2 := "hoge"
	id2, err2 := ParseULID(s2)
	if assert.Error(t, err2) {
		fmt.Println(err2)
	}
	// assert.Nil(t, id2) これは失敗
	assert.Equal(t, ulid.ULID{}, id2) // 空のULIDが返ってくる
}
