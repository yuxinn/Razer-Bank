package main

import (

	// "fmt"

	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/gorilla/mux"
	// "strings"
	// "time"
)

type FWDStruct struct {
	Keywords *FWDWords `json:"keywords"`
	Image    string    `json:"base64image"`
}

type FWDWords struct {
	NRIC string `json:"idNum"`
	Name string `json:"name"`
}

type Token struct {
	Token string `json:"token"`
}

type UserVerified struct {
	Verified bool   `json:"status"`
	Email    string `json:"email"`
}

type Error struct {
	Error string `json:"error"`
}

func TestPoint(w http.ResponseWriter, r *http.Request) {

	if (*r).Method == "OPTIONS" {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Alive")
}

func Mambu(w http.ResponseWriter, r *http.Request) {

	if (*r).Method == "OPTIONS" {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	selection := ""
	mambuUser := os.Getenv("")
	mambuPass := os.Getenv("")
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
	json.NewEncoder(w).Encode(result)
}

func FWDKyc(w http.ResponseWriter, r *http.Request) {

	if (*r).Method == "OPTIONS" {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	selection := ""
	fwdKey := os.Getenv("")
	client := http.Client{}
	req, _ := http.NewRequest("POST", selection, r.Body)
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
	json.NewEncoder(w).Encode(result)
}

func RedirectToRazer(w http.ResponseWriter, r *http.Request) {

	if (*r).Method == "OPTIONS" {
		return
	}
	u := ""
	http.Redirect(w, r, u, http.StatusTemporaryRedirect)
}

func Callback(w http.ResponseWriter, r *http.Request) {

	if (*r).Method == "OPTIONS" {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	code := r.URL.Query()["code"][0]
	fmt.Println(code)
	selection := ""
	client := http.Client{}
	data := url.Values{}
	data.Set("client_id", "")
	data.Set("client_secret", "")
	data.Set("grant_type", "authorization_code")
	data.Set("code", code)
	data.Set("redirect_uri", "")
	req, _ := http.NewRequest("POST", selection, strings.NewReader(data.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(req)
	if err != nil {
		panic(err)

	}
	if resp.StatusCode == http.StatusBadRequest {

	}
	if resp.Body != nil {
		var result map[string]interface{}
		json.NewDecoder(resp.Body).Decode(&result)
		fmt.Println(result["access_token"])
		fmt.Println(result["expires_in"])
		fmt.Println(result["id_token"])
		fmt.Println(result["scope"])
		fmt.Println(result["token_type"])
		// json.NewEncoder(w).Encode(result)
		selection := ""
		client := http.Client{}
		ureq, _ := http.NewRequest("GET", selection, nil)
		ureq.Header.Set("Authorization", "Bearer "+result["access_token"].(string))
		uresp, err := client.Do(ureq)
		if err != nil {
			panic(err)

		}
		if uresp.StatusCode == http.StatusBadRequest {

		}
		var uresult map[string]interface{}
		json.NewDecoder(uresp.Body).Decode(&uresult)
		if CheckeEmailExist(uresult["email"].(string)) {
			nric, mambukey := SeeIfExisting((uresult["email"].(string)))
			jwttoken := GenerateJWT(uresult["email"].(string), uresult["first_name"].(string), uresult["last_name"].(string), nric, mambukey)
			u := "https://bank.ntucbee.click/verify?token=" + jwttoken
			http.Redirect(w, r, u, http.StatusTemporaryRedirect)
			return
		} else {
			_ = AddUser(uresult["email"].(string))
			nric, mambukey := SeeIfExisting((uresult["email"].(string)))
			fn := ""
			ln := ""
			if uresult["first_name"] != nil {
				fn = uresult["first_name"].(string)
			}
			if uresult["last_name"] != nil {
				ln = uresult["last_name"].(string)
			}
			jwttoken := GenerateJWT(uresult["email"].(string), fn, ln, nric, mambukey)
			u := "https://bank.ntucbee.click/verify?token=" + jwttoken
			http.Redirect(w, r, u, http.StatusTemporaryRedirect)
			return
		}

	}

}

func Authorization(w http.ResponseWriter, r *http.Request) {

	if (*r).Method == "OPTIONS" {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	code := r.URL.Query()["token"][0]
	if VerifyJWT(code) {
		returnToken := &Token{Token: code}
		json.NewEncoder(w).Encode(returnToken)
		return
	}
	w.WriteHeader(http.StatusBadRequest)
}

func GetClientVerified(w http.ResponseWriter, r *http.Request) {

	if (*r).Method == "OPTIONS" {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	code := r.URL.Query()["email"][0]
	state := RetrieveClientState(code)
	if state {
		returnToken := &UserVerified{Verified: state, Email: code}
		json.NewEncoder(w).Encode(returnToken)
		return
	}
	returnToken := &UserVerified{Verified: state, Email: code}
	json.NewEncoder(w).Encode(returnToken)
	return
}

func RegisterClientBank(w http.ResponseWriter, r *http.Request) {

	if (*r).Method == "OPTIONS" {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	tokenKey := r.Header.Get("X-RBank-Token")
	var input map[string]interface{}
	json.NewDecoder(r.Body).Decode(&input)
	frontNric := FWDKycAh(input["icFront"].(string))
	frontScanIcNum := (frontNric["vision"].(map[string]interface{})["extract"]).(map[string]interface{})["idNum"].(string)
	frontScanIcBool := (frontNric["vision"].(map[string]interface{})["extract"]).(map[string]interface{})["isIdNum"].(bool)
	// frontScanIcName := (frontNric["vision"].(map[string]interface{})["extract"]).(map[string]interface{})["name"].(string)
	frontScanQuality := (frontNric["qualityCheck"].(map[string]interface{})["finalDecision"]).(bool)

	backNric := FWDKycAh(input["icBack"].(string))
	backScanIcNum := (backNric["vision"].(map[string]interface{})["extract"]).(map[string]interface{})["idNum"].(string)
	backScanIcBool := (backNric["vision"].(map[string]interface{})["extract"]).(map[string]interface{})["isIdNum"].(bool)
	// backScanIcAddress := (backNric["vision"].(map[string]interface{})["extract"]).(map[string]interface{})["address"].(string)
	// backScanIcPostal := (backNric["vision"].(map[string]interface{})["extract"]).(map[string]interface{})["postalCode"].(string)
	backScanQuality := (backNric["qualityCheck"].(map[string]interface{})["finalDecision"]).(bool)
	inputNric := input["nric"].(string)
	inputAddress := input["address"].(string)
	inputPostal := input["postal"].(string)
	inputcountry := input["country"].(string)
	inputFirstName := input["firstName"].(string)
	inputLastName := input["lastName"].(string)
	inputPreferredLanguage := input["preferredLanguage"].(string)
	inputBranchKey := "8a8e878e71c7a4d70171ca4bd43e10a0"
	if (!frontScanQuality && !backScanQuality) || (!frontScanIcBool || !backScanIcBool) || (frontScanIcNum != backScanIcNum) || (inputNric != backScanIcNum) || (frontScanIcNum != inputNric) {
		err := &Error{Error: "Bad NRIC Quality. Please reupload."}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}
	result := MambuNumberFive(inputFirstName, inputLastName, inputPreferredLanguage, inputBranchKey, inputNric, inputAddress, inputPostal, inputcountry)
	email := GetEmailFromJWT(tokenKey)
	mambukey := result["client"].(map[string]interface{})["encodedKey"].(string)
	_, _ = UpdateMambuKey(email, mambukey)
	json.NewEncoder(w).Encode(result)
	return
}

func GetAllSavingssOfClient(w http.ResponseWriter, r *http.Request) {

	if (*r).Method == "OPTIONS" {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	id := r.URL.Query()["mambukey"][0]
	selection := ""
	mambuUser := os.Getenv("")
	mambuPass := os.Getenv("")
	client := http.Client{}
	req, _ := http.NewRequest("GET", selection, nil)
	req.SetBasicAuth(mambuUser, mambuPass)
	resp, err := client.Do(req)
	if err != nil {
		panic(err)

	}
	if resp.StatusCode == http.StatusBadRequest {

	}
	var result []map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	newResult := make([]map[string]interface{}, 0)
	for _, i := range result {
		if i["name"].(string) == "Digital Account" {
			temp := make(map[string]interface{}, 0)
			temp["accountState"] = i["accountState"].(string)
			temp["id"] = i["id"].(string)
			temp["balance"] = i["balance"].(string)
			temp["availableBalance"] = i["availableBalance"].(string)
			temp["overdraftLimit"] = i["overdraftLimit"].(string)
			temp["name"] = i["name"].(string)
			newResult = append(newResult, temp)
		}
	}
	json.NewEncoder(w).Encode(newResult)
}

func GetAllTransactionsOfClient(w http.ResponseWriter, r *http.Request) {

	if (*r).Method == "OPTIONS" {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	id := r.URL.Query()["bankaccountid"][0]
	selection := ""
	mambuUser := os.Getenv("")
	mambuPass := os.Getenv("")
	client := http.Client{}
	req, _ := http.NewRequest("GET", selection, nil)
	req.SetBasicAuth(mambuUser, mambuPass)
	resp, err := client.Do(req)
	if err != nil {
		panic(err)

	}
	if resp.StatusCode == http.StatusBadRequest {

	}
	var result []map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	transactionResult := make([]map[string]interface{}, 0)
	for _, i := range result {
		temp := make(map[string]interface{}, 0)
		temp["amount"] = i["amount"].(string)
		temp["balance"] = i["balance"].(string)
		temp["comment"] = i["comment"].(string)
		temp["entryDate"] = i["entryDate"].(string)
		temp["transactionId"] = i["transactionId"].(float64)
		temp["type"] = i["type"].(string)
		transactionResult = append(transactionResult, temp)

	}
	json.NewEncoder(w).Encode(transactionResult)
}

func DepositeForclient(w http.ResponseWriter, r *http.Request) {

	if (*r).Method == "OPTIONS" {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	id := r.URL.Query()["bankaccountid"][0]
	var input map[string]interface{}
	json.NewDecoder(r.Body).Decode(&input)
	amount := fmt.Sprintf("%f", input["amount"].(float64))
	selection := ""
	var data = []byte(`{
		"amount": ` + amount + `,
		"notes": "Deposit into savings account",
		"type": "DEPOSIT",
		"method": "bank",
		"customInformation": [
			{
				"value": "unique identifier for receipt",
				"customFieldID": "IDENTIFIER_TRANSACTION_CHANNEL_I"
			}
		]
	}`)
	mambuUser := os.Getenv("USERMAMBU")
	mambuPass := os.Getenv("PASSMAMBU")
	client := http.Client{}
	req, _ := http.NewRequest("POST", selection, bytes.NewBuffer(data))
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
	output := make(map[string]interface{})
	output["amountDeposit"] = result["amount"].(string)
	output["newBalance"] = result["balance"].(string)
	json.NewEncoder(w).Encode(output)
}

func GetAccountDetails(w http.ResponseWriter, r *http.Request) {

	if (*r).Method == "OPTIONS" {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	id := r.URL.Query()["bankaccountid"][0]
	selection := ""
	mambuUser := os.Getenv("")
	mambuPass := os.Getenv("")
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
	output := make(map[string]interface{}, 0)
	output["id"] = result["id"].(string)
	output["accountState"] = result["accountState"].(string)
	output["availableBalance"] = result["availableBalance"].(string)
	output["balance"] = result["balance"].(string)
	output["name"] = result["name"].(string)
	output["overdraftLimit"] = result["overdraftLimit"].(string)
	json.NewEncoder(w).Encode(output)
}

func CreateSavingsAccountsOfClient(w http.ResponseWriter, r *http.Request) {

	if (*r).Method == "OPTIONS" {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	id := r.URL.Query()["mambukey"][0]
	selection := ""
	var data = []byte(`{
		"savingsAccount": {
			"name": "Digital Account",
			"accountHolderType": "CLIENT",
			"accountHolderKey": "` + id + `",
			"accountState": "APPROVED",
			"productTypeKey": "",
			"accountType": "CURRENT_ACCOUNT",
			"currencyCode": "SGD",
			"allowOverdraft": "true",
			"overdraftLimit": "500",
			"overdraftInterestSettings": {
				"interestRate": 5
			},
				"interestSettings": {
			"interestRate": "1.25"
		}
		}
	
	}`)
	mambuUser := os.Getenv("USERMAMBU")
	mambuPass := os.Getenv("PASSMAMBU")
	client := http.Client{}
	req, _ := http.NewRequest("POST", selection, bytes.NewBuffer(data))
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
	json.NewEncoder(w).Encode(result)

}

func TransferBetweenClients(w http.ResponseWriter, r *http.Request) {
	if (*r).Method == "OPTIONS" {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	id := r.URL.Query()["bankaccountid"][0]
	var input map[string]interface{}
	json.NewDecoder(r.Body).Decode(&input)
	amount := fmt.Sprintf("%f", input["amount"].(float64))
	selection := ""
	var data = []byte(`{
		"type": "TRANSFER",
		"amount": "` + amount + `",
		"notes": "Transfer to Expenses Account",
		"toSavingsAccount": "` + input["target"].(string) + `",
		"method":"bank"
	}`)
	mambuUser := os.Getenv("")
	mambuPass := os.Getenv("")
	client := http.Client{}
	req, _ := http.NewRequest("POST", selection, bytes.NewBuffer(data))
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
	output := make(map[string]interface{}, 0)
	output["amount"] = result["amount"].(string)
	output["balance"] = result["balance"].(string)
	output["transactionId"] = result["transactionId"].(float64)
	output["transferTo"] = input["target"].(string)
	output["entryDate"] = result["entryDate"].(string)
	json.NewEncoder(w).Encode(result)
}

func CreateLoanAccountsOfClient(w http.ResponseWriter, r *http.Request) {

	if (*r).Method == "OPTIONS" {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	id := r.URL.Query()["mambukey"][0]
	var input map[string]interface{}
	json.NewDecoder(r.Body).Decode(&input)
	amount := fmt.Sprintf("%f", input["amount"].(float64))
	selection := ""
	var data = []byte(`{
		"loanAccount": {
			"accountHolderType": "CLIENT",
			"accountHolderKey": "` + id + `",
			"productTypeKey": "8a8e867271bd280c0171bf768cc31a89",
			"assignedBranchKey": "8a8e878e71c7a4d70171ca4bd43e10a0",
			"loanName": "Student Loan",
			"loanAmount": ` + amount + `,
			"interestRate": "2",
			"arrearsTolerancePeriod": "0",
			"gracePeriod": "0",
			"repaymentInstallments": "10",
			"repaymentPeriodCount": "1",
			"periodicPayment": "0",
			"repaymentPeriodUnit": "WEEKS",
			"disbursementDetails": {
				"customInformation": [
					{
						"value": "unique identifier for this transaction",
						"customFieldID": "IDENTIFIER_TRANSACTION_CHANNEL_I"
					}
				]
			}
		}
	}`)
	mambuUser := os.Getenv("")
	mambuPass := os.Getenv("")
	client := http.Client{}
	req, _ := http.NewRequest("POST", selection, bytes.NewBuffer(data))
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
	
	loanId := result["loanAccount"].(map[string]interface{})["id"].(string)
	disburseURL := ""
	sclient := http.Client{}
	var sdata = []byte(`{
		"type": "DISBURSEMENT",
		"method":"bank",
		"customInformation": [
			{
				"value": "unique identifier for transaction",
				"customFieldID": "IDENTIFIER_TRANSACTION_CHANNEL_I"
			}
		]
	}`)
	sreq, _ := http.NewRequest("POST", disburseURL, bytes.NewBuffer(sdata))
	sreq.SetBasicAuth(mambuUser, mambuPass)
	sresp, serr := sclient.Do(sreq)
	if serr != nil {
		panic(serr)

	}
	if sresp.StatusCode == http.StatusBadRequest {

	}
	output := make(map[string]interface{}, 0)
	output["loanId"] = loanId
	output["loanAmount"] = input["amount"].(float64)
	json.NewEncoder(w).Encode(output)
}

func GetAllLoansOfClient(w http.ResponseWriter, r *http.Request) {

	if (*r).Method == "OPTIONS" {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	id := r.URL.Query()["mambukey"][0]
	selection := ""
	mambuUser := os.Getenv("")
	mambuPass := os.Getenv("")
	client := http.Client{}
	req, _ := http.NewRequest("GET", selection, nil)
	req.SetBasicAuth(mambuUser, mambuPass)
	resp, err := client.Do(req)
	if err != nil {
		panic(err)

	}
	if resp.StatusCode == http.StatusBadRequest {

	}
	var result []map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	output := make([]map[string]interface{}, 0)
	for _, i := range result {
		temp := make(map[string]interface{}, 0)
		temp["accountState"] = i["accountState"].(string)
		temp["id"] = i["id"].(string)
		temp["loanName"] = i["loanName"].(string)
		temp["loanAmount"] = i["loanAmount"].(string)
		temp["feesDue"] = i["feesDue"].(string)
		temp["feesPaid"] = i["feesPaid"].(string)
		temp["feesBalance"] = i["feesBalance"].(string)
		temp["principalBalance"] = i["principalBalance"].(string)
		temp["principalDue"] = i["principalDue"].(string)
		temp["principalPaid"] = i["principalPaid"].(string)
		output = append(output, temp)
	}
	json.NewEncoder(w).Encode(output)
}

func GetLoanDetails(w http.ResponseWriter, r *http.Request) {

	if (*r).Method == "OPTIONS" {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	id := r.URL.Query()["loanaccountid"][0]
	selection := ""
	mambuUser := os.Getenv("")
	mambuPass := os.Getenv("")
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
	temp := make(map[string]interface{}, 0)
	temp["accountState"] = result["accountState"].(string)
	temp["id"] = result["id"].(string)
	temp["loanName"] = result["loanName"].(string)
	temp["loanAmount"] = result["loanAmount"].(string)
	temp["feesDue"] = result["feesDue"].(string)
	temp["feesPaid"] = result["feesPaid"].(string)
	temp["feesBalance"] = result["feesBalance"].(string)
	temp["principalBalance"] = result["principalBalance"].(string)
	temp["principalDue"] = result["principalDue"].(string)
	temp["principalPaid"] = result["principalPaid"].(string)
	json.NewEncoder(w).Encode(temp)
}

func CreateDepositAccountsOfClient(w http.ResponseWriter, r *http.Request) {

	if (*r).Method == "OPTIONS" {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	id := r.URL.Query()["mambukey"][0]
	selection := ""
	var data = []byte(`{
		"savingsAccount" : {
		"notes": "",
		"name": "Fixed Deposit Account",
		"accountHolderType": "CLIENT",
		"accountHolderKey": "` + id + `",
		"accountState": "APPROVED",
		"productTypeKey": "",
		"assignedBranchKey": "",
		"accountType": "FIXED_DEPOSIT",
		"currencyCode": "SGD",
		"interestSettings": {
			"interestRate": "2"
			}
		}
	}`)
	mambuUser := os.Getenv("")
	mambuPass := os.Getenv("")
	client := http.Client{}
	req, _ := http.NewRequest("POST", selection, bytes.NewBuffer(data))
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
	json.NewEncoder(w).Encode(result)
}

func GetAllDepositOfClient(w http.ResponseWriter, r *http.Request) {

	if (*r).Method == "OPTIONS" {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	id := r.URL.Query()["mambukey"][0]
	selection := ""
	mambuUser := os.Getenv("")
	mambuPass := os.Getenv("")
	client := http.Client{}
	req, _ := http.NewRequest("GET", selection, nil)
	req.SetBasicAuth(mambuUser, mambuPass)
	resp, err := client.Do(req)
	if err != nil {
		panic(err)

	}
	if resp.StatusCode == http.StatusBadRequest {

	}
	var result []map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	newResult := make([]map[string]interface{}, 0)
	for _, i := range result {
		if i["name"].(string) == "Fixed Deposit Account" {
			temp := make(map[string]interface{}, 0)
			temp["accountState"] = i["accountState"].(string)
			temp["id"] = i["id"].(string)
			temp["name"] = i["name"].(string)
			temp["balance"] = i["balance"].(string)
			temp["availableBalance"] = i["availableBalance"].(string)
			temp["interestRate"] = (i["interestSettings"].(map[string]interface{})["interestRate"]).(string)
			newResult = append(newResult, temp)
		}
	}
	json.NewEncoder(w).Encode(newResult)
}

func GetDepositDetails(w http.ResponseWriter, r *http.Request) {

	if (*r).Method == "OPTIONS" {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	id := r.URL.Query()["depositid"][0]
	selection := ""
	mambuUser := os.Getenv("")
	mambuPass := os.Getenv("")
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
	temp := make(map[string]interface{}, 0)
	temp["accountState"] = result["accountState"].(string)
	temp["id"] = result["id"].(string)
	temp["name"] = result["name"].(string)
	temp["balance"] = result["balance"].(string)
	temp["availableBalance"] = result["availableBalance"].(string)
	temp["interestRate"] = (result["interestSettings"].(map[string]interface{})["interestRate"]).(string)
	json.NewEncoder(w).Encode(temp)
}

func GetAllAccountsOfClient(w http.ResponseWriter, r *http.Request) {

	if (*r).Method == "OPTIONS" {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	id := r.URL.Query()["mambukey"][0]
	selection := ""
	mambuUser := os.Getenv("")
	mambuPass := os.Getenv("")
	client := http.Client{}
	req, _ := http.NewRequest("GET", selection, nil)
	req.SetBasicAuth(mambuUser, mambuPass)
	resp, err := client.Do(req)
	if err != nil {
		panic(err)

	}
	if resp.StatusCode == http.StatusBadRequest {

	}
	var result []map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	savingsResult := make([]map[string]interface{}, 0)
	depositResult := make([]map[string]interface{}, 0)
	for _, i := range result {
		if i["name"].(string) == "Digital Account" {
			temp := make(map[string]interface{}, 0)
			temp["accountState"] = i["accountState"].(string)
			temp["id"] = i["id"].(string)
			temp["balance"] = i["balance"].(string)
			temp["availableBalance"] = i["availableBalance"].(string)
			temp["overdraftLimit"] = i["overdraftLimit"].(string)
			temp["name"] = i["name"].(string)
			savingsResult = append(savingsResult, temp)
		} else if i["name"].(string) == "Fixed Deposit Account" {
			temp := make(map[string]interface{}, 0)
			temp["accountState"] = i["accountState"].(string)
			temp["id"] = i["id"].(string)
			temp["name"] = i["name"].(string)
			temp["balance"] = i["balance"].(string)
			temp["availableBalance"] = i["availableBalance"].(string)
			temp["interestRate"] = (i["interestSettings"].(map[string]interface{})["interestRate"]).(string)
			depositResult = append(depositResult, temp)
		}
	}

	sselection := ""
	sclient := 
	
	.Client{}
	sreq, _ := http.NewRequest("GET", sselection, nil)
	sreq.SetBasicAuth(mambuUser, mambuPass)
	sresp, serr := sclient.Do(sreq)
	if serr != nil {
		panic(serr)

	}
	if sresp.StatusCode == http.StatusBadRequest {

	}
	var sresult []map[string]interface{}
	json.NewDecoder(sresp.Body).Decode(&sresult)
	loansResult := make([]map[string]interface{}, 0)
	for _, i := range sresult {
		temp := make(map[string]interface{}, 0)
		temp["accountState"] = i["accountState"].(string)
		temp["id"] = i["id"].(string)
		temp["loanName"] = i["loanName"].(string)
		temp["loanAmount"] = i["loanAmount"].(string)
		temp["feesDue"] = i["feesDue"].(string)
		temp["feesPaid"] = i["feesPaid"].(string)
		temp["feesBalance"] = i["feesBalance"].(string)
		temp["principalBalance"] = i["principalBalance"].(string)
		temp["principalDue"] = i["principalDue"].(string)
		temp["principalPaid"] = i["principalPaid"].(string)
		loansResult = append(loansResult, temp)

	}
	outputResult := make(map[string]interface{}, 0)
	outputResult["savings"] = savingsResult
	outputResult["deposit"] = depositResult
	outputResult["loans"] = loansResult
	json.NewEncoder(w).Encode(outputResult)
}

func CreateTypesAccountsOfClient(w http.ResponseWriter, r *http.Request) {

	if (*r).Method == "OPTIONS" {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	id := r.URL.Query()[""][0]
	accountType := r.URL.Query()["create"][0]

	if accountType == "savings" {
		selection := "" + id
		client := http.Client{}
		req, _ := http.NewRequest("POST", selection, r.Body)
		req.Header.Set("Content-Type", "application/json")
		resp, err := client.Do(req)
		if err != nil {
			panic(err)

		}
		if resp.StatusCode == http.StatusBadRequest {

		}
		var result map[string]interface{}
		json.NewDecoder(resp.Body).Decode(&result)
		json.NewEncoder(w).Encode(result)
		return
	} else if accountType == "deposit" {
		selection := "" + id
		client := http.Client{}
		req, _ := http.NewRequest("POST", selection, r.Body)
		req.Header.Set("Content-Type", "application/json")
		resp, err := client.Do(req)
		if err != nil {
			panic(err)

		}
		if resp.StatusCode == http.StatusBadRequest {

		}
		var result map[string]interface{}
		json.NewDecoder(resp.Body).Decode(&result)
		json.NewEncoder(w).Encode(result)
		return
	} else if accountType == "loans" {
		selection := "" + id
		client := http.Client{}
		req, _ := http.NewRequest("POST", selection, r.Body)
		req.Header.Set("Content-Type", "application/json")
		resp, err := client.Do(req)
		if err != nil {
			panic(err)

		}
		if resp.StatusCode == http.StatusBadRequest {

		}
		var result map[string]interface{}
		json.NewDecoder(resp.Body).Decode(&result)
		json.NewEncoder(w).Encode(result)
		return
	} else {

	}

}

func DepositMoneyToFDAccount(w http.ResponseWriter, r *http.Request) {

	if (*r).Method == "OPTIONS" {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	id := r.URL.Query()["depositid"][0]
	var input map[string]interface{}
	json.NewDecoder(r.Body).Decode(&input)
	amount := fmt.Sprintf("%f", input["amount"].(float64))
	selection := ""
	var data = []byte(`{
		"amount": ` + amount + `,
		"notes": "Deposit into Fixed Deposit account",
		"type": "DEPOSIT",
		"method": "bank",
		"customInformation": [
			{
				"value": "unique identifier for receipt",
				"customFieldID": "IDENTIFIER_TRANSACTION_CHANNEL_I"
			}
		]
	}`)
	mambuUser := os.Getenv("")
	mambuPass := os.Getenv("")
	client := http.Client{}
	req, _ := http.NewRequest("POST", selection, bytes.NewBuffer(data))
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
	json.NewEncoder(w).Encode(result)
}

func ServeService() http.Handler {
	router := mux.NewRouter()
	router.HandleFunc("/auth/xx/login", RedirectToRazer)
	router.HandleFunc("/auth/xx/callback", Callback)
	router.HandleFunc("/auth/verify", Authorization)
	router.HandleFunc("/bank/client/register", RegisterClientBank).Methods("POST", "OPTIONS")
	router.HandleFunc("/bank/client/verify", GetClientVerified).Methods("GET", "OPTIONS")
	router.HandleFunc("/bank/client/savings/all", GetAllSavingssOfClient).Methods("GET", "OPTIONS")
	router.HandleFunc("/bank/client/savings", TransferBetweenClients).Methods("PUT", "OPTIONS")
	router.HandleFunc("/bank/client/savings/transaction/all", GetAllTransactionsOfClient).Methods("GET", "OPTIONS")
	router.HandleFunc("/bank/client/savings/transaction", DepositeForclient).Methods("POST", "OPTIONS")
	router.HandleFunc("/bank/client/savings/transaction", GetAccountDetails).Methods("GET", "OPTIONS")
	router.HandleFunc("/mambu", Mambu).Methods("GET", "OPTIONS")
	router.HandleFunc("/fwd", FWDKyc).Methods("POST", "OPTIONS")

	router.HandleFunc("/bank/client/loan", GetLoanDetails).Methods("GET", "OPTIONS")
	router.HandleFunc("/bank/client/loan/all", GetAllLoansOfClient).Methods("GET", "OPTIONS")

	router.HandleFunc("/bank/client/deposit", GetDepositDetails).Methods("GET", "OPTIONS")
	router.HandleFunc("/bank/client/deposit", DepositMoneyToFDAccount).Methods("PUT", "OPTIONS")
	router.HandleFunc("/bank/client/deposit/all", GetAllDepositOfClient).Methods("GET", "OPTIONS")

	router.HandleFunc("/bank/client/accounts/all", GetAllAccountsOfClient).Methods("GET", "OPTIONS")

	router.HandleFunc("/bank/client/loan", CreateLoanAccountsOfClient).Methods("POST", "OPTIONS")
	router.HandleFunc("/bank/client/deposit", CreateDepositAccountsOfClient).Methods("POST", "OPTIONS")
	router.HandleFunc("/bank/client/savings", CreateSavingsAccountsOfClient).Methods("POST", "OPTIONS")
	router.HandleFunc("/bank/client/accounts", CreateTypesAccountsOfClient).Methods("POST", "OPTIONS")
	return router
}

func main() {
	http.ListenAndServe(":8000", ServeService())
}
