package main
import(
	"fmt"
	. "hashtree"
)
func main(){
	/*
	root := make(map[int]*HashNode)
	root[0] = &HashNode{Key:0}
	root[0].AddValueRecursive([]int{2,3})
	root[0].AddValueRecursive([]int{2,4})
	root[0].AddValueRecursive([]int{2,4,13})
	root[0].AddValueRecursive([]int{})
	fmt.Println(root[0].GetValueRecursive([]int{2,4}))
	fmt.Println(root[0].GetValueRecursive([]int{2,4,13}))
	fmt.Println(root[0].GetValueRecursive([]int{2,3}))*/
	root:=&HashNode{}
	root.AddValueRecursive([]int{0,2,3})
	fmt.Println(root.GetValueRecursive([]int{0,2,3}))
}
