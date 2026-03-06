module loglint/cmd/golangci-plugin

go 1.26

require (
	github.com/golangci/plugin-module-register v0.1.2
	golang.org/x/tools v0.42.0
	loglint v0.0.0
)

replace loglint => ../..