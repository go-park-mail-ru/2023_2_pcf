package balance

import (
	"AdHub/internal/pkg/entities"
	"AdHub/internal/pkg/entities/mock_entities"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestBalanceUseCase_BalanceCreate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_entities.NewMockBalanceRepoInterface(ctrl)

	useCase := New(mockRepo)

	fakeBalance := &entities.Balance{
		Total_balance:     1000.0,
		Reserved_balance:  100.0,
		Available_balance: 900.0,
	}

	mockRepo.EXPECT().Create(gomock.Eq(fakeBalance)).Return(fakeBalance, nil)

	createdBalance, err := useCase.BalanceCreate(fakeBalance)
	assert.NoError(t, err)
	assert.Equal(t, fakeBalance, createdBalance)
}

func TestBalanceUseCase_BalanceRead(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_entities.NewMockBalanceRepoInterface(ctrl)

	useCase := New(mockRepo)

	fakeBalance := &entities.Balance{
		Id:                1,
		Total_balance:     1000.0,
		Reserved_balance:  100.0,
		Available_balance: 900.0,
	}

	mockRepo.EXPECT().Read(1).Return(fakeBalance, nil)

	balance, err := useCase.BalanceRead(1)
	assert.NoError(t, err)
	assert.Equal(t, fakeBalance, balance)
}

func TestBalanceUseCase_BalanceRemove(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_entities.NewMockBalanceRepoInterface(ctrl)

	useCase := New(mockRepo)

	mockRepo.EXPECT().Remove(1).Return(nil)

	err := useCase.BalanceRemove(1)
	assert.NoError(t, err)
}

func TestBalanceUseCase_BalanceUP(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_entities.NewMockBalanceRepoInterface(ctrl)

	useCase := New(mockRepo)

	fakeBalance := &entities.Balance{
		Id:                1,
		Total_balance:     1000.0,
		Reserved_balance:  100.0,
		Available_balance: 900.0,
	}

	mockRepo.EXPECT().Read(1).Return(fakeBalance, nil)
	mockRepo.EXPECT().Update(gomock.Eq(fakeBalance)).Return(nil)

	err := useCase.BalanceUP(200.0, 1)
	assert.NoError(t, err)
	assert.Equal(t, 1200.0, fakeBalance.Total_balance)
	assert.Equal(t, 1100.0, fakeBalance.Available_balance)
}

func TestBalanceUseCase_BalanceDown(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_entities.NewMockBalanceRepoInterface(ctrl)

	useCase := New(mockRepo)

	fakeBalance := &entities.Balance{
		Id:                1,
		Total_balance:     1000.0,
		Reserved_balance:  100.0,
		Available_balance: 900.0,
	}

	mockRepo.EXPECT().Read(1).Return(fakeBalance, nil)
	mockRepo.EXPECT().Update(gomock.Eq(fakeBalance)).Return(nil)

	err := useCase.BalanceDown(200.0, 1)
	assert.NoError(t, err)
	assert.Equal(t, 800.0, fakeBalance.Total_balance)
	assert.Equal(t, 700.0, fakeBalance.Available_balance)
}

func TestBalanceUseCase_BalanceReserve(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_entities.NewMockBalanceRepoInterface(ctrl)

	useCase := New(mockRepo)

	fakeBalance := &entities.Balance{
		Id:                1,
		Total_balance:     1000.0,
		Reserved_balance:  100.0,
		Available_balance: 900.0,
	}

	mockRepo.EXPECT().Read(1).Return(fakeBalance, nil)
	mockRepo.EXPECT().Update(gomock.Eq(fakeBalance)).Return(nil)

	err := useCase.BalanceReserve(50.0, 1)
	assert.NoError(t, err)
	assert.Equal(t, 150.0, fakeBalance.Reserved_balance)
}
