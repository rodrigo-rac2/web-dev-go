package handlers

import (
	"github.com/rodrigo-rac2/web-dev-go/s3/bootstrap/pkg/render"
	"net/http"
)

// Home is the handler for the home page
func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderHtml(w, "home.page.tmpl")
}

func About(w http.ResponseWriter, r *http.Request) {
	render.RenderHtml(w, "about.page.tmpl")
}
