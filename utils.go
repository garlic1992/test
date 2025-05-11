package main

// QuickSort 快速排序
func QuickSort(nums []int, a int, b int) {
	if a+1 >= b {
		return
	}
	left, right := a, b
	flag := true //判断当前应该移动left指针还是right指针，flag==true时移动right，flag==false时移动left
	for left < right {
		if nums[left] > nums[right] {
			nums[left], nums[right] = nums[right], nums[left]
			flag = !flag //如果发生交换，则更换移动的指针
		}
		if flag {
			right--
		} else {
			left++
		}
	}
	QuickSort(nums, a, left-1)
	QuickSort(nums, left+1, b)
}

// MergeSort 归并排序
func MergeSort(nums []int) {
	temp := make([]int, len(nums))
	for k := 1; k < len(nums); k *= 2 {
		for i := 0; i < len(nums); i = i + 2*k {
			if i+k < len(nums) {
				//获取两个有序数组
				a := nums[i : i+k]
				var b []int
				if i+2*k < len(nums) {
					b = nums[i+k : i+2*k]
				} else {
					b = nums[i+k:]
				}
				//合并两个有序数组a[]和b[]
				pa := 0
				pb := 0
				j := i
				for pa < len(a) && pb < len(b) {
					if a[pa] <= b[pb] {
						temp[j] = a[pa]
						pa++
					} else {
						temp[j] = b[pb]
						pb++
					}
					j++
				}
				for pa < len(a) {
					temp[j] = a[pa]
					pa++
					j++
				}
				for pb < len(b) {
					temp[j] = b[pb]
					pb++
					j++
				}
			} else {
				for j := i; j < len(nums); j++ {
					temp[j] = nums[j]
				}
			}
		}
		for i := 0; i < len(nums); i++ {
			nums[i] = temp[i]
		}
	}
}

// Down 下沉操作，调整node至二叉树的合适位置
// node的左右孩子（如果存在）：2*node+1,2*node+2
func Down(heap []int, node int, heapSize int, less func(int, int) bool) {
	for 2*node+1 < heapSize {
		//取node的左右孩子中的较大孩子，其位置保存为maxChild
		maxChild := 2*node + 1
		if maxChild+1 < heapSize {
			if heap[maxChild+1] > heap[maxChild] {
				maxChild++
			}
		}
		if less(heap[node], heap[maxChild]) { //将node与其较大的孩子maxChild交换位置，并继续循环上述操作
			heap[node], heap[maxChild] = heap[maxChild], heap[node]
			node = maxChild
		} else {
			return
		}
	}
}

// Up 上浮操作，调整node至二叉树的合适位置
// node的父节点（如果存在）：(node-1)/2
func Up(heap []int, node int, less func(int, int) bool) {
	for node > 0 {
		father := (node - 1) / 2
		if less(heap[father], heap[node]) {
			heap[node], heap[father] = heap[father], heap[node]
			node = father
		} else {
			return
		}
	}
}

// CreateHeap 建堆
// 堆的最后一个非叶节点的序号：heapSize/2-1
// 建堆的时间复杂度为O(n)
func CreateHeap(heap []int) {
	n := len(heap)
	for i := n/2 - 1; i >= 0; i-- {
		Down(heap, i, n, Less)
	}
}

// Less 此处为大顶堆（或从小到大排序），若想实现小顶堆（或从大到小排序）则改为i>j
func Less(i, j int) bool {
	return i < j
}

// HeapSort 堆排序
func HeapSort(heap []int) {
	CreateHeap(heap)
	for i := len(heap) - 1; i > 0; i-- {
		heap[0], heap[i] = heap[i], heap[0]
		Down(heap, 0, i, Less)
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

//
