package flagslice

import "strings"

// FlagSlice provides a []string flag.Var type to handle duplicate options.
type FlagSlice []string

// String satisfies the flag.Value interface.
func (f *FlagSlice) String() string {
	return ""
}

// Set handles duplicate Set calls by appending to the slice.
func (f *FlagSlice) Set(value string) error {
	*f = append(*f, strings.TrimSpace(value))
	return nil
}
