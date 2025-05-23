// Developed by DevPros with Codex as supporting tool
// DO NOT EDIT MANUALLY

package main

import (
	"context"
	"log"
	"os"

	"github.com/devprosvn/VNPrider/pkg/node"
)

var runNodeFn = node.Run
var logFatal = log.Fatal

func main() {
	log.Println("vnprider-node starting...")
	ctx := context.Background()
	if err := runNodeFn(ctx); err != nil {
		if os.Getenv("TESTING") != "" {
			panic(err)
		}
		logFatal(err)
	}
}
