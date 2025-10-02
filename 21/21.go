package day21

func Solve(file string, part int) int {
	values, equations := ParseInput(file)
	target, log := "humn", [][3]string{}

	for !equations.IsEmpty() {
		equation := equations.Pop()

		if value, solved := equation.Solve(values); solved {
			equation.Trace(&target, &log, values)
			values.Add(equation[0], value)
		} else {
			equations.Push(equation)
		}
	}

	if part == 2 {
		return FindHumanValue(log)
	}

	return values.Get("root")
}
