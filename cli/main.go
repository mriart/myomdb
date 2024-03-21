// An omdb CLI to get information (ratings, year...) of a movie.
// Usage:  Usage: cli [-p] [-h] title
// Credits to OMDB. Requires an API key.
// Marc Riart, 202403

package main

import (
	"flag"
	"fmt"

	"github.com/mriart/myomdb/movie"
)

const usage = `Usage: cli [-p] [-h] title
	-p	Show the movie poster in ascii art
	-h	Display help`

func main() {
	title, showPoster, ok := getArgs()
	if !ok {
		return
	}

	m := movie.Movie{}
	err := m.GetMovie(title)
	if err != nil {
		fmt.Println("Something went wrong in GetMovie")
		return
	}

	if showPoster {
		err = m.PrintPoster()
		if err != nil {
			fmt.Println("Something went wrong in PrinterPoster")
			return
		}
	}

	m.PrintMovie()
}

// Get arguments from the command line
func getArgs() (title string, showPoster bool, ok bool) {
	p := flag.Bool("p", false, "Print the movie poster")
	h := flag.Bool("h", false, "Display help")
	flag.Parse()

	// If -h, display help and exit
	if *h {
		fmt.Println(usage)
		return "", false, false
	}

	// Get list of arguments, that is, the title of the movie
	titles := flag.Args()
	if len(titles) < 1 {
		fmt.Printf("Specify a movie title\n%s\n", usage)
		return "", false, false
	}

	return titles[0], *p, true
}
