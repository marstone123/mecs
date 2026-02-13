package mecs

//remove entity from the system
func (w *World) RemoveEntity(index EntityId) {
	w.toRemove = append(w.toRemove, int(index))
}

func (w *World) removeEntitys() {
	for _, index := range w.toRemove {
		for keys := range w.components {
			w.components[keys].Remove(int(index))
		}
	}
	w.toRemove = make([]int, 0)
}
//adds entity right when you say it (its deangures beacuse it can change the data you have like list of components and made you run twice on one object or be in infinity loop) try to avoid it
func (w *World) AddEntityNow(comps ...Component) {
	if len(w.freeindexes) != 0 {
		index := w.freeindexes[len(w.freeindexes)-1]
		for comp := range comps {
			w.setComponent(EntityId(index), comps[comp])
		}
		w.freeindexes = w.freeindexes[:len(w.freeindexes)-1]
	} else {
		index := w.lestindex
		w.lestindex++
		for comp := range comps {
			w.setComponent(EntityId(index), comps[comp])
		}
	}
}
//adds entity after the LateUpdate systems correct use world.AddEntity(&ForMutable{},ForConst{},...)
func (w *World) AddEntity(comps ...Component) {
	w.toAdd = append(w.toAdd, comps)
}

func (w *World) addEntitysNow() {
	for i := range w.toAdd {
		w.AddEntityNow(w.toAdd[i]...)
	}
	w.toAdd = make([][]Component, 0)
}