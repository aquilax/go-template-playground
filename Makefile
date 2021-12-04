GOPHERJS_GOROOT="$(go env GOROOT)"

public/tgo.js: gopherjs/tgo.go
	gopherjs -m build -o public/tgo.js gopherjs/tgo.go
