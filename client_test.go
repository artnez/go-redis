package yarc

import (
	"flag"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	client := NewClient("localhost:6379", &Options{
		KeyPrefix:    "yarctest",
		KeyDelimiter: ":",
	})
	GlobalInit(client)

	flag.Parse()
	os.Exit(m.Run())
}
