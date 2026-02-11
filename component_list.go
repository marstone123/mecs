package mecs

type ComponentsList struct {
	Components []Component
	entitys    []int
}

func (c *ComponentsList) GetComponentEntity(index int) EntityId {
	return EntityId(c.entitys[index])
}

func (w *World) GetComponents(components_type string) *ComponentsList {
	if _, ok := w.components[components_type]; !ok {
		return nil
	}
	return &ComponentsList{
		Components: w.components[components_type].GetDenseValues(),
		entitys:    w.components[components_type].GetDenseList(),
	}
}
func (w *World) GetComponent(entity EntityId, component_type string) Component {
	if _, ok := w.components[component_type]; !ok {
		return nil
	}
	if w.components[component_type].sparse_list[entity] != -1 {
		return w.components[component_type].Get(int(entity))
	}
	return nil
}

func (w *World) HasComponent(entity EntityId, component_type string) bool {
	if _, ok := w.components[component_type]; !ok {
		return false
	}
	return w.components[component_type].Has(int(entity))
}