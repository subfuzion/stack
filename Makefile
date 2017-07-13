.PHONY: vendor clean test build-image run

PKG := github.com/subfuzion/stack/cmd
TARGET := bin/stackcli
IMAGE := subfuzion/stackcli

build:
	go build -o $(TARGET) $(PKG)

vendor: vendor.conf
	vndr

clean:
	rm -f $(TARGET)

test:
	go test -v -timeout 5m

image:
	docker build -t $(IMAGE) .

run:
	docker run -t --rm -v /var/run/docker:/var/run/docker:ro $(IMAGE) $${@}

