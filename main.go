package main

import (
	asciiart "asii-art-web/ascii-art"
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

type PageData struct {
	Txt          string
	Style        string
	Out          string
	ErrorMessage string
}

var tpl *template.Template

// formatInput makes sure the program can read the input
func formatInput(s string) (string, bool) {
	out := ""
	allGood := true
	for _, char := range s {
		if (char > 31 && char < 127) || char == '\n' { // Put printable characters and line feed to output
			out += string(char)
		} else if char >= 127 { // Characters outside ASCII not accepted
			allGood = false
		}
	}
	return out, allGood
}

// checkBanner makes sure the selected banner style is available
func checkBanner(banner string) bool {
	if banner != "standard" && banner != "shadow" && banner != "thinkertoy" {
		return false
	}
	return true
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	data := PageData{
		Style: "standard", // Default value. Nothing else is needed for Get method (no text to convert)
	}

	if r.URL.Path != "/" {
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}

	if r.Method == http.MethodPost {

		userInput := r.FormValue("usertext")
		style := r.FormValue("banner")

		if !checkBanner(style) {
			http.Error(w, "400 Bad request: Banner style incorrect", http.StatusBadRequest)
			return
		}

		input, ok := formatInput(userInput)
		if !ok {
			data.ErrorMessage = "Unable to convert all characters" // Message to user, still covert what we can
		}

		out, err := asciiart.CreateArt(strings.Split(input, "\n"), style) // Create the art
		if err != nil {
			http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
			return
		}

		// Put values to page data struct
		data.Txt = userInput
		data.Style = style
		data.Out = out

	} else if r.Method != http.MethodGet {
		http.Error(w, "405 Method not Allowed", http.StatusMethodNotAllowed)
		return
	}

	// Execute the template with the page data struct
	tpl.Execute(w, data)
}

func main() {

	var err error
	tpl, err = template.ParseFiles("templates/index.html")
	if err != nil {
		fmt.Println("Error parsing index template:", err.Error())
	}

	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fileServer))

	http.HandleFunc("/", homeHandler)

	fmt.Println("Server is running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
