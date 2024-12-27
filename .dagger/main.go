package main

import (
	"context"
	"dagger/digital-wellbeing/internal/dagger"
	"fmt"
)

const (
	GOLANGCILINT_VERSION = "v1.61.0"
	GO_VERSION           = "1.22.5"
	SYFT_VERSION         = "v1.9.0"
	GORELEASER_VERSION   = "v2.3.2"
)

type DigitalWellbeing struct {
	// Local or remote directory with source code, defaults to "./"
	Source *dagger.Directory
}

// LintReport Executes the Linter and writes the linting results to a file golangci-lint-report.sarif
func (m *DigitalWellbeing) LintReport(ctx context.Context) *dagger.File {
	report := "golangci-lint-report.sarif"
	return m.lint(ctx).WithExec([]string{
		"golangci-lint", "run",
		"--out-format", "sarif:" + report,
		"--issues-exit-code", "0",
	}).File(report)
}

// Lint Run the linter golangci-lint
func (m *DigitalWellbeing) Lint(ctx context.Context) (string, error) {
	return m.lint(ctx).WithExec([]string{"golangci-lint", "run"}).Stderr(ctx)
}

func (m *DigitalWellbeing) lint(_ context.Context) *dagger.Container {
	fmt.Println("👀 Running linter and printing results to file golangci-lint.txt.")
	linter := dag.Container().
		From("golangci/golangci-lint:"+GOLANGCILINT_VERSION+"-alpine").
		WithMountedCache("/lint-cache", dag.CacheVolume("/lint-cache")).
		WithEnvVariable("GOLANGCI_LINT_CACHE", "/lint-cache").
		WithMountedDirectory("/src", m.Source).
		WithWorkdir("/src")
	return linter
}

// Returns a container that echoes whatever string argument is provided
func (m *DigitalWellbeing) ContainerEcho(stringArg string) *dagger.Container {
	return dag.Container().From("alpine:latest").WithExec([]string{"echo", stringArg})
}

// Returns lines that match a pattern in the files of the provided Directory
func (m *DigitalWellbeing) GrepDir(ctx context.Context, directoryArg *dagger.Directory, pattern string) (string, error) {
	return dag.Container().
		From("alpine:latest").
		WithMountedDirectory("/mnt", directoryArg).
		WithWorkdir("/mnt").
		WithExec([]string{"grep", "-R", pattern, "."}).
		Stdout(ctx)
}
