package service

import (
	"context"
	"testing"

	"github.com/AndresBC-Dev/parfum/entity"
)

type RepositoryMock struct{}

func (rm *RepositoryMock) Create(perfume entity.Parfum) error {
	return nil
}

func TestCreate(t *testing.T) error {
	testCases := []struct {
		Name          string
		ExpectedError error
	}{
		{
			Name:          "Create Perfume",
			ExpectedError: nil,
		},
	}

	repo := &RepositoryMock{}
	var s perfumeService = NewPerfumeService(repo)
	ctx := context.Background()

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()

			_, err := s.Create(ctx)

			if err != tc.ExpectedError {
				t.Errorf("expected %v, got %v", tx.expected, err)
			}
		})

	}
}
