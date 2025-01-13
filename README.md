# ascii-art-web

Ascii Art Web runs a web server that serves a web page where users can input text and choose a banner style to render the text as ASCII art. 

## Author

- [Björn Österman](https://01.gritlab.ax/git/bosterma)

## Usage

Run the program in its root directory with:
```bash
go run .
```

Then navigate to http://localhost:8080

## Program Description

The web page includes:

- A text input area for users to enter the text they want converted to ASCII art.
- Radio buttons to select one of three banner styles: "Standard", "Shadow", or "Thinkertoy".
- A "Generate ASCII Art" button to submit the form and generate the ASCII art.

The web server is a program written in Go where the main logic is handled by the `homeHandler` function, which processes the user's input:

- If the server is accessed through any path other than the root (`/`), it returns a 404 Not Found error.
- If the submitted banner style is invalid, the program returns a 400 Bad Request error.
- Invalid characters are removed from the user-provided text using the `formatInput` function, and the ASCII art is generated with the `CreateArt` function that uses code from the ascii-art exercise. If an error occurs during this process, a 500 Internal Server Error is returned.
- The ASCII art result is then displayed on the web page inside a `<pre>` tag, allowing proper formatting and alignment for the ASCII characters.

The `main` function sets up the server to handle:

- Static files, such as stylesheets, from the `/static/` directory.
- The root path (`/`), where the form and results are displayed.

The server runs on `http://localhost:8080`.