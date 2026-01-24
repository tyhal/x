package flop

import "github.com/spf13/cobra"

type toggleFlagOption[T any] struct {
	flag string
	f    func() T
}

// Toggle is like a bool flag but doesn't have the bool as a parameter
func Toggle[T any](flag string, f func() T) FlagWith[T] {
	return &toggleFlagOption[T]{flag: flag, f: f}
}

func (tfo toggleFlagOption[T]) Flag() string {
	return tfo.flag
}

func (tfo toggleFlagOption[T]) Get(cmd *cobra.Command) T {
	if b, err := cmd.Flags().GetBool(tfo.flag); err == nil && b {
		return tfo.f()
	}
	var t T
	return t
}
