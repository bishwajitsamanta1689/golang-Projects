package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Azure struct{
	Database struct {
		Host string `json:"host"`
		Port string `json:"port"`
	}`json:"database"`
	Credentials struct {
		StorageAccountName string `json:"storage_account_name"`
		SecondaryAccessKey string `json:"secondary_access_key"`
	}`json:"credentials"`
	Secrets []struct {
		Labels string `json:"labels"`
		Azure struct {
			StorageAccount string `json:"storage_account_name"`
			PrimaryAccessKey string `json:"primary_access_key"`
			SecondaryAccessKey string `json:"secondary_access_key"`
		}`json:"Azure"`
	}`json:"secrets"`
}

func LoadConfiguration(filename string) (Azure,error) {

	var keys Azure
	file, err:= os.Open(filename)
	if err != nil {
		panic(err)
	}
	jsonParser:= json.NewDecoder(file)
	err = jsonParser.Decode(&keys)
	return keys, err
}

func main(){
	config, _:= LoadConfiguration("filetointerpolate.json")
	fmt.Println("Database Host=", config.Database.Host)
	fmt.Println("Secondary Access Key=", config.Credentials.SecondaryAccessKey)
	array := config.Secrets
	for _, keysStruct:= range array{
		fmt.Println("label=", keysStruct.Labels)
	}
	for _, keys:= range array{
		fmt.Println("Primary Access Key=", keys.Azure.PrimaryAccessKey)
	}

}
