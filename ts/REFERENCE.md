# Abhi TypeScript SDK Reference

Complete API reference for the Abhi TypeScript SDK.


## AbhiSDK

### Constructor

```ts
new AbhiSDK(options?: object)
```

Create a new SDK client instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `options` | `object` | SDK configuration options. |
| `options.base` | `string` | Base URL for API requests. |
| `options.prefix` | `string` | URL prefix appended after base. |
| `options.suffix` | `string` | URL suffix appended after path. |
| `options.headers` | `object` | Custom headers for all requests. |
| `options.feature` | `object` | Feature configuration. |
| `options.system` | `object` | System overrides (e.g. custom fetch). |


### Static Methods

#### `AbhiSDK.test(testopts?, sdkopts?)`

Create a test client with mock features active.

```ts
const client = AbhiSDK.test()
```

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `testopts` | `object` | Test feature options. |
| `sdkopts` | `object` | Additional SDK options merged with test defaults. |

**Returns:** `AbhiSDK` instance in test mode.


### Instance Methods

#### `Anime(data?: object)`

Create a new `Anime` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `AnimeEntity` instance.

#### `Download(data?: object)`

Create a new `Download` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `DownloadEntity` instance.

#### `Fun(data?: object)`

Create a new `Fun` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `FunEntity` instance.

#### `Game(data?: object)`

Create a new `Game` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `GameEntity` instance.

#### `Logo(data?: object)`

Create a new `Logo` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `LogoEntity` instance.

#### `Tool(data?: object)`

Create a new `Tool` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `ToolEntity` instance.

#### `options()`

Return a deep copy of the current SDK options.

**Returns:** `object`

#### `utility()`

Return a copy of the SDK utility object.

**Returns:** `object`

#### `direct(fetchargs?: object)`

Make a direct HTTP request to any API endpoint.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `fetchargs.path` | `string` | URL path with optional `{param}` placeholders. |
| `fetchargs.method` | `string` | HTTP method (default: `GET`). |
| `fetchargs.params` | `object` | Path parameter values for `{param}` substitution. |
| `fetchargs.query` | `object` | Query string parameters. |
| `fetchargs.headers` | `object` | Request headers (merged with defaults). |
| `fetchargs.body` | `any` | Request body (objects are JSON-serialized). |
| `fetchargs.ctrl` | `object` | Control options (e.g. `{ explain: true }`). |

**Returns:** `Promise<{ ok, status, headers, data } | Error>`

#### `prepare(fetchargs?: object)`

Prepare a fetch definition without sending the request. Accepts the
same parameters as `direct()`.

**Returns:** `Promise<{ url, method, headers, body } | Error>`

#### `tester(testopts?, sdkopts?)`

Alias for `AbhiSDK.test()`.

**Returns:** `AbhiSDK` instance in test mode.


---

## AnimeEntity

```ts
const anime = client.Anime()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `data` | `Record<string, any>` | No |  |
| `status` | `string` | No |  |

### Operations

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.Anime().load()
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `AnimeEntity` instance with the same client and
options.

#### `client()`

Return the parent `AbhiSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## DownloadEntity

```ts
const download = client.Download()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `download_url` | `string` | No |  |
| `status` | `string` | No |  |

### Operations

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.Download().load()
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `DownloadEntity` instance with the same client and
options.

#### `client()`

Return the parent `AbhiSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## FunEntity

```ts
const fun = client.Fun()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `fact` | `string` | No |  |
| `status` | `string` | No |  |

### Operations

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.Fun().load()
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `FunEntity` instance with the same client and
options.

#### `client()`

Return the parent `AbhiSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## GameEntity

```ts
const game = client.Game()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `data` | `any[]` | No |  |
| `status` | `string` | No |  |

### Operations

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.Game().list()
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `GameEntity` instance with the same client and
options.

#### `client()`

Return the parent `AbhiSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## LogoEntity

```ts
const logo = client.Logo()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `logo_url` | `string` | No |  |
| `status` | `string` | No |  |

### Operations

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.Logo().load()
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `LogoEntity` instance with the same client and
options.

#### `client()`

Return the parent `AbhiSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## ToolEntity

```ts
const tool = client.Tool()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `audio_url` | `string` | No |  |
| `original_url` | `string` | No |  |
| `short_url` | `string` | No |  |
| `status` | `string` | No |  |
| `url` | `string` | Yes |  |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.Tool().create({
  url: /* string */,
})
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.Tool().load()
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `ToolEntity` instance with the same client and
options.

#### `client()`

Return the parent `AbhiSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## Features

| Feature | Version | Description |
| --- | --- | --- |
| `test` | 0.0.1 | In-memory mock transport for testing without a live server |


Features are activated via the `feature` option:

```ts
const client = new AbhiSDK({
  feature: {
    test: { active: true },
  }
})
```

