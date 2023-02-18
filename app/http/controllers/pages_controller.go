package controllers

import (
	"gos/pkg/view"
	"net/http"
)

type PagesController struct {
}

func (*PagesController) Home(w http.ResponseWriter, r *http.Request) {
	view.Render(w, "pages.home", nil)
}
