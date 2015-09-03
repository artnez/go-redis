package yarc

var client *Client

func GlobalInit(globalClient *Client) {
	client = globalClient
}

func Hash(key string) *HashType {
	return client.Hash(key)
}

func SortedSet(key string) *SortedSetType {
	return client.SortedSet(key)
}
