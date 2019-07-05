package main

var infinite = 999999999999

func makeNodeMap(m maze) [][]int {
	nm := make([][]int, m.Width)
	for x := 0; x < m.Width; x++ {
		nm[x] = make([]int, m.Height)
		for y := 0; y < m.Height; y++ {
			nm[x][y] = infinite
		}
	}
	return nm
}

func constructPath(cameFrom map[string]node, goal node) string {
	path := ""
	current := goal
	for prev, ok := cameFrom[current.toString()]; ok; prev, ok = cameFrom[current.toString()] {
		if current[0] == prev[0]-1 {
			path = "W" + path
		} else if current[0] == prev[0]+1 {
			path = "E" + path
		} else if current[1] == prev[1]-1 {
			path = "N" + path
		} else if current[1] == prev[1]+1 {
			path = "S" + path
		}
		current = prev
	}
	return path
}

func hScore(a node, b node) int {
	dx := a[0] - b[0]
	if dx < 0 {
		dx = -dx
	}
	dy := a[1] - b[1]
	if dy < 0 {
		dy = -dy
	}
	return dx + dy
}

func astar(m maze, start node, goal node) string {
	closedSet := map[string]node{}
	openSet := map[string]node{start.toString(): start}
	cameFrom := map[string]node{}
	gScore := makeNodeMap(m)
	gScore[start[0]][start[1]] = 0
	fScore := makeNodeMap(m)
	fScore[start[0]][start[1]] = hScore(start, goal)

	for len(openSet) > 0 {
		lowest := infinite
		current := node{-1, -1}
		for _, n := range openSet {
			if fScore[n[0]][n[1]] < lowest {
				current = n
				lowest = fScore[n[0]][n[1]]
			}
		}

		if current.toString() == goal.toString() {
			return constructPath(cameFrom, goal)
		}

		delete(openSet, current.toString())
		closedSet[current.toString()] = current

		for _, n := range m.neighbors(current) {
			_, inClosedSet := closedSet[n.toString()]
			if !inClosedSet {
				score := gScore[current[0]][current[1]] + 1
				openSet[n.toString()] = n
				if score < gScore[n[0]][n[1]] {
					// improved path
					cameFrom[n.toString()] = current
					gScore[n[0]][n[1]] = score
					fScore[n[0]][n[1]] = score + hScore(n, goal)
				}
			}
		}
	}

	return ""
}
