package main

import "fmt"

type TreeNode struct {
	Value     int
	LeftNode  *TreeNode
	RightNode *TreeNode
}

func NewTreeNode(val int, left, right *TreeNode) *TreeNode {
	return &TreeNode{Value: val, LeftNode: left, RightNode: right}
}

func (t *TreeNode) Insert(val int) {
	if val < t.Value {
		if t.LeftNode == nil {
			t.LeftNode = NewTreeNode(val, nil, nil)
		} else {
			t.LeftNode.Insert(val)
		}
	} else if val > t.Value {
		if t.RightNode == nil {
			t.RightNode = NewTreeNode(val, nil, nil)
		} else {
			t.RightNode.Insert(val)
		}
	}
}

func Search(node *TreeNode, val int) *TreeNode {
	if node == nil || node.Value == val {
		return node
	} else if val < node.Value {
		return Search(node.LeftNode, val)
	} else if val > node.Value {
		return Search(node.RightNode, val)
	}
	return nil

}

func Traverse(t *TreeNode) {
	if t != nil {
		Traverse(t.LeftNode)
		fmt.Println(t.Value)
		Traverse(t.RightNode)
	}
}

func Delete(t *TreeNode, value int) *TreeNode {
	/*
	   If the node being deleted has no children, simply delete it.
	   If the node being deleted has one child, delete it and plug the child into the spot where the deleted node was.
	   When deleting a node with two children, replace the deleted node with the successor node.
	   The successor node is the child node whose value is the least of all values that are greater than the deleted node.
	   If the successor node has a right child, after plugging the successor node into the spot of the deleted node,
	   take the right child of the successor node and turn it into the left child of the parent of the successor node.
	*/
	if t == nil {
		return nil
	} else if value < t.Value {
		t.LeftNode = Delete(t.LeftNode, value)
		return t
	} else if value > t.Value {
		t.RightNode = Delete(t.RightNode, value)
		return t
	} else if t.Value == value {
		if t.LeftNode == nil {
			return t.RightNode
		} else if t.RightNode == nil {
			return t.LeftNode
		} else {
			fmt.Printf("%v\n", t)
			t.RightNode = lift(t.RightNode, t)
			fmt.Println(t, t.RightNode)
			return t
		}
	}
	return nil

}

func lift(node, nodeToDelete *TreeNode) *TreeNode {
	fmt.Println(node, nodeToDelete)
	if node.LeftNode != nil {

		node.LeftNode = lift(node.LeftNode, nodeToDelete)
		fmt.Println("Lf", node)
		return node
	} else {
		fmt.Println("RF", node)
		nodeToDelete.Value = node.Value
		return node.RightNode
	}

}

func main() {
	rootNode := NewTreeNode(50, nil, nil)
	rootNode.Insert(25)
	rootNode.Insert(75)
	rootNode.Insert(11)
	rootNode.Insert(33)
	rootNode.Insert(30)
	rootNode.Insert(40)
	rootNode.Insert(61)
	rootNode.Insert(89)
	rootNode.Insert(52)
	rootNode.Insert(82)
	rootNode.Insert(95)

	Traverse(rootNode)
	n := Search(rootNode, 14)
	fmt.Println(n)
	fmt.Println("====")
	Delete(rootNode, 50)
	Traverse(rootNode)

}
