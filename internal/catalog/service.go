package catalog

type Service struct {
	store CatalogStore
}

func NewService(store CatalogStore) *Service {
	return &Service{
		store: store,
	}
}

func (s *Service) ListLocations() ([]Location, error) {
	return s.store.ListLocations()
}

func (s *Service) ListMovies() ([]Movie, error) {
	return s.store.ListMovies()
}

func (s *Service) ListShows(movieId uint) ([]Show, error) {
	return s.store.ListShows(movieId)
}

func (s *Service) ListShowsByTheatre(theatreId uint) ([]Show, error) {
	return s.store.ListShowsByTheatre(theatreId)
}

func (s *Service) ListTheatres(locationId uint) ([]Theatre, error) {
	return s.store.ListTheatres(locationId)
}
