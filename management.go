package perman

func (p Perman) Serialize(flags []string) int64 {
	if len(flags) < 1 {
		return 0
	}

	res := int64(0)
	for _, flag := range flags {
		res |= p.Get(flag)
	}

	return res
}

func (p Perman) Deserialize(permission int64) []string {
	if permission == 0 {
		return []string{}
	}

	res := make([]string, 0)
	for key, value := range p.flags {
		if permission&value != 0 {
			res = append(res, key)
		}
	}
	return res
}
