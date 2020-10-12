package main

func main() {

}

func sortColors(nums []int) {
	var s0, s1 int
	for i := 0; i < len(nums); i++ {

		if nums[i] == 0 && i > s0 {
			swap(&nums[i], &nums[s0])
			s0++
		}
		if nums[i] == 1 && i > s1 {

			swap(&nums[i], &nums[s1])
			s1++

		}
	}
}
	func
	swap(a, b*int)
	{
		c := b
		b = a
		a = c
	}
