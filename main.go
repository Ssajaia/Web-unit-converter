package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

type PageData struct {
	Title   string
	Active  string
	Value   string
	From    string
	To      string
	Result  string
	HasForm bool
}

func main() {
	// Debug info
	log.Println("Starting server...")
	wd, _ := os.Getwd()
	log.Println("Working directory:", wd)

	// Load templates with Windows compatibility
	templates := template.New("").Funcs(template.FuncMap{})
	templates = template.Must(templates.ParseGlob(filepath.Join("templates", "*.html")))

	// Serve static files
	http.Handle("/static/", http.StripPrefix("/static/", 
		http.FileServer(http.Dir(filepath.Join(".", "static")))))

	// Routes
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/length", http.StatusSeeOther)
	})

	http.HandleFunc("/length", func(w http.ResponseWriter, r *http.Request) {
		handleConverterPage(w, r, templates, "length")
	})

	http.HandleFunc("/weight", func(w http.ResponseWriter, r *http.Request) {
		handleConverterPage(w, r, templates, "weight")
	})

	http.HandleFunc("/temperature", func(w http.ResponseWriter, r *http.Request) {
		handleConverterPage(w, r, templates, "temperature")
	})

	// Test route
	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Server is working!"))
	})

	log.Println("Server started on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleConverterPage(w http.ResponseWriter, r *http.Request, templates *template.Template, converterType string) {
	data := PageData{
		Title:   converterType,
		Active:  converterType,
		HasForm: true,
	}

	if r.Method == http.MethodPost {
		value := r.FormValue("value")
		from := r.FormValue("from")
		to := r.FormValue("to")

		data.Value = value
		data.From = from
		data.To = to

		// Add your conversion logic here
		data.Result = "42" // Test value
	}

	err := templates.ExecuteTemplate(w, converterType+".html", data)
	if err != nil {
		log.Printf("Template error: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error loading template"))
	}
}