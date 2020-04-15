package array

import (
	"errors"
	"fmt"
)

type Array struct {
	data   []int
	length uint
}

func NewArray(cap uint) *Array {
	if cap == 0 {
		return nil
	}
	return &Array{
		data:   make([]int, cap, cap),
		length: 0,
	}
}

func (this *Array) Len() uint {
	return this.length
}

func (this *Array) isIndexOutOfRange(index uint) bool {
	if index > uint(cap(this.data)) {
		return true
	}
	return false
}

func (this *Array) Find(index uint) (int, error) {
	if this.isIndexOutOfRange(index) {
		return 0, errors.New("out of index range")
	}
	return this.data[index], nil
}

func (this *Array) Insert(index uint, val int) error {
	if this.length == uint(cap(this.data)) {
		return errors.New("full array")
	}
	if this.isIndexOutOfRange(index) {
		return errors.New("out of index range")
	}
	for i := this.length; i > index; i-- {
		this.data[i] = this.data[i-1]
	}
	this.data[index] = val
	this.length++
	return nil
}

func (this *Array) Delete(index uint) (int, error) {
	if this.length == uint(0) {
		return 0, errors.New("empty array")
	}
	if this.isIndexOutOfRange(index) {
		return 0, errors.New("out of index range")
	}
	val := this.data[index]
	for i := index; i < this.length; i++ {
		this.data[i] = this.data[i+1]
	}
	this.length--
	return val, nil
}

func (this *Array) Print() {
	format := ""
	for i := uint(0); i < this.length; i++ {
		format += fmt.Sprintf("|%+v", this.data[i])
	}
	fmt.Println(format)
}
