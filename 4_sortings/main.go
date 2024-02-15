package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"runtime"
	"sort"
	"strings"
	"sync"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	// arr := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	var arr []int
	//
	for i := 20000; i > 10000; i-- {
		arr = append(arr, i)
	}
	fmt.Println("Start")
	sorts := []func([]int){selection_sort, bubble_sort, insertion_sort, counting_sort, quick_sort, sort.Ints, radix_sotr}
	var wg sync.WaitGroup
	for _, v := range sorts {
		wg.Add(1)
		go func(arr []int, v func([]int)) {
			defer wg.Done()
			cp := make([]int, len(arr))
			copy(cp, arr)
			start := time.Now()
			v(cp)
			end := time.Since(start)
			fmt.Printf("%v\t %s\n", end, strings.Split(runtime.FuncForPC(reflect.ValueOf(v).Pointer()).Name(), ".")[1])
		}(arr, v)
	}

	wg.Add(1)
	go func(arr []int) {
		defer wg.Done()
		c := make([]int, len(arr))
		copy(c, arr)
		start := time.Now()
		merge_sort(&c)
		end := time.Since(start)
		fmt.Println(end, "\t", "merge_sort")
	}(arr)
	wg.Wait()
	fmt.Println("Done")
	radix_sotr([]int{})

}

// 1 3 2 4 7 0|
// 1 3 2 4 0|7
// 1 3 2 0|4 7
// 1 0 2|3 4 7
// 1 0|2 3 4 7
// 0 1 2 3 4 7
// n + (n-1) + (n-2)...
// n*(n+1)/2 - 1 ~ n^2
// O(n^2)
func selection_sort(arr []int) {
	for j := 0; j < len(arr); j++ {
		min_idx := j
		for i := j + 1; i < len(arr); i++ {
			if arr[min_idx] > arr[i] {
				min_idx = i
			}
		}
		arr[j], arr[min_idx] = arr[min_idx], arr[j]
	}
}

// 3 1|5 4 2
// 1|3 5|4 2
// 1 3|5 4|2
// 1 3 4|5 2
// 1 3 4 2 5
// ...
// 1 2 3 4 5
// O(n^2)
func bubble_sort(arr []int) {
	for j := 0; j < len(arr); j++ {
		isSorted := true
		for i := 1; i < len(arr)-j; i++ {
			if arr[i] < arr[i-1] {
				isSorted = false
				arr[i], arr[i-1] = arr[i-1], arr[i]
			}
		}
		if isSorted {
			return
		}
	}
}

// 2 1 3 0 2
// 1 2 3 0 2
// 0 1 2 3 2
// 0 1 2 2 3
// O(n^2)
// better for small array or almost sorted array
func insertion_sort(arr []int) {
	for i := 1; i < len(arr); i++ {
		cur := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > cur {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = cur
	}
}

// O(n+m)
// m = max - min
func counting_sort(arr []int) {
	min := arr[0]
	max := arr[0]
	for i := 0; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
			continue
		}
		if arr[i] < min {
			min = arr[i]
		}
	}

	bucket := make([]int, max-min+1)
	for i := 0; i < len(arr); i++ {
		bucket[arr[i]-min]++
	}
	idx := 0
	for i, c := range bucket {
		for j := 0; j < c; j++ {
			arr[idx] = i + min
			idx++
		}
	}
}

// O(n*log(n))
// log2(n) = log2(2^k)
// log2(n) = k
func merge_sort(arr *[]int) {
	if len(*arr) <= 1 {
		return
	}
	n := len(*arr) / 2
	var left, right []int
	for i := 0; i < n; i++ {
		left = append(left, (*arr)[i])
	}
	for i := n; i < len(*arr); i++ {
		right = append(right, (*arr)[i])
	}
	merge_sort(&left)
	merge_sort(&right)
	*arr = merge(&left, &right)
}
func merge(l, r *[]int) []int {
	arr := make([]int, 0, len(*l)+len(*r))
	var l_i, r_i int
	for l_i < len(*l) && r_i < len(*r) {
		if (*l)[l_i] <= (*r)[r_i] {
			arr = append(arr, (*l)[l_i])
			l_i++
		} else {
			arr = append(arr, (*r)[r_i])
			r_i++
		}
	}
	for l_i < len(*l) {
		arr = append(arr, (*l)[l_i])
		l_i++
	}
	for r_i < len(*r) {
		arr = append(arr, (*r)[r_i])
		r_i++
	}
	return arr
}

// O(n^2) if pivot == 1
// O(n*log(n)) if pivot == random
func _quick_sort(arr []int, i, j int) {
	if i == j {
		return
	}

	pivot := partition(arr, i, j) // [i, pivot - 1] < pivot // [pivot + 1, j] >= pivot
	_quick_sort(arr, i, pivot)
	_quick_sort(arr, pivot+1, j)
}
func partition(arr []int, i, j int) int {
	pivot := rand.Int()%(j-i) + i
	arr[pivot], arr[i] = arr[i], arr[pivot]
	pivot = i
	s1_index := i
	s2_index := i
	for k := i + 1; k < j; k++ {
		if arr[k] >= arr[pivot] {
			s2_index++
		} else {
			s1_index++
			arr[s1_index], arr[k] = arr[k], arr[s1_index]
			s2_index++
		}
	}
	arr[pivot], arr[s1_index] = arr[s1_index], arr[pivot]
	// pivot = s1_index
	return s1_index
}
func quick_sort(arr []int) {
	_quick_sort(arr, 0, len(arr))
}

// O(n*d) d - symbol len
// memory O(n)
func radix_sotr(arr []int) {
	// 10 - система счисления
	max_symbol_len := 5
	buckets := make([][]int, 10)
	power_of_ten := 1
	for pow := 0; pow <= max_symbol_len; pow++ {
		for _, v := range arr {
			buckets[v/power_of_ten%10] = append(buckets[v/power_of_ten%10], v)
		}
		arr_idx := 0
		for i := 0; i < len(buckets); i++ {
			for j := 0; j < len(buckets[i]); j++ { // O(n)
				arr[arr_idx] = buckets[i][j]
				arr_idx++
				buckets[i] = make([]int, 0)
			}
		}
		power_of_ten *= 10
	}
}
