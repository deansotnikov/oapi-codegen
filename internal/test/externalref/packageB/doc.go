package packageB

//go:generate go run github.com/deansotnikov/oapi-codegen/cmd/oapi-codegen -generate types,skip-prune --package=packageB -o externalref.gen.go spec.yaml
