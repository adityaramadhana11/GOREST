package middleware

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"gorest/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// GetUser will return a single user by its id
func GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Unable to convert the string into int.  %v", err)
	}

	user, err := getUser(int64(id))

	if err != nil {
		log.Fatalf("Unable to get user. %v", err)
	}

	json.NewEncoder(w).Encode(user)
}

func getUser(id int64) (models.User, error) {
	db := createConnection()
	defer db.Close()

	var user models.User

	sqlStatement := `SELECT * FROM usermanagement WHERE userid=$1`

	row := db.QueryRow(sqlStatement, id)

	err := row.Scan(&user.ID, &user.Name, &user.Location, &user.Age)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return user, nil
	case nil:
		return user, nil
	default:
		log.Fatalf("Unable to scan the row. %v", err)
	}

	return user, err
}
