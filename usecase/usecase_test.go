package usecase_test

import (
	"errors"
	"testing"
	"unit-test-with-clean/dto"
	"unit-test-with-clean/usecase"
	"unit-test-with-clean/usecase/mocks"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGetFirstCustomerSuccess(t *testing.T) {
	expected := dto.Customer{ID: 1, Name: "John"}

	assert := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockCustomerRepository(ctrl)

	m.EXPECT().First(1).Return(dto.Customer{ID: 1, Name: "John"}, nil).Times(1)

	u := usecase.NewUsecase(m)
	output, _ := u.GetFirstCustomer(1)

	assert.Equal(expected, output)
}

func TestGetFirstCustomerError(t *testing.T) {
	assert := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockCustomerRepository(ctrl)

	dummyErrExpected := errors.New("404 user not found")
	m.EXPECT().First(1).Return(dto.Customer{}, dummyErrExpected).Times(1)

	u := usecase.NewUsecase(m)
	_, err := u.GetFirstCustomer(1)

	assert.EqualError(err, dummyErrExpected.Error())
}
