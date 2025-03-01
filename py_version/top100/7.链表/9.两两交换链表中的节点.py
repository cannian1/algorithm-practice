# 24. 两两交换链表中的节点 https://leetcode.cn/problems/swap-nodes-in-pairs/
from typing import Optional


class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


class Solution:
    """
    cur -> 1 -> 2 -> 3 -> 4
    cur -> 2
    cur -> 2 -> 1
    cur -> 2 -> 1 -> 3 -> 4
    2 -> 1 -> 3 -> 4
        cur ↗
    """

    def swapPairs(self, head: Optional[ListNode]) -> Optional[ListNode]:
        dummy_head = ListNode(next=head)
        cur = dummy_head

        # 必须有cur的下一个和下下个才能交换，否则说明已经交换结束
        while cur.next and cur.next.next:
            temp = cur.next  # 暂存节点保证能被指向
            temp1 = cur.next.next.next  # 暂存 1,3位置

            cur.next = cur.next.next
            cur.next.next = temp
            temp.next = temp1
            cur = cur.next.next
        return dummy_head.next
