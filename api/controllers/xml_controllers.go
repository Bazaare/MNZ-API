package controllers

import (
	"MNZ/api/models"
	"MNZ/api/responses"
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"strconv"
)

func (server *Server) GetXML(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	XMLId, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	url := fmt.Sprintf("https://raw.githubusercontent.com/MiddlewareNewZealand/evaluation-instructions/main/xml-api/%d.xml", XMLId)

	resp, err := models.XMLGet(url)

	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	bodyBytes, err := io.ReadAll(resp.Body)

	data := models.XML{}

	err = xml.Unmarshal(bodyBytes, &data)

	jsonData, err := json.Marshal(data) // Encode our Struct as JSON

	defer resp.Body.Close()

	if err != nil {
		responses.ERROR(w, resp.StatusCode, err)
	}

	var prettyJSON bytes.Buffer
	error := json.Indent(&prettyJSON, jsonData, "", "\t")
	if error != nil {
		log.Println("JSON parse error: ", error)
		return
	}
	fmt.Fprintf(w, "%s", string(prettyJSON.Bytes()))
}
