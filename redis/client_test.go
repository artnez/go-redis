package redis

func testClient() *Client {
	return NewClient("redis://localhost:6379", Options{
		KeyPrefix:    "redistest",
		KeyDelimiter: ":",
	})
}
