package day20

func Solve(file string, part int) int {
	key, cycles := GetModifiers(part)
	list, ref := GenerateList(file, key)
	for range cycles {
		for _, node := range ref {
			removed, start := list.Remove(node)
			location := list.Step(start, removed.value-1)
			list.InsertAfter(location, removed)
		}
	}
	return GetGrooveCoordinate(list)
}
