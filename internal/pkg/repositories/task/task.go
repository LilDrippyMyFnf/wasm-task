package pessoas

import (
	"github.com/jgbz/wasm-task/internal/pkg/repositories/config"
)

type PessoasRepository struct {
	instance *config.ConfigRepository
}

func NewPessoasRepository() *PessoasRepository {
	return &PessoasRepository{
		instance: config.GetConfigRepository(),
	}
}
