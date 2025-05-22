// Developed by DevPros with Codex as supporting tool
// DO NOT EDIT MANUALLY
package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/devprosvn/VNPrider/pkg/mnemonic"
)

func main() {
	flag.Parse()
	if len(flag.Args()) == 0 {
		fmt.Println("usage: vnprider <command>")
		return
	}
	cmd := flag.Arg(0)
	switch cmd {
	case "status":
		fmt.Println("node running")
	case "mnemonic":
		fmt.Println(mnemonic.GenerateMnemonic())
	case "validate":
		if len(flag.Args()) < 2 {
			fmt.Println("usage: vnprider validate <mnemonic>")
			return
		}
		if mnemonic.ValidateMnemonic(flag.Arg(1)) {
			fmt.Println("valid")
		} else {
			fmt.Println("invalid")
		}
	default:
		fmt.Println("unknown command")
	}
	log.Println("vnprider-cli done")
}
