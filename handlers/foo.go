package handlers

import (
	"net/http"

	"github.com/VladanT3/GoTH/views"
)

func Foo(w http.ResponseWriter, r *http.Request) error {
	return views.Index().Render(r.Context(), w)
}
