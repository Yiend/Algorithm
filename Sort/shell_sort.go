package Sort

/*
希尔排序是将组分段,进行插入排序
*/
func ShellSort(array []int) {
	var inc int
	for inc = 1; inc <= len(array)/9; inc = 3*inc + 1 {
	}
	for ; inc > 0; inc /= 3 {
		for i := inc + 1; i <= len(array); i += inc {
			temp := array[i-1]
			j := i
			for j > inc && (array[j-inc-1] > temp) {
				array[j-1] = array[j-inc-1]
				j -= inc
			}
			array[j-1] = temp
		}
	}
}
