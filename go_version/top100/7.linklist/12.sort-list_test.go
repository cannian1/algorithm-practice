package linklist

import (
	"reflect"
	"testing"
)

func Test_sortList(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{"空链表", []int{}, []int{}},
		{"单节点", []int{1}, []int{1}},
		{"已排序", []int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{"未排序", []int{4, 2, 1, 3}, []int{1, 2, 3, 4}},
		{"带有重复项的链表", []int{4, 2, 2, 3, 3}, []int{2, 2, 3, 3, 4}},
		{"带有负数的链表", []int{-1, 5, 3, 4, 0}, []int{-1, 0, 3, 4, 5}},
		{"两个节点", []int{2, 1}, []int{1, 2}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 将输入切片转为链表
			inputList := SliceToList(tt.input)

			// 调用被测试函数
			sortedList := sortList(inputList)

			// 将输出链表转为切片
			output := ListToSlice(sortedList)

			// 比较结果
			if !reflect.DeepEqual(output, tt.expected) {
				t.Errorf("expected %v, got %v", tt.expected, output)
			}
		})
	}
}

func Test_getMid(t *testing.T) {

	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{"空链表", []int{}, 0}, // 返回 0 表示无效节点
		{"单节点", []int{1}, 1},
		{"偶数个节点", []int{1, 2, 3, 4}, 2},
		{"奇数个节点", []int{1, 2, 3, 4, 5}, 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := SliceToList(tt.input)
			mid := getMid(head)
			if (mid == nil && tt.expected != 0) || (mid != nil && mid.Val != tt.expected) {
				t.Errorf("expected %d, got %d", tt.expected, func() int {
					if mid == nil {
						return 0
					}
					return mid.Val
				}())
			}
		})
	}
}

func Test_mergeTwoList(t *testing.T) {
	tests := []struct {
		name     string
		list1    []int
		list2    []int
		expected []int
	}{
		{"两个链表都为空", []int{}, []int{}, []int{}},
		{"第一个链表为空", []int{}, []int{1, 2, 3}, []int{1, 2, 3}},
		{"第二个链表为空", []int{1, 2, 3}, []int{}, []int{1, 2, 3}},
		{"不重叠", []int{1, 2, 3}, []int{4, 5, 6}, []int{1, 2, 3, 4, 5, 6}},
		{"重叠", []int{1, 3, 5}, []int{2, 4, 6}, []int{1, 2, 3, 4, 5, 6}},
		{"值重复", []int{1, 2, 2}, []int{2, 2, 3}, []int{1, 2, 2, 2, 2, 3}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list1 := SliceToList(tt.list1)
			list2 := SliceToList(tt.list2)
			merged := mergeTwoList(list1, list2)
			output := ListToSlice(merged)
			if !reflect.DeepEqual(output, tt.expected) {
				t.Errorf("expected %v, got %v", tt.expected, output)
			}
		})
	}
}
