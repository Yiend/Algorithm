package Sort

/*
 插入排序
 输入一个元素，检查数组列表中的每个元素，将其插入到一个已经排好序的数列中的适当位置，使数列依然有序，当最后一个元素放入合适位置时，该数组排序完毕。
*/
func InsertSort(array []int) {
	for i := 1; i < len(array); i++ {
		temp := array[i]
		j := i
		for j > 0 && (array[j-1] > temp) {
			array[j] = array[j-1]
			j--
			array[j] = temp
		}
	}
}
