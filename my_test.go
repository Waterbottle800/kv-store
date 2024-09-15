package main

import (
	"testing"
)

func TestSet(t *testing.T) {
	store := InitKV()
	t.Log("Setting key-value pairs in the store...")

	store.Set("abc-1", "value1")
	store.Set("abc-2", "value2")

	t.Log("Verifying stored values...")

	if val, exists := store.Get("abc-1"); !exists || val != "value1" {
		t.Fatalf("Expected abc-1 to be 'value1', but got %v", val)
	} else {
		t.Logf("Key: abc-1, Value: %s, Test Passed", val)
	}

	if val, exists := store.Get("abc-2"); !exists || val != "value2" {
		t.Fatalf("Expected abc-2 to be 'value2', but got %v", val)
	} else {
		t.Logf("Key: abc-2, Value: %s, Test Passed", val)
	}
}

func TestGet(t *testing.T) {
	store := InitKV()
	store.Set("abc-1", "value1")

	t.Log("Testing retrieval of existing and non-existing keys...")

	// Test existing key
	if val, exists := store.Get("abc-1"); !exists || val != "value1" {
		t.Fatalf("Expected abc-1 to be 'value1', but got %v", val)
	} else {
		t.Logf("Key: abc-1, Value: %s, Test Passed", val)
	}

	// Test non-existent key
	if _, exists := store.Get("xyz-1"); exists {
		t.Fatalf("Expected 'xyz-1' to not exist")
	} else {
		t.Log("Key: xyz-1 does not exist, Test Passed")
	}
}

func TestSearch(t *testing.T) {
	store := InitKV()

	t.Log("Setting keys for search tests...")
	store.Set("abc-1", "value1")
	store.Set("abc-2", "value2")
	store.Set("xyz-1", "value3")
	store.Set("xyz-2", "value4")

	t.Log("Testing prefix search...")
	prefixResult := store.Search("abc", "")
	if len(prefixResult) != 2 || prefixResult[0] != "abc-1" || prefixResult[1] != "abc-2" {
		t.Fatalf("Expected 'abc-1', 'abc-2', got %v", prefixResult)
	} else {
		t.Logf("Prefix search for 'abc' returned: %v, Test Passed", prefixResult)
	}

	t.Log("Testing suffix search...")
	suffixResult := store.Search("", "-1")
	if len(suffixResult) != 2 || suffixResult[0] != "abc-1" || suffixResult[1] != "xyz-1" {
		t.Fatalf("Expected 'abc-1', 'xyz-1', got %v", suffixResult)
	} else {
		t.Logf("Suffix search for '-1' returned: %v, Test Passed", suffixResult)
	}

	t.Log("Testing combined prefix and suffix search...")
	bothResult := store.Search("abc", "-1")
	if len(bothResult) != 1 || bothResult[0] != "abc-1" {
		t.Fatalf("Expected 'abc-1', got %v", bothResult)
	} else {
		t.Logf("Combined search for 'abc' and '-1' returned: %v, Test Passed", bothResult)
	}
}

