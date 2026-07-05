-- Typed models for the Abhi SDK (LuaLS annotations).
--
-- GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
-- params (op.<name>.points[].args.params[]). Field/param types come from the
-- canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
-- @voxgig/apidef VALID_CANON). Annotations only — no runtime effect. Do not
-- edit by hand.

---@class Anime
---@field data? table
---@field status? string

---@class AnimeLoadMatch
---@field data? table
---@field status? string

---@class Download
---@field download_url? string
---@field status? string

---@class DownloadLoadMatch
---@field download_url? string
---@field status? string

---@class Fun
---@field fact? string
---@field status? string

---@class FunLoadMatch
---@field fact? string
---@field status? string

---@class Game
---@field data? table
---@field status? string

---@class GameListMatch
---@field data? table
---@field status? string

---@class Logo
---@field logo_url? string
---@field status? string

---@class LogoLoadMatch
---@field logo_url? string
---@field status? string

---@class Tool
---@field audio_url? string
---@field original_url? string
---@field short_url? string
---@field status? string
---@field url string

---@class ToolLoadMatch
---@field audio_url? string
---@field original_url? string
---@field short_url? string
---@field status? string
---@field url? string

---@class ToolCreateData
---@field audio_url? string
---@field original_url? string
---@field short_url? string
---@field status? string
---@field url string

local M = {}

return M
