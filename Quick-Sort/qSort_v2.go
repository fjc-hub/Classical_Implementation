/**
 *	Difference from qSort_v1 is that this implementation
 *	abstracts Partition subroutine into a function
 *
 */

package quicksort

import "math/rand"

func qSort(arr []int) {
	if len(arr) < 2 {
		return
	}
	n := len(arr)
	rd := rand.Intn(n)
	arr[rd], arr[n-1] = arr[n-1], arr[rd]
	inPlace := partition(arr, 0, n-1) // partition and get a in-place index
	// reduce
	qSort(arr[:inPlace])
	qSort(arr[inPlace+1:])
}

// two-pointer partition, partition by arr[y]
// return the index of item now known to be in place
func partition(arr []int, x, y int) int {
	pivot, j := arr[y], x // mentain [x, j) < pivot
	for i := x; i < y; i++ {
		if arr[i] >= pivot {
			continue
		}
		arr[i], arr[j] = arr[j], arr[i]
		j++
	}
	arr[j], arr[y] = arr[y], arr[j]
	return j // return the index of pivot
}
