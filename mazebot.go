package main

import (
	"fmt"
	"log"
	"os"
)

func solve(m maze) string {
	return astar(m, m.StartingPosition, m.EndingPosition)
}

func main() {
	showHelp := false
	args := os.Args[1:]

	if len(args) == 0 || args[0] == "help" {
		showHelp = true
	} else if args[0] == "rand" {
		// get a random maze and solve it
		theMaze := getMaze("https://api.noopschallenge.com/mazebot/random")
		sol := solve(theMaze)
		theMaze.display(sol)

		// post the solution
		var result mazeResponse
		jsonPost("https://api.noopschallenge.com"+theMaze.MazePath, jsonType{"directions": sol}, &result)

		fmt.Printf("\n%v\n\n", result.Message)
	} else if args[0] == "race" {
		if len(args) != 2 {
			fmt.Printf("ERROR: Missing login or too many parameters\n\n")
			showHelp = true
		} else {
			var result mazeResponse
			// start the race
			respCode := jsonPost("https://api.noopschallenge.com/mazebot/race/start", jsonType{"login": args[1]}, &result)
			if respCode != 200 {
				log.Fatalln("Could not begin race")
			}

			// solve the mazes
			for mazePath := result.NextMaze; mazePath != ""; {
				theMaze := getMaze("https://api.noopschallenge.com" + mazePath)
				sol := solve(theMaze)

				var solResult mazeResponse
				jsonPost("https://api.noopschallenge.com"+theMaze.MazePath, jsonType{"directions": sol}, &solResult)

				fmt.Printf("Maze: %v : Solution Length: %v\n", theMaze.Name, len(sol))
				mazePath = solResult.NextMaze

				if solResult.Result == "finished" {
					fmt.Printf("\n%v\n\nCertificate: https://api.noopschallenge.com%v\n\n", solResult.Message, solResult.Certificate)
				}
			}
		}
	}

	if showHelp {
		fmt.Printf("Parameters:\n\n")
		fmt.Printf("\thelp\tThis help screen\n\n")
		fmt.Printf("\trand\tUse a random maze\n\n")
		fmt.Printf("\trace [login]\tRace mode using [login] to login\n\n")
	}
}
