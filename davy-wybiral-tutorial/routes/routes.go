package routes

import (
	"../middleware"
	"../models"
	"../sessions"
	"../view"
	"context"
	"github.com/gorilla/mux"
	"net/http"
)

var ctx context.Context

func NewRouter(context *context.Context) (*mux.Router) {
	r := mux.NewRouter()
	ctx = *context
	r.HandleFunc("/", middleware.AuthRequired(indexGetHandler)).Methods("GET")
	r.HandleFunc("/", middleware.AuthRequired(indexPostHandler)).Methods("POST")

	r.HandleFunc("/login", loginGetHandler).Methods("GET")
	r.HandleFunc("/login", loginPostHandler).Methods("POST")


	r.HandleFunc("/register", registerGetHandler).Methods("GET")
	r.HandleFunc("/register", registerPostHandler).Methods("POST")


	fs := http.FileServer(http.Dir("./static/"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static", fs))
	return r
}

func loginGetHandler(w http.ResponseWriter, r *http.Request) {
	view.ExecuteTemplate(w, "login.html", nil)
}

func loginPostHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := r.PostForm.Get("username")
	password := r.PostForm.Get("password")
	err := models.AuthenticateUser(ctx, username, password)
	if err != nil {
		switch err {
		case models.ErrUserNotFound:
			view.ExecuteTemplate(w, "login.html", "unknown user")
		case models.ErrInvalidLogin:
			view.ExecuteTemplate(w, "login.html", "invalid login")
		default:
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Internal server error"))
		}
		return
	}
	session, _ := sessions.Store.Get(r, "session")
	session.Values["username"] = username
	session.Save(r, w)
	http.Redirect(w, r, "/", 302)
}

func indexGetHandler(w http.ResponseWriter, r *http.Request) {
	comments, err := models.GetComments(ctx)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
		return
	}
	view.ExecuteTemplate(w, "index.html", comments)
}

func indexPostHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	comment := r.PostForm.Get("comment")
	err := models.PostComments(ctx, comment)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
	} else {
		http.Redirect(w, r, "/", 302)
	}
}

func registerGetHandler(w http.ResponseWriter, r *http.Request) {
	view.ExecuteTemplate(w, "register.html", nil)
}

func registerPostHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := r.PostForm.Get("username")
	password := r.PostForm.Get("password")
	err := models.RegisterUser(ctx, username, password)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
		return
	}
	http.Redirect(w, r, "/login", 302)
}