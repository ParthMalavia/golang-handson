package story

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl, _ = template.ParseFiles("story\\template.html")
	// if err != nil {
	// 	panic(err)
	// }
}

type Option struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

type Chapter struct {
	Title      string   `json:"title"`
	Paragraphs []string `json:"story"`
	Options    []Option `json:"options"`
}

type HandlerOption func(h *Handler)

func WithTemplate(filePath string) HandlerOption {
	return func(h *Handler) {
		t, _ := template.ParseFiles(filePath)
		h.t = t
	}
}

func WithPathFunc(pathFunc func(*http.Request) string) HandlerOption {
	return func(h *Handler) {
		h.pathFunc = pathFunc
	}
}

func defaultPath(r *http.Request) string {
	path := r.URL.Path
	if path == "" || path == "/" {
		path = "/story-arc"
	}
	return path[1:]
}

type Handler struct {
	S        Story
	t        *template.Template
	pathFunc func(*http.Request) string
}

func GetNewHandler(s Story, opts ...HandlerOption) Handler {
	h := Handler{s, tpl, defaultPath}
	for _, opt := range opts {
		opt(&h)
	}
	return h
}

type Story map[string]Chapter

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	// File name form current working dir
	// tpl, err := template.ParseFiles("story\\template.html")
	// if err != nil {
	// 	panic(err)
	// }

	path := h.pathFunc(r)
	if chapter, ok := h.S[path]; ok {
		err := h.t.Execute(w, chapter)
		if err != nil {
			log.Println("Error: ", err)
			http.Error(w, "Something went wrong...", http.StatusInternalServerError)
		}
		return
	}
	http.Error(w, "Chapter not found.", http.StatusNotFound)
}

func JsonToStory(r io.Reader) (Story, error) {
	d := json.NewDecoder(r)
	var story Story
	if err := d.Decode(&story); err != nil {
		return nil, err
	}
	return story, nil
}
