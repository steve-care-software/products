package repositories

import (
	"github.com/deepvalue-network/software/blockchain/domain/blocks"
	"github.com/deepvalue-network/software/libs/hash"
)

type block struct {
	blockRepository blocks.Repository
}

func createBlock(
	blockRepository blocks.Repository,
) Block {
	out := block{
		blockRepository: blockRepository,
	}

	return &out
}

// List returns the block hashes list
func (app *block) List() ([]hash.Hash, error) {
	return app.blockRepository.List()
}

// Retrieve retrieves a block by hash
func (app *block) Retrieve(hash hash.Hash) (blocks.Block, error) {
	return app.blockRepository.Retrieve(hash)
}
