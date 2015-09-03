package redis

import (
	"flag"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	Init("redis://localhost:6379", &Options{
		KeyPrefix:    "redistest",
		KeyDelimiter: ":",
	})

	flag.Parse()
	os.Exit(m.Run())
}
