package rating

type ServiceRating interface {
	SetRating(bookId, rating int) (bool, error)
}

type Service struct {
	Repository *Repository
}

func (s *Service) SetRating(bookId, rating int) (bool, error) {
	return s.Repository.SetRating(bookId, rating)
}
