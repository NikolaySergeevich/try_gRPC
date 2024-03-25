package handler

import (
	"trygrpc/internal/memstore"
)

type store interface {
	Add(item memstore.Item)
	DeleteByObjectID(objectID string)
	FindByObjectID(objectID string) (memstore.Item, bool)
	FindAll() []memstore.Item
}
