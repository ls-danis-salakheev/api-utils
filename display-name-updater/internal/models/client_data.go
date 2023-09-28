package models

import "strings"

func CreateClientArr(lines [][]string, offset int, clientCount int) []ClientDisplayNameData {
	clientDataArr := make([]ClientDisplayNameData, clientCount-offset)
	for i, line := range lines {
		if i == 0 {
			continue
		}
		newClientData := newClient(line)
		clientDataArr[i-offset] = newClientData
	}
	return clientDataArr
}

func newClient(csvLine []string) ClientDisplayNameData {
	return ClientDisplayNameData{
		ClientId: strings.TrimSpace(csvLine[0]),
		AdditionalInformation: map[string]string{
			"displayName": strings.TrimSpace(csvLine[1]),
		},
	}
}

type ClientDisplayNameData struct {
	ClientId              string            `json:"clientId"`
	AdditionalInformation map[string]string `json:"additional_information"`
}
