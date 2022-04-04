package structure

type Stack struct {
	nums []int
}

func NewStack() *Stack {
	return &Stack{
		nums: make([]int, 0),
	}
}

// 入栈
func (s *Stack) Push(x int) {
	s.nums = append(s.nums, x)
}

// 出栈
func (s *Stack) Pop() int {
	x := s.nums[len(s.nums)-1]
	s.nums = s.nums[:len(s.nums)-1]
	return x
}

// 栈的长度
func (s *Stack) Len() int {
	return len(s.nums)
}

// 栈是否为空
func (s *Stack) IsEmpty() bool {
	return len(s.nums) == 0
}
