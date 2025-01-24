// 207.课程表
// https://leetcode.cn/problems/course-schedule/

package graph

// 你这个学期必须选修 numCourses 门课程，记为 0 到 numCourses - 1 。
//
// 在选修某些课程之前需要一些先修课程。 先修课程按数组 prerequisites 给出，其中 prerequisites[i] = [ai, bi] ，表示如果要学习课程 ai 则 必须 先学习课程  bi 。
//
// 例如，先修课程对 [0, 1] 表示：想要学习课程 0 ，你需要先完成课程 1 。
// 请你判断是否可能完成所有课程的学习？如果可以，返回 true ；否则，返回 false 。

// 三色标记法
func canFinish(numCourses int, prerequisites [][]int) bool {
	// 创建邻接表
	// 索引代表课程编号（如索引 i 对应课程 i）
	// adjList[i] 存储所有依赖课程 i 的后续课程（即学完课程 i 才能学习的课程）
	adjList := make([][]int, numCourses)
	for _, prereq := range prerequisites {
		// 对每个先修条件 prereq = [a, b]（表示学课程 a 前需先学 b）：
		// pre = b（先修课程），course = a（目标课程）。
		// 将 a 添加到 adjList[b] 的列表中，表示 b 是 a 的先修条件，构建有向边 b → a。
		course, pre := prereq[0], prereq[1]
		adjList[pre] = append(adjList[pre], course)
	}

	// 颜色数组：0=白色，1=灰色，2=黑色
	color := make([]int, numCourses)

	// 深度优先搜索函数
	var dfs func(int) bool
	dfs = func(node int) bool {
		if color[node] == 1 {
			// 如果当前节点是灰色，说明存在环
			return false
		}
		if color[node] == 2 {
			// 如果当前节点是黑色，说明已访问，无需重复
			return true
		}

		// 将当前节点标记为灰色
		color[node] = 1
		// 遍历相邻节点
		for _, neighbor := range adjList[node] {
			if !dfs(neighbor) {
				return false
			}
		}
		// 将当前节点标记为黑色
		color[node] = 2
		return true
	}

	// 遍历所有课程，尝试进行 DFS
	for i := 0; i < numCourses; i++ {
		if color[i] == 0 {
			if !dfs(i) {
				return false
			}
		}
	}
	return true
}

// 拓扑排序法
func canFinish2(numCourses int, prerequisites [][]int) bool {
	// 创建一个邻接表和一个入度数组
	adjList := make([][]int, numCourses)
	inDegree := make([]int, numCourses)

	// 构建图并填充入度数组
	for _, prereq := range prerequisites {
		course, pre := prereq[0], prereq[1]
		adjList[pre] = append(adjList[pre], course)
		inDegree[course]++
	}

	// 将所有入度为 0 的课程加入队列
	queue := []int{}
	for i := 0; i < numCourses; i++ {
		if inDegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	// 执行拓扑排序
	takenCourses := 0
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		takenCourses++

		// 遍历当前课程的后续课程
		for _, next := range adjList[current] {
			inDegree[next]--
			if inDegree[next] == 0 {
				queue = append(queue, next)
			}
		}
	}

	// 如果已修课程等于总课程数，返回 true；否则返回 false
	return takenCourses == numCourses
}
