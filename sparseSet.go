package mecs

const MAXENTITYS = 1000_00

type SparseSet[T any] struct {
	values      []T
	dense_list  []int
	sparse_list []int
}

func NewSpareseSet[T any]() *SparseSet[T] {
	w := SparseSet[T]{
		values:      make([]T, 0),
		dense_list:  make([]int, 0),
		sparse_list: make([]int, MAXENTITYS),
	}

	for i := range w.sparse_list {
		w.sparse_list[i] = -1
	}
	return &w
}
func (set *SparseSet[T]) Remove(index int) {
	if index < 0 || index >= MAXENTITYS {
		return
	}
	real_index := set.sparse_list[index]
	if real_index == -1 {
		return
	}
	set.dense_list[real_index] = set.dense_list[len(set.dense_list)-1]
	set.values[real_index] = set.values[len(set.values)-1]
	set.sparse_list[set.dense_list[real_index]] = real_index

	set.dense_list = set.dense_list[:len(set.dense_list)-1]
	set.values = set.values[:len(set.values)-1]

	set.sparse_list[index] = -1
}

func (set *SparseSet[T]) Get(index int) T {
	if set.sparse_list[index] == -1 {
		panic("this value no longer exist")
	}
	return set.values[set.sparse_list[index]]
}

func (set *SparseSet[T]) Set(index int, value T) {
	if index < 0 || index >= MAXENTITYS {
		return
	}
	if set.sparse_list[index] == -1 {
		set.dense_list = append(set.dense_list, index)
		set.values = append(set.values, value)
		set.sparse_list[index] = len(set.dense_list) - 1
		return
	}
	set.values[set.sparse_list[index]] = value

}

func (set *SparseSet[T]) Has(index int) bool {
	return set.sparse_list[index] != -1
}

func (set *SparseSet[T]) GetDenseValues() []T {
	return set.values
}
func (set *SparseSet[T]) GetDenseList() []int {
	return set.dense_list
}