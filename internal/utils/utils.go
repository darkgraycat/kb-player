package utils

func IfThenElse[T any](cond bool, a, b T) T {
	if cond {
		return a
	}
	return b
}

func Match[K comparable, V any](k K, mapping map[K]V, def V) V {
	if v, ok := mapping[k]; ok {
		return v
	}
	return def
}

