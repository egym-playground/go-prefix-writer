[![Build Status](https://travis-ci.org/egymgmbh/go-prefix-writer.svg?branch=master)](https://travis-ci.org/egymgmbh/go-prefix-writer)
# go-prefix-writer
An io.Writer which adds a prefix to every line written.

## How to use
`prefixer.Writer` implements the `io.Writer` interface and can be constructed with the
`prefixer.New(io.Writer, func() string)` function which takes the target `io.Writer` as well as a
function which sets the prefix. The prefix is not a constant in order to allow dynamic prefixes
like timestamps.

## Example:
An example can be found in [example/main.go](example/main.go)
