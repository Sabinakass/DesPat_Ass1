package main

import (
	"fmt"
	"sort"
)

// SortingAlgorithms interface for future concrete strategies

type SortingAlgorithm interface {
	Sort(arr []int) []int
}

//BubbleSort is concrete strategy

type BubbleSort struct {
}

func (bs *BubbleSort) Sort(arr []int) []int {
	n := len(arr)
	for i := 0; i < n; i++ {
		for j := 0; j < n-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

//MergeSort is concrete strategy

type MergeSort struct {
}

func (ms *MergeSort) Sort(arr []int) []int {
	sort.Sort(sort.IntSlice(arr))
	return arr
}

//SimpleSort is concrete strategy

type SimpleSort struct {
}

func (qs *SimpleSort) Sort(arr []int) []int {
	sort.Ints(arr)
	return arr
}

//Context

type Sorter struct {
	SortingAlgorithm
}

func NewSorter(algo SortingAlgorithm) *Sorter {
	return &Sorter{SortingAlgorithm: algo}
}

func (s *Sorter) SortArr(arr []int) []int {
	return s.SortingAlgorithm.Sort(arr)
}

func main() {
	arr := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}

	bubbleSorter := NewSorter(&BubbleSort{})
	quickSorter := NewSorter(&SimpleSort{})
	mergeSorter := NewSorter(&MergeSort{})

	fmt.Println("Bubble Sort:", bubbleSorter.SortArr(arr))
	fmt.Println("Quick Sort:", quickSorter.SortArr(arr))
	fmt.Println("Merge Sort:", mergeSorter.SortArr(arr))
}
