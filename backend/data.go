package main

import (
	// "bytes"
	"bytes"
	"database/sql"
	"encoding/json"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

// Generic struct for organization data
type Organization struct {
	ID         int    `json:"organizationId"`
	Name       string `json:"organizationName"`
	OwnerEmail string `json:"organizationOwnerEmail"`
	PlanId     int    `json:"organizationPlanId"`
}

type User struct {
	Email      string `json:"email"`
	Name       string `json:"name"`
	Department string `json:"department"`
	JobTitle   string `json:"jobTitle"`
}

// Get DB settings from techchillaconfig
var dbCreds = GetDBConnectionSettingPostgresql()
var db *sql.DB

func InitDB() {
	var err error
	db, err = sql.Open("postgres", dbCreds)
	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}
}

func CheckDB() {
	if db == nil || db.Ping() != nil {
		InitDB()
	}
}

func MambuNumberFive(inputFirstName string, inputLastName string, inputPreferredLanguage string, inputBranchKey string, inputNric string, inputAddress string, inputPostal string, inputcountry string) map[string]interface{} {
	selection := "https://razerhackathon.sandbox.mambu.com/api/clients"
	mambuUser := os.Getenv("USERMAMBU")
	mambuPass := os.Getenv("PASSMAMBU")
	client := http.Client{}
	var data = []byte(`{
		"client": {
			"firstName": "` + inputFirstName + `",
			"lastName": "` + inputLastName + `",
			"preferredLanguage": "` + inputPreferredLanguage + `",
			"assignedBranchKey": "` + inputBranchKey + `"
		},
		"idDocuments": [
			{
				"identificationDocumentTemplateKey": "8a8e867271bd280c0171bf7e4ec71b01",
				"issuingAuthority": "Immigration Authority of Singapore",
				"documentType": "NRIC/Passport Number",
				"documentId": "` + inputNric + `"
			}
		],
		"addresses": [
				  {
			 "line1":"` + inputAddress + `",
			 "postcode":"` + inputPostal + `",
			 "country":"` + inputcountry + `"
		  }
			]
	}`)
	req, err := http.NewRequest("POST", selection, bytes.NewBuffer(data))
	if err != nil {
		panic(err)
	}
	req.SetBasicAuth(mambuUser, mambuPass)
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		panic(err)

	}
	if resp.StatusCode == http.StatusBadRequest {

	}
	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	return result
}

// Idk what I am doing here already. God save me.
func FWDKycAh(image string) map[string]interface{} {
	// w.Header().Set("Content-Type", "application/json")
	selection := "https://niw1itg937.execute-api.ap-southeast-1.amazonaws.com/Prod/verify"
	fwdKey := os.Getenv("FWDKEY")
	client := http.Client{}
	var data = []byte(`{"base64image":"` + image + `"}`)
	req, _ := http.NewRequest("POST", selection, bytes.NewBuffer(data))
	req.Header.Set("x-api-key", fwdKey)
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		panic(err)

	}
	if resp.StatusCode == http.StatusBadRequest {

	}
	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	return result
}

func CheckeEmailExist(email string) bool {
	CheckDB()
	statement := "select email from clients where email = $1"
	results := db.QueryRow(statement, email)
	outcome := ""
	err := results.Scan(&outcome)
	if err != nil {
		return false
	}
	return true
}

func SeeIfExisting(email string) (string, string) {
	CheckDB()
	statement := "select COALESCE (mambukey, '') from clients where email = $1"
	results := db.QueryRow(statement, email)
	outcome := ""
	_ = results.Scan(&outcome)
	if outcome == "" {
		return "", ""
	}
	selection := "https://razerhackathon.sandbox.mambu.com/api/clients/" + outcome + "?fullDetails=true"
	mambuUser := os.Getenv("USERMAMBU")
	mambuPass := os.Getenv("PASSMAMBU")
	client := http.Client{}
	req, _ := http.NewRequest("GET", selection, nil)
	req.SetBasicAuth(mambuUser, mambuPass)
	resp, err := client.Do(req)
	if err != nil {
		panic(err)

	}
	if resp.StatusCode == http.StatusBadRequest {

	}
	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	nricSlice := result["idDocuments"].([]interface{})[0]
	nric := nricSlice.(map[string]interface{})["documentId"].(string)
	return nric, outcome
}

func RetrieveClientState(email string) bool {
	CheckDB()
	statement := "select verified from clients where email = $1"
	results := db.QueryRow(statement, email)
	outcome := ""
	_ = results.Scan(&outcome)
	return outcome == "true"
}

func UpdateMambuKey(email string, mambykey string) (int64, error) {
	CheckDB()
	statement := "UPDATE clients SET mambukey = $1, verified = $2 WHERE email = $3"
	update, err := db.Exec(statement, mambykey, "true", email)
	if err != nil {
		panic(err)
	}
	rows, updateErr := update.RowsAffected()
	if updateErr != nil {
		panic(updateErr)
	}
	return rows, updateErr
}

func AddUser(email string) bool {
	CheckDB()
	statement := "INSERT INTO clients VALUES ($1,$2,null)"
	update, err := db.Exec(statement, email, "false")
	if err != nil {
		panic(err)
	}
	rows, updateErr := update.RowsAffected()
	if updateErr != nil {
		panic(updateErr)
	}
	return rows == 1
}

// func CreateOrganization(org Organization) (int64, error) {
// 	CheckDB()
// 	statement := "INSERT INTO organization (org_name, owner_email, plan_id) VALUES ($1,$2,$3)"
// 	update, err := db.Exec(statement, org.Name, org.OwnerEmail, org.PlanId)
// 	if err != nil {
// 		panic(err)
// 	}
// 	rows, updateErr := update.RowsAffected()
// 	if updateErr != nil {
// 		panic(updateErr)
// 	}
// 	return rows, updateErr
// }

// func GetUserByEmail(email string) *sql.Row {
// 	CheckDB()
// 	statement := "select * from users where email=$1"
// 	results := db.QueryRow(statement, email)
// 	return results

// }

// func GetUserByEntity(entity string) (*sql.Rows, error) {
// 	CheckDB()
// 	statement := "select * from users where email in (select email from user_entity where ent_id in (select ent_id from entity where entity_name = $1))"
// 	results, err := db.Query(statement, entity)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return results, err

// }

// func GetUserByOrganization(organization string) (*sql.Rows, error) {
// 	CheckDB()
// 	statement := "select * from users where email in (select email from user_entity where org_id in (select org_id from organization where org_name = $1))"
// 	results, err := db.Query(statement, organization)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return results, err

// }

// func CreateUser(user User) (int64, error) {
// 	CheckDB()
// 	statement := "INSERT INTO users (email, name, department,job_title) VALUES ($1,$2,$3,$4)"
// 	new, err := db.Exec(statement, user.Email, user.Name, user.Department, user.JobTitle)
// 	if err != nil {
// 		panic(err)
// 	}
// 	rows, updateErr := new.RowsAffected()
// 	if updateErr != nil {
// 		panic(updateErr)
// 	}
// 	return rows, updateErr
// }

// func CreateUsers(user []*User) (int64, error) {
// 	CheckDB()
// 	users := make([]interface{}, 0)
// 	statement := "INSERT INTO users (email, name, department,job_title) VALUES "
// 	count := 1
// 	for _, s := range user {
// 		users = append(users, s.Email, s.Name, s.Department, s.JobTitle)
// 		statement += fmt.Sprintf("($%d,$%d,$%d,$%d),", count, count+1, count+2, count+3)
// 		count += 4
// 	}

// 	statement = statement[:len(statement)-1]
// 	update, err := db.Exec(statement, users...)

// 	if err != nil {
// 		panic(err)
// 	}

// 	rows, updateErr := update.RowsAffected()
// 	if updateErr != nil {
// 		panic(updateErr)
// 	}

// 	return rows, updateErr
// }

// func UpdateUser(user User) (int64, error) {
// 	CheckDB()
// 	statement := "UPDATE users SET name = $1, department = $2, job_title = $3 WHERE email = $4"
// 	new, err := db.Exec(statement, user.Name, user.Department, user.JobTitle, user.Email)
// 	if err != nil {
// 		panic(err)
// 	}
// 	rows, updateErr := new.RowsAffected()
// 	if updateErr != nil {
// 		panic(updateErr)
// 	}
// 	return rows, updateErr
// }

// func DeleteUser(user User) (int64, error) {
// 	CheckDB()
// 	statement := "DELETE FROM users WHERE email = $1"
// 	new, err := db.Exec(statement, user.Email)
// 	if err != nil {
// 		panic(err)
// 	}
// 	rows, updateErr := new.RowsAffected()
// 	if updateErr != nil {
// 		panic(updateErr)
// 	}
// 	return rows, updateErr
// }
