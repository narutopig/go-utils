package utils

func BinarySearchI(val int, arr []int) (index int) {
	low := 0
	high := len(arr)

	for low <= high {
		mid := (low + high) / 2
		if arr[mid] == val {
			return index
		} else if arr[mid] < val {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return -1
}

func BinarySearchB(val byte, arr []byte) (index int) {
	low := 0
	high := len(arr)

	for low <= high {
		mid := (low + high) / 2
		if arr[mid] == val {
			return index
		} else if arr[mid] < val {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return -1
}

func BinarySearchF64(val float64, arr []float64) (index int) {
	low := 0
	high := len(arr)

	for low <= high {
		mid := (low + high) / 2
		if arr[mid] == val {
			return index
		} else if arr[mid] < val {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return -1
}

func BinarySearchF32(val float32, arr []float32) (index int) {
	low := 0
	high := len(arr)

	for low <= high {
		mid := (low + high) / 2
		if arr[mid] == val {
			return index
		} else if arr[mid] < val {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return -1
}

func BinarySearchS(val string, arr []string) (index int) {
	low := 0
	high := len(arr)

	for low <= high {
		mid := (low + high) / 2
		if arr[mid] == val {
			return index
		} else if arr[mid] < val {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return -1
}
