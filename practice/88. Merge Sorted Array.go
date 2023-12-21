package main

func merge(nums1 []int, m int, nums2 []int, n int) {
	if n == 0 {
		return
	}
	if m == 0 {
		for i := 0; i < n; i++ {
			nums1[i] = nums2[i]
		}
	}
	indexOne := 0
	indexTwo := 0
	arr := []int{}
	for indexOne < m || indexTwo < n {
		if nums1[indexOne] < nums2[indexTwo] {
			arr = append(arr, nums1[indexOne])
			indexOne++
		} else {
			arr = append(arr, nums2[indexTwo])
			indexTwo++
		}
	}

	for indexOne < m {
		arr = append(arr, nums1[indexOne])
		indexOne++
	}

	for indexTwo < n {
		arr = append(arr, nums2[indexTwo])
		indexTwo++
	}
	copy(nums1, arr)
}
