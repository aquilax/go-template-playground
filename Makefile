GOPHERJS_GOROOT="$(go env GOROOT)"

public/tgo.js: gopherjs/tgo.go
	gopherjs build -m -o public/tgo.js gopherjs/tgo.go
