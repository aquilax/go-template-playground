//go:build js && wasm

package main

import (
	"syscall/js"
)

func main() {
	js.Global().Set("renderTemplate", js.FuncOf(renderJS))
	js.Global().Set("compress", js.FuncOf(compressJS))
	js.Global().Set("decompress", js.FuncOf(decompressJS))

	select {}
}

func renderJS(this js.Value, args []js.Value) any {
	if len(args) < 2 {
		return []any{"", "invalid arguments"}
	}
	out, errStr := render(args[0].String(), args[1].String())
	return []any{out, errStr}
}

func compressJS(this js.Value, args []js.Value) any {
	if len(args) < 1 {
		return []any{"", "invalid arguments"}
	}
	out, errStr := compress(args[0].String())
	return []any{out, errStr}
}

func decompressJS(this js.Value, args []js.Value) any {
	if len(args) < 1 {
		return []any{"", "invalid arguments"}
	}
	out, errStr := decompress(args[0].String())
	return []any{out, errStr}
}
