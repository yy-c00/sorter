package sorter

import (
	"github.com/yy-c00/sorter/model"
)

//New returns an implementation of model.Sorter
func New() model.Sorter {
	return quickSorter{}
}

type quickSorter struct {}

func (q quickSorter) Sort(numbers *model.Numbers) {
	numbers.Sorted = make([]int, len(numbers.Unsorted))
	copy(numbers.Sorted, numbers.Unsorted)

	q.sort(numbers.Sorted, 0, len(numbers.Sorted) - 1)

	x := 1
	for i := 0; i < len(numbers.Sorted) - x; i++ {
		if numbers.Sorted[i] == numbers.Sorted[i + 1] {
			tmp := numbers.Sorted[i]
			numbers.Sorted = append(numbers.Sorted[:i], numbers.Sorted[i + 1:]...)
			numbers.Sorted = append(numbers.Sorted, tmp)
			x++
		}
	}
}

func (q quickSorter) sort(arr []int, start, end int) {
	if (end - start) < 1 {
		return
	}

	pivot := arr[end]
	splitIndex := start

	for i := start; i < end; i++ {
		if arr[i] < pivot {
			temp := arr[splitIndex]

			arr[splitIndex] = arr[i]
			arr[i] = temp

			splitIndex++
		}
	}

	arr[end] = arr[splitIndex]
	arr[splitIndex] = pivot

	q.sort(arr, start, splitIndex - 1)
	q.sort(arr, splitIndex + 1, end)
}