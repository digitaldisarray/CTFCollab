package handler

type Handler struct {
	//Queries *db.Queries
}

func NewHandler( /*q *db.Queries*/ ) *Handler {
	return &Handler{ /*Queries: q*/ }
}

func LoadHandler( /*dbUrl string*/ ) (*Handler, error) {
	// Connect to database

	// Make a handler object
	//queries := db.New(conn)
	handler := NewHandler( /*queries*/ )

	return handler, nil
}
