package api

import (
	"message-service/internal/model"
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

type Handlers struct {
	DB *sqlx.DB
}

func NewHandler(db *sqlx.DB) *Handlers {
	return &Handlers{DB: db}
}

func (h *Handlers) CreateMessage(c echo.Context) error {
	message := new(model.Message)
	if err := c.Bind(&message); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"Error": err.Error()})
	}
	query := `INSERT INTO message (content, status) VALUES ($1, $2) RETURNING id`
	err := h.DB.QueryRow(query, message.Content, message.Status).Scan(&message.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"Error": err.Error()})
	}
	return c.JSON(http.StatusOK, message)
}
