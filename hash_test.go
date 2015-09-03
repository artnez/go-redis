package yarc

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func getHash() *HashType {
	return client.Hash(client.Key("test", "hash"))
}

func TestHashGetSetDel(t *testing.T) {
	h := getHash()
	h.Set("foo", "bar")
	assert.Equal(t, "bar", h.Get("foo"))
	h.Del("foo")
	assert.Equal(t, "", h.Get("foo"))
}

func TestHashKeysAndValues(t *testing.T) {
	h := getHash()
	h.Set("foo1", "bar1")
	h.Set("foo2", "bar2")
	h.Set("foo3", "bar3")
	assert.EqualValues(t, []string{"foo1", "foo2", "foo3"}, h.Keys())
	assert.EqualValues(t, []string{"bar1", "bar2", "bar3"}, h.Vals())
	h.Del("foo1", "foo2", "foo3")
	assert.Equal(t, []string{}, h.Keys())
	assert.Equal(t, []string{}, h.Vals())
}

func TestHashUpdate(t *testing.T) {
	h := getHash()
	h.Update(map[string]string{
		"foo1": "bar1",
		"foo2": "bar2",
		"foo3": "bar3",
	})
	keys := h.Keys()
	sort.Strings(keys)
	vals := h.Vals()
	sort.Strings(vals)
	assert.EqualValues(t, []string{"foo1", "foo2", "foo3"}, keys)
	assert.EqualValues(t, []string{"bar1", "bar2", "bar3"}, vals)
	h.Del("foo1", "foo2", "foo3")
	assert.Equal(t, []string{}, h.Keys())
	assert.Equal(t, []string{}, h.Vals())
}

func TestHashLen(t *testing.T) {
	h := getHash()
	h.Update(map[string]string{
		"foo1": "bar1",
		"foo2": "bar2",
		"foo3": "bar3",
	})
	assert.Equal(t, int64(3), h.Len())
	h.Set("foo4", "bar4")
	assert.Equal(t, int64(4), h.Len())
	h.Del("foo1", "foo2", "foo3", "foo4")
}
