package grind75

// Leetcode 207. Course Schedule
// Grind 75 34

// canFinish There are a total of numCourses courses you have to take,
// labeled from 0 to numCourses - 1. You are given an array
// prerequisites where (prerequisites[i] = [ai, bi] indicates that you
// must take course bi first if you want to take course ai.
func canFinish(numCourses int, prerequisites [][]int) bool {
	// make a graph and graph and inDegree for topological sort
	graph := make(map[int][]int)
	inDegree := make(map[int]int)

	// prep of edges just doing x[0]=0 x[1]=0 ...
	for idx := 0; idx < numCourses; idx++ {
		inDegree[idx] = 0
	}

	// Filling Graph and incoming edges with data
	for _, prerequisite := range prerequisites {
		src, dst := prerequisite[1], prerequisite[0]
		graph[src] = append(graph[src], dst)
		inDegree[dst]++
	}

	// Topological sort of graph
	// 1. inDegree stor incomming edges
	var source []int
	for key := range inDegree {
		if inDegree[key] == 0 {
			source = append(source, key)
		}
	}

	// Storing courses that can be completed
	var courseSchedule []int
	for len(source) > 0 {
		vertex := source[0]
		source = source[1:]
		courseSchedule = append(courseSchedule, vertex)

		for _, neighbor := range graph[vertex] {
			inDegree[neighbor]--
			if inDegree[neighbor] == 0 {
				source = append(source, neighbor)
			}
		}
	}

	return len(courseSchedule) == numCourses
}

/* 3 ms solution on leetcode
/*

/*
func canFinish(numCourses int, prerequisites [][]int) bool {
	// map: course -> next courses
	next := make(map[int][]int)
	// indegree of nodes (courses)
	indegree := make([]int, numCourses)
	for _, p := range prerequisites {
		next[p[1]] = append(next[p[1]], p[0])
		indegree[p[0]]++
	}

	var count int
	var q []int
	// add all zero-indegree nodes to queue
	for c := 0; c < numCourses; c++ {
		if indegree[c] == 0 {
			q = append(q, c)
		}
	}

	for len(q) > 0 {
		// pop a node and add to output
		c := q[0]
		q = q[1:]
		count++
		// for each neighbor, decrease its indgree, and add to queue if indegree becomes 0
		for _, nei := range next[c] {
			indegree[nei]--
			if indegree[nei] == 0 {
				q = append(q, nei)
			}
		}
	}

	if count < numCourses {	// must have a cycle in this case
		return false
	}
	return true
}
*/
