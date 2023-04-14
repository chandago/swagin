# Integration tests

This directory contains integrations test to run with [Venom](https://github.com/intercloud/venom).

To run integration tests:

- Go to project root directory
- Build example application with `go build -o example ./examples`
- Start example application with `./example &`
- Run test in Venom with `venom run examples/test/*.yml`
