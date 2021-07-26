package handler

import (
	"github.com/m13rr0r/egame/infrastructure/storage"
)

type (
	Handler struct {
		DB *storage.Storage
	}
)
