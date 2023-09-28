package rest

import (
	"bytes"
	"display-name-updater/internal/models"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"net/http"
	"os"
)

var httpClient *http.Client

// env vars
var nwUrl *string
var encodedCreds string
var responseState []string

func init() {
	envErr := godotenv.Load()
	if envErr != nil {
		fmt.Println("Cannot load env vars")
		return
	}
	httpClient = &http.Client{}
	nwUrl = checkAndHandlePath(os.Getenv("NW_URL"))
	encodedCreds = prepareCreds()
}

func Update(clients []models.ClientDisplayNameData) {
	responseState = make([]string, len(clients)-1)
	for _, clientData := range clients {
		jsoned, err := json.Marshal(clientData)
		if err != nil {
			fmt.Printf("Could not jsoned clientData = %v\n", clientData)
			continue
		}

		buffer := bytes.NewBuffer(jsoned)
		putUrl := *nwUrl + clientData.ClientId
		request, err := http.NewRequest(http.MethodPut, putUrl, buffer)
		setCredentials(request, encodedCreds)
		setHeaders(request)

		response, err := httpClient.Do(request)
		if response == nil || err != nil || response.StatusCode != http.StatusOK {
			fmt.Printf("ClientData data could not be updated. Accepted response %v\n", response)
		}
		updateStateArr(response, clientData)
		fmt.Println("=====================================")
		fmt.Printf("Updated clients state: %v\n", responseState)
		fmt.Println("=====================================")
	}
}

func updateStateArr(response *http.Response, clientData models.ClientDisplayNameData) {
	var message string
	if response == nil {
		message = "response is nil"
	} else {
		message = response.Status
	}

	responseState = append(responseState, "response code = "+message+" for client id = "+clientData.ClientId)
}

func setCredentials(request *http.Request, encodedCreds string) {
	request.Header.Set("Authorization", "Basic "+encodedCreds)
}

func setHeaders(req *http.Request) {
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
}

func prepareCreds() string {
	adminName := os.Getenv("NW_ADMIN")
	adminPass := os.Getenv("NW_PASSWORD")
	data := []byte(adminName + ":" + adminPass)
	return base64.StdEncoding.EncodeToString(data)
}

func checkAndHandlePath(nwUrl string) *string {
	char := nwUrl[len(nwUrl)-1]
	slash := "/"
	if string(char) != slash {
		res := nwUrl + slash
		return &res
	}
	return &nwUrl
}
