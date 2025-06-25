/*
 * Copyright (c) 2025. Tadahiro Ishisaka All rights reserved.
 */

// バッチ生成の例

package batch

import (
	"fmt"
	"testing"

	"github.com/oklog/ulid"
	"github.com/stretchr/testify/assert"
)

func TestGenerateBatch(t *testing.T) {
	t.Run("generates valid ULIDs", func(t *testing.T) {
		count := 5
		result := GenerateBatch(count)

		assert.Equal(t, count, len(result), "batch size mismatch")
		for _, id := range result {
			assert.NotEqual(t, ulid.ULID{}, id, "expected valid ULID, got zero ULID")
			fmt.Println(id.String())
		}
	})

	t.Run("handles zero batch size", func(t *testing.T) {
		count := 0
		result := GenerateBatch(count)

		assert.Empty(t, result, "expected empty batch for zero count")
	})

	t.Run("handles negative batch size", func(t *testing.T) {
		count := -5
		result := GenerateBatch(count)

		assert.Empty(t, result, "expected empty batch for negative count")
	})
}
