image:
	# Test Packages
	go test -v ./...

	# Build Binary
	CGO_ENABLED=0 GOOS=linux go build -a -tags netgo -ldflags '-w' -o build/gogogo .

	# Create Docker Image
	docker build --rm --no-cache --pull=true -t baljanak/gogogo .

clean:
	rm -f build/gogogo
	docker image prune -f
