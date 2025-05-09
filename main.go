package main

import (
	"html/template"
	"net/http"
)

type Project struct {
	Title       string
	Description string
	Slug        string
	Image       string
}

var projects = []Project{
	{Title: "HK Aerial", Slug: "hk-aerial", Description: "Lead a team to design and manufacture an all-terrain payload delivery drone."},
	{Title: "Portfoli-Go", Slug: "portfolio", Description: "Self-hosted portfolio using Go"},
	{Title: "Homelab", Slug: "homelab", Description: "Home server running Proxmox hypervisor with LXC's, Docker, and Samba share ZFS pool via DAS (direct attached storage)."},
}

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/projects/", projectHandler)

	println("Server started at port :3000...")
	http.ListenAndServe(":3000", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	err := tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

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

	tmpl := template.Must(template.ParseFiles("templates/project.html"))
	tmpl.Execute(w, selected)
}

// Old main, non-dynamic style
/*
func handler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	err := tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	// Serve main page
	http.HandleFunc("/", handler)

	// Serve static files
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	log.Println("Server starting @ :3000...")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
*/
