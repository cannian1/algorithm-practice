# 206.反转链表
# https://leetcode.cn/problems/reverse-linked-list/

from typing import Optional


class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


class Solution:
    def reverseList(self, head: Optional[ListNode]) -> Optional[ListNode]:
        cur = head
        pre = None
        while cur:
            t = cur.next  # 保存 cur 的下一个节点，因为接下来要改变 cur.next
            cur.next = pre  # 反转
            # 更新 pre、cur 指针
            pre = cur
            cur = t
        return t


class Solution2:
    """头插法实现反转链表"""
    """ 
    X->None 
    1->2->3->4->None
    
    X->1->None
    2->3->4->None
    """
    def reverseList(self, head: Optional[ListNode]) -> Optional[ListNode]:
        # 创建虚拟头节点
        dummy = ListNode(next=None)
        # 遍历所有节点
        cur = head
        while cur:
            t = cur.next
            # 头插法
            cur.next = dummy.next
            dummy.next = cur
            cur = t

        return dummy.next


class Solution3:
    """递归法"""

    def reverseList(self, head: Optional[ListNode]) -> Optional[ListNode]:
        return self.reverse(head, None)

    def reverse(self, cur: ListNode, pre: Optional[ListNode]) -> Optional[ListNode]:
        if cur is None:
            return pre
        temp = cur.next
        cur.next = pre
        return self.reverse(temp, pre)
