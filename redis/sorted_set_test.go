package redis

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func getSortedSet() *SortedSetType {
	return client.SortedSet("test", "sortedset")
}

func TestSortedSetSet(t *testing.T) {
	s := getSortedSet()
	s.Set("foo1", 0.1)
	s.Set("foo2", 0.2)
	s.Set("foo3", 0.3)
	assert.Equal(t, s.Score("foo1"), 0.1)
	assert.Equal(t, s.Score("foo2"), 0.2)
	assert.Equal(t, s.Score("foo3"), 0.3)
	s.Del("foo1", "foo2", "foo3")
}

func TestSortedSetUpdate(t *testing.T) {
	s := getSortedSet()
	s.Update(map[string]float64{
		"foo1": 0.4,
		"foo2": 0.5,
		"foo3": 0.6,
	})
	assert.Equal(t, s.Score("foo1"), 0.4)
	assert.Equal(t, s.Score("foo2"), 0.5)
	assert.Equal(t, s.Score("foo3"), 0.6)
	s.Del("foo1", "foo2", "foo3")
}

func TestSortedSetLen(t *testing.T) {
	s := getSortedSet()
	s.Update(map[string]float64{
		"foo1": 0.4,
		"foo2": 0.5,
		"foo3": 0.6,
	})
	assert.Equal(t, int64(3), s.Len())
	s.Set("foo4", 0.7)
	assert.Equal(t, int64(4), s.Len())
	s.Del("foo1", "foo2", "foo3", "foo4")
}

func TestSortedSetAll(t *testing.T) {
	s := getSortedSet()
	s.Update(map[string]float64{
		"foo1": 0.1,
		"foo3": 0.3,
		"foo2": 0.2,
	})

	all := s.All()
	assert.Equal(t, []float64{0.1, 0.2, 0.3}, all.Scores())
	assert.Equal(t, []string{"foo1", "foo2", "foo3"}, all.Members())

	s.Del("foo1", "foo2", "foo3")
}
