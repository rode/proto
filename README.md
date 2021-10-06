# proto

[![test](https://github.com/rode/proto/actions/workflows/test.yaml/badge.svg)](https://github.com/rode/proto/actions/workflows/test.yaml)

Experimental Rode protobufs. 

## Development

This project uses [`buf`](https://docs.buf.build/installation) to manage protobuf definitions.

There are two configuration files for `buf`:

- `buf.yaml`: The settings for the linter and breaking change detection
- `buf.gen.yaml`: Manages `protoc` plugins as well as protobuf dependencies.

With `buf` installed, use the following Make targets to work locally: 

- `make lint`: Run the `buf` linter.
- `make breaking`: Check for any breaking changes against definitions in `main`.
- `make test`: Run all quality checks (lint and breaking change detection). 
- `make generate`: Output Go code from the protobuf definitions

Generated code is placed in `gen`, with Go code nested under `gen/go`.

### Ignoring Lint Violations

Exclude an entire file from a particular rule by setting `ignore_only` in `buf.yaml`. You can also use `buf` to generate
the configuration: `buf lint --error-format config-ignore-yaml`. 

Alternatively, for a violation in one file, use a leading comment:

```protobuf
// buf:lint:ignore FIELD_LOWER_SNAKE_CASE
message foo {}
```

To fix multiple violations, use a another comment line. A list of rule names is available [here](https://docs.buf.build/configuration/v1beta1/lint-rules/).

### Add a New `protoc` Plugin

To add a new `protoc` plugin, add it under the `plugins` list in `buf.gen.yaml`. 
Use the name without the `protoc-gen` prefix, e.g., `grpc-gateway` instead of `protoc-gen-grpc-gateway`. 

### Add an External Protobuf Dependency

1. Find the dependency in the [Buf Schema Registry](https://buf.build/explore)
1. Add it under the `dep` list in `buf.yaml`
1. Run `buf mod update` to update the `buf.lock` file

Note that the commit sha in `buf.lock` corresponds to the registry's own Git mirror and not the 
source repository. 
