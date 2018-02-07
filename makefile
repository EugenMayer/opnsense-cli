start:
	vagrant up opnsense
	docker-compose up -d

rm:
	vagrant destroy
	docker-compose down -v

build: prepare
	go build -o dist/opn-macos-amd64 opn.go
	env GOSS=linux go build -o dist/opn-linux-amd64 opn.go

release: prepare
	go build -o dist/opn-macos-amd64-${VERSION} opn.go
	env GOSS=linux go build -o dist/opn-linux-amd64-${VERSION} opn.go

prepare:
	glide install

init:
	brew install glide
	glide install
