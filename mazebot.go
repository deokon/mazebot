package main

import "fmt"

func solve(m maze) string {
	return astar(m, m.StartingPosition, m.EndingPosition)
}

func main() {
	theMaze := getMaze("https://api.noopschallenge.com/mazebot/random")
	sol := solve(theMaze)
	theMaze.display(sol)
	// fmt.Printf("Solution: %v\n", sol)

	var result solutionResponse
	jsonPost("https://api.noopschallenge.com"+theMaze.MazePath, jsonType{"directions": sol}, &result)
	fmt.Println(result)
}
