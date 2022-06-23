package perman

func (p Perman) Match(permission int64, flags []string) bool {
	if len(flags) < 1 {
		return true
	}

	for _, flag := range flags {
		if permission&p.Get(flag) != p.Get(flag) {
			return false
		}
	}
	return true
}

func (p Perman) MatchAll(permission int64, flags []string) bool {
	if len(flags) < 1 {
		return true
	}

	for _, flag := range flags {
		if permission&p.Get(flag) != p.Get(flag) {
			return false
		}
	}
	return true
}

func (p Perman) HasAll(permission int64, flags []string) bool {
	if len(flags) < 1 {
		return true
	}

	for _, flag := range flags {
		if permission&p.Get(flag) != p.Get(flag) {
			return false
		}
	}
	return true
}

func (p Perman) Some(permission int64, flags []string) bool {
	if len(flags) < 1 {
		return true
	}

	for _, flag := range flags {
		if permission&p.Get(flag) == p.Get(flag) {
			return true
		}
	}
	return false
}

func (p Perman) HasSome(permission int64, flags []string) bool {
	if len(flags) < 1 {
		return true
	}

	for _, flag := range flags {
		if permission&p.Get(flag) == p.Get(flag) {
			return true
		}
	}
	return false
}

func (p Perman) HasNone(permission int64, flags []string) bool {
	if len(flags) < 1 {
		return true
	}

	for _, flag := range flags {
		if permission&p.Get(flag) == p.Get(flag) {
			return false
		}
	}
	return true
}

func (p Perman) None(permission int64, flags []string) bool {
	if len(flags) < 1 {
		return true
	}

	for _, flag := range flags {
		if permission&p.Get(flag) == p.Get(flag) {
			return false
		}
	}
	return true
}

func (p Perman) Has(permission int64, flag string) bool {
	return permission&p.Get(flag) == p.Get(flag)
}

func (p Perman) Test(permission int64, flag string) bool {
	return permission&p.Get(flag) == p.Get(flag)
}
