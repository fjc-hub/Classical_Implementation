/* This function takes two parameters:
 * array (which is an array of ints)
 * n (the number of items in 'array')
 * and sorts the integers found in 'array' into ascending order.
 * The sorting algorithm used here is the widely-known "quick sort."
 * 
 * Reference: Duke C Programming, Part 4 in Coursera.org
 */
void sort(int * array, size_t n) {
  if (n <= 1) {
    return; //arrays of size 1 or smaller are trivially sorted
  }
  //i starts at -1: it is incremented before it is used to index the array
  int i = -1;
  //j starts at n-1 even though it is decremented before it
  //is used to index the array, since array[n-1] is our pivot.
  int j = n-1;
  //we just use array[n-1] here for simplicity. In a real implementation,
  //we might want more sophisticated pivot selection (see CLR)
  int pivot = array[n-1];
  while (1) {
    do { //scan from left for value >= pivot
      i = i+1;
    } while (array[i] < pivot);
    do {  //scan from right for value < pivot
      j = j-1;
    } while (j >= 0 && array[j] >= pivot );
    if (i >= j) {  //if i and j have crossed, data is partitioned.
      break;
    }
    //swap array[i] with array[j]
    int temp = array[i];
    array[i] = array[j];
    array[j] = temp;
  }
  //swap array[i] with the pivot (array[n-1])
  array[n-1] = array[i];
  array[i] = pivot;
  //recursively sort the partitioned halves: [0,i) and [i+1,n)
  sort(array, i);
  sort(&array[i+1], n-i-1);
}