#binary name
BINARY_NAME=fast

#installation directory
INSTALL_DIR=$(HOME)/bin

all: build

build:
	@go build -o fast

run: build
	~/bin/fast

install: build
	cp $(BINARY_NAME) $(INSTALL_DIR)