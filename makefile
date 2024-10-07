all:
	go build ./cmd/host

host:
	go build ./cmd/host

clean:
	rm -fv host