image:
	CGO_ENABLED=0 GOOS=linux go build -a -tags netgo -ldflags '-w' -o build/gogogo .
	docker build --no-cache --pull=true -t baljanak/gogogo .

clean:
	rm -f build/gogogo
