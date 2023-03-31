VERSION=$(shell cat VERSION)
BUILD_DATE=$(shell date -u +"%Y-%m-%dT%H:%M:%SZ")
GIT_HASH=$(shell git rev-parse HEAD)

TARGETS=linux.amd64 linux.386 linux.arm64 linux.mips64 windows.amd64.exe darwin.amd64 darwin.arm64 freebsd.amd64
BINARIES=$(addprefix bin/go-on-$(VERSION)., $(TARGETS)) $(addprefix bin/date-delta-$(VERSION)., $(TARGETS))
RELEASES=$(subst windows.amd64.tar.gz,windows.amd64.zip,$(foreach r,$(subst .exe,,$(TARGETS)),releases/go-on-$(VERSION).$(r).tar.gz))

LDFLAGS=-X main.version=$(VERSION) -X main.buildDate=$(BUILD_DATE) -X main.gitHash=$(GIT_HASH)

bin/go-on: bin
	go build -o $@ ./cmd/go-on

bin/date-delta: bin
	go build -o $@ ./cmd/date-delta

binaries: $(BINARIES)
releases: $(RELEASES)
	make $(RELEASES)

bin/go-on-$(VERSION).%:
	env GOARCH=$(subst .,,$(suffix $(subst .exe,,$@))) \
		GOOS=$(subst .,,$(suffix $(basename $(subst .exe,,$@)))) \
		CGO_ENABLED=0 \
		go build -ldflags "$(LDFLAGS)" -o $@ ./cmd/go-on

bin/date-delta-$(VERSION).%:
	env GOARCH=$(subst .,,$(suffix $(subst .exe,,$@))) \
		GOOS=$(subst .,,$(suffix $(basename $(subst .exe,,$@)))) \
		CGO_ENABLED=0 \
		go build -ldflags "$(LDFLAGS)" -o $@ ./cmd/date-delta

releases/2h5tils-$(VERSION).%.zip: bin/*-$(VERSION).%.exe
	mkdir -p releases
	zip -9 -j -r $@ README.md $<

releases/2h5tils-$(VERSION).%.tar.gz: bin/*-$(VERSION).%
	mkdir -p releases
	tar -cf $(basename $@) README.md && \
		tar -rf $(basename $@) --strip-components 1 $< && \
		gzip -9 $(basename $@)

deps-vendor:
	go mod vendor
deps-cleanup:
	go mod tidy

run-tests:
	go test -v ./cmd/...

bin:
	mkdir $@

report: report-staticcheck report-ineffassign report-vet 
report: report-cyclo
report: report-mispell
report: report-golangci-lint
report-cyclo:
	@echo '####################################################################'
	gocyclo ./cmd/date-delta/* ./cmd/go-on/*
report-mispell:
	@echo '####################################################################'
	misspell ./cmd/...
report-lint:
	@echo '####################################################################'
	golint ./cmd/...
report-ineffassign:
	@echo '####################################################################'
	ineffassign ./cmd/...
report-vet:
	@echo '####################################################################'
	go vet ./cmd/...
report-staticcheck:
	@echo '####################################################################'
	staticcheck ./cmd/...
report-golangci-lint:
	@echo '####################################################################'
	golangci-lint run

fetch-report-tools:
	go install github.com/fzipp/gocyclo/cmd/gocyclo@latest
	go install github.com/client9/misspell/cmd/misspell@latest
	go install github.com/gordonklaus/ineffassign@latest
	go install honnef.co/go/tools/cmd/staticcheck@latest
	go install golang.org/x/lint/golint@latest

.PHONY: bin/go-on bin/date-delta
