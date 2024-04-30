# 23.合并K个升序链表 https://leetcode.cn/problems/merge-k-sorted-lists/
from typing import Optional, List
import heapq


class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


# __lt__定义了对象的“小于”比较操作。当使用比较操作符<时，会调用这个方法。
ListNode.__lt__ = lambda a, b: a.val < b.val


class Solution:
    """最小堆 O(nlogk) k为 lists的长度，n为所有链表的节点数之和 O(k) 堆中至多有 k 个元素"""

    def mergeKLists(self, lists: List[Optional[ListNode]]) -> Optional[ListNode]:
        cur = dummy = ListNode()
        # h = [head for head in lists if head]  # 初始把所有链表的头节点入堆
        h = []
        for head in lists:
            if head:
                h.append(head)
        heapq.heapify(h)  # 堆化
        while h:
            node = heapq.heappop(h)  # 剩余节点中的最小节点(本身已经是升序链表，入堆的时候就存在当前最小元素)
            if node.next:  # 下一个节点不为空
                heapq.heappush(h, node.next)  # 下一个节点有可能是最小节点，入堆
            cur.next = node  # 合并到新链表中
            cur = cur.next
        return dummy.next


class Solution2:
    """分而治之"""

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

    def mergeKLists(self, lists: List[Optional[ListNode]]) -> Optional[ListNode]:
        lists_length = len(lists)
        if lists_length == 0:
            return None
        if lists_length == 1:
            return lists[0]
        # 类似归并
        left = self.mergeKLists(lists[:lists_length // 2])  # 合并左半部分
        right = self.mergeKLists(lists[lists_length // 2:])  # 合并右半部分
        return self.merge_two_lists(left, right)  # 最后把左半部分和右半部分合并
