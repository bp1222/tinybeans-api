package handlers
import (
    "github.com/bp1222/tinybeans-api/go/models"
    "github.com/labstack/echo/v4"
    "net/http"
)

// GetSheetData - Retrieve data from a sheet
func (c *Container) GetSheetData(ctx echo.Context) error {
    return ctx.JSON(http.StatusOK, models.HelloWorld {
        Message: "Hello World",
    })
}


// SheetUpdate - Submit a sheet update
func (c *Container) SheetUpdate(ctx echo.Context) error {
    return ctx.JSON(http.StatusOK, models.HelloWorld {
        Message: "Hello World",
    })
}
