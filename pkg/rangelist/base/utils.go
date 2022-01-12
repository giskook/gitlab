package base

// @BinarySearchInsert 查找
func BinarySearchInsert(sortedArray []int, target int) int {
	n := len(sortedArray)
	left, right := 0, n-1
	ret := n
	for left <= right {
		mid := (right-left)>>1 + left
		if target <= sortedArray[mid] {
			ret = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return ret
}

func Max(x, y int) int {
	if x > y {
		return x
	}

	return y
}

func Min(x, y int) int {
	if x < y {
		return x
	}

	return y
}