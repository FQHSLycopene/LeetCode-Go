package main

import "fmt"

func main() {
	a := findMedianSortedArraysLycopene([]int{1}, []int{2, 3, 4, 5, 6, 7, 8})
	fmt.Println(a)
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
