package main

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
	"sort"
	"strings"
	"os"

	"github.com/russross/blackfriday/v2"
)

type Project struct {
	Title       string
	Description string
	Slug        string
	Status      string
	HTMLContent template.HTML
}

// Projects Section
var projects = []Project{
	// {Title: "Template", Slug: "Template", Description: "WIP", Status: "wip"},
	{Title: "HK-Aerial", Slug: "hkaerial", Description: "Senior capstone design drone project", Status: "active"},
	{Title: "Portfoli-Go", Slug: "portfolio", Description: "Self-hosted portfolio using Go", Status: "active"},
	{Title: "Homelab", Slug: "homelab", Description: "Homelab for personal development", Status: "active"},
	{Title: "GBA SP Reshell + USB-C Mod", Slug: "gameboy", Description: "Restored an old GBA SP with a new shell and USB-C mod", Status: "active"},
	{Title: "Satisfaction75 Build", Slug: "satisfaction75", Description: "Build log for custom keyboard", Status: "active"},
	{Title: "DAS - Direct-Attached Storage", Slug: "das", Description: "Expansion of homelab storage", Status: "active"},
	{Title: "SEO Optimizer", Slug: "seo-optimizer", Description: "Created an SEO Optimizer for websites", Status: "active"},
	{Title: "DIY TV Ambilight", Slug: "ambilight", Description: "Replicated an ambilight system using open-sourced software and off-the-shelf products", Status: "active"},
	{Title: "Home Surveillance", Slug: "home-cam", Description: "WIP", Status: "wip"},
	//{Title: "Hand Tracked Mouse", Slug: "camera-mouse", Description: "WIP", Status: "wip"},
	{Title: "Riftbound TCG Assistant", Slug: "riftbound-assistant", Description: "Assistant for Riftbound TCG assists in decision making and learning about ML", Status: "wip"},
	{Title: "BO2 Interactive Zombies Mod", Slug: "bo2-mod", Description: "BO2 Zombies Mod allowing for interactive participation from streamer chat", Status: "active"},
	{Title: "Cyberdeckv0.11", Slug: "cyberdeck", Description: "WIP", Status: "wip"},
}

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/projects", projectsRouter)
	http.HandleFunc("/projects/", projectsRouter)
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/skills", skillsHandler)

	fmt.Println("New version detected...")
	fmt.Println("Server started at port :3000...")
	http.ListenAndServe(":3000", nil)
}

// Serves index page (main page)
func indexHandler(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	err := tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// Serves skills page
func skillsHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/skills" {
		http.NotFound(w, r)
		return
	}

	tmpl := template.Must(template.ParseFiles("templates/skills.html"))
	err := tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// Routes both /projects and /projects/{slug}
func projectsRouter(w http.ResponseWriter, r *http.Request) {

	path := strings.TrimSuffix(r.URL.Path, "/")

	// Archive page
	if path == "/projects" {
		sortedProjects := make([]Project, len(projects))
		copy(sortedProjects, projects)

		sort.SliceStable(sortedProjects, func(i, j int) bool {
			// Put WIP projects at the bottom
			if sortedProjects[i].Status != sortedProjects[j].Status {
				return sortedProjects[i].Status != "wip"
			}

			// Alphabetical inside each group
			return strings.ToLower(sortedProjects[i].Title) < strings.ToLower(sortedProjects[j].Title)
		})

		tmpl := template.Must(template.ParseFiles("templates/projects.html"))
		err := tmpl.Execute(w, sortedProjects)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	// Individual project page
	slug := strings.TrimPrefix(path, "/projects/")
	if slug == "" || slug == "/projects" {
		http.NotFound(w, r)
		return
	}

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
	mdBytes, err := os.ReadFile(mdPath)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	rendered := blackfriday.Run(mdBytes)
	selected.HTMLContent = template.HTML(rendered)

	tmpl := template.Must(template.ParseFiles("templates/project-single.html"))
	err = tmpl.Execute(w, selected)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
