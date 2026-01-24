package flop

import "github.com/spf13/cobra"

type stringFlagOption[T any] struct {
	flag string
	f    func(string) T
}

// String for string flags that map to a WithSomething(string) func
func String[T any](flag string, f func(string) T) FlagWith[T] {
	return &stringFlagOption[T]{flag: flag, f: f}
}

func (sfo stringFlagOption[T]) Flag() string {
	return sfo.flag
}

func (sfo stringFlagOption[T]) Get(cmd *cobra.Command) T {
	if s, err := cmd.Flags().GetString(sfo.flag); err == nil {
		return sfo.f(s)
	}
	var t T
	return t
}
