/**
 * Definition for singly-linked list.
You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order, and each of their nodes contains a single digit. Add the two numbers and return the sum as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.
 */

 type ListNode struct {
      Val int
      Next *ListNode
 }

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
   
    if l1 == nil && l2 == nil {
        return nil
    }
    
    if l1 != nil && l2 == nil {
        return l1
    }
    
    if l2 != nil && l1 == nil {
        return l1
    }
    
    
    var first *ListNode = new(ListNode)
    res := first
    carry := 0
    
    i := l1
    j := l2
    
    for true {
        if i == nil && j == nil {
          res.Val = carry
          carry = 0
        } else if i != nil && j == nil {
             res.Val = (i.Val + carry)%10
            carry = (i.Val +carry)/10
            i = i.Next
        } else if j != nil && i == nil {
             res.Val = (j.Val + carry)%10
            carry = (j.Val +carry)/10
            j = j.Next
        } else {
            res.Val = (i.Val+j.Val+carry)%10
            carry = (i.Val+j.Val+carry)/10
             i = i.Next
             j = j.Next
        }
        
        if i == nil && j == nil && carry == 0 {
            break
        }
        
        res.Next = new(ListNode)
        res = res.Next
        
    }

    return first
    
}
