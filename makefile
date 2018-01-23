start:
	vagrant up opnsense
	docker-compose up -d

rm:
	vagrant destroy
	docker-compose down -v

build: prepare
	go build -o dist/opn cmd/dw/dwaccount.go

release: prepare
	go build -o dist/opn-${VERSION} cmd/dw/dwaccount.go

prepare:
	glide install

init:
	brew install glide
	glide install
