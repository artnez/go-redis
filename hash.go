package yarc

import "github.com/garyburd/redigo/redis"

type Hash struct {
	Key    string
	client *Client
}

func (h *Hash) Len() int64 {
	reply, err := h.client.do("HLEN", h.Key)
	length, _ := redis.Int64(reply, err)
	return length
}

func (h *Hash) Keys() []string {
	reply, err := h.client.do("HKEYS", h.Key)
	keys, _ := redis.Strings(reply, err)
	return keys
}

func (h *Hash) Vals() []string {
	reply, err := h.client.do("HVALS", h.Key)
	keys, _ := redis.Strings(reply, err)
	return keys
}

func (h *Hash) Get(key string) string {
	reply, err := h.client.do("HGET", h.Key, key)
	result, _ := redis.String(reply, err)
	return result
}

func (h *Hash) Set(key, value string) {
	h.client.do("HSET", h.Key, key, value)
}

func (h *Hash) Del(keys ...string) {
	args := []interface{}{h.Key}
	for _, key := range keys {
		args = append(args, key)
	}
	h.client.do("HDEL", args...)
}

func (h *Hash) Update(values map[string]string) {
	args := []interface{}{h.Key}
	for key, value := range values {
		args = append(args, key, value)
	}
	h.client.do("HMSET", args...)
}
