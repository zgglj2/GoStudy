package while

func main() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	print(sum)
}
