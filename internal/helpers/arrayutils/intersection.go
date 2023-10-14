package arrayutils

// Take the intersection of two arrays of uints
func Intersection(nums1 []uint, nums2 []uint) []uint {
	allNums := make(map[uint]bool)
	for _, i := range nums1 {
		allNums[i] = false
	}
	for _, i := range nums2 {
		if _, value := allNums[i]; value {
			allNums[i] = true
		}
	}
	output := []uint{}
	for i, b := range allNums {
		if b {
			output = append(output, i)
		}
	}
	return output
}
