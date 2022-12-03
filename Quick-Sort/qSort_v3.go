/**
 *	3-way partition quick sort for cases of existing many duplicate key
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
	inPlace0, inPlace1 := partition(arr, 0, n-1) // partition and get a in-place index
	// reduce
	qSort(arr[:inPlace0])
	qSort(arr[inPlace1:])
}

// 3-way partition, pivot = arr[y]
// return the interval that now known to be in place, such that all entry equals to pivot
// for example, all arr[r0~r1) == pivot
func partition(arr []int, x, y int) (int, int) {
	// mentain: [x, lt) < pivot; [lt, i)==pivot; [i,gt) unknown; [gt, y) > pivot; [y,y]==pivot
	pivot, lt, gt := arr[y], x, y
	for i := x; i < gt; {
		if arr[i] > pivot {
			gt--
			arr[i], arr[gt] = arr[gt], arr[i]
		} else if arr[i] < pivot {
			arr[i], arr[lt] = arr[lt], arr[i]
			lt++
			i++
		} else {
			i++
		}
	}
	arr[gt], arr[y] = arr[y], arr[gt]
	return lt, gt // return the interval of pivot
}
