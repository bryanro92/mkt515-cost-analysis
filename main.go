package main

import (
	"context"
	"log"

	"github.com/bryanro92/mkt515-cost-analysis/pkg/analysis"
)

func main() {
	ctx := context.Background()

	// Enter into our program, exit on error
	if err := analysis.Run(ctx); err != nil {
		log.Fatal(err)
	}
}
