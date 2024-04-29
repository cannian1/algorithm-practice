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


class ListNode2:
    """官解"""
    """翻转一个子链表，并且返回新的头与尾"""

    def reverse(self, head: ListNode, tail: ListNode):
        prev = tail.next
        p = head
        while prev != tail:
            nex = p.next
            p.next = prev
            prev = p
            p = nex
        return tail, head

    def reverseKGroup(self, head: ListNode, k: int) -> ListNode:
        hair = ListNode(0)
        hair.next = head
        pre = hair

        while head:
            tail = pre
            # 查看剩余部分长度是否大于等于 k
            for i in range(k):
                tail = tail.next
                if not tail:
                    return hair.next
            nex = tail.next
            head, tail = self.reverse(head, tail)
            # 把子链表重新接回原链表
            pre.next = head
            tail.next = nex
            pre = tail
            head = tail.next

        return hair.next


head = ListNode(val=1,
                next=ListNode(val=2, next=ListNode(val=3, next=ListNode(val=4, next=ListNode(val=5, next=ListNode())))))
s = Solution()
s.reverseKGroup(head, 2)
res = []
while head:
    res.append(head.val)
    head = head.next
print(res)
