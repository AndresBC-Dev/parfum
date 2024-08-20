package service

import (
	"testing"

	"github.com/AndresBC-Dev/parfum/entity"
)

type RepositoryMock struct{}

func (rm *RepositoryMock) Create(perfume entity.Parfum) error {
	return nil
}

func TestCreate(t *testing.T) {
	testCases := []struct {
		Name          string
		Perfume       entity.Parfum
		ExpectedError error
	}{
		{
			Name:          "Create Perfume",
			Perfume:       entity.Parfum{Name: "Test Perfume"},
			ExpectedError: nil,
		},
	}

	repo := &RepositoryMock{}
	s := NewPerfumeService(repo)

	for _, tc := range testCases {
		tc := tc // evita un problema común en el uso de variables de loop
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()

			err := s.Create(tc.Perfume)
			if err != tc.ExpectedError {
				t.Errorf("expected %v, got %v", tc.ExpectedError, err)
			}
		})
	}
}
