// Developed by DevPros with Codex as supporting tool
// DO NOT EDIT MANUALLY
package internal

import "log"

// CheckErr terminates the program if err is not nil.
func CheckErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
