package sol

import (
	"reflect"
	"testing"
)

func CreateList(nums *[]int) *ListNode {
	arr := *nums
	var head, cur *ListNode
	for idx, val := range arr {
		if idx == 0 {
			head = &ListNode{Val: val}
			cur = head
		} else {
			cur.Next = &ListNode{Val: val}
			cur = cur.Next
		}
	}
	return head
}
func Test_reorderList(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "[1,2,3,4]",
			args: args{head: CreateList(&[]int{1, 2, 3, 4})},
			want: CreateList(&[]int{1, 4, 2, 3}),
		},
		{
			name: "[1,2,3,4,5]",
			args: args{head: CreateList(&[]int{1, 2, 3, 4, 5})},
			want: CreateList(&[]int{1, 5, 2, 4, 3}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reorderList(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reorderList() = %v, want %v", got, tt.want)
			}
		})
	}
}
