# Testing in Golang

## Testing Terms

* The name of the Test should have the same package name
* The name of the Test Program should have the name ending with the main program.
    * Example:- Program Name=> sum.go 
    * Test Name:- sum_test.go
* Package name for Testing.
    * import testing
* Function name for testing should start with `Test`
|Function|Description|
|-|-|
|`t.Errorf`|It Allows the Other Test cases to Run in case one gets failed|
|`t.Fatalf`|It stops Execution of other Test cases|