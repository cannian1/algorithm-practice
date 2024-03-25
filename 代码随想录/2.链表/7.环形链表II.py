# 142. 环形链表 II
# https://leetcode.cn/problems/linked-list-cycle-ii/

# Definition for singly-linked list.
class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None


class Solution:
    """快慢指针法"""

    # https://www.programmercarl.com/0142.%E7%8E%AF%E5%BD%A2%E9%93%BE%E8%A1%A8II.html
    def detectCycle(self, head: ListNode) -> ListNode:
        slow = head
        fast = head

        while fast and fast.next:
            slow = slow.next
            fast = fast.next.next

            # 如果有环，快慢指针终会相遇
            if slow == fast:
                # 将一个指针移回链表头
                slow = head
                while slow != fast:
                    slow = slow.next
                    fast = fast.next
                return slow
        return None


class Solution2:
    """集合法"""

    def detectCycle(self, head: ListNode) -> ListNode:
        visited = set()

        while head:
            if head in visited:
                return head
            visited.add(head)
            head = head.next

        return head
