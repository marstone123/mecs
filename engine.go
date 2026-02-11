package mecs

import (
	"reflect"
)

type EntityId int
type System func(world *World)
type Component any
type eventType int

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

// you need to keep every component of the same entity at the same position
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

func (w *World) AddSystem(event_type eventType,function System){
	if _,ok := w.systems[event_type];ok{
		w.systems[event_type] = append(w.systems[event_type], function)
	}else{
		w.systems[event_type] = make([]System, 1)
		w.systems[event_type][0] = function
	}
}

func (w *World) Start(){
	for i := range w.systems[EventProgramStart]{
		w.systems[EventProgramStart][i](w)
	}
}

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