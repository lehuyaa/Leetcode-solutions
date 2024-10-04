package main

import "math/rand"

type RandomizedSet struct {
	arr  []int
	set  map[int]int
	size int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		arr:  []int{},
		set:  make(map[int]int),
		size: 0,
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	_, ok := this.set[val]
	if ok {
		return false
	}
	this.set[val] = this.size
	this.arr = append(this.arr, val)
	this.size++
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	index, ok := this.set[val]
	if !ok {
		return false
	}
	lastItem := this.arr[this.size-1]
	this.arr[this.size-1] = this.arr[index]
	this.arr[index] = lastItem
	this.set[lastItem] = index
	this.arr = this.arr[:this.size-1]
	this.size--
	delete(this.set, val)
	return true
}

func (this *RandomizedSet) GetRandom() int {
	index := rand.Intn(this.size)
	return this.arr[index]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
