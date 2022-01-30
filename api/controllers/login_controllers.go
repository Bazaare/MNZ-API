package controllers

import (
	"MNZ/api/auth"
	"MNZ/api/responses"
	"net/http"
)

func (server *Server) Login(w http.ResponseWriter, r *http.Request) {
	token, err := auth.CreateToken()

	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
	}

	responses.JSON(w, http.StatusOK, token)
}
