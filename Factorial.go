package main

func factorialIteration (factNumber int) int{
	factResult := 1
	for i := 1; i<=factNumber; i++{
		factResult = factResult * i
	}
	return factResult
}

func factorialRecurs(factNumber int) int {
	if factNumber < 0 {
		return 0
	} else if factNumber == 0 {
		return 1
	} else {
		return factNumber * factorialRecurs(factNumber-1)
	}
}


func main() {
	number := 6
	println(factorialIteration(number))
	println(factorialRecurs(number))
	//factorialRecurs(number)
}
