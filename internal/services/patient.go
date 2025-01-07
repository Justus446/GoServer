package services

import(
	"database/sql"
	"go-server/internal/models"
	"go-server/internal/repositories"
)

func CreatePatient(db *sql.DB, patient models.Patient)(int, error){
	return repositories.CreatePatient(db, patient)
}

func GetPatient(db *sql.DB, id int)(models.Patient, error){
	return repositories.GetPatient(db, id)
}

func GetAllPatients(db *sql.DB)([]models.Patient, error){
	return repositories.GetAllPatients(db)
}