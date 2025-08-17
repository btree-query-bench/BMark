package listindex

import (
	"os"
	"testing"
)

func TestListIndex_SaveLoadSearchDelete(t *testing.T) {
	filename := "test_listindex.db"
	defer os.Remove(filename)

	idx := NewListIndex()
	for i := 0; i < 100; i++ {
		key := int64(i)
		val := int64(i * 10)
		err := idx.Insert(key, val)
		if err != nil {
			t.Fatalf("Insert failed: %v", err)
		}
	}

	println("Printing bevor")
	idx.Print()

	if err := idx.SaveTo(filename); err != nil {
		t.Fatalf("Save failed: %v", err)
	}

	loaded := NewListIndex()
	if err := loaded.LoadFrom(filename); err != nil {
		t.Fatalf("Load failed: %v", err)
	}
	println("Printing after")
	loaded.Print()

	// Verify the loaded data matches the original
	if len(idx.Data) != len(loaded.Data) {
		t.Errorf("Data length mismatch: got %d, want %d", len(loaded.Data), len(idx.Data))
	}
}
