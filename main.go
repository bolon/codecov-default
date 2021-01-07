package main

func compareAndSum(a int, b int) bool {
	if a > b {
		if a + b > 10 {
			return false
		} else {
			return true
		}
	}
	return false
}

func main() {
}