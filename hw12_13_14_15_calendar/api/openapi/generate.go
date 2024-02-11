package grpc

//nolint:lll
//go:generate oapi-codegen -package openapi -generate types -old-config-style -o ../../pkg/api/openapi/types.gen.go ./schema.yaml
//go:generate oapi-codegen -package openapi -generate spec -old-config-style -o ../../pkg/api/openapi/spec.gen.go ./schema.yaml
//go:generate oapi-codegen -package openapi -generate chi-server -old-config-style -o ../../pkg/api/openapi/server.gen.go ./schema.yaml
//go:generate oapi-codegen -package openapi -generate client -old-config-style -o ../../pkg/api/openapi/client.gen.go ./schema.yaml
