# UniverStar Tools

This is a tool set used in `UniverStar` project.

## Requirement

- `go` >= `1.25.7`

## Usage

1. Download the latest version:

   ```bash
   curl -L -o ust https://github.com/universtar-org/ust/releases/latest/download/ust
   chmod +x ust
   ```

2. Command list

   | Operation             | Usage                          |
   | --------------------- | ------------------------------ |
   | update repo data      | `ust update /path/to/data-dir` |
   | check repos           | `ust check /path/to/data-dir`  |
   | check user uniqueness | `ust unique [username]`        |

3. Flags
   - `--debug`: enable debug log
   - `--token` or `-t`: GitHub token to avoid rate limit
   - `--help`: show help information

## Acknowledgement

- [`cobra`](https://github.com/spf13/cobra): A Commander for modern Go CLI interactions.
- [`go-yaml`](https://github.com/goccy/go-yaml): YAML support for the Go language.
- [`tint`](https://github.com/lmittmann/tint): Log beautifier.
