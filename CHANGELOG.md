# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## v2.1.2

### Changed

* _Deprecation_ The flag `substreams-sink-kv inject --listen-addr=...` must now be used like `substreams-sink-kv inject --server-listen-addr=...`

### Fixed

* Added back missing flag `--server-api-prefix` to define API prefix when server mode is active.

## v2.1.1

### Added

* Made argument `<manifest>` optional again, in which case `.` is assumed.

* Added ability to provide `<manifest>` argument as a directory, enabling the possibility to do `substreams-sink-kv inject <endpoint> <dsn> .`

## v2.1.0

### Highlights

We change the arguments accepted by `substreams-sink-kv inject` where now the first argument must be the endpoint to consume from. This was previously the provided through the flag `-e, --endpoint`. This has been done to aligned with other sinks supported by StreamingFast team.

Before:

```bash
# Implicit default endpoint
substreams-sink-kv inject <dsn> substreams.spkg

# Explicit endpoint
substreams-sink-kv inject -e mainnet.eth.streamingfast.io:443 <dsn> substreams.spkg
```

After:

```bash
substreams-sink-kv inject mainnet.eth.streamingfast.io:443 <dsn> substreams.spkg
```

### Changed

- **Breaking** Flag `-e, --endpoint` has been replaced by a mandatory positional argument instead, see highlights for upgrade procedure.

## v2.0.0

### Highlights

This release drops support for Substreams RPC protocol `sf.substreams.v1` and switch to Substreams RPC protocol `sf.substreams.rpc.v2`. As a end user, right now the transition is seamless. All StreamingFast endpoints have been updated to to support the legacy Substreams RPC protocol `sf.substreams.v1` as well as the newer Substreams RPC protocol `sf.substreams.rpc.v2`.

Support for legacy Substreams RPC protocol `sf.substreams.v1` is expected to end by June 6 2023. What this means is that you will need to update to at least this release if you are running `substreams-sink-kv` in production. Otherwise, after this date, your current binary will stop working and will return errors that `sf.substreams.v1.Blocks` is not supported on the endpoint.

From a database and operator standpoint, this binary is **fully** backward compatible with your current schema. Updating to this binary will continue to sink just like if you used a prior release.

#### Retryable Errors

The errors coming from KV store are **not** retried anymore and will stop the binary immediately.

### Added

- Added `--infinite-retry` to never exit on error and retry indefinitely instead.

- Added `--development-mode` to run in development mode.

    > **Warning** You should use that flag for testing purposes, development mode drastically reduce performance you get from the server.

- Added `--final-blocks-only` to only deal with final (irreversible) blocks.

- Added `--undo-buffer-size` (defaults to 12) that deals with re-org handling, this will delayed your live block by this amount.

- Added `--live-block-time-delta` (defaults to 300s) that determine if a block is considered "live" or "historical". The if `time.Now() - block's timestamp` is lower or equal to `--live-block-time-delta` then the block is considered live.

## [0.1.3](https://github.com/streamingfast/substreams-sink-kv/releases/tag/v0.1.3)

### Added

* Added WASM Query Support
* `GenericService` example `./examples/generic-service`
* `WASMQueryService` example `./examples/wasm-query-service`
* Panic handling in wasm query engine

### Changed

* Removed support for GRPC `stream`, only support request response
* Remove `substream-sink-kv run` CLI and added `substream-sink-kv inject` & `substream-sink-kv serve`
* Move `rust` crate to its own repo `https://github.com/streamingfast/substreams-sink-rs`

## [0.1.2](https://github.com/streamingfast/substreams-sink-kv/releases/tag/v0.1.2)

### Changed

* Prevent hanging on connection error.

### Fixed

* Fixed wrong accepted module's output type `sf.substreams.kv.v1.KVOperations`, correct value is `sf.substreams.sink.kv.v1.KVOperations`.

## [0.1.1](https://github.com/streamingfast/substreams-sink-kv/releases/tag/v0.1.1)

* Initial release