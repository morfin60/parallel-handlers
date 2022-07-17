BINDIR			:= $(CURDIR)/bin
LDFLAGS			:= -w -s
GOFLAGS			:=
BINNAME			?= parallel-handlers

.PHONY: all
all: build

# ------------------------------------------------------------------------------
#  build
.PHONY: build
build: cmd/parallel-handlers/*.go
	GO111MODULE=on go build -o '$(BINDIR)'/$(BINNAME) ./cmd/parallel-handlers/*.go
