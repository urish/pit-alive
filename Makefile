bin/pit-alive: cmd/pit-alive/main.go
	  GOOS=linux GOARCH=arm go build -o $@ $<

install:
		cp ./bin/pit-alive /usr/local/bin
