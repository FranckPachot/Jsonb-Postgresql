package main

import (
	"Jsonb/internal/db"
	"fmt"
	"log"
	"time"

	"github.com/joho/godotenv"
)

func init() {
	start := time.Now()

	// Load .env
	if err := godotenv.Load(); err != nil {
		log.Fatal("❌ Error loading .env file")
	}
	fmt.Println("✅ .env loaded")

	// PostgreSQL Connection
	pgStart := time.Now()
	if err := db.ConnectPostgres(); err != nil {
		log.Fatalf("❌ Postgres connection failed: %v", err)
	}
	fmt.Println("✅ Connected to PostgreSQL in", time.Since(pgStart))

	// MongoDB Connection
	mongoStart := time.Now()
	if err := db.ConnectMongo(); err != nil {
		log.Fatalf("❌ MongoDB connection failed: %v", err)
	}
	fmt.Println("✅ Connected to MongoDB in", time.Since(mongoStart))

	fmt.Println("🔗 Total DB connection setup time:", time.Since(start))
}

func main() {
	// records := data.GenerateDummyData(100)
	// fmt.Println("✅ Generated 100 dummy records")

	// // Insert into PostgreSQL
	// pgInsertStart := time.Now()
	// if err := db.InsertPostgres(records); err != nil {
	// 	log.Fatal("❌ Postgres insert error:", err)
	// }
	// fmt.Println("📝 Inserted 100 records into PostgreSQL in", time.Since(pgInsertStart))

	// // Insert into MongoDB
	// mongoInsertStart := time.Now()
	// if err := db.InsertMongo(records); err != nil {
	// 	log.Fatal("❌ Mongo insert error:", err)
	// }
	// fmt.Println("📝 Inserted 100 records into MongoDB in", time.Since(mongoInsertStart))
	fmt.Println("🏁 All records inserted successfully!")
}
