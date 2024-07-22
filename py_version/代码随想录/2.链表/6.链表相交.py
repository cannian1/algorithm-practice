# 面试题 02.07. 链表相交
# https://leetcode.cn/problems/intersection-of-two-linked-lists-lcci/


# Definition for singly-linked list.
class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


class Solution:
    def getIntersectionNode(self, headA: ListNode, headB: ListNode) -> ListNode:
        lenA, lenB = 0, 0

        # 求链表A的长度
        cur = headA
        while cur:
            cur = cur.next
            lenA += 1

        # 求链表B的长度
        cur = lenB
        while cur:
            cur = cur.next
            lenB += 1

        # 重新指向头节点
        curA, curB = headA, headB
        # 让 curB 为最长链表的头，lenB为其长度
        if lenA > lenB:
            curA, curB = curB, curA
            lenA, lenB = lenB, lenA

        # 让 curA 与 curB 在同一起点（末尾元素对齐）
        for _ in range(lenB - lenA):
            curB = curB.next

        #  遍历curA 和 curB，遇到相同则直接返回
        while curA:
            if curA == curB:
                return curA
            else:
                curA = curA.next
                curB = curB.next
        return None


class Solution2:
    """代码复用+精简"""

    def getIntersectionNode(self, headA: ListNode, headB: ListNode) -> ListNode:
        dis = self.getLength(headA) - self.getLength(headB)

        if dis > 0:
            headA = self.moveForword(headA, dis)
        else:
            headB = self.moveForword(headB, -dis)

        # 将两个头向前移动，直到它们相交
        while headA and headB:
            if headA == headB:
                return headA
            headA = headA.next
            headB = headB.next

        return None

    def getLength(self, head: ListNode) -> int:
        length = 0
        while head:
            length += 1
            head = head.next
        return length

    def moveForword(self, head: ListNode, steps: int) -> ListNode:
        while steps > 0:
            head = head.next
            steps -= 1
        return head
