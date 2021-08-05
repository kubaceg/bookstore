package uuid

import uuid2 "github.com/google/uuid"

type V4Generator struct{}

func (v *V4Generator) Generate() (string, error) {
	uuid, err := uuid2.NewUUID()

	if err != nil {
		return "", err
	}

	return uuid.String(), err
}
