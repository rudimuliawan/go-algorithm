package sort

func ShellSort(data []int) []int {
	n := len(data)
	h := 1

	for h < n/3 {
		h = 3*h + 1
	}

	for h >= 1 {
		for i := h; i < n; i++ {
			for j := i; j >= h && data[j] < data[j-h]; j -= h {
				temp := data[j]
				data[j] = data[j-1]
				data[j-1] = temp
			}
		}

		h = h / 3
	}

	return data
}
