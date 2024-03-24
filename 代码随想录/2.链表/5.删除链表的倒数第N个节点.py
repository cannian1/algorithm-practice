# 19. 删除链表的倒数第 N 个结点
# https://leetcode.cn/problems/remove-nth-node-from-end-of-list/

from typing import Optional


class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


class Solution:
    def removeNthFromEnd(self, head: Optional[ListNode], n: int) -> Optional[ListNode]:
        dummy_head = ListNode(next=head)
        pre = dummy_head
        cur = dummy_head

        # 快慢指针，快指针比慢指针先走 n 步
        for i in range(n):
            cur = cur.next

        # 直到快指针下一步到链表尾部
        while cur.next:
            pre = pre.next
            cur = cur.next

        # 操作慢指针删除元素
        pre.next = pre.next.next

        return dummy_head.next
