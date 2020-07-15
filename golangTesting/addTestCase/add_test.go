package main

import (
	"io/ioutil"
	"testing"
)

// Test Driven Test Cases in Golang

type AddResult struct{
	x int
	y int
	expected int
}

var addResults = []AddResult{
	{1, 1,2},
	{2,3,5},
}

func TestAdd(t *testing.T){
	for _, test:= range addResults{
		result := Add(test.x, test.y)
		if result != test.expected {
			t.Fatal("Expected result not given")
		}
	}
}
func TestStringData(t *testing.T){
	data, err:= ioutil.ReadFile("testdata/test.data")
	if err != nil {
		t.Fatal("Could not Open the File")
	}
	if string(data) != "key = value" {
		t.Fatal("Expected Text not Matched")
	}
}
