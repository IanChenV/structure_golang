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
// 栈为先进后出
func (this *Stack) Push(x int) {
	this.nums = append(this.nums, x)
}

// 出栈
// 出栈顶元素
func (this *Stack) Pop() int {
	x := this.nums[len(this.nums)-1]
	this.nums = this.nums[:len(this.nums)-1]
	return x
}

// 栈的长度
func (this *Stack) Len() int {
	return len(this.nums)
}

// 判断栈是否为空的
func (this *Stack) IsEmpty() bool {
	return len(this.nums) == 0
}
