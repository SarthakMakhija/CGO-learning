package linkedlist_test

import (
	"reflect"
	"testing"
)
import "linkedlist"

func TestPutASingleKeyValue(t *testing.T) {
	defer linkedlist.Close()
	linkedlist.Put(10, 100)

	value := linkedlist.GetBy(10)
	if value != 100 {
		t.Fatalf("Expected %v, received %v", 100, value)
	}
}

func TestPutMultipleKeyValues(t *testing.T) {
	defer linkedlist.Close()
	linkedlist.Put(10, 100)
	linkedlist.Put(20, 200)
	linkedlist.Put(30, 300)

	value := linkedlist.GetBy(20)
	if value != 200 {
		t.Fatalf("Expected %v, received %v", 200, value)
	}
}

func TestGetAllValues(t *testing.T) {
	defer linkedlist.Close()
	linkedlist.Put(10, 100)
	linkedlist.Put(20, 200)
	linkedlist.Put(30, 300)

	values := linkedlist.GetAllValues()
	expected := []int{100, 200, 300}

	if !reflect.DeepEqual(expected, values) {
		t.Fatalf("Expected %v, received %v", expected, values)
	}
}
