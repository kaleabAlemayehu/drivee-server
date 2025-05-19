package utils

import "github.com/google/uuid"

func ParseNullableUUID(id uuid.UUID) *uuid.NullUUID {
	if id == uuid.Nil {
		return &uuid.NullUUID{
			Valid: false,
			UUID:  id,
		}
	} else {
		return &uuid.NullUUID{
			Valid: true,
			UUID:  id,
		}
	}
}
