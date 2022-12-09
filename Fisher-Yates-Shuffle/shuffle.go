/**
 *	1.This is a kind of Shuffle Algorithm, shuffle an array such that every element
 *	  has the same probability of appearing in each position
 *
 *	2.Interesting analogy from wikipedia of this algorithm: The algorithm effectively
 *    puts all the elements into a hat; it continually determines the next element by
 *    randomly drawing an element from the hat until no elements remain.
 *
 *  3.Reference: https://en.wikipedia.org/wiki/Fisher%E2%80%93Yates_shuffle#:~:text=become%20frustratingly%20common.-,The%20modern%20algorithm,-%5Bedit%5D
 */

package fisheryatesshuffle

import (
	"math/rand"
	"time"
)

var seed = time.Now().Unix()

func Shuffle(arr []interface{}) {
	rand.Seed(seed)
	for i := len(arr) - 1; i > 0; i-- {
		x := rand.Intn(i + 1)
		arr[x], arr[i] = arr[i], arr[x]
	}
}
