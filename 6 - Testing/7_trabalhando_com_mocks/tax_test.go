package tax

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type TaxRepositoryMock struct {
	mock.Mock
}

func (m *TaxRepositoryMock) SaveTax(tax float64) error {
	args := m.Called(tax)
	return args.Error(0) // posição do argumento de retorno
	// Caso fosse (error, int), então seria return args.Error(0), args.Int(1)
}

func TestCalculateTaxAndSave(t *testing.T) {
	repository := &TaxRepositoryMock{}
	repository.On("SaveTax", 10.0).Return(nil).Twice() // Define o comportamento da função X quando recebe o parametro Y e deve retornar Z
	repository.On("SaveTax", 0.0).Return(errors.New("error saving tax"))
	repository.On("SaveTax", mock.Anything).Return(errors.New("error saving tax")) // Passa qualquer valor por argumento

	err := CalculateTaxAndSave(1000.0, repository)
	assert.Nil(t, err)
	err = CalculateTaxAndSave(1000.0, repository)
	assert.Nil(t, err)

	err = CalculateTaxAndSave(0, repository)
	assert.Error(t, err, "error saving tax")

	repository.AssertExpectations(t)
	repository.AssertNumberOfCalls(t, "SaveTax", 3)
}
