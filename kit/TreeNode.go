package kit

// TreeNode :=
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// NULL replacement for nil
var NULL = -1

// SliceToTree :=
func SliceToTree(s []int) *TreeNode {
	length := len(s)
	if length == 0 {
		return nil
	}

	root := &TreeNode{
		Val: s[0],
	}

	queue := make([]*TreeNode, 1, length*2)
	queue[0] = root

	i := 1
	for i < len(s) {
		node := queue[0]
		queue = queue[1:]

		if i < length && s[i] != NULL {
			node.Left = &TreeNode{Val: s[i]}
			queue = append(queue, node.Left)
		}
		i++

		if i < length && s[i] != NULL {
			node.Right = &TreeNode{Val: s[i]}
			queue = append(queue, node.Right)
		}
		i++
	}

	return root
}

// TreeToSlice :=
func TreeToSlice(t *TreeNode) []int {
	result := make([]int, 0, 1024)

	queue := []*TreeNode{t}

	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			node := queue[i]
			if node == nil {
				result = append(result, NULL)
			} else {
				result = append(result, node.Val)
				queue = append(queue, node.Left, node.Right)
			}
		}
		queue = queue[size:]
	}

	// Skip over the last trail of nulls
	// [3 9 20 -1 -1 15 7 -1 -1 -1 -1] becomes
	// [3 9 20 -1 -1 15 7]
	i := len(result)
	for i > 0 && result[i-1] == NULL {
		i--
	}

	return result[:i]
}
