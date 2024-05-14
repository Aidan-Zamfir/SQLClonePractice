package core

import (
	"github.com/marianogappa/sqlparser/query"
	"github.com/rs/zerolog/log"
)

func ExecuteStatement(q query.Query) {
	switch q.Type {
	case query.Insert:
		log.Info().Msg("Insert here")
	case query.Select:
		log.Info().Msg("Select here")
	}
}