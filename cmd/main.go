package main

import (
	"context"
	"fmt"
	"os"

	"github.com/slingercode/go-http/internal"
)

func main() {
	ctx := context.Background()

	if err := internal.Run(ctx, os.Stderr); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
