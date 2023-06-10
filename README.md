# go-simple-template
This is a simple template/boilerplate for creating services in Go. It provides a basic structure and common functionalities to kickstart your service development.

### I use existing libs :
- [Godotenv](github.com/joho/godotenv) for read .env config
-	[Echo](github.com/labstack/echo/v4) for routing framework
-	[Zerolog](github.com/rs/zerolog) for logging
-	[Testify](github.com/stretchr/testify) for unit testing
-	[GORM](gorm.io/gorm) for database operation

### How To Use?
1. Begin by cloning the repository to your local machine.
2. Set up your database.
3. Create an environment file by duplicating the .env.example file and filling in the required values for each service according to your needs.
4. Install the necessary dependencies.
```
go mod tidy
```
5. Run the program
```
go run *.go
```
6. Test liveness
```
curl --location 'localhost:8080/ping'
````
