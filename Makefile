GOROOT := $(shell go env GOROOT)

.PHONY: all
all: web server

.PHONY: web
web: configure_wasm
	GOOS=js GOARCH=wasm go build -o build/web/game.wasm cmd/web/main.go

.PHONY: desktop
desktop:
	go build -o build/game cmd/desktop/main.go

.PHONE: run
run: desktop
	./build/game

.PHONY: server
server:
	go build -o build/server cmd/server/main.go

.PHONY: configure_wasm
configure_wasm:
	cp $(GOROOT)/misc/wasm/wasm_exec.js build/web/
