package internal

import (
	at "github.com/Matej-Chmel/go-any-to-string"
)

// Returns options for any-to-string library
// with ByteAsString and RuneAsString flags set.
func GetOptions() *at.Options {
	o := at.NewOptions()
	o.ByteAsString = true
	o.RuneAsString = true
	return o
}
