// (c) 2025 Tadahiro Ishisaka all rights reserved
// ULIDのソートサンプル

package sort

import (
	"github.com/oklog/ulid"
	"sort"
)

// SortULIDs はULIDのスライスを時系列順に並び替えます
func SortULIDs(ids []ulid.ULID) {
	sort.Slice(ids, func(i, j int) bool {
		return ids[i].Compare(ids[j]) < 0
	})
}

// FindLatestULID は与えられたULIDの中で最も新しいものを返します
func FindLatestULID(ids []ulid.ULID) (ulid.ULID, bool) {
	if len(ids) == 0 {
		return ulid.ULID{}, false
	}

	latest := ids[0]
	for _, id := range ids[1:] {
		if id.Compare(latest) > 0 {
			latest = id
		}
	}

	return latest, true
}
