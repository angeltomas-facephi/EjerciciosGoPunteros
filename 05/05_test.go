package main

import (
	"testing"
)

func TestArraySortPointers(t *testing.T) {
	arrTest := []int{3, 6, 11, 5, 2}
	arrPunterosOrdenadoTest := []*int{
		&arrTest[4],
		&arrTest[0],
		&arrTest[3],
		&arrTest[1],
		&arrTest[2],
	}

	arr := []int{3, 6, 11, 5, 2}
	arrPunterosOrdenado := arraySortPointers(&arr)

	if !compareArr(arrPunterosOrdenadoTest, arrPunterosOrdenado) {
		t.Error("test was incorrect, because the array not ordered")

	}

}

func compareArr(arrOriginal []*int, arrOrdenado []*int) bool {
	for i, element := range arrOriginal {
		if *element != *arrOrdenado[i] {
			return false
		}
	}
	return true
}
