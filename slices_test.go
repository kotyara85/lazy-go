package lazygo

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSlices(t *testing.T) {

	expBTrue := true
	expBFalse := false

	cases := []struct {
		testName    string
		slice       []string
		expectSlice []string
		b           *bool
		expectB     *bool
	}{
		{
			testName: "uniq slice",
			slice: func() []string {
				slice := []string{"test1", "test2", "test2"}
				SliceUniq(&slice)
				return slice
			}(),
			expectSlice: []string{"test1", "test2"},
		},
		{
			testName: "slice remove empty",
			slice: func() []string {
				slice := []string{"test1", "  ", "test2"}
				SliceRemoveEmpty(&slice)
				return slice
			}(),
			expectSlice: []string{"test1", "test2"},
		},
		{
			testName: "slice include",
			b: func() *bool {
				slice := SliceInclude([]string{"test1", "test2"}, "test1")
				return &slice
			}(),
			expectB: &expBTrue,
		},
		{
			testName: "slice include, false",
			b: func() *bool {
				slice := SliceInclude([]string{"test1", "test2"}, "test3")
				return &slice
			}(),
			expectB: &expBFalse,
		},
		{
			testName: "slice to interface",
			b: func() *bool {
				testInterface := []interface{}{"test1", "test2"}
				testSlice := []string{"test1", "test2"}
				ret := reflect.DeepEqual(testInterface, SliceToInterface(testSlice))
				return &ret
			}(),
			expectB: &expBTrue,
		},
		{
			testName: "slice to interface, false",
			b: func() *bool {
				testInterface := []interface{}{"test1"}
				testSlice := []string{"test1", "test2"}
				ret := reflect.DeepEqual(testInterface, SliceToInterface(testSlice))
				return &ret
			}(),
			expectB: &expBFalse,
		},
	}

	for _, c := range cases {
		t.Run(c.testName, func(t *testing.T) {
			var ran bool
			if c.expectSlice != nil && c.slice != nil {
				ran = true
				assert.Equal(t, c.expectSlice, c.slice)
			}

			if c.expectB != nil {
				ran = true
				assert.Equal(t, c.expectB, c.b)
			}
			// make sure we ran unit test
			assert.Equal(t, true, ran)
		})
	}
}
