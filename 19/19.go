package day19

import "fmt"

func Solve(file string) int {
	blueprints := ParseBlueprints(file)

	for _, bp := range blueprints {
		fmt.Printf("ID: %d\n", bp.id)
		fmt.Printf("Ore: %v\n", bp.ore)
		fmt.Printf("Clay: %v\n", bp.clay)
		fmt.Printf("Obsidian: %v\n", bp.obsidian)
		fmt.Printf("Geode: %v\n\n", bp.geode)
	}

	return 0
}
