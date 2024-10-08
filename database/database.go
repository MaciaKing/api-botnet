package database

import (
	"api-botnet/models"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	host := os.Getenv("POSTGRES_HOST")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")
	dbport := os.Getenv("POSTGRES_PORT")
	// sslmode := os.Getenv("SSLMODE")
	sslmode := "disable"
	dsn := "host=" + host + " user=" + user + password + " dbname=" + dbname + " port=" + dbport + " sslmode=" + sslmode

	fmt.Println(dsn)
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	log.Println("Database connection successful.")
}

func ResetDatabase(db *gorm.DB) error {
	// Drop todas las tablas
	err := db.Migrator().DropTable("bots", "victims") // Agrega todas las tablas aquí
	if err != nil {
		return err
	}

	// Vuelve a migrar las tablas
	err = db.AutoMigrate(&models.Bot{}, &models.Victim{}) // Pasa tus modelos aquí
	if err != nil {
		return err
	}

	return nil
}

func Migrate() {
	ResetDatabase(DB)
	if err := DB.AutoMigrate(&models.Bot{}); err != nil {
		log.Fatal("Failed to migrate Library model:", err)
	}
	createDefaultBoots()
	log.Println("Bot model migration successful.")

	if err := DB.AutoMigrate(&models.Victim{}); err != nil {
		log.Fatal("Failed to migrate Library model:", err)
	}
	log.Println("Victim model migration successful.")

}

func createDefaultBoots() {
	boots := []models.Bot{
		{Ip: "1.0.0.0"},
		// {Ip: "2.0.0.0"},
		// {Ip: "3.0.0.0"},
		// {Ip: "4.0.0.0"},
	}

	for _, bot := range boots {
		if err := DB.Create(&bot).Error; err != nil {
			log.Printf("Error creating book %v: %v\n", bot, err)
		}
	}
	log.Println("Default boots created.")
}
