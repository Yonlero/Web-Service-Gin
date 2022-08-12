# Web-Service-Gin
    I'm creating this API just to learn more about GoLang and the Gin library.
    And i'm tryng to bring java concepts to GoLang like the way to organize code and files

# Commands to setup the docker
- In the project root folder
    * docker-compose -f docker-compose.yml -p "Database_Containers" up -d

# Commands to setup database
- In the Db_Postgres container you'll run ( You can access the container shell if you run this in your terminal: docker exec -it Db_Postgres /bin/bash )
    * psql -d Db_Gin -U root -h localhost -W ( your password )
    * CREATE EXTENSION IF NOT EXISTS "uuid-ossp"; ( let's create the extension for UUID )
    * CREATE TABLE tb_albums (id uuid DEFAULT uuid_generate_v4 (), title VARCHAR NOT NULL, artist VARCHAR NOT NULL, price float NOT NULL, PRIMARY KEY (id));

# Commands to Run the API
- In the project root folder
    * go run .
- In the test folder
    * go test -v

# Endpoints
 - GET - http://localhost:8080/albums ( Get all albums )
 - GET - http://localhost:8080/albums/id ( Get an album by Id )
 - POST - http://localhost:8080/albums ( Create a new album )
 - PUT - http://localhost:8080/albums ( Update an existent album )
 - DELETE - http://localhost:8080/albums/id ( Delete an album )