package main

import "context"

// main is the CLI executable entrypoint.
func main()

// run parses CLI args and dispatches to command handlers.
func run(ctx context.Context, backend Backend, args []string) error
