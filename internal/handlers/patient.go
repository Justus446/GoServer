package handlers

import (
	"database/sql"
	"go-server/internal/models"
	"go-server/internal/services"
	"html/template"
	"net/http"
	"strconv"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func IndexHandler(db *sql.DB) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		patients, err := services.GetAllPatients(db)
		if err != nil{
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		templates.ExecuteTemplate(w, "index.html", patients)
	}
}


func CreatePatientHandler(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        if r.Method == http.MethodPost {
            patient := models.Patient{
                Name:  r.FormValue("name"),
                Age:   atoi(r.FormValue("age")),
                Email: r.FormValue("email"),
            }

            if _, err := services.CreatePatient(db, patient); err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
                return
            }
            http.Redirect(w, r, "/", http.StatusSeeOther)
            return
        }
        templates.ExecuteTemplate(w, "create.html", nil)
    }
}

func atoi(s string) int{
	i, _ := strconv.Atoi(s)

	return i
}