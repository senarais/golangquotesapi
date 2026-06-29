package api


type QuoteGenreParams struct {
	genre string
}

type QuoteResponse struct {
	code int
	quote string
}

type Error struct {
	code int
	message string
}