package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	var (
		args  = os.Args[1:]
		name  string
		tMood string
		happy = [...]string{
			"feels happy",
			"feels awnsome!",
		}
		sad = [...]string{
			"feels sad",
			"feels unhapy",
		}
		mad = [...]string{
			"feels mad",
			"feels hungry",
		}
		mood = [...][2]string{happy, mad, sad}
	)

	if len(args) == 2 {
		name, tMood = args[0], args[1]

		rand.Seed(time.Now().UnixNano())
		n := rand.Intn(len(mood) - 1)

		switch true {
		case tMood == "happy":
			fmt.Printf("%s %v\n", name, mood[0][n])
		case tMood == "mad":
			fmt.Printf("%s %v\n", name, mood[1][n])
		default:
			fmt.Printf("%s %v\n", name, mood[2][n])
		}
	} else {
		fmt.Println("[Please, write a name and a feeling]")
	}
}
