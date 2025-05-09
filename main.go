package main

import (
	"github.com/russross/blackfriday/v2"
	"html/template"
	"io/ioutil"
	"net/http"
	"path/filepath"
)

type Project struct {
	Title string
	Slug  string
}

var projects = []Project{
	{Title: "Template", Slug: "template"},
	{Title: "HK-Aerial", Slug: "hkaerial"},
	{Title: "Portfoli-Go", Slug: "portfolio"},
	{Title: "Homelab", Slug: "homelab"},
	{Title: "Satisfaction75 Build", Slug: "satisfaction75"},
	{Title: "DAS - Direct-Attached Storage", Slug: "das"},
	{Title: "Home Surveillance", Slug: "home-cam"},
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

	mdPath := filepath.Join("content", slug+".md")
	mdBytes, err := ioutil.ReadFile(mdPath)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	htmlContent := blackfriday.Run(mdBytes)

	tmpl := template.Must(template.ParseFiles("templates/project-deprecated.html"))
	tmpl.Execute(w, struct {
		Title   string
		Content template.HTML
	}{
		Title:   slug,
		Content: template.HTML(htmlContent),
	})
}

// Implementation using structs. Deprecated
/*type Project struct {
	Title        string
	Description  string
	Slug         string
	Image        string
	Video        string
	Content      string
	Technologies []string
}

var projects = []Project{
	{
		Title:        "Template",
		Slug:         "Template",
		Description:  "Template",
		Content:      "",
		Technologies: []string{},
		Image:        "",
		Video:        "",
	},
	{
		Title:        "HK Aerial - WIP",
		Slug:         "hk-aerial",
		Description:  "Lead a team to design and manufacture an all-terrain payload delivery drone.",
		Content:      "WIP",
		Technologies: []string{},
		Image:        "/static/images/hkaerial.jpg",
		Video:        "https://www.youtube.com/embed/A8Q2BG28Pes?si=UdRDsQPxgsCypOam",
	},
	{
		Title:        "Portfoli-Go - WIP",
		Slug:         "portfolio",
		Description:  "Self-hosted portfolio using Go",
		Content:      "",
		Technologies: []string{},
		Image:        "",
		Video:        "",
	},
	{
		Title:        "Homelab - WIP",
		Slug:         "homelab",
		Description:  "Home server running Proxmox hypervisor with LXC's, Docker, and Samba share ZFS pool via DAS (direct attached storage).",
		Content:      "",
		Technologies: []string{},
		Image:        "",
		Video:        "",
	},
	{
		Title:        "Satisfaction75 Build - WIP",
		Slug:         "satisfaction75",
		Description:  "Somewhat in-depth build log on friend's keyboard, Satisfaction75.",
		Content:      "",
		Technologies: []string{},
		Image:        "",
		Video:        "",
	},
	{
		Title:        "Direct-Attached Storage - WIP",
		Slug:         "das",
		Description:  "Implementing a DAS to my homelab to expand storage",
		Content:      "",
		Technologies: []string{},
		Image:        "",
		Video:        "",
	},
	{
		Title:        "ESP-32 Cam + Object Detection - WIP",
		Slug:         "home-cam",
		Description:  "Utilizing an ESP-32 Cam module to develop a home surveillance camera with object detection capabilities",
		Content:      "",
		Technologies: []string{},
		Image:        "",
		Video:        "",
	},
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

	tmpl := template.Must(template.ParseFiles("templates/project-deprecated.html"))
	tmpl.Execute(w, selected)
}*/

// Original implementation, non-dynamic. Deprecated
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
