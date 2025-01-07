// Database operations for patient model
// Patient DB quries

package repositories

import (
	"database/sql"
	"go-server/internal/models"
)


func CreatePatient(db *sql.DB, patient models.Patient)(int, error){
	query := `INSERT INTO patients (name, age, email) VALUES ($1, $2, $3) RETURNING id`
	var id int
	err := db.QueryRow(query, patient.Name, patient.Age, patient.Email).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil 
}


func GetPatient(db *sql.DB, id int)(models.Patient, error){
	query := `SELECT name, age, email FROM patients WHERE id=$1`
	var patient models.Patient
	err := db.QueryRow(query, id).Scan(&patient.Name, &patient.Age, &patient.Email)
	if err != nil {
		return models.Patient{}, err
	}
	return patient, nil
}

func GetAllPatients(db *sql.DB)([]models.Patient, error){
	query := `SELECT id, name, age, email FROM patients`
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var patients []models.Patient
	for rows.Next(){
		var patient models.Patient
		err := rows.Scan(&patient.Id, &patient.Name, &patient.Age, &patient.Email)
		if err != nil {
			return nil, err
		}
		patients = append(patients, patient)
	}
	return patients,nil
}