package main

func main() {

	val := sum(1, 2, 3, 4, 5, 6, 7, 8, 9)
	println(val)

	str := multiple(1, 2, 2.3, 4, 5, 6)
	println(str)
}

func sum(value ...int) int { // this is variadic function which accept n no of parameter

	total := 0
	for _, val := range value {
		total += val
	}

	return total

}

func multiple(value ...float32) float32 { // this is variadic function which accept n no of parameter

	mult := 1
	for _, val := range value {
		mult = mult * int(val)
	}

	return float32(mult)
}
