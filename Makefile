start:
	vagrant up opnsense
	docker-compose up -d

rm:
	vagrant destroy
	docker-compose down -v

build:
	go build -o dist/opn-macos-amd64 opn.go
	env GOSS=linux go build -o dist/opn-linux-amd64 opn.go

release:
	go build -o dist/opn-macos-amd64-${VERSION} opn.go
	env GOSS=linux go build -o dist/opn-linux-amd64-${VERSION} opn.go
