package db

import (
	"Jsonb/data"
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type JSONData struct {
	ID   uint                   `gorm:"primaryKey"`
	Data map[string]interface{} `gorm:"type:jsonb"`
}

var PostgresDB *gorm.DB

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
	// Auto-migrate here:
	if err := db.AutoMigrate(&JSONData{}); err != nil {
		return err
	}
	PostgresDB = db
	fmt.Println("✅ Connected to PostgreSQL")
	return nil
}

func InsertPostgres(records []data.DummyData) error {
	for _, record := range records {
		// Convert DummyData to JSONData model
		item := JSONData{
			Data: map[string]interface{}{
				"id":      record.ID,
				"name":    record.Name,
				"age":     record.Age,
				"address": record.Address,
				"tags":    record.Tags,
			},
		}

		if err := PostgresDB.Create(&item).Error; err != nil {
			return err
		}
	}
	return nil
}

func QueryPostgres(minAge int) ([]JSONData, error) {
	var results []JSONData

	err := PostgresDB.
		Where("CAST(data->>'age' AS INTEGER) > ?", minAge).
		Find(&results).Error

	return results, err
}

func UpdatePostgresCity(minAge int, newCity string) error {
	// JSONB Set Example: updates the 'address->city' field in 'data'
	return PostgresDB.Exec(`
		UPDATE json_data
		SET data = jsonb_set(data, '{address,city}', to_jsonb(?::text), false)
		WHERE (data->>'age')::int > ?`, newCity, minAge).Error
}

func DeletePostgresByAge(maxAge int) error {
	return PostgresDB.Exec(`
		DELETE FROM json_data
		WHERE (data->>'age')::int < ?`, maxAge).Error
}

func CreatePostgres(record data.DummyData) error {
	item := JSONData{
		Data: map[string]interface{}{
			"id":      record.ID,
			"name":    record.Name,
			"age":     record.Age,
			"address": record.Address,
			"tags":    record.Tags,
		},
	}

	return PostgresDB.Create(&item).Error
}

func FindPostgresByID(jsonID int) (JSONData, error) {
	var result JSONData

	err := PostgresDB.
		Where("CAST(data->>'id' AS INTEGER) = ?", jsonID).
		First(&result).Error

	return result, err
}
