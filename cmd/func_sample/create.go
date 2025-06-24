package func_sample

import (
	"github.com/oklog/ulid"
	"math/rand"
	"time"
)

// CreateULID は指定された時刻に基づいて新しいULIDを生成する関数です。
// エントロピーソースにより一意性を保証するための仕組みを含んでいます。
func CreateULID(t time.Time) ulid.ULID {
	// エントロピーソースの作成
	// ここでは簡易にmath/randを用いているが、厳密さを求めるなら乱数発生器は工夫した方が良い
	entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
	// ULIDの作成
	i := ulid.MustNew(ulid.Timestamp(t), entropy)
	return i
}

// ParseULID は文字列からULIDを解析し、ULID型とエラーメッセージを返します。
func ParseULID(s string) (ulid.ULID, error) {
	return ulid.Parse(s)
}
