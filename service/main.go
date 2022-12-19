package main

import (
	"github.com/bp1222/tinybeans-api/go/handlers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

    //todo: handle the error!
	c, _ := handlers.NewContainer()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())


	// GetSheetData - Retrieve data from a sheet
	e.GET("/spreadsheets/:SpreadsheetId/sheets/:SheetId/sheetdata", c.GetSheetData)

	// SheetUpdate - Submit a sheet update
	e.POST("/spreadsheets/:SpreadsheetId/sheets/:SheetId/update", c.SheetUpdate)


	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}