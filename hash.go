package yarc

import "github.com/garyburd/redigo/redis"

type HashType struct {
	Key    string
	client *Client
}

func (h *HashType) Len() int64 {
	reply, err := h.client.do("HLEN", h.Key)
	length, _ := redis.Int64(reply, err)
	return length
}

func (h *HashType) Keys() []string {
	reply, err := h.client.do("HKEYS", h.Key)
	keys, _ := redis.Strings(reply, err)
	return keys
}

func (h *HashType) Vals() []string {
	reply, err := h.client.do("HVALS", h.Key)
	keys, _ := redis.Strings(reply, err)
	return keys
}

func (h *HashType) Get(key string) string {
	reply, err := h.client.do("HGET", h.Key, key)
	result, _ := redis.String(reply, err)
	return result
}

func (h *HashType) Set(key, value string) {
	h.client.do("HSET", h.Key, key, value)
}

func (h *HashType) Del(keys ...string) {
	args := []interface{}{h.Key}
	for _, key := range keys {
		args = append(args, key)
	}
	h.client.do("HDEL", args...)
}

func (h *HashType) Update(values map[string]string) {
	args := []interface{}{h.Key}
	for key, value := range values {
		args = append(args, key, value)
	}
	h.client.do("HMSET", args...)
}
