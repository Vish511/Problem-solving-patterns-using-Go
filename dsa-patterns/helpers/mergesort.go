package helpers

func MergeSort(arr []int) []int {
	if len(arr) == 1 {
		return arr
	}
	var result []int
	var mid int = len(arr) / 2
	var left = arr[0:mid]
	var right = arr[mid:]
	sortedLeft := MergeSort(left)
	sortedRight := MergeSort(right)
	result = append(result, MergeHelper(sortedLeft, sortedRight)...)
	return MergeHelper(sortedLeft, sortedRight)
}

func MergeHelper(arr1, arr2 []int) []int {
	var i, j int
	var result = make([]int, 0, len(arr1)+len(arr2))
	for i < len(arr1) && j < len(arr2) {
		if arr1[i] < arr2[j] {
			result = append(result, arr1[i])
			i++
		} else {
			result = append(result, arr2[j])
			j++
		}
	}

	for i < len(arr1) {
		result = append(result, arr1[i])
		i++
	}
	for j < len(arr2) {
		result = append(result, arr2[j])
		j++
	}
	return result

}
