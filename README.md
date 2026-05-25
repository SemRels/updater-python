# updater-python

PyPI package updater plugin for SemRel.

Updates Python package versions for projects published to PyPI.

## Documentation

- SemRel docs (planned): <https://github.com/SemRels/semrel/tree/main/docs/plugins/updater-python>
- Plugin template: <https://github.com/SemRels/plugin-template>
- Registry: <https://registry.semrel.io>

## Repository Layout

~~~text
cmd/plugin/              Plugin entry point
internal/plugin/         Business logic scaffold
internal/grpc/           gRPC transport scaffold
proto/v1                 Symlink to the SemRel protobuf contract
.github/workflows/       CI, release, and security automation
~~~

## Development

~~~bash
go build ./cmd/plugin
go test ./...
~~~

## Configuration Example

~~~yaml
plugins:
  - name: updater-python
    type: updater
    config:
      pyproject_file: pyproject.toml
      version_files:
        - src/package/__init__.py
      repository_url: https://upload.pypi.org/legacy/
~~~

## Status

This repository is bootstrapped from SemRels/plugin-template and is ready for implementation.
