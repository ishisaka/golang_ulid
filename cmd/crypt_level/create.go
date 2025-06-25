/*
 * Copyright (c) 2025. Tadahiro Ishisaka All rights reserved.
 */

// 暗号グレードのエントロピーを使用してULIDを生成するサンプル

package crypt_level

import (
	"crypto/rand"
	"time"

	"github.com/oklog/ulid"
)

// CreateULID は指定された時刻に基づいて一意のULIDを生成します。
// 入力時刻がゼロ値の場合は空のULIDを返します。
// ULIDに生成においては暗号グレードのエントロピーを使用します。
func CreateULID(t time.Time) ulid.ULID {
	if t.IsZero() {
		return ulid.ULID{}
	}
	// エントロピーソースの作成
	entropy := ulid.Monotonic(rand.Reader, 0)
	// ULIDの作成
	i := ulid.MustNew(ulid.Timestamp(t), entropy)
	return i
}
