/*
 * Copyright (c) 2025. Tadahiro Ishisaka All rights reserved.
 */

// ULID作成のためのシンプルなサンプル

package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/oklog/ulid"
)

// ExampleULID は、新しいULIDを生成して返します。
// エントロピーソースには引数の時間に基づくモノトニックな乱数を使用します。
// ULIDのタイムスタンプは指定された時間を基準に設定されます。
func ExampleULID() ulid.ULID {
	t := time.Unix(1000000, 0)
	// エントロピーソースの作成
	entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
	// ULIDの作成
	i := ulid.MustNew(ulid.Timestamp(t), entropy)
	return i

}

func main() {
	i := ExampleULID()
	fmt.Println(i)
	// Output: 0000XSNJG0MQJHBF4QX1EFD6Y3
}
