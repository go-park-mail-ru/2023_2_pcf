package rate

import (
	"AdHub/survey/pkg/entities"
)

type RateUseCase struct {
	repo entities.RateRepoInterface
}

func New(r entities.RateRepoInterface) *RateUseCase {
	return &RateUseCase{
		repo: r,
	}
}

func (uc *RateUseCase) RateCreate(rate *entities.Rate) (*entities.Rate, error) {
	return uc.repo.Create(rate)
}

func (uc *RateUseCase) RateRead(id int) (*entities.Response, error) {
	rates, err := uc.repo.Read(id)
	if err != nil {
		return nil, err
	}
	var (
		count int
		sum   int
	)
	for _, rate := range rates {
		count += 1
		sum += rate.Rate
	}

	return &entities.Response{
		Count:     count,
		Avg_rate:  float64(sum) / float64(count),
		Survey_id: id,
	}, nil
}
