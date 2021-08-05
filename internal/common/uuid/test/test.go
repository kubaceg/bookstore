package test

import "github.com/stretchr/testify/mock"

type UuidGeneratorMock struct {
	mock.Mock
}

func (t *UuidGeneratorMock) Generate() (string, error) {
	args := t.Called()

	return args.String(0), args.Error(1)
}
