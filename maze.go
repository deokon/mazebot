package main

import "fmt"

type node []int

func (n node) toString() string {
	return fmt.Sprintf("%d,%d", n[0], n[1])
}

type maze struct {
	Name             string
	MazePath         string
	StartingPosition node
	EndingPosition   node
	Message          string
	Map              [][]string
	Width            int
	Height           int
}

type solutionResponse struct {
	Result                 string
	Message                string
	ShortestSolutionLength string
	YourSolutionlength     string
	Elapsed                int
	NextMaze               string
}

var mazeCharacters = map[string]rune{
	"X:    ": '\u25EF',
	"X:X   ": '\u2575',
	"X: X  ": '\u2576',
	"X:  X ": '\u2577',
	"X:   X": '\u2574',
	"X:XX  ": '\u2514',
	"X:X X ": '\u2502',
	"X:X  X": '\u2518',
	"X: XX ": '\u250C',
	"X: X X": '\u2500',
	"X:  XX": '\u2510',
	"X:XXX ": '\u251C',
	"X:XX X": '\u2534',
	"X:X XX": '\u2524',
	"X: XXX": '\u252C',
	"X:XXXX": '\u253C',
}

func (theMaze maze) isWall(x, y int) bool {
	if x < -1 || y < -1 || x > theMaze.Width || y > theMaze.Width {
		return false
	}
	if x == -1 || y == -1 || x == theMaze.Width || y == theMaze.Width {
		return true
	}
	return theMaze.Map[y][x] == "X"
}

func (theMaze maze) neighbors(n node) (ne []node) {
	if !theMaze.isWall(n[0]-1, n[1]) {
		ne = append(ne, node{n[0] - 1, n[1]})
	}
	if !theMaze.isWall(n[0], n[1]-1) {
		ne = append(ne, node{n[0], n[1] - 1})
	}
	if !theMaze.isWall(n[0]+1, n[1]) {
		ne = append(ne, node{n[0] + 1, n[1]})
	}
	if !theMaze.isWall(n[0], n[1]+1) {
		ne = append(ne, node{n[0], n[1] + 1})
	}
	return
}

func (theMaze maze) wallShape(x, y int) string {
	shape := ""
	if theMaze.isWall(x, y) {
		shape += "X:"
	} else {
		shape += " :"
	}
	if theMaze.isWall(x, y-1) {
		shape += "X"
	} else {
		shape += " "
	}
	if theMaze.isWall(x+1, y) {
		shape += "X"
	} else {
		shape += " "
	}
	if theMaze.isWall(x, y+1) {
		shape += "X"
	} else {
		shape += " "
	}
	if theMaze.isWall(x-1, y) {
		shape += "X"
	} else {
		shape += " "
	}
	return shape
}

func (theMaze maze) display() {
	fmt.Printf("---------------------------\nName: %v\nPath: %v\n%v\n---------------------------\n", theMaze.Name, theMaze.MazePath, theMaze.Message)

	for y := -1; y <= theMaze.Height; y++ {
		for x := -1; x <= theMaze.Width; x++ {
			c, ok := mazeCharacters[theMaze.wallShape(x, y)]
			if ok {
				fmt.Printf("%c", c)
			} else {
				fmt.Printf(" ")
			}
		}
		fmt.Println()
	}
}

func getMaze(url string) (theMaze maze) {
	// get a random maze and set the width and height
	jsonGet(url, &theMaze)
	theMaze.Height = len(theMaze.Map)
	theMaze.Width = len(theMaze.Map[0])

	return
}
