// (c) 2025 Tadahiro Ishisaka all rights reserved.
// バッチ生成の例

package batch

import (
	"github.com/oklog/ulid"
	"math/rand"
	"time"
)

// GenerateBatch は指定された数のULIDを生成します
func GenerateBatch(count int) []ulid.ULID {
	if count <= 0 {
		return make([]ulid.ULID, 0)
	}
	result := make([]ulid.ULID, count)
	now := time.Now()
	// 乱数作成はもっと厳密にするべき
	e := ulid.Monotonic(rand.New(rand.NewSource(now.UnixNano())), 0)
	for i := range count {
		result[i] = ulid.MustNew(ulid.Timestamp(now), e)
	}

	return result
}
