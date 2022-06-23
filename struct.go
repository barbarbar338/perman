package perman

import "math"

type Perman struct {
	flags map[string]int64
}

func Factory(flags []string) Perman {
	p := Perman{
		flags: make(map[string]int64),
	}
	for i, flag := range flags {
		p.flags[flag] = int64(math.Pow(2, float64(i)))
	}
	return p
}
