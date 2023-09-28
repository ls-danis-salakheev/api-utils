package mapper

import (
	"bytes"
	. "display-name-updater/internal/models"
	"encoding/json"
	"fmt"
)

func MapToNwClient(body []byte, err error) *NwClient {
	if body != nil {
		var res NwClient
		err = json.Unmarshal(body, &res)
		if err != nil {
			fmt.Printf("Could not parse the body = %v\n, err = %v\n", body, err)
			return nil
		}
		return &res
	}
	return nil
}

func ToBytes(existingNwClient NwClient, newClientData ClientDisplayNameData) *bytes.Buffer {
	jsoned, err := json.Marshal(existingNwClient)
	if err != nil {
		fmt.Printf("Could not jsoned clientData = %v\n, err = %v\n", newClientData, err)
		return nil
	}
	return bytes.NewBuffer(jsoned)
}
