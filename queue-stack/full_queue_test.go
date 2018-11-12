package main

import (
	"sync"
	"testing"
)

func TestPutAndGet(t *testing.T) {
	q := New(10)

	q.Put("test")
	if q.Len() != 1 {
		t.Errorf("Expected: %d. Got: %d\n", 1, q.Len())
	}

	results, err := q.Get(1)
	result := results[0]
	if err != nil {
		t.Error(err)
	}
	if result != "test" {
		t.Errorf("Expected: '%s' Got: '%s'\n", "test", result)
	}
	if !q.Empty() {
		t.Error("Queue is not empty")
	}

	q.Put("test2")
	if q.Len() != 1 {
		t.Errorf("Expected: %d. Got: %d\n", 1, q.Len())
	}

	results, err = q.Get(1)
	result = results[0]
	if err != nil {
		t.Error(err)
	}
	if result != "test2" {
		t.Errorf("Expected: '%s'. Got: '%s'\n", "test2", result)
	}
	if !q.Empty() {
		t.Error("Queue is not empty")
	}
}

func TestAddEmptyPut(t *testing.T) {
	q := New(10)

	q.Put()

	if q.Len() != 0 {
		t.Errorf("Expected len: %d, Got: %d\n", 0, q.Len())
	}
}

func TestGetNonPositiveNumber(t *testing.T) {
	q := New(10)

	q.Put(`test`)
	result, err := q.Get(0)
	if err != nil {
		t.Error(err)
	}

	if len(result) != 0 {
		t.Errorf("Expected len: %d. Got: %d\n", 0, len(result))
	}
}

func TestEmpty(t *testing.T) {
	q := New(10)

	if !q.Empty() {
		t.Error("Expected empty queue")
	}

	q.Put("test")
	if q.Empty() {
		t.Error("Expected non-empty queue")
	}
}

func TestGetEmpty(t *testing.T) {
	q := New(10)

	go func() {
		q.Put("a")
	}()

	result, err := q.Get(2)
	if err != nil {
		t.Error(err)
	}

	if len(result) != 1 {
		t.Errorf("Expected: %d. Got: %d\n", 1, len(result))
	}
	if result[0] != "a" {
		t.Errorf("Expected: %s. Got: %s\n", "a", result[0])
	}
}

func TestMultipleGetEmpty(t *testing.T) {
	q := New(10)
	var wg sync.WaitGroup
	wg.Add(2)
	results := make([][]interface{}, 2)

	go func() {
		wg.Done()
		local, err := q.Get(1)
		if err != nil {
			t.Error(err)
		}

		results[0] = local
		wg.Done()
	}()

	go func() {
		wg.Done()
		local, err := q.Get(1)
		if err != nil {
			t.Error(err)
		}

		results[1] = local
		wg.Done()
	}()

	wg.Wait()
	wg.Add(2)

	q.Put("a", "b", "c")
	wg.Wait()

	if len(results[0]) != 1 && len(results[1]) != 1 {
		if (results[0][0] == "a" && results[1][0] == "b") ||
			(results[0][0] == "b" && results[1][0] == `a`) {
			t.Error("The array should be a, b or b, a")
		}
	}
}

func TestEmptyGetWithDispose(t *testing.T) {
	q := New(10)
	var wg sync.WaitGroup
	wg.Add(1)

	var err error

	go func() {
		wg.Done()
		_, err = q.Get(1)
		wg.Done()
	}()

	wg.Wait()
	wg.Add(1)

	q.Dispose()

	wg.Wait()

	_, ok := err.(error)
	if !ok {
		t.Errorf("%v is not an error\n", err)
	}
}

func TestGetPutDisposed(t *testing.T) {
	q := New(10)

	q.Dispose()

	_, err := q.Get(1)
	_, ok := err.(error)
	if !ok {
		t.Errorf("%v is not an error\n", err)
	}

	err = q.Put("a")
	_, ok = err.(error)
	if !ok {
		t.Errorf("%v is not an error\n", err)
	}
}
