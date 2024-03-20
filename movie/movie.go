// Movie type struct & methods to complement the main package.
// Credits to OMDB. Requires an API key.
// Marc Riart, 202403

package movie

import (
	"encoding/json"
	"fmt"
	"image"
	"io"
	"net/http"

	"github.com/qeesung/image2ascii/convert"
)

const apiKey = "8564cbe7"

// Type definition for Movie. It gets the data from ODMB
type Movie struct {
	Title    string `json:"Title"`
	Year     string `json:"Year"`
	Runtime  string `json:"Runtime"`
	Genre    string `json:"Genre"`
	Director string `json:"Director"`
	Actors   string `json:"Actors"`
	Plot     string `json:"Plot"`
	Awards   string `json:"Awards"`
	Poster   string `json:"Poster"`
	Ratings  []struct {
		Source string `json:"Source"`
		Value  string `json:"Value"`
	} `json:"Ratings"`
	Metascore  string `json:"Metascore"`
	ImdbRating string `json:"imdbRating"`
	Type       string `json:"Type"`
}

// Method to fill the struct, getting from the OMDB site
func (m *Movie) GetMovie(title string) error {
	res, err := http.Get("http://www.omdbapi.com/?t=" + title + "&apikey=" + apiKey)
	if err != nil {
		fmt.Println("Error connectiong to OMDB for " + title)
		return err
	}
	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)
	err = json.Unmarshal([]byte(body), m)
	if err != nil {
		fmt.Println("Error unmarshalling " + title)
		return err
	}

	return nil
}

// Print the movie information
func (m *Movie) PrintMovie() {
	fmt.Println("Title:\t", m.Title)
	fmt.Println("Year:\t", m.Year)
	for _, v := range m.Ratings {
		fmt.Println(shorten(v.Source), "\b:\t", v.Value)
	}
	fmt.Println("Awards:\t", m.Awards)
	fmt.Println("Genre:\t", m.Genre)
	fmt.Println(shorten("Director"), "\b:\t", m.Director)
	fmt.Println("Actors:\t", m.Actors)
}

// Print in ascci art the poster. It takes the URL from OMDB, and gets the poster from the internet (media AWS)
func (m *Movie) PrintPoster() error {
	// Get the poster image from url
	url := m.Poster
	res, err := http.Get(url)
	if err != nil {
		fmt.Println("Error connectiong to poster url" + url)
		return err
	}
	defer res.Body.Close()

	// Get the image into internal format
	img, _, err := image.Decode(res.Body)
	if err != nil {
		fmt.Println("Error decoding image:", err)
		return err
	}

	// Convert the image data to an ASCII art representation
	convertOptions := convert.DefaultOptions
	asciiArt := convert.NewImageConverter().Image2ASCIIString(img, &convertOptions)

	// Print the ASCII art to the console
	fmt.Println(asciiArt)

	return nil
}

// Shorten the receiver string in order to tabulate well, visually appealing
func shorten(source string) string {
	switch source {
	case "Internet Movie Database":
		return "IMDb"
	case "Rotten Tomatoes":
		return "Rotten"
	case "Metacritic":
		return "Meta"
	case "Director":
		return "Dir"
	default:
		return source
	}
}
