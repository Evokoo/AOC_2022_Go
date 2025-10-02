package day20

func Solve(file string) int {
	list, ref := GenerateList(file)

	for _, node := range ref {
		removed, start := list.Remove(node)
		location := list.Step(start, removed.value-1)
		list.InsertAfter(location, removed)
	}

	arr := list.ToArray()
	offset := IndexOf(arr, 0)
	sum := 0
	for _, n := range []int{1000, 2000, 3000} {
		sum += arr[(n+offset)%len(arr)]
	}

	return sum
}
