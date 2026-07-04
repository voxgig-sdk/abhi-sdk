// Typed models for the Abhi SDK.
//
// GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
// params (op.<name>.points[].args.params[]). Field/param types come from the
// canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
// @voxgig/apidef VALID_CANON). Do not edit by hand.

export interface Anime {
  data?: Record<string, any>
  status?: string
}

export type AnimeLoadMatch = Partial<Anime>

export interface Download {
  download_url?: string
  status?: string
}

export type DownloadLoadMatch = Partial<Download>

export interface Fun {
  fact?: string
  status?: string
}

export type FunLoadMatch = Partial<Fun>

export interface Game {
  data?: any[]
  status?: string
}

export type GameListMatch = Partial<Game>

export interface Logo {
  logo_url?: string
  status?: string
}

export type LogoLoadMatch = Partial<Logo>

export interface Tool {
  audio_url?: string
  original_url?: string
  short_url?: string
  status?: string
  url: string
}

export type ToolLoadMatch = Partial<Tool>

export type ToolCreateData = Partial<Tool>

