package structure

type TreeNode struct {
	Val      int
	L        int
	RightVal int
	Left     *TreeNode
	Right    *TreeNode
}

// εεΊιε
func buildTree(nums []int, left, right int) *TreeNode {
	if left == right {
		return &TreeNode{
			Val:      nums[left],
			LeftVal:  left,
			RightVal: right,
		}
	}

	return &TreeNode{}
}
