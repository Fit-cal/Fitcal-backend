package interactor

import (
	"errors"
	"fitcal-backend/domain/entities"
	mock_inputport "fitcal-backend/domain/repository/inputport/mock"
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/golang/mock/gomock"
)

func TestGetUsers(t *testing.T) {
	err := errors.New("error")
	cases := []struct {
		name          string
		prepareMockFn func(ma *mock_inputport.MockUserRepositoryInputPort)
		expected      entities.Users
		expectedErr   error
	}{
		{
			name: "success",
			prepareMockFn: func(ma *mock_inputport.MockUserRepositoryInputPort) {
				ma.EXPECT().GetUsers().Return(entities.Users{{UserId: "test", FirstName: "test", LastName: "test", Email: "test"}}, nil)
			},
			expected:    entities.Users{{UserId: "test", FirstName: "test", LastName: "test", Email: "test"}},
			expectedErr: nil,
		},
		{
			name: "fail",
			prepareMockFn: func(ma *mock_inputport.MockUserRepositoryInputPort) {
				ma.EXPECT().GetUsers().Return(entities.Users{{UserId: "test", FirstName: "test", LastName: "test", Email: ""}}, err)
			},
			expected:    entities.Users{{UserId: "test", FirstName: "test", LastName: "test", Email: ""}},
			expectedErr: err,
		},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mockRepo := mock_inputport.NewMockUserRepositoryInputPort(ctrl)
			interactor := UserInteractor{
				userRepository: mockRepo,
			}
			tt.prepareMockFn(mockRepo)
			actual, err := interactor.userRepository.GetUsers()
			assert.Equal(t, tt.expected, actual)
			assert.Equal(t, tt.expectedErr, err)
		})
	}
}

func TestSearchusers(t *testing.T) {
	testKeyword := "test"
	testUsers := entities.Users{
		{
			UserId:    "testUserID",
			FirstName: "testFirstName",
			LastName:  "testLastName",
			Email:     "test@email.com",
		},
	}
	cases := []struct {
		name          string
		keyword       string
		prepareMockFn func(ma *mock_inputport.MockUserRepositoryInputPort, keyword string)
		expected      entities.Users
		expectedErr   error
	}{
		{
			name:    "success",
			keyword: testKeyword,
			prepareMockFn: func(ma *mock_inputport.MockUserRepositoryInputPort, keyword string) {
				ma.EXPECT().SearchUsers(keyword).Return(testUsers, nil)
			},
			expected:    testUsers,
			expectedErr: nil,
		},
		{
			name:    "fail",
			keyword: testKeyword,
			prepareMockFn: func(ma *mock_inputport.MockUserRepositoryInputPort, keyword string) {
				ma.EXPECT().SearchUsers(keyword).Return(nil, errors.New("error"))
			},
			expected:    nil,
			expectedErr: errors.New("error"),
		},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mockRepo := mock_inputport.NewMockUserRepositoryInputPort(ctrl)
			interactor := UserInteractor{
				userRepository: mockRepo,
			}
			tt.prepareMockFn(mockRepo, testKeyword)
			actual, err := interactor.SearchUsers(tt.keyword)
			assert.Equal(t, tt.expected, actual)
			assert.Equal(t, tt.expectedErr, err)
		})
	}
}
