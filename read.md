Running Postgres in a docker container

1.docker pull postgres  
2.docker run --name pg-json-test -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=pass -e POSTGRES_DB=testdb -p 5432:5432 -d postgres
3.docker ps
4.(optional) docker exec -it pg-json-test psql -U postgres -d testdb  
4.1. \l -> list all databases

Running Mongo in a docker container
1.docker run --name mongo-json-test -p 27017:27017 -d mongo
2.docker ps
3.(optional) docker exec -it mongo-json-test mongo testdb
3.1. show dbs

Run the app

1. go run main.go
