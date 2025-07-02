package main

import (
	"html/template"
	"io/ioutil"
	"net/http"
	"path/filepath"

	"github.com/russross/blackfriday/v2"
)

type Project struct {
	Title       string
	Description string
	Slug        string
	HTMLContent template.HTML
}

// Projects Section
var projects = []Project{
	//{Title: "Template", Slug: "template", Description: ""},
	{Title: "HK-Aerial", Slug: "hkaerial", Description: "Senior design drone project"},
	{Title: "Portfoli-Go", Slug: "portfolio", Description: "Self-hosted portfolio using Go"},
	{Title: "Homelab", Slug: "homelab", Description: "Homelab for personal development"},
	{Title: "GBA SP Reshell + USB-C Mod", Slug: "gameboy", Description: "Restored an old GBA SP with a new shell and USB-C mod"},
	{Title: "Satisfaction75 Build", Slug: "satisfaction75", Description: "Build log for custom keyboard"},
	{Title: "DAS - Direct-Attached Storage", Slug: "das", Description: "Expansion of homelab storage"},
	{Title: "SEO Optimizer", Slug: "seo-optimizer", Description: "Created an SEO Optimizer for websites"},
	{Title: "DIY TV Ambilight", Slug: "ambilight", Description: "Replicated an ambilight system using open-sourced software and off-the-shelf products"},
	{Title: "Home Surveillance", Slug: "home-cam", Description: "WIP"},
	{Title: "Hand Tracked Mouse", Slug: "camera-mouse", Description: "WIP"},
}

// Main function serving
func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/projects/", projectHandler)

	println("Server started at port :3000...")
	http.ListenAndServe(":3000", nil)
}

// Serves index page (main page)
func indexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	err := tmpl.Execute(w, projects)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// Handles projects listings
func projectHandler(w http.ResponseWriter, r *http.Request) {
	slug := r.URL.Path[len("/projects/"):]

	var selected *Project
	for _, p := range projects {
		if p.Slug == slug {
			selected = &p
			break
		}
	}

	if selected == nil {
		http.NotFound(w, r)
		return
	}

	mdPath := filepath.Join("content", slug+".md")
	mdBytes, err := ioutil.ReadFile(mdPath)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	rendered := blackfriday.Run(mdBytes)
	selected.HTMLContent = template.HTML(rendered)

	tmpl := template.Must(template.ParseFiles("templates/project.html"))
	tmpl.Execute(w, selected)
}
