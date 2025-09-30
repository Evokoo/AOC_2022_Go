package day18

func Solve(file string, part int) int {
	lavaDroplet := NewDroplet(file)
	if part == 2 {
		return lavaDroplet.ExternalSurfaceArea()
	}
	return lavaDroplet.TotalSurfaceArea()
}
