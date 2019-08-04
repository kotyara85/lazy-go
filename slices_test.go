package lazygo

import "testing"

func TestSlices(t *testing.T) {
	// SliceUniq
	slice := []string{"test1", "test2", "test2"}
	SliceUniq(&slice)
	if len(slice) != 2 {
		t.Error("SliceUniq test failed")
	}
	// SliceRemoveEmpty
	slice = []string{"test1", "", "test2"}
	SliceRemoveEmpty(&slice)
	if len(slice) != 2 {
		t.Error("SliceRemoveEmpty test failed")
	}
	// SliceRemoveEmpty
	slice = []string{"test1", "test2"}
	if !SliceInclude(slice, "test1") {
		t.Error("SliceInclude test failed")
	}
	if SliceInclude(slice, "test3") {
		t.Error("SliceInclude test failed - shouldn't return true")
	}

}