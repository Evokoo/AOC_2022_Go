package day18

func Solve(file string) int {
	lavaDroplet := NewDroplet(file)
	return lavaDroplet.TotalSurfaceArea()
}
