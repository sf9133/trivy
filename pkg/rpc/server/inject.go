// +build wireinject

package server

import (
	"github.com/sf9133/fanal/cache"
	"github.com/google/wire"
)

func initializeScanServer(localArtifactCache cache.LocalArtifactCache) *ScanServer {
	wire.Build(ScanSuperSet)
	return &ScanServer{}
}

func initializeDBWorker(cacheDir string, quiet bool) dbWorker {
	wire.Build(DBWorkerSuperSet)
	return dbWorker{}
}
