package store

import (
	"log"

	"github.com/c12s/meridian/internal/domain"
)

type namespaceNeo4jStore struct {
	txManager *TxManager
}

func NewNamespaceNeo4jStore(txManager *TxManager) domain.NamespaceStore {
	if txManager == nil {
		log.Fatalln("tx manager while initializing namespace neo4j store")
	}
	return &namespaceNeo4jStore{
		txManager: txManager,
	}
}

// Add implements domain.NamespaceStore.
func (n *namespaceNeo4jStore) Add(namespace domain.Namespace, parent domain.Namespace) error {
	panic("unimplemented")
}

// Get implements domain.NamespaceStore.
func (n *namespaceNeo4jStore) Get(id string) (domain.Namespace, error) {
	panic("unimplemented")
}

// GetHierarchy implements domain.NamespaceStore.
func (n *namespaceNeo4jStore) GetHierarchy(rootId string) (domain.NamespaceTree, error) {
	panic("unimplemented")
}

// Remove implements domain.NamespaceStore.
func (n *namespaceNeo4jStore) Remove(id string) error {
	panic("unimplemented")
}
