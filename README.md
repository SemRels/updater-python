# updater-python

[![Latest Release](https://img.shields.io/github/v/release/SemRels/updater-python?label=version&color=blue)](https://github.com/SemRels/updater-python/releases/latest)
[![CI](https://github.com/SemRels/updater-python/actions/workflows/ci.yml/badge.svg)](https://github.com/SemRels/updater-python/actions/workflows/ci.yml)
[![License](https://img.shields.io/github/license/SemRels/updater-python)](LICENSE)
[![OpenSSF Scorecard](https://api.scorecard.dev/projects/github.com/SemRels/updater-python/badge)](https://scorecard.dev/viewer/?uri=github.com/SemRels/updater-python)
[![REUSE status](https://api.reuse.software/badge/github.com/SemRels/updater-python)](https://api.reuse.software/info/github.com/SemRels/updater-python)

Updates the Python package version in the selected packaging backend.

This plugin is distributed as the standalone Go binary `semrel-plugin-updater-python`. Semrel executes the binary as a subprocess, provides plugin configuration through `SEMREL_PLUGIN_*` environment variables, provides release context through `SEMREL_*` environment variables, reads standard output, and treats exit code `0` as success and any non-zero exit code as failure. Install the binary in `~/.semrel/plugins/` or anywhere on your `$PATH`.

## Installation

### Binary

```bash
go install github.com/SemRels/updater-python/cmd/plugin@latest
```

### Docker

Pre-built, multi-platform images (linux/amd64, linux/arm64) are published to the GitHub Container Registry on every release:

```bash
docker pull ghcr.io/semrels/updater-python:latest
```

Images are signed with [cosign](https://github.com/sigstore/cosign) and include a full SBOM attestation. Verify the signature:

```bash
cosign verify ghcr.io/semrels/updater-python:latest \
  --certificate-identity-regexp 'https://github.com/SemRels/updater-python/.github/workflows/container.yml.*' \
  --certificate-oidc-issuer https://token.actions.githubusercontent.com
```


## Configuration

```yaml
plugins:
  - uses: "@semrel/updater-python"
    phase: pre-tag
    args:
      file: "pyproject.toml"
      backend: "pyproject"
```

## Options

Every option can be set as an entry under `args` in `.semrel.yaml`. semrel passes each entry to the plugin as the environment variable shown in the table, so you can also set that variable directly (for example in your CI environment) instead of using `args`.

| Argument | Environment variable | Required | Description | Default |
| --- | --- | --- | --- | --- |
| `file` | `SEMREL_PLUGIN_FILE` | Optional | Path to the Python packaging file to update. | pyproject.toml |
| `backend` | `SEMREL_PLUGIN_BACKEND` | Optional | Backend format used to locate and update the version. | pyproject |

## `SEMREL_*` release context used

| Variable | Description |
| --- | --- |
| `SEMREL_VERSION` | Resolved release version for the current run. |
| `SEMREL_NEXT_VERSION` | Next version computed by semrel for the release. |
| `SEMREL_DRY_RUN` | Whether semrel is running in dry-run mode. |

## Example behavior

The plugin updates the version in `pyproject.toml`, `setup.cfg`, or `version.py` depending on the selected backend.

## License

Apache-2.0
