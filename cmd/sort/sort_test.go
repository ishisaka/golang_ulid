/*
 * Copyright (c) 2025. Tadahiro Ishisaka All rights reserved.
 *
 * ULIDのソートサンプル
 */

package sort

import (
	"math/rand"
	"testing"
	"time"

	"github.com/oklog/ulid"
	"github.com/stretchr/testify/assert"
)

func TestSortULIDs(t *testing.T) {
	t.Run("empty slice", func(t *testing.T) {
		ids := []ulid.ULID{}
		SortULIDs(ids)
		assert.Empty(t, ids, "expected empty slice")
	})

	t.Run("single element", func(t *testing.T) {
		id := ulid.MustNew(ulid.Now(), rand.New(rand.NewSource(time.Now().UnixNano())))
		ids := []ulid.ULID{id}
		SortULIDs(ids)
		assert.Len(t, ids, 1, "expected single element slice to remain unchanged")
		assert.Equal(t, id, ids[0], "expected element to remain unchanged")
	})

	t.Run("unsorted slice", func(t *testing.T) {
		id1 := ulid.MustNew(ulid.Now(), rand.New(rand.NewSource(time.Now().UnixNano())))
		id2 := ulid.MustNew(ulid.Now()-1, rand.New(rand.NewSource(time.Now().UnixNano())))
		id3 := ulid.MustNew(ulid.Now()-2, rand.New(rand.NewSource(time.Now().UnixNano())))

		ids := []ulid.ULID{id1, id3, id2}
		expected := []ulid.ULID{id3, id2, id1}

		SortULIDs(ids)

		assert.Equal(t, expected, ids, "expected slice to be sorted")
	})
}

func TestFindLatestULID(t *testing.T) {
	t.Run("empty slice", func(t *testing.T) {
		ids := []ulid.ULID{}
		latest, found := FindLatestULID(ids)
		assert.False(t, found, "expected no ulid to be found in empty slice")
		assert.Equal(t, ulid.ULID{}, latest, "expected empty ULID when no ULID is found")
	})

	t.Run("single element", func(t *testing.T) {
		id := ulid.MustNew(ulid.Now(), rand.New(rand.NewSource(time.Now().UnixNano())))
		ids := []ulid.ULID{id}
		latest, found := FindLatestULID(ids)
		assert.True(t, found, "expected a ULID to be found")
		assert.Equal(t, id, latest, "expected only element to be found as latest")
	})

	t.Run("multiple elements", func(t *testing.T) {
		id1 := ulid.MustNew(ulid.Now()-2, rand.New(rand.NewSource(time.Now().UnixNano())))
		id2 := ulid.MustNew(ulid.Now()-1, rand.New(rand.NewSource(time.Now().UnixNano())))
		id3 := ulid.MustNew(ulid.Now(), rand.New(rand.NewSource(time.Now().UnixNano())))

		ids := []ulid.ULID{id1, id2, id3}
		latest, found := FindLatestULID(ids)

		assert.True(t, found, "expected a ULID to be found")
		assert.Equal(t, id3, latest, "expected the latest ULID to be found")
	})
}
