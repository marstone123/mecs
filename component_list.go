package mecs

//a list of components
type ComponentsList struct {
	Components []Component
	entitys    []int
}
//this give you the EntityID of component by their index
func (c *ComponentsList) GetEntityByIndex(index int) EntityId {
	return EntityId(c.entitys[index])
}
//you get all components of spasific type (the name of the type in string) from a component list
func (w *World) GetAllComponents(components_type string) *ComponentsList {
	if _, ok := w.components[components_type]; !ok {
		return nil
	}
	return &ComponentsList{
		Components: w.components[components_type].GetDenseValues(),
		entitys:    w.components[components_type].GetDenseList(),
	}
}
//you get component by his type (the name of the type as string) and entityId
func (w *World) GetComponent(entity EntityId, component_type string) Component {
	if _, ok := w.components[component_type]; !ok {
		return nil
	}
	if w.components[component_type].sparse_list[entity] != -1 {
		return w.components[component_type].Get(int(entity))
	}
	return nil
}
//chack if entity have some componet (again the component_type is the name of the type in string)
func (w *World) HasComponent(entity EntityId, component_type string) bool {
	if _, ok := w.components[component_type]; !ok {
		return false
	}
	return w.components[component_type].Has(int(entity))
}