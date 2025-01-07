package main
 import(
	"go-server/internal/handlers"
	"go-server/pkg/db"
	"net/http"
	"log"

 )

 func main(){
	database := db.ConnectPostgres()
	defer database.Close()

	http.HandleFunc("/", handlers.IndexHandler(database))
	http.HandleFunc("/createPatient", handlers.CreatePatientHandler(database))

	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
	
 }