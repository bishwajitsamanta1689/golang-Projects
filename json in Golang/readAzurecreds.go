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
	}`json:"Database"`
	Credentials struct {
		StorageAccountName string `json:"storage_account_name"`
		SecondaryAccessKey string `json:"secondary_access_key"`
	}`json:credentials`
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
	fmt.Println("Secondary Access Key=", config.Credentials.SecondaryAccessKey)
}
