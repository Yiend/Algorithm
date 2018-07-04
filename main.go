package main

import (
	"Algorithm/Sort"
	"fmt"
	"strings"
)

func main()  {
    array:=[]int{1, 5, 3, 6, 10, 55, 9, 2, 87, 12, 34, 75, 33, 47,-1,16,-22,122,108,54,-14,102,96,22}

	fmt.Println("希尔排序")
    Sort.ShellSort(array)
	fmt.Println(ArrayToString(array))
    fmt.Println("-------------------------------------")

	fmt.Println("插入排序")
	Sort.InsertSort(array)
	fmt.Println(ArrayToString(array))
	fmt.Println("-------------------------------------")

	fmt.Println("选择排序")
	Sort.SelectionSort(array)
	fmt.Println(ArrayToString(array))
	fmt.Println("-------------------------------------")

	fmt.Println("直接插入排序")
	Sort.StraightInsertionSort(array)
	fmt.Println(ArrayToString(array))
	fmt.Println("-------------------------------------")

	fmt.Println("冒泡排序算法")
	Sort.BubbleSort(array)
	fmt.Println(ArrayToString(array))
	fmt.Println("-------------------------------------")

	fmt.Println("快速排序")
	Sort.QuickSort(array)
	fmt.Println(ArrayToString(array))
	fmt.Println("-------------------------------------")

	fmt.Println("堆排序")
	Sort.HeapSort(array)
	fmt.Println(ArrayToString(array))
	fmt.Println("-------------------------------------")

	fmt.Println("归并排序")
	Sort.MergeSort(array)
	fmt.Println(ArrayToString(array))
	fmt.Println("-------------------------------------")
}


func ArrayToString(arr interface{})string  {
	return strings.Replace(strings.Trim(fmt.Sprint(arr), "[]"), " ", ",", -1)
}