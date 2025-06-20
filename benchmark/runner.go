package benchmark

import (
	"Jsonb/data"
	"Jsonb/internal/db"
	"fmt"
	"time"
)

func RunBenchmark() {
	records := data.GenerateDummyData(10000)
	fmt.Println("✅ Generated 10000 dummy records for benchmarking")

	// 🔍 Benchmark PostgreSQL Insert
	pgInsertStart := time.Now()
	if err := db.InsertPostgres(records); err != nil {
		fmt.Println("❌ Postgres insert error:", err)
	} else {
		fmt.Println("📝 PostgreSQL Insert Time:", time.Since(pgInsertStart))
	}

	// 🔍 Benchmark MongoDB Insert
	mongoInsertStart := time.Now()
	if err := db.InsertMongo(records); err != nil {
		fmt.Println("❌ Mongo insert error:", err)
	} else {
		fmt.Println("📝 MongoDB Insert Time:", time.Since(mongoInsertStart))
	}

	// 🔍 Benchmark PostgreSQL Read
	pgReadStart := time.Now()
	_, err := db.FindPostgresByID(10)
	if err != nil {
		fmt.Println("❌ Postgres read error:", err)
	} else {
		fmt.Println("🔍 PostgreSQL Read Time:", time.Since(pgReadStart))
	}

	// 🔍 Benchmark MongoDB Read
	mongoReadStart := time.Now()
	_, err = db.FindMongoByID(10)
	if err != nil {
		fmt.Println("❌ Mongo read error:", err)
	} else {
		fmt.Println("🔍 MongoDB Read Time:", time.Since(mongoReadStart))
	}

	// 🔍 Benchmark PostgreSQL Update
	pgUpdateStart := time.Now()
	if err := db.UpdatePostgresCity(30, "BenchCity"); err != nil {
		fmt.Println("❌ Postgres update error:", err)
	} else {
		fmt.Println("✏️ PostgreSQL Update Time:", time.Since(pgUpdateStart))
	}

	// 🔍 Benchmark MongoDB Update
	mongoUpdateStart := time.Now()
	if err := db.UpdateMongoCity(30, "BenchCity"); err != nil {
		fmt.Println("❌ Mongo update error:", err)
	} else {
		fmt.Println("✏️ MongoDB Update Time:", time.Since(mongoUpdateStart))
	}

	// 🔍 Benchmark PostgreSQL Delete
	pgDeleteStart := time.Now()
	if err := db.DeletePostgresByAge(20); err != nil {
		fmt.Println("❌ Postgres delete error:", err)
	} else {
		fmt.Println("🗑️ PostgreSQL Delete Time:", time.Since(pgDeleteStart))
	}

	// 🔍 Benchmark MongoDB Delete
	mongoDeleteStart := time.Now()
	if err := db.DeleteMongoByAge(20); err != nil {
		fmt.Println("❌ Mongo delete error:", err)
	} else {
		fmt.Println("🗑️ MongoDB Delete Time:", time.Since(mongoDeleteStart))
	}
}
