# 21. 合并两个有序链表 https://leetcode.cn/problems/merge-two-sorted-lists/
from typing import Optional


class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


class Solution:
    """创建一个虚拟头节点，循环将两个链表中较小的加入它后面，最后把剩余部分合并"""
    def mergeTwoLists(self, list1: Optional[ListNode], list2: Optional[ListNode]) -> Optional[ListNode]:
        dummy = cur = ListNode()

        while list1 and list2:
            if list1.val < list2.val:
                cur.next = list1
                list1 = list1.next
            else:
                cur.next = list2
                list2 = list2.next
            cur = cur.next
        # 将剩余尾部合并
        if list1:
            cur.next = list1
        else:
            cur.next = list2
        return dummy.next
