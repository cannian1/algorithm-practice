# 225. 用队列实现栈
# https://leetcode.cn/problems/implement-stack-using-queues/
from collections import deque


class MyStack:

    def __init__(self):
        self.que = deque()

    def push(self, x: int) -> None:
        self.que.append(x)

    def pop(self) -> int:
        if self.empty():
            return None
        # 把队列里的 size - 1 个元素出队再入队，剩下的那个元素就是栈该弹出的元素了
        for i in range(len(self.que) - 1):
            self.que.append(self.que.popleft())
        return self.que.popleft()

    def top(self) -> int:
        # 写法一：
        # if self.empty():
        #     return None
        # return self.que[-1]

        # 写法二：
        if self.empty():
            return None
        for i in range(len(self.que) - 1):
            self.que.append(self.que.popleft())
        temp = self.que.popleft()
        self.que.append(temp)
        return temp

    def empty(self) -> bool:
        return not self.que
