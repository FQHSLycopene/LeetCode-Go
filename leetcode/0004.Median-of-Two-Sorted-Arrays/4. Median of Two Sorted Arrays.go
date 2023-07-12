package leetcode

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// 假设 nums1 的长度小
	if len(nums1) > len(nums2) {
		return findMedianSortedArrays(nums2, nums1)
	}
	low, high, k, nums1Mid, nums2Mid := 0, len(nums1), (len(nums1)+len(nums2)+1)>>1, 0, 0
	for low <= high {
		// nums1:  ……………… nums1[nums1Mid-1] | nums1[nums1Mid] ……………………
		// nums2:  ……………… nums2[nums2Mid-1] | nums2[nums2Mid] ……………………
		nums1Mid = low + (high-low)>>1 // 分界限右侧是 mid，分界线左侧是 mid - 1
		nums2Mid = k - nums1Mid
		if nums1Mid > 0 && nums1[nums1Mid-1] > nums2[nums2Mid] { // nums1 中的分界线划多了，要向左边移动
			high = nums1Mid - 1
		} else if nums1Mid != len(nums1) && nums1[nums1Mid] < nums2[nums2Mid-1] { // nums1 中的分界线划少了，要向右边移动
			low = nums1Mid + 1
		} else {
			// 找到合适的划分了，需要输出最终结果了
			// 分为奇数偶数 2 种情况
			break
		}
	}
	midLeft, midRight := 0, 0
	if nums1Mid == 0 {
		midLeft = nums2[nums2Mid-1]
	} else if nums2Mid == 0 {
		midLeft = nums1[nums1Mid-1]
	} else {
		midLeft = max(nums1[nums1Mid-1], nums2[nums2Mid-1])
	}
	if (len(nums1)+len(nums2))&1 == 1 {
		return float64(midLeft)
	}
	if nums1Mid == len(nums1) {
		midRight = nums2[nums2Mid]
	} else if nums2Mid == len(nums2) {
		midRight = nums1[nums1Mid]
	} else {
		midRight = min(nums1[nums1Mid], nums2[nums2Mid])
	}
	return float64(midLeft+midRight) / 2
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}

func findMedianSortedArraysLycopene(nums1 []int, nums2 []int) float64 {
	l := len(nums1) + len(nums2)
	if l == 0 {
		return 0
	}
	if l == 1 {
		return getKth(nums1, nums2, 0)
	}
	if (l)%2 == 0 {
		return (getKth(nums1, nums2, l/2-1) + getKth(nums1, nums2, l/2)) / 2
	} else {
		return getKth(nums1, nums2, l/2)
	}
	return 0
}

func getKth(nums1 []int, nums2 []int, k int) float64 {
	if len(nums1) == 0 {
		return float64(nums2[k])
	}
	if len(nums2) == 0 {
		return float64(nums1[k])
	}

	if k == 0 {
		if nums1[0] > nums2[0] {
			return float64(nums2[0])
		}
		return float64(nums1[0])
	} else if k == 1 {
		if nums1[0] < nums2[0] {
			return getKth(nums1[1:], nums2, k-1)
		}
		return getKth(nums1, nums2[1:], k-1)
	}
	n1, n2 := k/2-1, k/2-1
	if k/2-1 >= len(nums1) {
		n1 = len(nums1) - 1
	}
	if k/2-1 >= len(nums2) {
		n2 = len(nums2) - 1
	}
	if nums1[n1] > nums2[n2] {
		return getKth(nums1, nums2[n2+1:], k-n2-1)
	} else {
		return getKth(nums1[n1+1:], nums2, k-n1-1)
	}

}
