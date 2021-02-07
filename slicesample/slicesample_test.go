package slicesample

import "testing"

type testObj struct {
	name   string
	method func()
}

func TestSliceTest(t *testing.T) {
	SliceTest()
}

func TestSliceTest1(t *testing.T) {
	SliceTest1()
}

func TestSliceTest2(t *testing.T) {
	SliceTest2()
}

func TestSuit(t *testing.T) {

	testGroup := []testObj{
		{"sliceCopy", SliceCopy},
		{"sliceChange", SliceChange},
		{"sliceScale", SliceScale},
	}

	for _, tc := range testGroup {
		t.Run(tc.name, func(t *testing.T) {
			tc.method()
		})
	}

}
