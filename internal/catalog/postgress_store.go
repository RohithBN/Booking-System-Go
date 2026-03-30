package catalog

import "gorm.io/gorm"

type PostgressStore struct {
	db *gorm.DB
}

func NewPostgressStore(db *gorm.DB) *PostgressStore {
	return &PostgressStore{
		db: db,
	}
}

func (s *PostgressStore) ListLocations() ([]Location, error) {
	var locations []Location
	result := s.db.Find(&locations)
	return locations, result.Error
}

func (s *PostgressStore) ListMovies() ([]Movie, error) {
	var movies []Movie
	result := s.db.Find(&movies)
	return movies, result.Error
}

func (s *PostgressStore) ListShows(movieId uint) ([]Show, error) {
	var shows []Show
	result := s.db.Preload("Movie").Preload("Theatre").Where("movie_id = ?", movieId).Find(&shows)
	return shows, result.Error
}

func (s *PostgressStore) ListShowsByTheatre(theatreId uint) ([]Show, error) {
	var shows []Show
	result := s.db.Preload("Movie").Preload("Theatre").Where("theatre_id = ?", theatreId).Find(&shows)
	return shows, result.Error
}

func (s *PostgressStore) ListTheatres(locationId uint) ([]Theatre, error) {
	var theatres []Theatre
	result := s.db.Where("location_id = ?", locationId).Find(&theatres)
	return theatres, result.Error
}
