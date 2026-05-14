package domain

import (
	"time"
)

type TideService struct {
	repo HexagramRepository
}

func NewTideService(repo HexagramRepository) *TideService {
	return &TideService{repo: repo}
}

type Result struct {
	Hexagram Hexagram
	Data     HexagramData
	Upper    HexagramData
	Lower    HexagramData
}

func (s *TideService) DailyDraw() (Result, error) {
	h := DrawDaily(time.Now())
	return s.getResult(h)
}

func (s *TideService) LuckyDraw() (Result, error) {
	h := DrawLucky()
	return s.getResult(h)
}

func (s *TideService) getResult(h Hexagram) (Result, error) {
	data, err := s.repo.GetHexagramData(h)
	if err != nil {
		return Result{}, err
	}

	upper, err := s.repo.GetTrigramData(h.GetUpperTrigram())
	if err != nil {
		return Result{}, err
	}

	lower, err := s.repo.GetTrigramData(h.GetLowerTrigram())
	if err != nil {
		return Result{}, err
	}

	return Result{
		Hexagram: h,
		Data:     data,
		Upper:    upper,
		Lower:    lower,
	}, nil
}
