# Abhi SDK

Free grab-bag of anime, downloads, fun, games, logo generation, and small utilities like URL shortening and TTS

> TypeScript, Python, PHP, Golang, Ruby, Lua SDKs, a CLI, an interactive REPL, and an MCP server for AI agents — all generated from one OpenAPI spec by [@voxgig/sdkgen](https://github.com/voxgig/sdkgen).

## About Abhi API

The Abhi API is a free, all-in-one collection of small public endpoints maintained by [Abhishek Suresh](https://github.com/AbhishekSuresh2). It bundles anime imagery, social-media downloaders, jokes and trivia, mini-games, text-to-logo generators, and a handful of utility tools behind a single host at `https://abhi-api.vercel.app`.

**What you get from the API:**

- Anime image and character endpoints (Itachi, Naruto, Nezuko, Miku, Waifu, and more)
- Spotify and TikTok media downloaders
- Fun content: facts, jokes (dark / dev / general), memes, questions, quotes, roasts
- Mini-games: dare, math, riddle, slots, truth
- Logo generators that render styled text (cyberMask, glitch, gold, leaves, matrix, neon, neonDream, retroWave, shadow)
- Utilities: Bitly and TinyURL shorteners, translation, text-to-speech

The project is published as a single Vercel-hosted service; the catalogue listing on [freepublicapis.com](https://freepublicapis.com/abhi-api) reports CORS as disabled. License, authentication, and rate-limit terms are not documented on the API homepage or docs page.

## Try it

**TypeScript**
```bash
npm install abhi
```

**Python**
```bash
pip install abhi-sdk
```

**PHP**
```bash
composer require voxgig/abhi-sdk
```

**Golang**
```bash
go get github.com/voxgig-sdk/abhi-sdk/go
```

**Ruby**
```bash
gem install abhi-sdk
```

**Lua**
```bash
luarocks install abhi-sdk
```

## 30-second quickstart

### TypeScript

```ts
import { AbhiSDK } from 'abhi'

const client = new AbhiSDK({})

```

See the [TypeScript README](ts/README.md) for the
full guide, or scroll down for the same example in other languages.

## What's in the box

| Surface | Use it for | Path |
| --- | --- | --- |
| **SDK** (TypeScript, Python, PHP, Golang, Ruby, Lua) | App integration | `ts/` `py/` `php/` `go/` `rb/` `lua/` |
| **CLI** | Scripts, CI, ops, one-off API calls | `go-cli/` |
| **MCP server** | AI agents (Claude, Cursor, Cline) | `go-mcp/` |

## Use it from an AI agent (MCP)

The generated MCP server exposes every operation in this SDK as an
[MCP](https://modelcontextprotocol.io) tool that Claude, Cursor or Cline
can call directly. Build and register it:

```bash
cd go-mcp && go build -o abhi-mcp .
```

Then add it to your agent's MCP config (Claude Desktop, Cursor, etc.):

```json
{
  "mcpServers": {
    "abhi": {
      "command": "/abs/path/to/abhi-mcp"
    }
  }
}
```

## Entities

The API exposes 6 entities:

| Entity | Description | API path |
| --- | --- | --- |
| **Anime** | Anime character and image endpoints under `/api/anime/*` (e.g. `astatus`, `couplepp`, `itori`, `itachi`, `loli`, `miku`, `naruto`, `nezuko`, `waifu`). | `/api/anime` |
| **Download** | Social-media media downloaders under `/api/download/*` for Spotify (`/api/download/spotify?url=`) and TikTok (`/api/download/tiktok?url=`). | `/api/download` |
| **Fun** | Lightweight fun content under `/api/fun/*`: `facts`, `jdark`, `jdev`, `jgeneral`, `meme`, `question`, `quotes`, `roast`. | `/api/facts` |
| **Game** | Mini-game prompts and outcomes split across `/api/game/*` and `/api/games/*`: `dare`, `math`, `riddle`, `slots`, `truth`. | `/api/games` |
| **Logo** | Styled text-to-logo image generators under `/api/logo/*` taking a `text` query parameter (`cyberMask`, `glitch`, `gold`, `leaves`, `matrix`, `neon`, `neonDream`, `retroWave`, `shadow`). | `/api/logo` |
| **Tool** | Utility endpoints under `/api/tool/*`: URL shortening via `bitly` and `tinyurl`, plus `translate` and `tts` (text-to-speech). | `/api/shorten` |

Each entity supports the following operations where available: **load**,
**list**, **create**, **update**, and **remove**.

## Quickstart in other languages

### Python

```python
from abhi_sdk import AbhiSDK

client = AbhiSDK({})


# Load a specific anime
anime, err = client.Anime(None).load(
    {"id": "example_id"}, None
)
```

### PHP

```php
<?php
require_once 'abhi_sdk.php';

$client = new AbhiSDK([]);


// Load a specific anime
[$anime, $err] = $client->Anime(null)->load(
    ["id" => "example_id"], null
);
```

### Golang

```go
import sdk "github.com/voxgig-sdk/abhi-sdk/go"

client := sdk.NewAbhiSDK(map[string]any{})

```

### Ruby

```ruby
require_relative "Abhi_sdk"

client = AbhiSDK.new({})


# Load a specific anime
anime, err = client.Anime(nil).load(
  { "id" => "example_id" }, nil
)
```

### Lua

```lua
local sdk = require("abhi_sdk")

local client = sdk.new({})


-- Load a specific anime
local anime, err = client:Anime(nil):load(
  { id = "example_id" }, nil
)
```

## Unit testing in offline mode

Every SDK ships a test mode that swaps the HTTP transport for an
in-memory mock, so unit tests run offline.

### TypeScript

```ts
const client = AbhiSDK.test()
const result = await client.Anime().load({ id: 'test01' })
// result.ok === true, result.data contains mock data
```

### Python

```python
client = AbhiSDK.test(None, None)
result, err = client.Anime(None).load(
    {"id": "test01"}, None
)
```

### PHP

```php
$client = AbhiSDK::test(null, null);
[$result, $err] = $client->Anime(null)->load(
    ["id" => "test01"], null
);
```

### Golang

```go
client := sdk.TestSDK(nil, nil)
result, err := client.Anime(nil).Load(
    map[string]any{"id": "test01"}, nil,
)
```

### Ruby

```ruby
client = AbhiSDK.test(nil, nil)
result, err = client.Anime(nil).load(
  { "id" => "test01" }, nil
)
```

### Lua

```lua
local client = sdk.test(nil, nil)
local result, err = client:Anime(nil):load(
  { id = "test01" }, nil
)
```

## How it works

Every SDK call runs the same five-stage pipeline:

1. **Point** — resolve the API endpoint from the operation definition.
2. **Spec** — build the HTTP specification (URL, method, headers, body).
3. **Request** — send the HTTP request.
4. **Response** — receive and parse the response.
5. **Result** — extract the result data for the caller.

A feature hook fires at each stage (e.g. `PrePoint`, `PreSpec`,
`PreRequest`), so features can inspect or modify the pipeline without
forking the SDK.

### Features

| Feature | Purpose |
| --- | --- |
| **TestFeature** | In-memory mock transport for testing without a live server |

Pass custom features via the `extend` option at construction time.

### Direct and Prepare

For endpoints the entity model doesn't cover, use the low-level methods:

- **`direct(fetchargs)`** — build and send an HTTP request in one step.
- **`prepare(fetchargs)`** — build the request without sending it.

Both accept a map with `path`, `method`, `params`, `query`,
`headers`, and `body`. See the [How-to guides](#how-to-guides) below.

## How-to guides

### Make a direct API call

When the entity interface does not cover an endpoint, use `direct`:

**TypeScript:**
```ts
const result = await client.direct({
  path: '/api/resource/{id}',
  method: 'GET',
  params: { id: 'example' },
})
console.log(result.data)
```

**Python:**
```python
result, err = client.direct({
    "path": "/api/resource/{id}",
    "method": "GET",
    "params": {"id": "example"},
})
```

**PHP:**
```php
[$result, $err] = $client->direct([
    "path" => "/api/resource/{id}",
    "method" => "GET",
    "params" => ["id" => "example"],
]);
```

**Go:**
```go
result, err := client.Direct(map[string]any{
    "path":   "/api/resource/{id}",
    "method": "GET",
    "params": map[string]any{"id": "example"},
})
```

**Ruby:**
```ruby
result, err = client.direct({
  "path" => "/api/resource/{id}",
  "method" => "GET",
  "params" => { "id" => "example" },
})
```

**Lua:**
```lua
local result, err = client:direct({
  path = "/api/resource/{id}",
  method = "GET",
  params = { id = "example" },
})
```

## Per-language documentation

- [TypeScript](ts/README.md)
- [Python](py/README.md)
- [PHP](php/README.md)
- [Golang](go/README.md)
- [Ruby](rb/README.md)
- [Lua](lua/README.md)

## Using the Abhi API

- Upstream: [https://abhi-api.vercel.app](https://abhi-api.vercel.app)
- API docs: [https://abhi-api.vercel.app/docs](https://abhi-api.vercel.app/docs)

---

Generated from the Abhi API OpenAPI spec by [@voxgig/sdkgen](https://github.com/voxgig/sdkgen).
