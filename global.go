package yarc

var client *Client

func GlobalInit(globalClient *Client) {
	client = globalClient
}

func Key(parts ...string) string {
	return client.Key(parts...)
}

func Hash(key string) *HashType {
	return client.Hash(key)
}

func SortedSet(key string) *SortedSetType {
	return client.SortedSet(key)
}
