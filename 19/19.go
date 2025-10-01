package day19

import "fmt"

func Solve(file string) int {
	blueprints := ParseBlueprints(file)

	simulate(blueprints[0], 24)

	return 0
}

func simulate(blueprint Blueprint, time int) {
	queue := NewQueue(NewState(time))
	target := blueprint.MaxResourceCap()
	geodes := 0

	for !queue.IsEmpty() {
		current := queue.Pop()

		if current.time == 0 {
			if current.materials[GEODE] > geodes {
				geodes = current.materials[GEODE]
			}
			continue
		}

		// Attempt to build a robot
		for _, resource := range [4]Resource{ORE, CLAY, OBSIDAN, GEODE} {
			if resource == GEODE || current.robots[resource] < target[resource] {
				next := current.CopyState()
				bot, err := next.robots.Build(resource, &next.materials, blueprint)

				if err == nil {
					next.robots.Collect(&next.materials)
					next.robots.Deploy(bot)
					next.time--
					queue.Push(next)
				}

			}
		}
		// Waiting state
		wait := current.CopyState()
		wait.robots.Collect(&wait.materials)
		wait.time--
		queue.Push(wait)
	}

	fmt.Println(geodes)

	panic("")
}
