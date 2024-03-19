# LeetCode 203. 移除链表元素
# https://leetcode.cn/problems/remove-linked-list-elements

from typing import Optional


# Definition for singly-linked list.
class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


class Solution:
    def removeElements(self, head: Optional[ListNode], val: int) -> Optional[ListNode]:
        # 虚拟头节点指向链表头节点
        dummy = ListNode(next=head)
        cur = dummy
        while cur.next:
            if cur.next.val == val:
                cur.next = cur.next.next
            else:
                cur = cur.next

        return dummy.next


list_node0 = ListNode(val=7)
list_node1 = ListNode(val=7)
list_node2 = ListNode(val=6)
list_node3 = ListNode(val=7)

list_node0.next = list_node1
list_node1.next = list_node2
list_node2.next = list_node3

sol = Solution()
result_head = sol.removeElements(list_node0, 7)
result = []
while result_head:
    result.append(result_head.val)
    result_head = result_head.next
print(result)
