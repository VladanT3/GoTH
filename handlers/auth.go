package handlers

import (
	"net/http"

	"github.com/VladanT3/GoTH/views/auth"
)

func GetLogin(w http.ResponseWriter, r *http.Request) error {
	return Render(w, r, auth.Login())
}
