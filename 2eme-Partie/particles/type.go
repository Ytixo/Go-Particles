package particles

import "container/list"

// System définit un système de particules.
type System struct {
	Content          *list.List
	SpawnAccumulator float64
}
