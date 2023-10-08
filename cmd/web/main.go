package main

import (
	"context"
	"errors"
	"io"
	"log"

	"syscall/js"

	"github.com/senior-sigan/alyoep/game"
)

// see: https://github.com/tarndt/wasmws

type WebSocket struct {
	ctx       context.Context
	ctxCancel context.CancelFunc

	URL string
	ws  js.Value

	readCh chan io.Reader
	errCh  chan error
	openCh chan struct{}

	cleanup []func()
}

func init() {
	ctx, cancel := context.WithCancel(context.Background())
	URL := "ws://localhost:8000/ws"
	ws := WebSocket{
		ctx:       ctx,
		ctxCancel: cancel,
		URL:       URL,
		ws:        js.Global().Get("WebSocket").New(URL),

		cleanup: make([]func(), 0, 3),

		readCh: make(chan io.Reader, 8),
		openCh: make(chan struct{}),
		errCh:  make(chan error, 1),
	}

	ws.addHandler(ws.handleOpen, "open")
	ws.addHandler(ws.handleClose, "close")
	ws.addHandler(ws.handleError, "error")
	ws.addHandler(ws.handleMessage, "message")
}

// handleOpen is a callback for JavaScript to notify Go when the websocket is open:
// See: https://developer.mozilla.org/en-US/docs/Web/API/WebSocket/onopen
func (ws *WebSocket) handleOpen(_ js.Value, _ []js.Value) {
	println("Websocket: Open JS callback!")
	close(ws.openCh)
}

// handleClose is a callback for JavaScript to notify Go when the websocket is closed:
// See: https://developer.mozilla.org/en-US/docs/Web/API/WebSocket/onclose
func (ws *WebSocket) handleClose(_ js.Value, _ []js.Value) {
	println("Websocket: Close JS callback!")
	ws.ctxCancel()
}

// handleError is a callback for JavaScript to notify Go when the websocket is in an error state:
// See: https://developer.mozilla.org/en-US/docs/Web/API/WebSocket/onerror
func (ws *WebSocket) handleError(_ js.Value, args []js.Value) {
	println("Websocket: Error JS Callback")
	errMsg := "Unknown error"
	if len(args) > 0 {
		errMsg = args[0].String()
	}

	select {
	case ws.errCh <- errors.New(errMsg):
	default:
	}
}

func (ws *WebSocket) handleMessage(_ js.Value, args []js.Value) {
	println("Websocket: New Message JS Callback")
}

func (ws *WebSocket) addHandler(handler func(this js.Value, args []js.Value), event string) {
	jsHandler := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		handler(this, args)
		return nil
	})
	cleanup := func() {
		ws.ws.Call("removeEventListener", event, jsHandler)
		jsHandler.Release()
	}
	ws.ws.Call("addEventListener", event, jsHandler)
	ws.cleanup = append(ws.cleanup, cleanup)
}

func main() {
	log.Println("Hello from web game")
	game.RunApp()
}
