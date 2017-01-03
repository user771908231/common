package utils

func IsSliceContainsInt32(sl []int32, i int32) bool {
	for _, s := range sl {
		if s == i {
			return true
		}
	}
	return false
}
