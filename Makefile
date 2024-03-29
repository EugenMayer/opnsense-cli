init:
	go mod tidy
	go mod verify
	go mod vendor

update:
	go get -u
	go mod tidy

start:
	vagrant up opnsense
	docker-compose up -d

rm:
	vagrant destroy
	docker-compose down -v

build: init
	go build -o dist/opn-macos-amd64 opn.go
	env GOSS=linux go build -o dist/opn-linux-amd64 opn.go
	env GOSS=windows go build -o dist/opn-windows-amd64 opn.go
release:
	go build -o dist/opn-macos-amd64-${VERSION} opn.go
	env GOSS=linux go build -o dist/opn-linux-amd64-${VERSION} opn.go
