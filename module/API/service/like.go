package service

import . "github.com/04Akaps/metting/db/HDD/types"

func (s *service) LikeSomeOne(from, to string) error {
	return nil
}

func (s *service) RefuseLike(from, to string) error {
	return nil
}
func (s *service) AcceptedLike(from, to string) error {
	return nil
}

func (s *service) GetLikedList(from string) ([]*User, error) {
	return nil, nil
}
