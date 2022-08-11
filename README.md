# Web-Service-Gin
    I'm creating this API just to learn more about GoLang and the Gin library.
    I'm tryng bring java concepts to GoLang like the way to organize code and files
# Commands to Run the API
- In the project root folder
 - go run .
- In the test folder
 - go test -v

# Command to setup the docker
- In the project root folder
 - docker-compose -f docker-compose.yml -p "Database_Containers" up -d

# Endpoints
 - GET - http://localhost:8080/albums ( Get all albums )
 - GET - http://localhost:8080/albums/id ( Get an album by Id )
 - POST - http://localhost:8080/albums ( Create a new album )
 - PUT - http://localhost:8080/albums ( Update an existent album )
 - DELETE - http://localhost:8080/albums/id ( Delete an album )