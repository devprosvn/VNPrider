// Developed by DevPros with Codex as supporting tool
// DO NOT EDIT MANUALLY
package internal

// Package internal provides utility helpers for the command packages.

// CheckErr panics if err is not nil.
func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
