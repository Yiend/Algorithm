package Sort

/*
堆排序
堆排序是一种树形选择排序，是对直接选择排序的有效改进。
堆的定义如下：具有n个元素的序列（h1,h2,...,hn),当且仅当
满足（hi>=h2i,hi>=2i+1）或（hi<=h2i,hi<=2i+1）(i=1,2,...,n/2)
时称之为堆。在这里只讨论满足前者条件的堆。

由堆的定义可以看出，堆顶元素（即第一个元素）必为最大项。完全二叉树可以很直观地表示堆的结构。堆顶为根，其它为左子树、右子树。
初始时把要排序的数的序列看作是一棵顺序存储的二叉树，调整它们的存储顺序，使之成为一个堆，这时堆的根节点的数最大。然后将根节点与堆的最后一个节点交换。
然后对前面(n-1)个数重新调整使之成为堆。依此类推，直到只有两个节点的堆，并对它们作交换，最后得到有n个节点的有序序列。
从算法描述来看，堆排序需要两个过程，一是建立堆，二是堆顶与堆的最后一个元素交换位置。所以堆排序有两个函数组成。一是建堆的渗透函数，二是反复调用渗透函数实现排序的函数。
*/
func HeapSort(array []int) {
	buildMaxHeap(array) //创建大顶推（初始状态看做：整体无序）
	for i := len(array) - 1; i > 0; i-- {
		//将堆顶元素依次与无序区的最后一位交换（使堆顶元素进入有序区）
		temp := array[0]
		array[0] = array[i]
		array[i] = temp

		maxHeapify(array, 0, i) //重新将无序区调整为大顶堆
	}
}

//创建大顶推（根节点大于左右子节点）
func buildMaxHeap(array []int) {
	//根据大顶堆的性质可知：数组的前半段的元素为根节点，其余元素都为叶节点 从最底层的最后一个根节点开始进行大顶推的调整
	for i := len(array)/2 - 1; i >= 0; i-- {
		maxHeapify(array, i, len(array)) //调整大顶堆
	}
}

//大顶推的调整过程
//currentIndex：待调整元素在数组中的位置（即：根节点）
//heapSize：堆中所有元素的个数
func maxHeapify(array []int, currentIndex, heapSize int) () {

	left := 2*currentIndex + 1  //左子节点在数组中的位置
	right := 2*currentIndex + 2 //右子节点在数组中的位置
	large := currentIndex       //记录此根节点、左子节点、右子节点 三者中最大值的位置

	//与左子节点进行比较
	if left < heapSize && array[left] > array[large] {
		large = left
	}

	//与右子节点进行比较
	if right < heapSize && array[right] > array[large] {
		large = right
	}

	//如果 currentIndex != large 则表明 large 发生变化（即：左右子节点中有大于根节点的情况）
	if currentIndex != large {
		//将左右节点中的大者与根节点进行交换（即：实现局部大顶堆）
		temp := array[currentIndex]
		array[currentIndex] = array[large]
		array[large] = temp

		//以上次调整动作的large位置（为此次调整的根节点位置），进行递归调整
		maxHeapify(array, large, heapSize)
	}
}
