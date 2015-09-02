package yarc

import (
	"log"
	"strings"

	"github.com/garyburd/redigo/redis"
)

type Logger func(cmd string, args ...interface{})

func DefaultLogger(cmd string, args ...interface{}) {
	log.Printf("[yarc] %s %s", cmd, args)
}

type Options struct {
	KeyPrefix    string
	KeyDelimiter string
	Logger       Logger
}

type Client struct {
	pool    *redis.Pool
	options *Options
}

func NewClient(address string, options *Options) *Client {
	pool := &redis.Pool{
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", address)
		},
	}
	return NewClientWithPool(pool, options)
}

func NewClientWithPool(pool *redis.Pool, options *Options) *Client {
	return &Client{pool, options}
}

func (c *Client) Key(parts ...string) string {
	if c.options.KeyPrefix != "" {
		parts = append([]string{c.options.KeyPrefix}, parts...)
	}
	return strings.Join(parts, c.options.KeyDelimiter)
}

func (c *Client) Hash(key string) *Hash {
	return &Hash{Key: key, client: c}
}

func (c *Client) SortedSet(key string) *SortedSet {
	return &SortedSet{Key: key, client: c}
}

func (c *Client) do(cmd string, args ...interface{}) (interface{}, error) {
	conn := c.pool.Get()
	if c.options.Logger != nil {
		c.options.Logger(cmd, args...)
	}
	return conn.Do(cmd, args...)
}

/*
func encode(value interface{}) []byte {
	b, err := json.Marshal(value)
	if err != nil {
		panic(err)
	}
	return b
}

func decode(value []byte, dest interface{}) {
	json.Unmarshal(value, dest)
}
*/
