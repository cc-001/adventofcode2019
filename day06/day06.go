package day06

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Node struct {
	val string
	depth int
	parent string
	parent_p *Node
	children []*Node
}

type Tree struct {
	head *Node
}

func (t Tree) insert_r(node *Node, cur *Node) bool {
	if (*cur).val == node.parent {
		cur.children = append(cur.children, node)
		return true
	}
	for _, i := range cur.children {
		if t.insert_r(node, i) {
			return true
		}
	}
	return false
}

func (t Tree) insert(node *Node) bool {
	return t.insert_r(node, t.head)
}

func (t Tree) set_depth_r(node *Node, depth int) {
	node.depth = depth
	for _, i := range(node.children) {
		t.set_depth_r(i, depth + 1)
	}
}

func (t Tree) checksum_r(node *Node, cs *int)  {
	*cs += node.depth
	for _, i := range(node.children) {
		t.checksum_r(i, cs)
	}
}
func (t Tree) checksum() int {
	t.set_depth_r(t.head, 1)
	cs := 0
	t.checksum_r(t.head, &cs)
	return cs
}

func (t Tree) find_r(node *Node, val string) *Node {
	if node.val == val {
		return node
	}
	for _, i := range node.children {
		ret := t.find_r(i, val)
		if ret != nil {
			return ret
		}
	}
	return nil
}

func (t Tree) find(val string) *Node {
	return t.find_r(t.head, val)
}

func (t Tree) path_count(start_val string, finish_val string) int {
	start := t.find(start_val)
	if start == nil {
		return -1
	}
	start.depth = 0

	work := make([]*Node, 0)
	work = append(work, start)
	visited := make(map[string]int)

	for {
		if len(work) == 0 {
			return -1
		}

		test := work[0]
		work = work[1:]
		if _, ok := visited[test.val]; ok {
			continue
		}
		visited[test.val] = test.depth

		if test.val == finish_val {
			return test.depth - 2
		} else {
			// this is stupid, too lazy
			parent := t.find_r(t.head, test.parent)
			if parent != nil {
				parent.depth = test.depth + 1
				work = append(work, parent)
			}
			for _, i := range test.children {
				i.depth = test.depth + 1
				work = append(work, i)
			}
		}
	}
	return -1
}

func Day06_parse(fileName string) *Tree {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	trees := make([]*Tree, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		strs := strings.Split(scanner.Text(), ")")
		added := false
		new_node := new(Node)
		new_node.val = strs[1]
		new_node.parent = strs[0]
		for _, tree := range trees {
			if tree.insert(new_node) {
				added = true
				break
			}
		}
		if added == false {
			new_tree := new(Tree)
			new_tree.head = new_node
			trees = append(trees, new_tree)
		}
	}

	fmt.Printf("len(trees)=%d\n", len(trees))
	no_parent_count := 0
	for {
		if (len(trees) == 1) {
			break;
		}

		to_merge := trees[0]
		trees = trees[1:]
		err := true
		for _, tree := range trees {
			if tree.insert(to_merge.head) {
				err = false
				break;
			}
		}
		if err {
			no_parent_count++
			trees = append(trees, to_merge)
			if no_parent_count > 1 {
				return nil
			}
		}
	}
	return trees[0]
}