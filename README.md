# updater-python

Python updater plugin for Semantic Release.

Updates Python package metadata and versions during Semantic Release.

## Documentation

- Docs (coming soon): <https://github.com/SemRels/semrel/tree/main/docs/plugins/updater-python>
- Template source: <https://github.com/SemRels/plugin-template>

## Repository Layout

`	ext
cmd/plugin/              Plugin entry point
internal/plugin/         Business logic scaffold
internal/grpc/           gRPC transport scaffold
proto/v1                 Symlink to the SemRel protobuf contract
.github/workflows/       CI, release, and security automation
`

## Development

`ash
go build ./cmd/plugin
go test ./...
`

## Configuration Example

`yaml
plugins:
  - name: updater-python
    type: updater
    config:
      pyproject_file: pyproject.toml
      version_files:
        - src/example/__init__.py
      build_backend: hatchling
`

## Status

This repository is bootstrapped from SemRels/plugin-template and is ready for implementation.