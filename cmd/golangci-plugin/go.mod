module github.com/GrimVeqsum/GoLinter/cmd/golangci-plugin

go 1.26

require (
	github.com/GrimVeqsum/GoLinter v0.0.0
	github.com/golangci/plugin-module-register v0.1.2
	golang.org/x/tools v0.42.0
)

replace github.com/GrimVeqsum/GoLinter => ../..
