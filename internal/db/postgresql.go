package db

import (
	"Jsonb/data"
	"encoding/json"
	"fmt"
	"os"

	"gorm.io/datatypes"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var PostgresDB *gorm.DB

// Model
type JSONData struct {
	ID   uint           `gorm:"primaryKey"`
	Data datatypes.JSON `gorm:"type:jsonb"`
}

func ConnectPostgres() error {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
		os.Getenv("POSTGRES_PORT"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	if err := db.AutoMigrate(&JSONData{}); err != nil {
		return err
	}

	sql := `CREATE INDEX IF NOT EXISTS idx_json_data_gin ON json_data USING GIN (data);`
	if err := db.Exec(sql).Error; err != nil {
		return err
	}

	fmt.Println("✅ GIN index on JSONB 'data' created (or already exists)")

	PostgresDB = db
	fmt.Println("✅ Connected and AutoMigrated PostgreSQL")
	return nil
}

// Create Single Record
func CreatePostgres(record data.DummyData) error {
	jsonBytes, err := json.Marshal(record)
	if err != nil {
		return err
	}

	item := JSONData{
		Data: datatypes.JSON(jsonBytes),
	}

	return PostgresDB.Create(&item).Error
}

// Insert Multiple Records
func InsertPostgres(records []data.DummyData) error {
	for _, record := range records {
		if err := CreatePostgres(record); err != nil {
			return err
		}
	}
	return nil
}

// Find by JSON field (data.id)
func FindPostgresByID(jsonID int) (JSONData, error) {
	var result JSONData
	idStr := fmt.Sprintf("%d", jsonID) // Convert int to string because data->>'id' is text
	err := PostgresDB.
		Where("data->>'id' = ?", idStr). // No CAST — use string directly for index scan
		First(&result).Error
	return result, err
}

// Query by condition (e.g. age > 30)
func QueryPostgres(minAge int) ([]JSONData, error) {
	var results []JSONData
	err := PostgresDB.
		Where("CAST(data->>'age' AS INTEGER) > ?", minAge).
		Find(&results).Error
	return results, err
}

// Update JSON field (change city where age > 40)
func UpdatePostgresCity(minAge int, newCity string) error {
	return PostgresDB.Exec(`
		UPDATE json_data
		SET data = jsonb_set(data, '{address,city}', to_jsonb(?::text), false)
		WHERE (data->>'age')::int > ?`, newCity, minAge).Error
}

// Delete by condition (e.g. age < 25)
func DeletePostgresByAge(maxAge int) error {
	return PostgresDB.Exec(`
		DELETE FROM json_data
		WHERE (data->>'age')::int < ?`, maxAge).Error
}
