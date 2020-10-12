package main

func main() {
	sortColors([]int{1, 0, 0})
}

func sortColors(nums []int) {
	var s0, s1 int
	for i := 0; i < len(nums); i++ {
		switch nums[i] {
		case 2:
		case 1:
			swap(&nums[i], &nums[s1+s0])
			s1++
		case 0:
			swap(&nums[i], &nums[s0])
			if nums[i] == 1 {
				swap(&nums[i], &nums[s1+s0])
			}
			s0++
		}
	}
}
func swap(a, b *int) {
	c := *b
	*b = *a
	*a = c
}
