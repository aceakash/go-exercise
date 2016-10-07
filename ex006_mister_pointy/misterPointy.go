package misterPointy

func add(a, b int) int {
	return a + b
}

func addPointers(a, b *int) int {
	return add(*a, *b)
}

func addReturnPointer(a, b int) *int {
	res := add(a, b)
	return &res
}

func addPointersReturnPointer(a, b *int) *int {
	res := addPointers(a, b)
	return &res
}

func addModifyArgument(a int, b *int) {
	*b = add(a, *b)
}

func addPointerToPointer(a, b **int) int {
	return add(**a, **b)
}
