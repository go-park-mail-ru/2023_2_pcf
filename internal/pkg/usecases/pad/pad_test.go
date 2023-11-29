package pad

//
//func setupPadUseCase(t *testing.T) (*PadUseCase, *mock_entities.MockPadRepoInterface, *gomock.Controller) {
//	ctrl := gomock.NewController(t)
//	mockRepo := mock_entities.NewMockPadRepoInterface(ctrl)
//	uc := New(mockRepo)
//	return uc, mockRepo, ctrl
//}
//
//func TestPadCreate(t *testing.T) {
//	uc, mockRepo, ctrl := setupPadUseCase(t)
//	defer ctrl.Finish()
//
//	pad := &entities.Pad{ /* initialize with test data */ }
//	mockRepo.EXPECT().Create(pad).Return(pad, nil)
//
//	result, err := uc.PadCreate(pad)
//
//	assert.NoError(t, err)
//	assert.Equal(t, pad, result)
//}
//
//func TestPadReadList(t *testing.T) {
//	uc, mockRepo, ctrl := setupPadUseCase(t)
//	defer ctrl.Finish()
//
//	pads := []*entities.Pad{ /* initialize with test data */ }
//	mockRepo.EXPECT().Read(1).Return(pads, nil)
//
//	result, err := uc.PadReadList(1)
//
//	assert.NoError(t, err)
//	assert.Equal(t, pads, result)
//}
//
//func TestPadRead(t *testing.T) {
//	uc, mockRepo, ctrl := setupPadUseCase(t)
//	defer ctrl.Finish()
//
//	pad := &entities.Pad{ /* initialize with test data */ }
//	mockRepo.EXPECT().Get(1).Return(pad, nil)
//
//	result, err := uc.PadRead(1)
//
//	assert.NoError(t, err)
//	assert.Equal(t, pad, result)
//}
//
//func TestPadRemove(t *testing.T) {
//	uc, mockRepo, ctrl := setupPadUseCase(t)
//	defer ctrl.Finish()
//
//	mockRepo.EXPECT().Remove(1).Return(nil)
//
//	err := uc.PadRemove(1)
//
//	assert.NoError(t, err)
//}
//
//func TestPadUpdate(t *testing.T) {
//	uc, mockRepo, ctrl := setupPadUseCase(t)
//	defer ctrl.Finish()
//
//	pad := &entities.Pad{ /* initialize with test data */ }
//	mockRepo.EXPECT().Update(pad).Return(nil)
//
//	err := uc.PadUpdate(pad)
//
//	assert.NoError(t, err)
//}
