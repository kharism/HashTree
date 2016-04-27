package hashtree
import(
	"errors"
	//"fmt"
	"strconv"
)

type HashNode struct{
	Key int
	Val int
	Childs map[int]*HashNode
}
func (c *HashNode) AddValueRecursive(key []int){
	c.Val+=1
	if len(key)==0{
		return
	}
	childNode,ok:=c.Childs[key[0]]
	if !ok{
		//fmt.Println("Node",c.Key,"adding",key)
		childNode = &HashNode{Key:key[0]}
		if c.Childs==nil{
			c.Childs = make(map[int]*HashNode)
		}
		
		c.Childs[key[0]] = childNode
		//fmt.Println("Adding ",len(c.Childs))
	}
	if len(key)>1{
		childNode.AddValueRecursive(key[1:])
	} else if len(key)==1{
		//fmt.Println("Node",c.Key,"Added",c.Val)
		childNode.Val+=1
	}
}
func (c *HashNode) GetValueRecursive(key []int)(int,error){
	if len(key)==0{
		return c.Val,nil
	}
	//fmt.Println("Node",c.Key,"Searched")
	childNode,ok:=c.Childs[key[0]]
	if ok{
		return childNode.GetValueRecursive(key[1:])
	} else{
		return 0,errors.New("No Key Error for Node"+strconv.Itoa(c.Key))
	}
}
