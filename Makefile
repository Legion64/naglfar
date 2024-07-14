.PHONY: build
build:
ifeq (,$(wildcard ./build/))
	mkdir build
endif
	go build -o ./build/ ./cmd/auth/main.go

.PHONY: run
run:
	go run ./cmd/auth/main.go