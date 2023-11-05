package badgo

// func mergeSort(arr []int, b *bigbuffer) []int {
// 	if len(arr) <= 1 {
// 		return arr
// 	}
// 	if b == nil {
// 		b = newBigBuffer(len(arr))
// 	}
// 	mid := len(arr) / 2
// 	s1 := mergeSort(arr[:mid], b)
// 	s2 := mergeSort(arr[mid:], b)
// 	return merge(s1, s2, b)
// }

// type bigbuffer []int

// func newBigBuffer(size int) *bigbuffer {
// 	l := size * bits.Len(uint(size))
// 	b := make(bigbuffer, l)
// 	return &b
// }

// func (b *bigbuffer) slice(size int) []int {
// 	s := (*b)[:size]
// 	(*b) = (*b)[size:]
// 	return s
// }

// func merge(arr1, arr2 []int, b *bigbuffer) []int {
// 	size1, size2 := len(arr1), len(arr2)
// 	out := b.slice(size1 + size2)
// 	var idx1, idx2, index int
// 	for idx1 < size1 && idx2 < size2 {
// 		if arr1[idx1] < arr2[idx2] {
// 			out[index] = arr1[idx1]
// 			idx1++
// 		} else {
// 			out[index] = arr2[idx2]
// 			idx2++
// 		}
// 		index++
// 	}
// 	for idx1 < size1 {
// 		out[index] = arr1[idx1]
// 		idx1++
// 		index++
// 	}
// 	for idx2 < size2 {
// 		out[index] = arr2[idx2]
// 		idx2++
// 		index++
// 	}
// 	return out
// }

func mergeSort2(in []int) []int {
	ch := make(chan int, len(in))
	mergeSortCh(in, ch)

	out := make([]int, 0, len(in))
	for i := range ch {
		out = append(out, i)
	}
	return out
}

func mergeSortCh(in []int, sorted chan int) {
	defer close(sorted)
	if len(in) <= 1 {
		if len(in) == 1 {
			sorted <- in[0]
		}
		return
	}
	mid := len(in) / 2
	sorted1 := make(chan int, mid)
	sorted2 := make(chan int, len(in)-mid)

	// Sort each half concurrently
	go mergeSortCh(in[:mid], sorted1)
	go mergeSortCh(in[mid:], sorted2)

	// Merge the channels
	mergeCh(sorted1, sorted2, sorted)
}

// mergeCh takes two sorted channels of ints in1, in2 and outputs a sorted stream of ints into out
func mergeCh(in1, in2, out chan int) {
	i1, ok1 := <-in1
	i2, ok2 := <-in2
	for ok1 && ok2 {
		if i1 < i2 {
			out <- i1
			i1, ok1 = <-in1
		} else {
			out <- i2
			i2, ok2 = <-in2
		}
	}
	for ok1 {
		out <- i1
		i1, ok1 = <-in1
	}
	for ok2 {
		out <- i2
		i2, ok2 = <-in2
	}
}

func mergeSort(in []int) []int {
	if len(in) <= 1 {
		return in
	}
	mid := len(in) / 2
	sorted1 := mergeSort(in[:mid])
	sorted2 := mergeSort(in[mid:])
	return merge(sorted1, sorted2)
}

// merge zips two sorted slices together
func merge(in1, in2 []int) []int {
	size1, size2 := len(in1), len(in2)
	out := make([]int, size1+size2)

	var idx1, idx2, index int
	for idx1 < size1 && idx2 < size2 {
		if in1[idx1] < in2[idx2] {
			out[index] = in1[idx1]
			idx1++
		} else {
			out[index] = in2[idx2]
			idx2++
		}
		index++
	}
	for idx1 < size1 {
		out[index] = in1[idx1]
		idx1++
		index++
	}
	for idx2 < size2 {
		out[index] = in2[idx2]
		idx2++
		index++
	}
	return out
}
