package main

import (
	"html/template"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func cryptePassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func generateHTML(w http.ResponseWriter, data interface{}, filenames ...string) {
	var files []string
	for _, filename := range filenames {
		files = append(files, "templates/"+filename+".html")
	}
	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(w, filenames[0], data)
}
