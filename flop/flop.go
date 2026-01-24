// Package flop provides a way to map cobra flags to functions using the With pattern that modify an underlying options
// ref: https://github.com/uber-go/guide/blob/master/style.md#functional-options
package flop

import (
	"github.com/spf13/cobra"
)

// FlagWith represents a link from a cobra flag to a function returning an
type FlagWith[T any] interface {
	Flag() string
	Get(cmd *cobra.Command) T
}

// With should be used as the container for collecting multiple FlagWith instances using its Get method to retrieve the related options
type With[T any] []FlagWith[T]

// Get collects all the options that have been set for the given flags
func (fos With[T]) Get(cmd *cobra.Command) []T {
	opts := make([]T, 0, len(fos))
	for _, opt := range fos {
		if cmd.Flags().Changed(opt.Flag()) {
			opts = append(opts, opt.Get(cmd))
		}
	}
	return opts
}
