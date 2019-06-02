all: bin/kubeselect

bin:
	mkdir -p bin

bin/kubeselect: $(shell find . -name '*.go') go.mod go.sum
	cd cmd/kubeselect && go build -o ../../$@

clean:
	rm -rf bin dist

snapshot:
	goreleaser release --skip-publish --skip-validate --rm-dist --snapshot

test:
	go test -v ./...

install: bin/kubeselect
	cp bin/kubeselect ~/bin/

.PHONY: clean all test install snapshot
