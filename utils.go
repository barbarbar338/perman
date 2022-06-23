package perman

func (p Perman) Keys() []string {
	keys := make([]string, 0)
	for key := range p.flags {
		keys = append(keys, key)
	}
	return keys
}

func (p Perman) Values() []int64 {
	values := make([]int64, 0)
	for flag := range p.flags {
		values = append(values, p.flags[flag])
	}
	return values
}

func (p Perman) Get(flag string) int64 {
	value, ok := p.flags[flag]
	if ok {
		return value
	} else {
		return int64(0)
	}
}

func (p Perman) Add(permission int64, key string) int64 {
	oldValues := p.Deserialize(permission)
	newValues := append(oldValues, key)
	return p.Serialize(newValues)
}

func (p Perman) Remove(permission int64, key string) int64 {
	oldValues := p.Deserialize(permission)
	newValues := make([]string, 0)
	for _, value := range oldValues {
		if value != key {
			newValues = append(newValues, value)
		}
	}
	return p.Serialize(newValues)
}

func (p Perman) Full() int64 {
	return p.Serialize(p.Keys())
}
