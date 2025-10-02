package day21

func Solve(file string, part int) int {
	values, equations := ParseInput(file)
	tracer := NewTracer("humn")

	for !equations.IsEmpty() {
		equation := equations.Pop()

		if value, solved := equation.Solve(values); solved {
			tracer.Log(equation, values)
			values.Add(equation[0], value)
		} else {
			equations.Push(equation)
		}
	}

	if part == 2 {
		return FindHumanValue(tracer.log)
	}

	return values.Get("root")
}
