package middleware

import (
	"encoding/json"
	"gorest/models"
	"log"
	"net/http"
)

// GetAllUser will return all the usermanagement
func GetAllUser(w http.ResponseWriter, r *http.Request) {
	usermanagement, err := getAllusermanagement()

	if err != nil {
		log.Fatalf("Unable to get all user. %v", err)
	}

	json.NewEncoder(w).Encode(usermanagement)
}

func getAllusermanagement() ([]models.User, error) {
	db := createConnection()

	defer db.Close()

	var usermanagement []models.User

	sqlStatement := `SELECT * FROM usermanagement`

	rows, err := db.Query(sqlStatement)
	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	defer rows.Close()

	for rows.Next() {
		var user models.User

		err = rows.Scan(&user.ID, &user.Name, &user.Location, &user.Age)
		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		usermanagement = append(usermanagement, user)

	}

	return usermanagement, err
}
