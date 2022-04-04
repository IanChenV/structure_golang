package structure

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var NULL = -1 << 63

// 利用数组转二叉树
func Ints2TreeNode(nums []int) *TreeNode {
	size := len(nums)
	if size == 0 {
		return nil
	}

	// 根节点
	root := &TreeNode{Val: nums[0]}

	// 生成队列
	queue := make([]*TreeNode, 0, 2*size)
	queue[0] = root

	index := 1
	for index < size {
		node := queue[0]
		queue = queue[1:]

		if index < size && nums[index] != NULL {
			node.Left = &TreeNode{Val: nums[index]}
			queue = append(queue, node.Right)
		}
		index++

		if index < size && nums[index] != NULL {
			node.Right = &TreeNode{Val: nums[index]}
			queue = append(queue, node.Right)
		}
		index++
	}
	return root
}

// 二叉树节点的查找
func GetTragetNode(root *TreeNode, target int) *TreeNode {
	if root == nil || root.Val == target {
		return root
	}

	node := GetTragetNode(root.Left, target)
	if node != nil {
		return node
	}
	return GetTragetNode(root.Right, target)
}

// 根据前序遍历和中序遍历转换二叉树
func PreIn2Tree(pre []int, in []int) *TreeNode {
	if len(pre) != len(in) {
		panic("preIn2Tree中两个切片的长度不相等")
	}

	if len(in) == 0 {
		return nil
	}

	root := &TreeNode{
		Val: pre[0],
	}

	if len(in) == 1 {
		return root
	}

	rootIndex := indexOf(root.Val, in)
	root.Left = PreIn2Tree(pre[1:rootIndex+1], in[:rootIndex])
	root.Right = PreIn2Tree(pre[rootIndex+1:], in[rootIndex+1:])
	return root
}

// 根据后序遍历和中序遍历转换二叉树
func InPost2Tree(in []int, post []int) *TreeNode {
	if len(post) != len(in) {
		panic("inPost2Tree 中两个切片的长度不相等")
	}
	if len(in) == 0 {
		return nil
	}

	root := &TreeNode{
		Val: post[len(post)-1],
	}

	if len(in) == 1 {
		return root
	}

	rootIndex := indexOf(root.Val, in)
	root.Left = InPost2Tree(in[:rootIndex], post[:rootIndex])
	root.Right = InPost2Tree(in[rootIndex+1:], post[rootIndex:len(post)-1])
	return root
}

// 寻找根节点的下标
func indexOf(val int, nums []int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			return i
		}
	}
	msg := fmt.Sprintf("%d 不存在于 %v中", val, nums)
	panic(msg)
}

// 二叉树的前序遍历
func Tree2Preorder(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}

	res := []int{root.Val}
	res = append(res, Tree2Preorder(root.Left)...)
	res = append(res, Tree2Preorder(root.Right)...)
	return res
}

// 二叉树的后序遍历
func Tree2Inorder(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}

	res := Tree2Inorder(root.Left)
	res = append(res, root.Val)
	res = append(res, Tree2Inorder(root.Right)...)
	return res
}

// 二叉树的后序遍历
func Tree2Postorder(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}

	res := Tree2Postorder(root.Left)
	res = append(res, Tree2Postorder(root.Right)...)
	res = append(res, root.Val)
	return res
}

// 树的节点是否相等
func (tn *TreeNode) Equal(a *TreeNode) bool {
	if tn == nil && a == nil {
		return true
	}

	if tn == nil || a == nil || tn.Val != a.Val {
		return false
	}

	return tn.Equal(a.Left) && tn.Equal(a.Right)
}

// 把TreeNode转换为[]int
// 层序化遍历
func Tree2ints(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	res := []int{}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			res = append(res, queue[i].Val)
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}

			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		queue = queue[size:]
	}
	return res
}

func T2s(head *TreeNode, array *[]int) {
	*array = append(*array, head.Val)
	if head.Left != nil {
		T2s(head.Left, array)
	}

	if head.Right != nil {
		T2s(head.Right, array)
	}
}
