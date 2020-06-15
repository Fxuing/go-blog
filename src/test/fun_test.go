package test

import (
	"testing"
)

func TestSlice(t *testing.T) {
	//a := []int{1, 2, 3, 4, 5, 1, 6, 2, 3, 4, 1, 2, 3, 5, 1, 2, 2, 2, 1, 3, 4,}
	a := []string{"a", "a", "a", "a", "a", "B", "B", "B", "B", "B", "B", "B"}
	m := split("B", a)

	for k, v := range m {
		t.Log(k, v)
	}

}

func split(char interface{}, data interface{}) map[interface{}][]interface{} {
	m := make(map[interface{}][]interface{})
	switch data.(type) {
	case []int:
		for _, t := range data.([]int) {
			if t == char {
				m[t] = append(m[t], t)
			}
		}
	case []string:
		for _, t := range data.([]string) {
			if t == char {
				m[t] = append(m[t], t)
			}
		}
	}
	return m
}
