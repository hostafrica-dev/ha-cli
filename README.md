# HostAfrica VPS CLI

Command-line interface for the HostAfrica VPS API.

## Install

```bash
git clone https://github.com/hostafrica/ha-cli
cd ha-cli
go build -o ha .
```

## Configuration

Set the server URL and auth token via flags or environment variables:

| Flag       | Env var     | Description              |
|------------|-------------|--------------------------|
| `--server` | `HA_SERVER` | API server base URL      |
| `--token`  | `HA_TOKEN`  | Bearer token             |

```bash
export HA_SERVER=https://api.hostafrica.com
export HA_TOKEN=your-bearer-token
```

## Usage

```bash
./ha --help
./ha list-vps-services
./ha get-vps-details --service_id 12345
./ha start-vps --service_id 12345
./ha stop-vps --service_id 12345
./ha reboot --service_id 12345
./ha create-backup --service_id 12345 --mode snapshot
./ha create-snapshot --service_id 12345
```

All commands print pretty-printed JSON to stdout.

## Shell Completion

The CLI supports tab completion for all commands and flags via Cobra's built-in completion support.

**Zsh:**
```zsh
source <(./ha completion zsh)
```

**Bash:**
```bash
source <(./ha completion bash)
```

**Fish:**
```fish
./ha completion fish | source
```

For a permanent setup, add the `source` line to your shell's rc file (e.g. `~/.zshrc`).

## Notes

- `generated.go` is auto-generated from the HostAfrica OpenAPI spec. Do not edit it manually.
- The source for the generator lives in the [ha-proxy-api](https://github.com/hostafrica/ha-proxy-api) repository.
