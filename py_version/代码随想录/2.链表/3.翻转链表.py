# 206.反转链表
# https://leetcode.cn/problems/reverse-linked-list/

from typing import Optional


class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


class Solution:
    """ 双指针法 """

    """
    None    1->2->3->4->None
    pre     cur
    
    None    1->2->3->4->None
    pre     cur t
    
    None <- 1   2->3->4->None
    pre  <-cur  t
    
    None <- 1   2->3->4->None
           pre  cur
    """



    def reverseList(self, head: Optional[ListNode]) -> Optional[ListNode]:
        cur = head
        pre = None
        while cur:
            t = cur.next  # 保存 cur 的下一个节点，因为接下来要改变 cur.next
            cur.next = pre  # 反转
            # 更新 pre、cur 指针
            pre = cur
            cur = t
        return pre


class Solution2:
    """头插法实现反转链表"""
    """ 
    X->None 
    1->2->3->4->None
    cur
    
    X->None 
    1->2->3->4->None
    cur t
    
    X->None 
    1 ↗  2->3->4->None
    cur  t
    
    X->1->None 
    2->3->4->None
    cur
    
    X->1->None
    2->3->4->None
    """

    def reverseList(self, head: Optional[ListNode]) -> Optional[ListNode]:
        # 创建虚拟头节点
        dummy = ListNode(next=None)
        # 遍历所有节点
        cur = head  # 指向链表头
        while cur:
            t = cur.next  # 保存 cur 下一个元素的位置
            # 头插法
            cur.next = dummy.next  # cur 指向虚拟头节点后的元素
            dummy.next = cur  # 虚拟头节点指向 cur
            cur = t  # 将之前暂存的 cur 下一个元素的位置交给 cur，以便下一次循环头插到虚拟头节点后

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
