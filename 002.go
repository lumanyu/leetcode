//import "container/list

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func convertLink (l *ListNode) *ListNode{
    l_oper := l
    var tmp *ListNode = nil
    l_first := l
    l_second := l.Next
    for l_first != nil && l_second != nil { //must save two pointer, move at the same time
        tmp = l_second.Next
        l_second.Next = l_first
        
        l_first = l_second
        l_second = tmp
    }
    l.Next = nil
    return l_oper
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    //fmt.Printf("%v", l1)
    //l1 = convertLink(l1)
    //fmt.Printf("%v", l1)
    //l2 = convertLink(l2)
    
    add := 0
    var head *ListNode = nil
    var last *ListNode = nil
    for l1 != nil || l2 !=nil || add != 0 {
        n := new(ListNode)
        if l1 == nil && l2 ==nil {
            n.Val = 1
            add = 0
        } else if l1 == nil {
            n.Val = (l2.Val + add) % 10
            if l2.Val + add >= 10 {
                add = 1
            } else {
                add = 0
            }
            l2 = l2.Next
        } else if l2 ==nil {
            n.Val = (l1.Val + add) % 10
            if l1.Val + add >= 10 {
                add = 1
            } else {
                add = 0
            }
            l1 = l1.Next
        } else {
            n.Val = (l1.Val + l2.Val + add) % 10
            if l1.Val + l2.Val + add >= 10 {
                add = 1
            } else {
                add = 0
            }
            l1 = l1.Next
            l2 = l2.Next
        }
        
        if head != nil {
            last.Next = n
            last = n
        } else {
            head = n
            last = n
        }
    }
    
    return head
    //return convertLink(head)
    
}
