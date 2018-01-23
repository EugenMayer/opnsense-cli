start:
	vagrant up opnsense
	docker-compose up -d

rm:
	vagrant destroy
	docker-compose down -v

build: prepare
	go build -o dist/opn opn.go

release: prepare
	go build -o dist/opn-${VERSION} opn.go

prepare:
	glide install

init:
	brew install glide
	glide install
