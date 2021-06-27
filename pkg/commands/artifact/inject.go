// +build wireinject

package artifact

import (
	"context"
	"time"

	"github.com/google/wire"

	"github.com/sf9133/fanal/analyzer"
	"github.com/sf9133/fanal/analyzer/config"
	"github.com/sf9133/fanal/cache"
	"github.com/aquasecurity/trivy/pkg/scanner"
	"github.com/aquasecurity/trivy/pkg/vulnerability"
)

func initializeDockerScanner(ctx context.Context, imageName string, artifactCache cache.ArtifactCache,
	localArtifactCache cache.LocalArtifactCache, timeout time.Duration, disableAnalyzers []analyzer.Type,
	configScannerOption config.ScannerOption) (scanner.Scanner, func(), error) {
	wire.Build(scanner.StandaloneDockerSet)
	return scanner.Scanner{}, nil, nil
}

func initializeArchiveScanner(ctx context.Context, filePath string, artifactCache cache.ArtifactCache,
	localArtifactCache cache.LocalArtifactCache, timeout time.Duration, disableAnalyzers []analyzer.Type,
	configScannerOption config.ScannerOption) (scanner.Scanner, error) {
	wire.Build(scanner.StandaloneArchiveSet)
	return scanner.Scanner{}, nil
}

func initializeFilesystemScanner(ctx context.Context, dir string, artifactCache cache.ArtifactCache,
	localArtifactCache cache.LocalArtifactCache, disableAnalyzers []analyzer.Type, configScannerOption config.ScannerOption) (
	scanner.Scanner, func(), error) {
	wire.Build(scanner.StandaloneFilesystemSet)
	return scanner.Scanner{}, nil, nil
}

func initializeRepositoryScanner(ctx context.Context, url string, artifactCache cache.ArtifactCache,
	localArtifactCache cache.LocalArtifactCache, disableAnalyzers []analyzer.Type, configScannerOption config.ScannerOption) (
	scanner.Scanner, func(), error) {
	wire.Build(scanner.StandaloneRepositorySet)
	return scanner.Scanner{}, nil, nil
}

func initializeResultClient() vulnerability.Client {
	wire.Build(vulnerability.SuperSet)
	return vulnerability.Client{}
}
