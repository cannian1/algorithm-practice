# 25. K 个一组翻转链表 https://leetcode.cn/problems/reverse-nodes-in-k-group/
from typing import Optional


# 给你链表的头节点 head ，每 k 个节点一组进行翻转，请你返回修改后的链表。

class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


class Solution:
    def reverseKGroup(self, head: Optional[ListNode], k: int) -> Optional[ListNode]:
        n = 0
        cur = head
        #  统计链表长度
        while cur:
            n += 1
            cur = cur.next

        p0 = dummy = ListNode(next=head)
        pre = None
        cur = head

        while n >= k:
            n -= k
            # 每k个元素反转
            for _ in range(k):
                temp = cur.next
                cur.next = pre
                pre = cur
                cur = temp

            nxt = p0.next
            p0.next.next = cur
            p0.next = pre
            p0 = nxt
        return dummy.next
