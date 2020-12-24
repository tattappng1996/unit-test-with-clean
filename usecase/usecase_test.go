package usecase_test

import (
	"testing"
	"unit-test-with-clean/usecase"
	"unit-test-with-clean/usecase/mocks"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestEiei(t *testing.T) {
	testcases := []usecase.TestcaseStruct{
		usecase.TestcaseStruct{
			Number: 1,
			Name:   "eiei1",
			ExpectData: usecase.Expect{
				ID: 1,
			},
		},
		usecase.TestcaseStruct{
			Number: 2,
			Name:   "eiei2",
			ExpectData: usecase.Expect{
				ID: 2,
			},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.Name, func(tt *testing.T) {
			assert := assert.New(tt)
			ctrl := gomock.NewController(tt)
			defer ctrl.Finish()

			m := mocks.NewMockUsecase(ctrl)
			m.EXPECT().TestNaja(tc).Return(usecase.Expect{}, nil).Times(1)

			u := usecase.NewUsecase(m)
			expectData := u.TestNaja(tc)

			assert.Equal(usecase.Expect{
				ID: tc.Number,
			}, expectData)
		})
	}
}
