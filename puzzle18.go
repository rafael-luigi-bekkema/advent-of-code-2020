package main

func p18parse(expr string) (int, int) {
	var i, total int
	var op byte
	for {
		var n1 int
		if expr[i] == '(' {
			var newpos int
			n1, newpos = p18parse(expr[i+1:])
			i += newpos + 1
		} else {
			n1 = int(expr[i]) - 48
		}
		if op == 0 {
			total = n1
		} else {
			switch op {
			case '+':
				total += n1
			case '*':
				total *= n1
			}
		}

		if i >= len(expr)-3 || expr[i+1] == ')' {
			i++
			break
		}

		op = expr[i+2]

		i += 4
		if i >= len(expr) {
			break
		}
	}
	return total, i
}

func p18calc(expr string) int {
	n1, _ := p18parse(expr)
	return n1
}

func puzzle18(data []string) int {
	var sum int
	for _, line := range data {
		sum += p18calc(line)
	}
	return sum
}

func Puzzle18() int {
	data := readFile("input/input18")
	return puzzle18(data)
}

type p18op struct {
	value int
	op    byte
}

func p18parseb(expr string) (int, int) {
	var i int
	var op byte
	var calcs []p18op
	for {
		var n1 int
		if expr[i] == '(' {
			var newpos int
			n1, newpos = p18parseb(expr[i+1:])
			i += newpos + 1
		} else {
			n1 = int(expr[i]) - 48
		}
		calcs = append(calcs, p18op{n1, op})

		if i >= len(expr)-3 || expr[i+1] == ')' {
			i++
			break
		}

		op = expr[i+2]

		i += 4
		if i >= len(expr) {
			break
		}
	}
	for {
		var changed bool
		for i, op := range calcs {
			if op.op == '+' {
				prevop := calcs[i-1]
				calcs[i-1].value = prevop.value + op.value
				calcs = append(calcs[:i], calcs[i+1:]...)
				changed = true
				break
			}
		}
		if !changed {
			break
		}
	}
	var total int
	for i, op := range calcs {
		if i == 0 {
			total = op.value
		} else {
			total *= op.value
		}
	}
	return total, i
}

func p18calcb(expr string) int {
	n1, _ := p18parseb(expr)
	return n1
}

func puzzle18b(data []string) int {
	var sum int
	for _, line := range data {
		sum += p18calcb(line)
	}
	return sum
}

func Puzzle18b() int {
	data := readFile("input/input18")
	return puzzle18b(data)
}
