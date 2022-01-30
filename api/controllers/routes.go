package controllers

import "MNZ/api/middlewares"

func (server *Server) initialiseRoutes() {
	server.Router.HandleFunc("/xml/{id:[0-9]}", middlewares.SetMiddlewareAuthentication(server.GetXML)).Methods("GET")
	server.Router.HandleFunc("/xml/login", middlewares.SetMiddlewareJSON(server.Login)).Methods("POST")
}
