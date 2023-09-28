package rest

import (
	"display-name-updater/internal/mapper"
	"display-name-updater/internal/models"
	"encoding/base64"
	"fmt"
	"github.com/joho/godotenv"
	"io"
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

func Update(clients []models.ClientDisplayNameData, fieldName string) {
	responseState = make([]string, len(clients)-1)
	for _, clientData := range clients {

		// fetch existing client data from NW
		existingNwClient := Get(clientData.ClientId)
		if existingNwClient == nil {
			continue
		}
		// update display name
		existingNwClient.DisplayName = clientData.AdditionalInformation[fieldName]
		// body construction
		buffer := mapper.ToBytes(*existingNwClient, clientData)
		if buffer == nil {
			continue
		}
		request, err := http.NewRequest(http.MethodPut, *nwUrl+clientData.ClientId, buffer)
		setCredentials(request, encodedCreds)
		setHeaders(request)
		// update client data with new display name
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

func Get(clientId string) *models.NwClient {
	request, err := http.NewRequest(http.MethodGet, *nwUrl+clientId, nil)
	setCredentials(request, encodedCreds)
	setHeaders(request)

	response, err := httpClient.Do(request)
	if response == nil || err != nil || response.StatusCode != http.StatusOK {
		fmt.Printf("ClientData data could not be updated. Accepted response %v\n", response)
		return nil
	}
	var reader = &response.Body
	defer closeReader(reader)
	body, err := io.ReadAll(*reader)
	if err != nil {
		fmt.Println("Could not read response body")
		return nil
	}
	return mapper.MapToNwClient(body, err)
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

func closeReader(body *io.ReadCloser) {
	err := (*body).Close()
	if err != nil {
		fmt.Println("Could not close response body")
	}
}
