package quicksort

func QuickSort(arr []int) {
  qs(arr, 0, len(arr) - 1)
}

func qs(arr []int, low int, high int) {
  if low >= high {
    return
  }

  pivotIdx := partition(arr, low, high)
  
  qs(arr, low, pivotIdx - 1)
  qs(arr, pivotIdx + 1, high)
}

func partition(arr []int, lo int, hi int) int {
	pivot := arr[hi]
	idx := lo - 1

	for i := lo; i < hi; i++ {
		if arr[i] <= pivot {
      idx++
			tmp := arr[i]
			arr[i] = arr[idx]
			arr[idx] = tmp
		}
	}

	idx++
	arr[hi] = arr[idx]
	arr[idx] = pivot

	return idx
}

