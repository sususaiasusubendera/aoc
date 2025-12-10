package main

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func best12(s string) string {
	keep := 12
	drop := len(s) - keep
	stack := make([]byte, 0, keep)
	for i := 0; i < len(s); i++ {
		d := s[i]

		for len(stack) > 0 && drop > 0 && stack[len(stack)-1] < d {
			stack = stack[:len(stack)-1] // pop
			drop--
		}

		stack = append(stack, d)
	}

	return string(stack[:keep])
}
