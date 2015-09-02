cover:
	go test -coverprofile cover.out	
	go tool cover -html=cover.out -o cover.html
	open cover.html
	sleep 1
	rm -f cover.out
	rm -f cover.html

.PHONY: cover
