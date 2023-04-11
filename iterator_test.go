package utils

import (
	"reflect"
	"strconv"
	"testing"
)

func TestBasic(t *testing.T) {
	integerArray := []int{1, 2, 3, 4, 5}
	integerIter := IntoIterator(integerArray)
	integerResult := Collect(integerIter)

	if !reflect.DeepEqual(integerResult, integerArray) {
		t.Errorf("Expected: %v, got: %v", integerArray, integerResult)
	}

	stringArray := []string{"1", "2", "3"}
	stringIter := IntoIterator(stringArray)
	stringResult := Collect(stringIter)

	if !reflect.DeepEqual(stringResult, stringArray) {
		t.Errorf("Expected: %v, got: %v", stringArray, stringResult)
	}
}

func TestFold(t *testing.T) {
	in := []int{1, 2, 3}
	expected := 6
	iter := IntoIterator(in)
	res := Fold(0, func(a int, b int) int {
		return a + b
	}, iter)

	if res != expected {
		t.Errorf("Expected: %v, got: %v", expected, res)
	}
}

func TestReduce(t *testing.T) {
	in := []int{1, 2, 3}
	expected := 6
	iter := IntoIterator(in)
	res := Reduce(func(a int, b int) int {
		return a + b
	}, iter)

	if res != expected {
		t.Errorf("Expected: %v, got: %v", expected, res)
	}
}

func TestFirstWhere(t *testing.T) {
	in := []int{1, 2, 3, 4, 5, 6, 7}
	expected := 2
	res, err := FirstWhere(func(a int) bool {
		return a%2 == 0
	}, IntoIterator(in))

	if err != nil {
		t.Errorf("Should find the element %v", expected)
	}

	if res != expected {
		t.Errorf("Expected: %v, got: %v", expected, res)
	}

	res, err = FirstWhere(func(a int) bool {
		return a > 10
	}, IntoIterator(in))

	if err != NotFound {
		t.Errorf("Should not find the element")
	}
}

func TestContain(t *testing.T) {
	in := []int{1, 2, 3, 4, 5, 6, 7}
	res := Contain(2, IntoIterator(in))

	if !res {
		t.Errorf("Should contain the element %v", 2)
	}

	res = Contain(20, IntoIterator(in))

	if res {
		t.Errorf("Should not contain the element %v", 20)
	}
}

func TestFilter(t *testing.T) {
	in := []int{1, 2, 3, 4, 5, 6, 7, 8}
	expected := []int{2, 4, 6, 8}
	res := Collect(Filter(func(a int) bool {
		return a%2 == 0
	}, IntoIterator(in)))

	if !reflect.DeepEqual(res, expected) {
		t.Errorf("Expected: %v, got: %v", expected, res)
	}
}

func TestMapping(t *testing.T) {
	in := []int{1, 2, 3}
	expected := []string{"1", "2", "3"}

	res := Collect(Map(func(x int) string {
		return strconv.Itoa(x)
	}, IntoIterator(in)))

	if !reflect.DeepEqual(res, expected) {
		t.Errorf("Expected: %v, got: %v", expected, res)
	}
}

func TestFilterAndMapping(t *testing.T) {
	in := []int{1, 2, 3, 4, 5, 6, 7, 8}
	expected := []string{"2", "4", "6", "8"}
	res := Collect(Map(func(a int) string {
		return strconv.Itoa(a)
	}, Filter(func(a int) bool {
		return a%2 == 0
	}, IntoIterator(in))))

	if !reflect.DeepEqual(res, expected) {
		t.Errorf("Expected: %v, got: %v", expected, res)
	}
}
