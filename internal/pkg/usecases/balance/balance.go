package balance

import (
	"AdHub/internal/pkg/entities"
)

type BalanceUseCase struct {
	repo entities.BalanceRepoInterface
}

func New(r entities.BalanceRepoInterface) *BalanceUseCase {
	return &BalanceUseCase{
		repo: r,
	}
}

func (uc *BalanceUseCase) BalanceCreate(ad *entities.Balance) (*entities.Balance, error) {
	return uc.repo.Create(ad)
}

func (uc *BalanceUseCase) BalanceRead(id int) (*entities.Balance, error) {
	return uc.repo.Read(id)
}

func (uc *BalanceUseCase) BalanceRemove(id int) error {
	return uc.repo.Remove(id)
}

func (uc *BalanceUseCase) BalanceUP(sum float64, id int) error {
	Bal, err := uc.repo.Read(id)
	if err != nil {
		return err
	}

	Bal.Total_balance += sum
	Bal.Available_balance += sum

	return uc.repo.Update(Bal)
}

func (uc *BalanceUseCase) BalanceDown(sum float64, id int) error {
	Bal, err := uc.repo.Read(id)
	if err != nil {
		return err
	}

	Bal.Total_balance -= sum
	Bal.Available_balance -= sum

	return uc.repo.Update(Bal)
}

func (uc *BalanceUseCase) BalanceReserve(sum float64, id int) error {
	Bal, err := uc.repo.Read(id)
	if err != nil {
		return err
	}

	Bal.Reserved_balance += sum

	return uc.repo.Update(Bal)
}
