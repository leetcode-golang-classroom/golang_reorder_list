package sol

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	arr := []*ListNode{}
	cur := head
	for cur != nil {
		arr = append(arr, cur)
		cur = cur.Next
	}
	aLen := len(arr)
	max_idx := aLen / 2
	for idx := 0; idx < max_idx; idx++ {
		arr[idx].Next = arr[aLen-idx-1]
		arr[aLen-idx-1].Next = arr[idx+1]
	}
	arr[max_idx].Next = nil
	return head
}
