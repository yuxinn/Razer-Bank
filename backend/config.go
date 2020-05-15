package main

import (
	"fmt"
	"net/http"
)

const (
	phost     = "database-2.ccq5zpa8lmak.ap-southeast-1.rds.amazonaws.com"
	pport     = 5432
	puser     = "postgres"
	ppassword = "iloveeggtoast123$$"
)

func GetJWTKey() []byte {
	return jwtKey
}

func GetDBConnectionSettingPostgresql() string {
	conf := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=require", phost, pport, puser, ppassword, "fintech")
	return conf
}

func SetupResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "*")
	(*w).Header().Set("Access-Control-Expose-Headers", "Content-Disposition")
}
