package Sort

/*
快速排序
算法思想简单描述：
快速排序是对冒泡排序的一种本质改进。它的基本思想是通过一趟扫描后，使得排序序列的长度能大幅度地减少。在冒泡排序中，一次扫描只能确保最大数值的数移到正确位置，而待排序序列的长度可能只减少1。
快速排序通过一趟扫描，就能确保某个数（以它为基准点吧）的左边各数都比它小，右边各数都比它大。然后又用同样的方法处理它左右两边的数，直到基准点的左右只有一个元素为止。
它是由C.A.R.Hoare于1962年提出的。

显然快速排序可以用递归实现，当然也可以用栈化解递归实现。下面的函数是用递归实现的，有兴趣的朋友可以改成非递归的。快速排序是不稳定的。最理想情况算法时间复杂度O(nlog2n)，最坏O(n2)
想到了快速排序，于是自己就用C#实现了快速排序的算法：

快速排序的基本思想：
分治法，即，分解，求解，组合 .

分解：
在 无序区R[low..high]中任选一个记录作为基准(通常选第一个记录，并记为keyValue,其下标为keyValuePosition),以此为基准划分成两个较小的 子区间R[low,keyValuePosition- 1]和R[keyValuePosition+ 1 , high],
并使左边子区间的所有记录均小于等于基准记录,右边子区间的所有记录均大于等于基准记录,基准记录无需参加后续的排序。而划分的关键是要求出 基准记录所在的位置keyValuePosition.

求解：
通过递归调用快速排序对左、右子区间R[low..keyValuePosition-1]和R[keyValuePosition+1..high]快速排序

组合：
当"求解"步骤中的两个递归调用结束时，其左、右两个子区间已有序。对快速排序而言，"组合"步骤无须做什么，可看作是空操作。

具体过程：
设序列为R[low,high]，从其中选第一个为基准，设为keyValue,然后设两个指针i和j,分别指向序列R[low,high]的起始和结束位置上：
      1），将i逐渐增大，直到找到大于keyValue的关键字为止；
      2），将j逐渐减少，直到找到小于等于keyValue的关键字为止；
      3），如果i<j，即R[i,j]的元素数大于1，则交换R[i]和R[j]；
      4），将基准记录keyValue放到合适的位置上，即i和j同时指向的位置（或者同时指向的位置-1）,则此位置为新的keyValuePosition。

*/

func QuickSort(array []int) {
	lowIndex := 0               //数组的起始位置（从0开始）
	highIndex := len(array) - 1 //数组的终止位置
	QuickSortFunc(array, lowIndex, highIndex)
}

func QuickSortFunc(array []int, low, high int) {
	var keyValuePosition int //记录关键值的下标
	//当传递的目标数组含有两个以上的元素时，进行递归调用。（即：当传递的目标数组只含有一个元素时，此趟排序结束）
	if (low < high) {
		keyValuePosition = keyValuePositionFunc(array, low, high); //获取关键值的下标（快排的核心）

		QuickSortFunc(array, low, keyValuePosition-1);  //递归调用，快排划分出来的左区间
		QuickSortFunc(array, keyValuePosition+1, high); //递归调用，快排划分出来的右区间
	}
}

func keyValuePositionFunc(array []int, low, high int) int {
	leftIndex := low   //记录目标数组的起始位置（后续动态的左侧下标）
	rightIndex := high //记录目标数组的结束位置（后续动态的右侧下标）

	keyValue := array[low]; //数组的第一个元素作为关键值
	var temp int
	//当 （左侧动态下标 == 右侧动态下标） 时跳出循环
	for leftIndex < rightIndex {

		//左侧动态下标逐渐增加，直至找到大于keyValue的下标
		for leftIndex < rightIndex && array[leftIndex] <= keyValue {
			leftIndex++
		}
		//右侧动态下标逐渐减小，直至找到小于或等于keyValue的下标
		for leftIndex < rightIndex && array[leftIndex] > keyValue {
			rightIndex--
		}

		//如果leftIndex < rightIndex，则交换左右动态下标所指定的值；当leftIndex==rightIndex时，跳出整个循环
		if leftIndex < rightIndex {
			temp = array[leftIndex]
			array[leftIndex] = array[rightIndex]
			array[rightIndex] = temp
		}

	}
	//当左右两个动态下标相等时（即：左右下标指向同一个位置），此时便可以确定keyValue的准确位置
	temp = keyValue

	//当keyValue < 左右下标同时指向的值，将keyValue与rightIndex - 1指向的值交换，并返回rightIndex - 1
	if temp < array[rightIndex] {
		array[low] = array[rightIndex-1]
		array[rightIndex-1] = temp
		return rightIndex - 1
	} else { //当keyValue >= 左右下标同时指向的值，将keyValue与rightIndex指向的值交换，并返回rightIndex
		array[low] = array[rightIndex]
		array[rightIndex] = temp
		return rightIndex
	}
}
