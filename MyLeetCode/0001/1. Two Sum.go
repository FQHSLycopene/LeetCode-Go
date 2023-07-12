package myleetcode

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, num := range nums {
		if i2, ok := m[target-num]; ok {
			return []int{i, i2}
		}
		m[num] = i
	}
	return nil
}
