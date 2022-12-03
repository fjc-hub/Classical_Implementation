/**
 *	The qSort function is a implementation of classical algorithm -- quick sort.
 *
 *	It is non-deterministic version, it introduces random number to select pivot in partition routine.
 * 	The partition subroutine is "2-way" version.
 *
 *	Another Optimize method
 *	- You can RANDOM SHUFFLE the array at every qSort call
 *	- Insertion sort when len(arr) <= 10
 *	- 3-way partitioning
 *
 *	Test case: leetcode.912
 *
 *	Reference: https://www.coursera.org/learn/algorithms-part1/supplement/efbDN/lecture-slides
 *
 */

package quicksort

import "math/rand"

func qSort(arr []int) {
	if len(arr) <= 1 {
		return
	}
	n, l, r := len(arr), 0, len(arr)-1
	rd := rand.Intn(n)
	pivot := arr[rd]
	arr[rd], arr[r] = arr[r], arr[rd]
	for {
		for ; l < n && arr[l] < pivot; l++ {
		}
		for ; r >= 0 && arr[r] >= pivot; r-- {
		} // Observe '>=', avoid dead loop
		if l >= r {
			break
		}
		arr[l], arr[r] = arr[r], arr[l]
	}
	// arr[l] is IN PLACE(就绪辣,不用再移动了): [0, l-1] < pivot, [l, n-1] >= pivot
	arr[l], arr[n-1] = arr[n-1], arr[l]
	qSort(arr[:l])   // reduce
	qSort(arr[l+1:]) // reduce
}
