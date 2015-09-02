package yarc

var client = NewClient("localhost:6379", &Options{
	KeyPrefix:    "yarctest",
	KeyDelimiter: ":",
})
