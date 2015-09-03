package yarc

import "github.com/garyburd/redigo/redis"

type SortedSetItems []SortedSetItem

func (s SortedSetItems) Scores() []float64 {
	scores := []float64{}
	for _, item := range s {
		scores = append(scores, item.Score)
	}
	return scores
}

func (s SortedSetItems) Members() []string {
	members := []string{}
	for _, item := range s {
		members = append(members, item.Member)
	}
	return members
}

type SortedSetItem struct {
	Member string
	Score  float64
}

type SortedSet struct {
	Key    string
	client *Client
}

func (s *SortedSet) Len() int64 {
	reply, err := s.client.do("ZCARD", s.Key)
	length, _ := redis.Int64(reply, err)
	return length
}

func (s *SortedSet) Score(member string) float64 {
	reply, err := s.client.do("ZSCORE", s.Key, member)
	score, _ := redis.Float64(reply, err)
	return score
}

func (s *SortedSet) All() SortedSetItems {
	reply, err := s.client.do("ZRANGE", s.Key, 0, -1, "WITHSCORES")
	values, err := redis.Values(reply, err)
	if err != nil {
		return SortedSetItems{}
	}
	items := make([]SortedSetItem, len(values)/2)
	for i := range items {
		values, _ = redis.Scan(values, &items[i].Member, &items[i].Score)
	}
	return SortedSetItems(items)
}

func (s *SortedSet) Set(member string, score float64) {
	s.client.do("ZADD", s.Key, score, member)
}

func (s *SortedSet) Del(members ...string) {
	args := []interface{}{s.Key}
	for _, member := range members {
		args = append(args, member)
	}
	s.client.do("ZREM", args...)
}

func (s *SortedSet) Update(values map[string]float64) {
	args := []interface{}{s.Key}
	for member, score := range values {
		args = append(args, score, member)
	}
	s.client.do("ZADD", args...)
}
