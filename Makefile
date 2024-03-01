build:
	go build -o tx-relay
run: build
	./tx-relay start
clean:
	rm tx-relay