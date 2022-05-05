# golang_reorder_list

You are given the head of a singly linked-list. The list can be represented as:

```
L0 → L1 → … → Ln - 1 → Ln
```

*Reorder the list to be on the following form:*

```
L0 → Ln → L1 → Ln - 1 → L2 → Ln - 2 → …

```

You may not modify the values in the list's nodes. Only nodes themselves may be changed.

## Examples

**Example 1:**

![https://assets.leetcode.com/uploads/2021/03/04/reorder1linked-list.jpg](https://assets.leetcode.com/uploads/2021/03/04/reorder1linked-list.jpg)

```
Input: head = [1,2,3,4]
Output: [1,4,2,3]

```

**Example 2:**

![https://assets.leetcode.com/uploads/2021/03/09/reorder2-linked-list.jpg](https://assets.leetcode.com/uploads/2021/03/09/reorder2-linked-list.jpg)

```
Input: head = [1,2,3,4,5]
Output: [1,5,2,4,3]

```

**Constraints:**

- The number of nodes in the list is in the range `[1,$5*10^4$]`.
- `1 <= Node.val <= 1000`

## 解析

題目要把原本的單向鏈結傳列照著一個新順序的重新接合

其中接合的順序仔細察看其位置關係

每次接合的點 每兩一個數其 index 相加剛好是 len(list) -1

而接在第二個結點的下一個結點 index 恰巧是第一個結點 index +1

所以總共會有幾個結點組 會是 len(list)/2 組

而最後一個解點 index 是 len(list)/2

這個結點不需要接點只要把 Next 設定為 nil

## 程式碼

```go
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode)  *ListNode{
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
    max_idx := aLen/2
    for idx := 0 ; idx < max_idx; idx++ {
        arr[idx].Next = arr[aLen - idx -1]
        arr[aLen - idx -1].Next = arr[idx+1]
    }
    arr[max_idx].Next = nil
    return head
}
```

## 困難點

1. 需要找出重新排列的關係式
2. 每次重新排序的接點方式需要去了解

## Solve Point

- [x]  Understand what problem would like to solve
- [x]  Analysis Complexity