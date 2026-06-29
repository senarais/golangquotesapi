package service

import (
	"github.com/senarais/golangquotesapi/internal/database"
	"github.com/senarais/golangquotesapi/internal/model"
	"math/rand/v2"
)

type QuotesService struct {}

func(q *QuotesService) GetRandomQuote() model.Quotes {
	quoteData := database.QuotesData{}
	data := quoteData.GetData()

	// Generate random number
	randomNumber := rand.IntN(len(data))

	response := data[randomNumber]

	return response
}