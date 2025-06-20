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

PS C:\Users\faizm\Desktop\Jsonb Postgresql> go run cmd/main.go
✅ .env loaded
✅ Connected and AutoMigrated PostgreSQL
✅ Connected to PostgreSQL in 87.1631ms
✅ Connected to MongoDB
✅ Connected to MongoDB in 19.736ms
🔗 Total DB connection setup time: 107.4083ms
✅ Generated 100 dummy records for benchmarking
📝 PostgreSQL Insert Time: 435.8252ms
📝 MongoDB Insert Time: 2.5203ms
🔍 PostgreSQL Read Time: 3.1459ms
🔍 MongoDB Read Time: 3.7764ms
✏️ PostgreSQL Update Time: 5.1816ms
✏️ MongoDB Update Time: 3.2175ms
🗑️ PostgreSQL Delete Time: 4.3106ms
🗑️ MongoDB Delete Time: 2.6258ms
🏁 Benchmark tests completed!
