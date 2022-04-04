package structure

type Queue struct {
	nums []int
}

func NewQueue() *Queue {
	return &Queue{
		nums: make([]int, 0),
	}
}

// 返回队列的长度
func (q *Queue) Len() int {
	return len(q.nums)
}

// 判断队列是否为空
func (q *Queue) IsEmpty() bool {
	return len(q.nums) == 0
}

//入队列
func (q *Queue) Push(val int) {
	q.nums = append(q.nums, val)
}

// 出队列
func (q *Queue) Pop() int {
	x := q.nums[0]
	q.nums = q.nums[1:]
	return x
}
