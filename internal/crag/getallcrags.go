package crag

//GetAllCragsQueryHandler Contains the dependencies of the Handler
type GetAllCragsQueryHandler interface {
	Handle() ([]QueryResult, error)
}

type getAllCragsQueryHandler struct {
	repo Repository
}

//NewGetAllCragsQueryHandler Handler constructor
func NewGetAllCragsQueryHandler(repo Repository) GetAllCragsQueryHandler {
	return getAllCragsQueryHandler{repo: repo}
}

//Handle Handles the query
func (h getAllCragsQueryHandler) Handle() ([]QueryResult, error) {

	res, err := h.repo.GetAll()
	if err != nil {
		return nil, err
	}
	var crags []QueryResult
	for _, crag := range res {
		crags = append(crags, QueryResult{ID: crag.ID, Name: crag.Name, Desc: crag.Desc, Country: crag.Country, CreatedAt: crag.CreatedAt})
	}
	return crags, nil
}
