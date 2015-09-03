package redis

var client *Client

func Init(address string, options *Options) {
	client = NewClient(address, options)
}

func Hash(key string) *HashType {
	return client.Hash(key)
}

func SortedSet(key string) *SortedSetType {
	return client.SortedSet(key)
}
