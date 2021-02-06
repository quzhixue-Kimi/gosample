package stringsample

import (
	"reflect"
	"testing"
)

type TestCase struct {
	str  string
	sep  string
	want []string
}

func setupTestCase(t *testing.T) func(t *testing.T) {
	t.Log("begin test")
	return func(t *testing.T) {
		t.Log("executing...")
	}
}

func tearDown(t *testing.T) func(t *testing.T) {
	t.Log("end")
	return func(t *testing.T) {
		t.Log("the end")
	}
}

func TestSuiteSplit(t *testing.T) {
	var testGroup = map[string]TestCase{
		"case1": TestCase{"aabcde", "b", []string{"aa", "cde"}},
		"case2": TestCase{"b:c:d:e", ":", []string{"b", "c", "d", "e"}},
		"case3": TestCase{"hello", "h", []string{"", "ello"}},
	}

	teardownTestCase := tearDown(t)
	defer teardownTestCase(t)

	for name, tc := range testGroup {
		t.Run(name, func(t *testing.T) {
			setupTestCase(t)
			defer teardownTestCase(t)
			got := Split(tc.str, tc.sep)
			if !reflect.DeepEqual(tc.want, got) {
				t.Fatalf("got:%#v, want:%#v\n", got, tc.want)
			}
		})
	}
}

func BenchmarkSplit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Split("a:b:c:d:e", ":")
	}
}

func benchmarkFib(b *testing.B, n int) {
	for i := 0; i < b.N; i++ {
		Fib(n)
	}
}

func BenchmarkFib1(b *testing.B) {
	benchmarkFib(b, 1)
}

func BenchmarkFib20(b *testing.B) {
	benchmarkFib(b, 20)
}
