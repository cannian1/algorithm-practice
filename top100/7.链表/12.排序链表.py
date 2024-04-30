# 148.排序链表 https://leetcode.cn/problems/sort-list/
from typing import Optional


class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


class Solution:
    """归并排序"""

    def sortList(self, head: Optional[ListNode]) -> Optional[ListNode]:
        return self.merge_sort(head)

    def merge_sort(self, head: Optional[ListNode]) -> Optional[ListNode]:
        """归并排序"""
        # 如果链表为空或只有一个节点，无需排序直接返回
        if not head or not head.next:
            return head
        # 获取链表的中间节点，分别对左右子链表进行排序
        mid = self.get_mid(head)
        right_sorted = self.merge_sort(mid.next)  # 排序右子链表
        mid.next = None  # 断开两段子链表
        left_sorted = self.merge_sort(head)  # 排序左子链表
        return self.merge_two_lists(left_sorted, right_sorted)  # 两个子链表必然有序，合并两个有序的链表

    def get_mid(self, head: Optional[ListNode]) -> Optional[ListNode]:
        """获取链表中间节点"""
        if not head:
            return head
        slow, fast = head, head.next
        while fast and fast.next:
            fast = fast.next.next
            slow = slow.next
        return slow

    def merge_two_lists(self, list1: Optional[ListNode], list2: Optional[ListNode]) -> Optional[ListNode]:
        """合并两个有序链表"""
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
