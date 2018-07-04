package Sort

/*
冒泡排序算法

*/
func BubbleSort(array []int) {
	for i := 0; i < len(array); i++ {
		for j := i + 1; j < len(array); j++ {
			if array[j] < array[i] {
				temp := array[j]
				array[j] = array[i]
				array[i] = temp
			}
		}
	}
}
