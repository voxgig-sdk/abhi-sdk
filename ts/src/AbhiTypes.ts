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

export interface AnimeLoadMatch {
  data?: Record<string, any>
  status?: string
}

export interface Download {
  download_url?: string
  status?: string
}

export interface DownloadLoadMatch {
  download_url?: string
  status?: string
}

export interface Fun {
  fact?: string
  status?: string
}

export interface FunLoadMatch {
  fact?: string
  status?: string
}

export interface Game {
  data?: any[]
  status?: string
}

export interface GameListMatch {
  data?: any[]
  status?: string
}

export interface Logo {
  logo_url?: string
  status?: string
}

export interface LogoLoadMatch {
  logo_url?: string
  status?: string
}

export interface Tool {
  audio_url?: string
  original_url?: string
  short_url?: string
  status?: string
  url: string
}

export interface ToolLoadMatch {
  audio_url?: string
  original_url?: string
  short_url?: string
  status?: string
  url?: string
}

export interface ToolCreateData {
  audio_url?: string
  original_url?: string
  short_url?: string
  status?: string
  url: string
}

