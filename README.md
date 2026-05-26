# updater-python

Updates the Python package version in the selected packaging backend.

This plugin is distributed as the standalone Go binary `semrel-plugin-updater-python`. Semrel executes the binary as a subprocess, provides plugin configuration through `SEMREL_PLUGIN_*` environment variables, provides release context through `SEMREL_*` environment variables, reads standard output, and treats exit code `0` as success and any non-zero exit code as failure. Install the binary in `~/.semrel/plugins/` or anywhere on your `$PATH`.

## Installation

```bash
go install github.com/SemRels/updater-python/cmd/plugin@latest
```

## Configuration

```yaml
plugins:
  - name: updater-python
    path: ~/.semrel/plugins/semrel-plugin-updater-python
    env:
      SEMREL_PLUGIN_FILE: "pyproject.toml"
      SEMREL_PLUGIN_BACKEND: "pyproject"
```

## `SEMREL_PLUGIN_*` variables

| Name | Required | Description | Default |
| --- | --- | --- | --- |
| `SEMREL_PLUGIN_FILE` | Optional | Path to the Python packaging file to update. | pyproject.toml |
| `SEMREL_PLUGIN_BACKEND` | Optional | Backend format used to locate and update the version. | pyproject |

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
