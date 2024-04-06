# 232. 用栈实现队列
# https://leetcode.cn/problems/implement-queue-using-stacks/


class MyQueue:

    def __init__(self):
        """
        in 主要负责 push，out 主要负责 pop
        """
        self.stack_in = []
        self.stack_out = []

    def push(self, x: int) -> None:
        """有新元素进来，就往 in 里面 push"""
        self.stack_in.append(x)

    def pop(self) -> int:
        """从队列前面移除一个元素并返回"""
        if self.empty():
            return None
        else:
            for i in range(len(self.stack_in)):
                # in 列表最后一个元素加入 out列表
                self.stack_out.append(self.stack_in.pop())
            return self.stack_out.pop()

    def peek(self) -> int:
        """返回队列开头的元素"""
        ans = self.pop()  # in 列表全部逆序导出到 out 列表，然后弹出元素
        self.stack_in.append(ans) # 再将弹出的元素放入 in 列表
        return ans

    def empty(self) -> bool:
        # 不在两个列表里
        return not (self.stack_in or self.stack_out)

# Your MyQueue object will be instantiated and called as such:
# obj = MyQueue()
# obj.push(x)
# param_2 = obj.pop()
# param_3 = obj.peek()
# param_4 = obj.empty()
