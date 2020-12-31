package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"
)

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/form", formHandler)

	fmt.Printf("Starting server at port 8080\n")

	db := connectDb()
	migrateDb(db)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	layoutPath := buildLayoutPath("index")
	templatePath := buildTemplatePath("index")

	layoutContent := loadFile(layoutPath)
	templateContent := loadFile(templatePath)

	regex := regexp.MustCompile(`<>.*<\/>`)
	matches := regex.FindAllString(layoutContent, -1)

	content := layoutContent

	for _, v := range matches {
		fmt.Println(v)

		if strings.Contains(v, "yield") {
			vRegex := regexp.MustCompile(v)
			fmt.Println(vRegex)
			content = vRegex.ReplaceAllString(content, templateContent)
			fmt.Println("YIELDED")
		}
	}

	fmt.Fprintf(w, content)
}

func buildTemplatePath(name string) string {
	full := "./templates/" + name + ".html"
	return full
}

func buildLayoutPath(name string) string {
	full := "./layouts/" + name + ".html"
	return full
}

func loadFile(full string) string {

	b, err := ioutil.ReadFile(full)

	if err != nil {
		fmt.Print(err)
	}

	return string(b)
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "POST request successful")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address = %s\n", address)
}
