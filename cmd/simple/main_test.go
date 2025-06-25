/*
 * Copyright (c) 2025. Tadahiro Ishisaka All rights reserved.
 */

// ULID作成のためのシンプルなサンプル

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExampleULID(t *testing.T) {
	gotULID := ExampleULID()
	assert.NotNil(t, gotULID)
	assert.NotEmpty(t, gotULID.String())
	assert.Equal(t,
		gotULID.String(),
		"0000XSNJG0MQJHBF4QX1EFD6Y3")
}
