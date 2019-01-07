NAME := tree-go
SRCS := $(shell find ./ -type f -name '*.go')
VERSION := v1.0.0
REVISION := $(shell git rev-parse --short HEAD)
PKGPATH := github.com/mi-bear/tree-go
LDFLAGS := -ldflags \
	'-s -w \
	-X "$(PKGPATH)/cmd.Version=$(VERSION)" \
	-X "$(PKGPATH)/cmd.Revision=$(REVISION)" \
	-extldflags "-static"'

DIST_DIRS := find * -type d -exec

$(NAME): $(SRCS)
	go build -a -tags netgo -installsuffix netgo $(LDFLAGS) -o $(NAME)

.PHONY: install
install:
	go install $(LDFLAGS)

.PHONY: clean
clean:
	rm -rf $(NAME)
