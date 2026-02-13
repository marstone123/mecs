package mecs

import (
	"reflect"
)

//Id you use to represend a single entity
type EntityId int
//function that you can register as system to the engine
type System func(world *World)
//interface you use to save and get components
type Component any

type eventType int
//this is all the events you can register your system into
const (
	_ = eventType(iota)
	EventProgramStart
	EventUpdate
	EventLateUpdate
	EventDraw
)

func nameOf(a any) string{
	t := reflect.TypeOf(a)
	points := ""
	for t.Kind() == reflect.Ptr {
		t = t.Elem()
		points += "*"
	}
	return points + t.Name()
}

// every world is one instance of the engine
type World struct {
	components  map[string]*SparseSet[Component]
	freeindexes []int
	toAdd [][]Component
	lestindex   int
	toRemove []int
	systems map[eventType][]System
}

func NewWorld() World {
	return World{
		components:  make(map[string]*SparseSet[Component], 0),
		freeindexes: make([]int, 0),
		lestindex:   0,
		systems: make(map[eventType][]System),
		toRemove: make([]int, 0),
		toAdd: make([][]Component, 0),
	}
}
func (w *World) setComponent(index EntityId,component Component){
	compname := nameOf(component)
	if _,ok := w.components[compname];!ok{
		w.components[compname] = NewSpareseSet[Component]()
	}
	w.components[compname].Set(int(index),component)
}
//here you add system to the world and choose when you will update that system (Update,LateUpdate,Draw or when the program start)
func (w *World) AddSystem(event_type eventType,function System){
	if _,ok := w.systems[event_type];ok{
		w.systems[event_type] = append(w.systems[event_type], function)
	}else{
		w.systems[event_type] = make([]System, 1)
		w.systems[event_type][0] = function
	}
}
//run all the systems that register on EventProgramStart
func (w *World) Start(){
	for i := range w.systems[EventProgramStart]{
		w.systems[EventProgramStart][i](w)
	}
}
//run all the systems from Update and LateUpdate and after that add and remove all the entity you ask for
func (w *World) Update(){
	for i := range w.systems[EventUpdate]{
		w.systems[EventUpdate][i](w)
	}
	for i := range w.systems[EventLateUpdate]{
		w.systems[EventLateUpdate][i](w)
	}
	w.removeEntitys()
	w.addEntitysNow()
}
//run all the draw systems
func (w *World) Draw(){
	for i := range w.systems[EventDraw]{
		w.systems[EventDraw][i](w)
	}
}


/*
func (w *World) PrintWorldComponents(){
	for key,value := range w.components{
		fmt.Println(key," : ",value.dense_list)
	}
}*/