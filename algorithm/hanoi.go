package main

import (
	"fmt"
	"container/list"
)
type BinaryTree struct{
	Data [3]string
	Left *BinaryTree
	Right *BinaryTree
}
func NewBinaryTree(data [3]string) *BinaryTree {
   return &BinaryTree{Data: data}
}
func right (data [3]string)[3]string{
	meta :=data[1]
	data[1]=data[0]
	data[0]=meta
	return data
}
func left (data [3]string)[3]string{
	meta :=data[1]
	data[1]=data[2]
	data[2]=meta
	return data
}
// 中序遍历-非递归
func (bt *BinaryTree) InOrderNoRecursion() []interface{} {
   t := bt
   stack := list.New()
   res := make([] interface{}, 0)
   for t != nil || stack.Len() != 0 {
      for t != nil {
         stack.PushBack(t)
         t = t.Left
      }
      if stack.Len() != 0 {
         v := stack.Back()
         t = v.Value.(*BinaryTree)
         res = append(res, t.Data)//visit
         t = t.Right
         stack.Remove(v)
      }
   }
   return res
}
//汉诺塔用二叉树的中序遍历的解法，先构建二叉树，然后用中序遍历
func (bt *BinaryTree) hanoi(num int){
	sli :=[]*BinaryTree{}
	stack :=list.New()
	stack.PushBack(bt)
	for i:=0;i<num;i++{
		for stack.Len()!=0{
			v := stack.Back()
         	 	t := v.Value.(*BinaryTree)
	        	data := t.Data
			stack.Remove(v)
			fmt.Println(data)
			leftdata := left(data)
			rightdata :=right(data)
			t.Left = NewBinaryTree(leftdata)
			t.Right=NewBinaryTree(rightdata)
			leftB := t.Left
			rightB := t.Right
			sli = append(sli,leftB)
			sli = append(sli,rightB)

		}
		for _,v := range sli{
			stack.PushBack(v)
		}
		sli=[]*BinaryTree{}
	}

}

func main() {
 	data :=[3]string{"A","B","C"}
	bt :=NewBinaryTree(data)
	num :=15
	bt.hanoi(num)

	res:=bt.InOrderNoRecursion()
	fmt.Println(res)
}
