package main

import (
	"fmt"
	"math/rand"
	"time"
	"os"
)
type Node struct {
	Value int
}

type Tree struct{
	nodes []Node
}
type TEXT interface{ //使用接口 传入指针改变原数组
	Init()
	down(i, n int)
	up(j int)
	Pop() Node
	Push(node Node)
	Remove(node Node)
	print_tree()
}

var init_nums int = 10

// 用于构建结构体数组为最小堆，需要调用down函数
func (tree *Tree)Init(nodes []Node) {
	for _,node := range nodes{
		tree.Push(node)
	}
}

// 需要down（下沉）的元素在数组中的索引为i，n为heap的长度，将该元素下沉到该元素对应的子树合适的位置，从而满足该子树为最小堆的要求
func (tree *Tree)down(i, n int) {
	if i >= n || 2 * i >= n { //无左子节点
		return
	}
	switch{
	case 2 * i + 1 >= n: //无右子节点
		if tree.nodes[i].Value > tree.nodes[2 * i].Value {
			tree.nodes[i], tree.nodes[2 * i] = tree.nodes[2 * i], tree.nodes[i]
		}
	case tree.nodes[2 * i].Value > tree.nodes[2 * i + 1].Value: //右子节点更小
		if tree.nodes[i].Value > tree.nodes[2 * i + 1].Value {
			tree.nodes[i], tree.nodes[2 * i + 1] = tree.nodes[2 * i + 1], tree.nodes[i]
			tree.down( 2 * i + 1, n) //递归
		}
	case tree.nodes[2 * i].Value <= tree.nodes[2 * i + 1].Value: //左子节点更小
		if tree.nodes[i].Value > tree.nodes[2 * i].Value {
			tree.nodes[i], tree.nodes[2 * i] = tree.nodes[2 * i], tree.nodes[i]
			tree.down(2 * i, n) //递归
		}
	}
}

// 用于保证插入新元素(j为元素的索引，数组末尾插入，堆底插入)的结构体数组之后仍然是一个最小堆
func (tree *Tree)up(j int) {
	if j < 0 {
		fmt.Fprintf(os.Stderr, "\n[Error] Index should be no less than zero!\n")
		os.Exit(3)
	}
	if j == 0 || j / 2 == 0{
		return
	}
	if tree.nodes[j].Value < tree.nodes[j / 2].Value { //小于父节点
		tree.nodes[j], tree.nodes[j / 2] = tree.nodes[j / 2], tree.nodes[j]
		tree.up(j / 2)
	}
}

// 弹出最小元素，并保证弹出后的结构体数组仍然是一个最小堆
func (tree *Tree)Pop() Node {
	if len(tree.nodes) <= 1 {
		fmt.Fprintf(os.Stderr, "\n[Error] No element to be poped!\n")
		os.Exit(2)
	}
	res := tree.nodes[1]//调用pop需要检测是否有节点 即len(nodes) >= 2
	tree.Remove(res)
	return res
}

// 保证插入新元素时，结构体数组仍然是一个最小堆，需要调用up函数
func (tree *Tree)Push(node Node) {
	tree.nodes = append(tree.nodes, node)
	tree.up(len(tree.nodes) - 1)
}

// 移除数组中指定索引的元素，保证移除后结构体数组仍然是一个最小堆
func (tree *Tree)Remove(node Node) {
	idx, last := -1, -1
	idx_arr := make([]int, 0)
	for index, node_tmp := range tree.nodes{
		if node_tmp.Value == node.Value{
			idx = index
			idx_arr = append(idx_arr, idx)
			break
		}
	}
	if idx == -1 {
		fmt.Fprintf(os.Stderr, "\n [Error]node can not be found!\n")
		os.Exit(1)
	}
	for _, value := range idx_arr{
		last = len(tree.nodes) - 1
		tree.nodes[value], tree.nodes[last] = tree.nodes[last], tree.nodes[value] //特定元素交换到最后
		tree.nodes = tree.nodes[:last] //删除
		tree.down(value, len(tree.nodes)) //恢复堆
	}
}

func test(tree *Tree){
	fmt.Printf("--------test of pop--------\n")
	for i := 1; i <= init_nums / 2; i++{
		pop_node := tree.Pop()
		fmt.Printf("Number %d test of function Pop(), pop node's value is %d\n", i, pop_node.Value)
		//tree.print_tree()  
	}
	rdm := rand.New(rand.NewSource(time.Now().UnixNano()))
	fmt.Printf("--------test of push--------\n")
	for i := 1; i <= init_nums / 2; i++{
		push_num := rdm.Intn(100)
		tree.Push( Node{Value:push_num} )
		fmt.Printf("Number %d test of function push(), push num value is %d, current top node's value is %d\n", i, push_num, tree.nodes[1].Value)
		//tree.print_tree()
	}
	fmt.Printf("--------test of remove--------\n")
	for i := 1; i <= init_nums / 2 && len(tree.nodes) >= 3; i++{
		remove_idx := 1
		if remove_idx < 1 {
			fmt.Fprintf(os.Stderr, "remove_idx should no less than 1")
			os.Exit(5)
		}
		tree.Remove(tree.nodes[remove_idx])
		fmt.Printf("Number %d test of function Remove(), remove the num of %d node, current top of tree is %d\n", i, remove_idx, tree.nodes[1].Value) //第1个元素只是保证下标从1开始
	}
	//tree.Remove(Node{Value:-1})   //test error func
}

func (tree *Tree)print_tree(){
	for _,value := range tree.nodes[1:] {
		fmt.Println(value.Value)
	}
}

func main() {
	test_arr := make([]Node, 1)
	test_arr[0] = Node{Value : -1} //保证下标从1开始
	for rdom, i:= rand.New(rand.NewSource(time.Now().UnixNano())), 0; i < init_nums; i++ {
		test_arr = append(test_arr,Node{ Value: rdom.Intn(100)} )
	}
	var tree Tree
	tree.Init(test_arr)
	fmt.Printf("current nodes is:\n")
	tree.print_tree()
	test(&tree)
}
