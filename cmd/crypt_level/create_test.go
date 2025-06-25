/*
 * Copyright (c) 2025. Tadahiro Ishisaka All rights reserved.
 */

// 暗号グレードのエントロピーを使用してULIDを生成するサンプル

package crypt_level

import (
	"fmt"
	"testing"
	"time"

	"github.com/oklog/ulid"
	"github.com/stretchr/testify/assert"
)

func TestCreateULID(t *testing.T) {
	tests := []struct {
		name string
		time time.Time
	}{
		{
			name: "ZeroTime",
			time: time.Time{},
		},
		{
			name: "CurrentTime",
			time: time.Now(),
		},
		{
			name: "PastTime",
			time: time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			name: "FutureTime",
			time: time.Date(2050, 1, 1, 0, 0, 0, 0, time.UTC),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ulidValue := CreateULID(tt.time)
			fmt.Printf("ULID: %v time: %v\n", ulidValue, tt.time)
			assert.NotEmpty(t, ulidValue.String(), "ULID should not be empty")

			parsedULID, err := ulid.Parse(ulidValue.String())
			assert.NoError(t, err, "ULID should be valid")

			if !tt.time.IsZero() {
				// ulid.Timestamp()がtime.Timeがゼロ値の時の処理がよろしくないので
				// ゼロ値の場合は以下の検査を除外
				assert.Equal(t, ulid.Timestamp(tt.time), ulidValue.Time(), "ULID timestamp mismatch")
			}
			assert.Equal(t, parsedULID, ulidValue, "Parsed ULID should match original")
		})
	}
}
