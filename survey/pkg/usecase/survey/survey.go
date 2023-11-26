package survey

import (
	"AdHub/survey/pkg/entities"
)

type SurveyUseCase struct {
	repo entities.SurveyRepoInterface
}

func New(r entities.SurveyRepoInterface) *SurveyUseCase {
	return &SurveyUseCase{
		repo: r,
	}
}

func (uc *SurveyUseCase) SurveyCreate(s *entities.Survey) (*entities.Survey, error) {
	return uc.repo.Create(s)
}

func (uc *SurveyUseCase) SurveyRead(id int) (*entities.Survey, error) {
	return uc.repo.Read(id)
}

func (uc *SurveyUseCase) SurveyRemove(id int) error {
	return uc.repo.Remove(id)
}

func (uc *SurveyUseCase) SurveyUpdate(s *entities.Survey) error {
	return uc.repo.Update(s)
}

func (uc *SurveyUseCase) SurveyList() (*[]entities.Survey, error) {
	return uc.repo.ReadList()
}
