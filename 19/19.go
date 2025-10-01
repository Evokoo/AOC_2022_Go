package day19

import "fmt"

func Solve(file string) int {
	blueprints := ParseBlueprints(file)

	simulate(blueprints[0], 24)

	return 0
}

func simulate(blueprint Blueprint, time int) {
	queue := NewQueue(State{time, NewRobotMap(), NewMaterialList()})
	geodeCount := 0

	for !queue.IsEmpty() {
		current := queue.Pop()

		if current.time == 0 {
			if current.materials.geode > geodeCount {
				geodeCount = current.materials.geode
			}
			continue
		}

		//Attempt to build robots
		for _, rt := range RobotTypes {
			robotsCopy := current.robots.Clone()
			materialCopy := current.materials
			newRobot, err := robotsCopy.Build(rt, &materialCopy, blueprint)
			robotsCopy.Collect(&materialCopy)

			if err == nil {
				robotsCopy.Deploy(newRobot)
				queue.Push(State{current.time - 1, robotsCopy, materialCopy})
			}
		}

		//Don't build any
		current.robots.Collect(&current.materials)
		queue.Push(State{current.time - 1, current.robots, current.materials})
	}

	fmt.Println(geodeCount)

	panic("")
}
