package sorter

import (
	"github.com/yy-c00/sorter/model"
	"reflect"
	"testing"
)

func init() {
	arr := make([]int, 1_000)

	for i := 0; i < len(arr); i++ {
		arr[i] = ((len(arr) - i) * (i + 1)) / 2
	}

	benchmark = model.Numbers{Unsorted: arr}
}

var (
	benchmark model.Numbers
	sorter = New()
	table = []struct{model.Numbers; expected []int} {
		{
			model.Numbers{Unsorted: []int{-2, 3, 4, 6, 21, 20}},
			[]int{-2, 3, 4, 6, 20, 21},
		},
		{
			model.Numbers{Unsorted: []int{2, 4, 5, -1, 6, 22, 31}},
			[]int{-1, 2, 4, 5, 6, 22, 31},
		},
	}
)

func TestQuickSorter_Sort(t *testing.T) {
	for i := 0; i < len(table); i++ {
		sorter.Sort(&table[i].Numbers)
		if !reflect.DeepEqual(table[i].expected, table[i].Sorted) {
			t.Errorf("Expected: %v Got: %v", table[i].expected, table[i].Sorted)
		}
		t.Log(table[i].Sorted)
	}
}

func BenchmarkQuickSorter_Sort(b *testing.B) {
	sorter.Sort(&benchmark)
}