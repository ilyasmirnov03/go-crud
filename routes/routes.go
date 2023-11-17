package routes

import (
	"l33ass/crud/database"
	"l33ass/crud/services"
	"math/rand"
	"net/http"
	"text/template"
)

// Sets up app routes and its handle functions
func SetupRouter() {
	http.Handle("/build/", http.StripPrefix("/build/", http.FileServer(http.Dir("_dist"))))
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("public/assets"))))

	tmpl := make(map[string]*template.Template)
	tmpl["landing"] = template.Must(template.ParseFiles("templates/landing.html", "templates/base/base.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		posts := database.GetPosts()
		tmpl["landing"].ExecuteTemplate(w, "base", posts)
	})
	http.HandleFunc("/random", func(w http.ResponseWriter, r *http.Request) {
		number := rand.Int31()
		tmpl := template.Must(template.ParseFiles("templates/random.html"))
		tmpl.Execute(w, number)
	})
	http.HandleFunc("/notification", func(w http.ResponseWriter, r *http.Request) {
		services.SendPushNotification()
	})
}
