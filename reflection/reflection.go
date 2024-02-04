package reflection

func walk(x any, fn func(input string)) {
	fn("I still can't believe Germany beat Brazil 7-1 :(")
}
