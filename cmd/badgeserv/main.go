package main

import (
	"os"

	"github.com/yctomwang/badgeserv/pkg/entrypoint"
)

func main() {
	exitCode := entrypoint.Entrypoint(os.Stdout, os.Stderr)
	os.Exit(exitCode)
}
