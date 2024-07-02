package store

import (
	"log"

	"github.com/c12s/meridian/internal/domain"
)

type appNeo4jStore struct {
	txManager *TxManager
}

func NewAppNeo4jStore(txManager *TxManager) domain.AppStore {
	if txManager == nil {
		log.Fatalln("tx manager while initializing app neo4j store")
	}
	return &appNeo4jStore{
		txManager: txManager,
	}
}

// Add implements domain.AppStore.
func (a *appNeo4jStore) Add(app domain.App) error {
	panic("unimplemented")
}

// FindByNamespace implements domain.AppStore.
func (a *appNeo4jStore) FindByNamespace(namespaceId string) ([]domain.App, error) {
	panic("unimplemented")
}

// Remove implements domain.AppStore.
func (a *appNeo4jStore) Remove(id string) error {
	panic("unimplemented")
}
