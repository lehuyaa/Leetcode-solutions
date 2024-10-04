package main

type MyHashMap struct {
	hashMap []int
}

func Constructor() MyHashMap {
	return MyHashMap{
		hashMap: make([]int, 1000001),
	}
}

func (this *MyHashMap) Put(key int, value int) {
	this.hashMap[key] = value + 1
}

func (this *MyHashMap) Get(key int) int {
	return this.hashMap[key] - 1
}

func (this *MyHashMap) Remove(key int) {
	this.hashMap[key] = 0
}
