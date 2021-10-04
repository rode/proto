MAKEFLAGS += --silent
.PHONY: lint breaking test generate

lint:
	buf lint

breaking:
	buf breaking --against '.git#branch=main'

test: lint breaking

generate:
	buf generate
