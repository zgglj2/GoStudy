package main

func main() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	println(sum)

	sum = 1
	for {
		sum += sum
		if sum > 1000 {
			break
		}
		println(sum)
	}
}
