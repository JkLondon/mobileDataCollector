package main

import (
	"log"
	"net/http"

	"github.com/JkLondon/mobileDataCollector/internal/models"
	"github.com/JkLondon/mobileDataCollector/pkg/sqlite"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	app := fiber.New()

	db, err := sqlx.Connect("sqlite3", "./data.db")
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(sqlite.QueryCreateOrderTable)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	app.Get("/storage", func(c *fiber.Ctx) error {
		result := models.ExperimentData{}
		result.Data = make([]models.ExperimentItem, 0)
		err := db.SelectContext(c.Context(), &result.Data, sqlite.QueryGetData)
		if err != nil {
			c.Status(http.StatusInternalServerError).SendString(err.Error())
		}
		return c.JSON(result)
	})

	app.Post("/data", func(c *fiber.Ctx) error {
		item := new(models.ExperimentItem)
		if err := c.BodyParser(item); err != nil {
			return c.Status(http.StatusBadRequest).SendString(err.Error())
		}
		_, err = db.ExecContext(c.Context(), sqlite.QueryInsertData, item.SSID, item.RSSI, item.Time)
		if err != nil {
			c.Status(http.StatusInternalServerError).SendString(err.Error())
		}
		log.Printf("new req with data %v\n", item)
		return c.SendStatus(http.StatusOK)
	})

	app.Listen(":3000")
}
