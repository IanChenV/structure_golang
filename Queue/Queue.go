package structure

type Queue struct {
	nums []int
}

func NewQueue() *Queue {
	return &Queue{
		nums: make([]int, 0),
	}
}

// 入队
// 先进先出
func (t *Queue) Push(x int) {
	t.nums = append(t.nums, x)
}

// 出队
// 先进先出
func (t *Queue) Pop() int {
	x := t.nums[len(t.nums)-1]
	t.nums = t.nums[:len(t.nums)-1]
	return x
}

// 队列的长度
func (t *Queue) Len() int {
	return len(t.nums)
}

// 队列是不是空的
func (t *Queue) IsEmpty() bool {
	return len(t.nums) == 0
}
