# Using Go Modules

* List of Steps for using Go Modules
    * Create a empty Directory
    * go mod init github.com/bishwajitsamanta1689/go-modules
    * go get github.com/aws/aws-sdk-go
    * go get go.uber.org/zap (Logging Purpose)
    
* Storing Credentials
   * aws-vault is a great way to store credentials
   * aws-vault add LA
        * Here LA is the profile name
        * The data will get stored under home directory/aws/config