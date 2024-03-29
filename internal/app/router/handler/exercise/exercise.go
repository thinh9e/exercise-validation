package exercise

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"

	"github.com/gorilla/mux"
	"github.com/npthinhdev/valexer/internal/app/types"
	"github.com/npthinhdev/valexer/internal/pkg/parse"
)

type (
	val map[string]interface{}
	// Handler is web handler
	Handler struct{}
)

var tmlp = template.Must(template.ParseGlob("web/template/*.html"))

// New return new web handler
func New() *Handler {
	return &Handler{}
}

// Get render get page
func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	ctx := val{"Title": "Exercise"}
	keys := mux.Vars(r)
	id := keys["id"]
	if len(id) < 1 {
		log.Println("url is not found")
		w.WriteHeader(http.StatusNotFound)
		return
	}
	apiURL := fmt.Sprintf("http://localhost:8080/api/exercise/%s", id)
	body := parse.Get(apiURL)
	if string(body) == "null" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	exer := types.Exercise{}
	_ = json.Unmarshal(body, &exer)
	ctx["Exer"] = exer
	err := tmlp.ExecuteTemplate(w, "exercise.html", ctx)
	if err != nil {
		log.Println(err)
	}
}

// Post render post page
func (h *Handler) Post(w http.ResponseWriter, r *http.Request) {
	ctx := val{"Title": "Exercise"}
	keys := mux.Vars(r)
	id := keys["id"]
	if len(id) < 1 {
		log.Println("url is not found")
		w.WriteHeader(http.StatusNotFound)
		return
	}
	apiURL := fmt.Sprintf("http://localhost:8080/api/exercise/%s", id)
	body := parse.Get(apiURL)
	if string(body) == "null" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	exer := types.Exercise{}
	_ = json.Unmarshal(body, &exer)
	ctx["Exer"] = exer
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
	}
	fileContent, err := parse.FormFile(r)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	result := parse.GetTesting(id, fileContent)
	ctx["Result"] = string(result)
	tmlp.ExecuteTemplate(w, "exercise.html", ctx)
}

// CreateGET render create page
func (h *Handler) CreateGET(w http.ResponseWriter, r *http.Request) {
	ctx := val{"Title": "Exercise"}
	err := tmlp.ExecuteTemplate(w, "create.html", ctx)
	if err != nil {
		log.Println(err)
	}
}

// CreatePOST render create page
func (h *Handler) CreatePOST(w http.ResponseWriter, r *http.Request) {
	apiURL := "http://localhost:8080/api/exercise"
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
	}
	fileContent, err := parse.FormFile(r)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	data := url.Values{}
	data.Set("title", r.Form.Get("title"))
	data.Set("description", r.Form.Get("description"))
	data.Set("testcase", string(fileContent))
	_ = parse.Post(apiURL, data)
	http.Redirect(w, r, "/admin", http.StatusFound)
}

// UpdateGET render edit page
func (h *Handler) UpdateGET(w http.ResponseWriter, r *http.Request) {
	ctx := val{"Title": "Update exercise"}
	keys := mux.Vars(r)
	id := keys["id"]
	if len(id) < 1 {
		log.Println("url is not found")
		w.WriteHeader(http.StatusNotFound)
		return
	}
	apiURL := fmt.Sprintf("http://localhost:8080/api/exercise/%s", id)
	body := parse.Get(apiURL)
	if string(body) == "null" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	exer := types.Exercise{}
	_ = json.Unmarshal(body, &exer)
	ctx["Exer"] = exer
	err := tmlp.ExecuteTemplate(w, "update.html", ctx)
	if err != nil {
		log.Println(err)
	}
}

// UpdatePOST render edit page
func (h *Handler) UpdatePOST(w http.ResponseWriter, r *http.Request) {
	keys := mux.Vars(r)
	id := keys["id"]
	apiURL := fmt.Sprintf("http://localhost:8080/api/exercise/%s", id)
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
	}
	data := url.Values{}
	data.Set("title", r.Form.Get("title"))
	data.Set("description", r.Form.Get("description"))
	data.Set("testcase", r.Form.Get("testcase"))
	_ = parse.Put(apiURL, data)
	http.Redirect(w, r, "/admin", http.StatusFound)
}

// Delete render delete page
func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	keys := mux.Vars(r)
	id := keys["id"]
	if len(id) < 1 {
		log.Println("url is not found")
		w.WriteHeader(http.StatusNotFound)
		return
	}
	apiURL := fmt.Sprintf("http://localhost:8080/api/exercise/%s", id)
	_ = parse.Delete(apiURL)
	http.Redirect(w, r, "/admin", http.StatusFound)
}
