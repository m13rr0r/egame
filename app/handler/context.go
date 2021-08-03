package handler

import (
	"sync"

	"github.com/m13rr0r/egame/domain"
	"github.com/m13rr0r/egame/infrastructure/storage"
)

type (
	Context struct {
		Storage *storage.Storage
		ChanSize int
		ChunkSize int
		EventPool *sync.Pool
		EventChan chan *domain.Event
	}
)
