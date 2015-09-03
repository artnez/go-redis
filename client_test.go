package yarc

import (
	"flag"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	Init("redis://localhost:6379", &Options{
		KeyPrefix:    "yarctest",
		KeyDelimiter: ":",
	})

	flag.Parse()
	os.Exit(m.Run())
}
