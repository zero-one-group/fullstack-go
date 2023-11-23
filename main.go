package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/zero-one-group/fullstack-go/controllers"
	"github.com/zero-one-group/fullstack-go/models"
	"github.com/zero-one-group/fullstack-go/templates"
	"github.com/zero-one-group/fullstack-go/views"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Setup a database connection
	cfg := models.DefaultPostgresConfig()
	db, err := models.Open(cfg)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Setup our model services
	userService := models.UserService{
		DB: db,
	}

	// Setup our controllers
	usersC := controllers.Users{
		UserService: &userService,
	}
	usersC.Templates.New = views.Must(views.ParseFS(
		templates.FS, "signup.html", "tailwind.html"))

	usersC.Templates.SignIn = views.Must(views.ParseFS(
		templates.FS, "signin.html", "tailwind.html"))

	// Setup our routing
	r.Get("/", controllers.StaticHandler(
		views.Must(views.ParseFS(templates.FS, "home.html", "tailwind.html"))))
	r.Get("/contact", controllers.StaticHandler(
		views.Must(views.ParseFS(templates.FS, "contact.html", "tailwind.html"))))
	r.Get("/faq", controllers.FAQ(
		views.Must(views.ParseFS(templates.FS, "faq.html", "tailwind.html"))))
	r.Get("/signup", usersC.New)
	r.Post("/signup", usersC.Create)

	r.Get("/signin", usersC.SignIn)
	r.Post("/signin", usersC.ProcessSignIn)

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})

	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", r)
}
