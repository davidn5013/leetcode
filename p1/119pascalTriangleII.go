package p1

func getRow(rowIndex int) []int {
	line := make([]int, rowIndex+1)
	line[0] = 1
	for k := 0; k < rowIndex; k++ {
		line[k+1] = line[k] * (rowIndex - k) / (k + 1)
	}
	return line
}
