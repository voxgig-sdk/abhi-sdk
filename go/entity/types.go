// Typed models for the Abhi SDK.
//
// GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
// params (op.<name>.points[].args.params[]). Field/param types come from the
// canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
// @voxgig/apidef VALID_CANON). Do not edit by hand.
package entity

import "encoding/json"

// Anime is the typed data model for the anime entity.
type Anime struct {
	Data *map[string]any `json:"data,omitempty"`
	Status *string `json:"status,omitempty"`
}

// AnimeLoadMatch is the typed request payload for Anime.LoadTyped.
type AnimeLoadMatch struct {
	Data *map[string]any `json:"data,omitempty"`
	Status *string `json:"status,omitempty"`
}

// Download is the typed data model for the download entity.
type Download struct {
	DownloadUrl *string `json:"download_url,omitempty"`
	Status *string `json:"status,omitempty"`
}

// DownloadLoadMatch is the typed request payload for Download.LoadTyped.
type DownloadLoadMatch struct {
	DownloadUrl *string `json:"download_url,omitempty"`
	Status *string `json:"status,omitempty"`
}

// Fun is the typed data model for the fun entity.
type Fun struct {
	Fact *string `json:"fact,omitempty"`
	Status *string `json:"status,omitempty"`
}

// FunLoadMatch is the typed request payload for Fun.LoadTyped.
type FunLoadMatch struct {
	Fact *string `json:"fact,omitempty"`
	Status *string `json:"status,omitempty"`
}

// Game is the typed data model for the game entity.
type Game struct {
	Data *[]any `json:"data,omitempty"`
	Status *string `json:"status,omitempty"`
}

// GameListMatch is the typed request payload for Game.ListTyped.
type GameListMatch struct {
	Data *[]any `json:"data,omitempty"`
	Status *string `json:"status,omitempty"`
}

// Logo is the typed data model for the logo entity.
type Logo struct {
	LogoUrl *string `json:"logo_url,omitempty"`
	Status *string `json:"status,omitempty"`
}

// LogoLoadMatch is the typed request payload for Logo.LoadTyped.
type LogoLoadMatch struct {
	LogoUrl *string `json:"logo_url,omitempty"`
	Status *string `json:"status,omitempty"`
}

// Tool is the typed data model for the tool entity.
type Tool struct {
	AudioUrl *string `json:"audio_url,omitempty"`
	OriginalUrl *string `json:"original_url,omitempty"`
	ShortUrl *string `json:"short_url,omitempty"`
	Status *string `json:"status,omitempty"`
	Url string `json:"url"`
}

// ToolLoadMatch is the typed request payload for Tool.LoadTyped.
type ToolLoadMatch struct {
	AudioUrl *string `json:"audio_url,omitempty"`
	OriginalUrl *string `json:"original_url,omitempty"`
	ShortUrl *string `json:"short_url,omitempty"`
	Status *string `json:"status,omitempty"`
	Url *string `json:"url,omitempty"`
}

// ToolCreateData is the typed request payload for Tool.CreateTyped.
type ToolCreateData struct {
	AudioUrl *string `json:"audio_url,omitempty"`
	OriginalUrl *string `json:"original_url,omitempty"`
	ShortUrl *string `json:"short_url,omitempty"`
	Status *string `json:"status,omitempty"`
	Url string `json:"url"`
}

// asMap turns a typed request/data struct into the map[string]any the
// runtime op pipeline consumes, honouring the json tags above.
func asMap(v any) map[string]any {
	out := map[string]any{}
	b, err := json.Marshal(v)
	if err != nil {
		return out
	}
	_ = json.Unmarshal(b, &out)
	return out
}

// typedFrom decodes a runtime value (a map[string]any produced by the op
// pipeline) into a typed model T via a JSON round-trip. On any error it
// returns the zero value of T; the op's own (value, error) tuple carries the
// real error.
func typedFrom[T any](v any) T {
	var out T
	if v == nil {
		return out
	}
	b, err := json.Marshal(v)
	if err != nil {
		return out
	}
	_ = json.Unmarshal(b, &out)
	return out
}

// typedSliceFrom decodes a runtime list value ([]any of maps) into a typed
// slice []T via a JSON round-trip, for list ops.
func typedSliceFrom[T any](v any) []T {
	var out []T
	if v == nil {
		return out
	}
	b, err := json.Marshal(v)
	if err != nil {
		return out
	}
	_ = json.Unmarshal(b, &out)
	return out
}
