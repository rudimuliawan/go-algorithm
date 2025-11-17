package sort

func SelectionSort(data []int) []int {
	n := len(data)

	for i := 0; i < n; i++ {
		min_ := i
		for j := i + 1; j < n; j++ {
			if data[j] < data[min_] {
				min_ = j
			}
		}

		temp := data[min_]
		data[min_] = data[i]
		data[i] = temp
	}

	return data
}
