package persist

import (
	"os"
	"reflect"
	"testing"
)

type TestData struct {
	Name string
	Age  int
}

func TestSaveLoad_WhenSavingAndLoading_ShouldPreserveData(t *testing.T) {
	// Given
	original := TestData{Name: "Alice", Age: 30}
	filename := "testdata.gob"
	defer os.Remove(filename) // Clean up after test

	// When
	if err := Save(filename, original); err != nil {
		t.Fatalf("Save failed: %v", err)
	}

	var loaded TestData
	if err := Load(filename, &loaded); err != nil {
		t.Fatalf("Load failed: %v", err)
	}

	// Then
	if !reflect.DeepEqual(original, loaded) {
		t.Errorf("Loaded data does not match original.\nGot: %+v\nWant: %+v", loaded, original)
	}
}
