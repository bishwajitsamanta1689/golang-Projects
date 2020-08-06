package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

/*
	Golang Get Requests
	Flow Step #1 - Get the URL
         Step #2 - Request for the URL using New Request Function
         Step #3 - Get a Response Back from the URL
		 Step #4 - Take the Response to a Body
		 Step #5 - Converts the bytes data to a String Format.
 */

func main(){
	url:= "https://reqres.in/api/users"
	req, _:= http.NewRequest("GET",url,nil)
	res, _:= http.DefaultClient.Do(req)
	defer res.Body.Close()
	body,_:= ioutil.ReadAll(res.Body)
	fmt.Println (string(body))
}
