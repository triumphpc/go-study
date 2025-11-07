package main

import (
	"math/rand"
	"time"
)

type RandomizedSet struct {
	dict map[int]int
	list []int
}

func Constructor() RandomizedSet {
	rand.Seed(time.Now().UnixNano())
	return RandomizedSet{
		dict: make(map[int]int),
		list: make([]int, 0),
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, exists := this.dict[val]; exists {
		return false
	}
	this.dict[val] = len(this.list)
	this.list = append(this.list, val)
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	index, exists := this.dict[val]
	if !exists {
		return false
	}
	lastElement := this.list[len(this.list)-1]
	this.list[index] = lastElement
	this.dict[lastElement] = index
	this.list = this.list[:len(this.list)-1]
	delete(this.dict, val)
	return true
}

func (this *RandomizedSet) GetRandom() int {
	return this.list[rand.Intn(len(this.list))]
}
