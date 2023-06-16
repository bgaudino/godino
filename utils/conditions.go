package utils

func All(conditions ...bool) bool {
	for _, c := range conditions {
		if !c {
			return false
		}
	}
	return true
}

func Any(conditions ...bool) bool {
	for _, c := range conditions {
		if c {
			return true
		}
	}
	return false
}
