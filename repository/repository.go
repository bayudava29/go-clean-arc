package repository

import (
	"context"

	"github.com/bayudava29/go-clean-arc/config"
)

type Repository interface {
	GetData(c context.Context)
}

type repository struct {
	hbaseConn []config.HbaseConnection
}

func InitRepository(hbase []config.HbaseConnection) Repository {
	return &repository{
		hbaseConn: hbase,
	}
}

func (repository *repository) GetData(c context.Context) {
	// Instruction for GetData Repository
}
