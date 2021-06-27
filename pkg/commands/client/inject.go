// +build wireinject

package client

import (
	"context"
	"time"

	"github.com/google/wire"

	"github.com/sf9133/fanal/analyzer"
	"github.com/sf9133/fanal/analyzer/config"
	"github.com/sf9133/fanal/cache"
	"github.com/aquasecurity/trivy/pkg/rpc/client"
	"github.com/aquasecurity/trivy/pkg/scanner"
	"github.com/aquasecurity/trivy/pkg/vulnerability"
)

func initializeDockerScanner(ctx context.Context, imageName string, artifactCache cache.ArtifactCache, customHeaders client.CustomHeaders,
	url client.RemoteURL, timeout time.Duration, disabled []analyzer.Type, configScannerOption config.ScannerOption) (
	scanner.Scanner, func(), error) {
	wire.Build(scanner.RemoteDockerSet)
	return scanner.Scanner{}, nil, nil
}

func initializeArchiveScanner(ctx context.Context, filePath string, artifactCache cache.ArtifactCache,
	customHeaders client.CustomHeaders, url client.RemoteURL, timeout time.Duration, disabled []analyzer.Type,
	configScannerOption config.ScannerOption) (scanner.Scanner, error) {
	wire.Build(scanner.RemoteArchiveSet)
	return scanner.Scanner{}, nil
}

func initializeResultClient() vulnerability.Client {
	wire.Build(vulnerability.SuperSet)
	return vulnerability.Client{}
}
