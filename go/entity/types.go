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

// AnimeLoadMatch mirrors the anime fields as an all-optional match
// filter (Go analog of Partial<Anime>).
type AnimeLoadMatch struct {
	Data *map[string]any `json:"data,omitempty"`
	Status *string `json:"status,omitempty"`
}

// Download is the typed data model for the download entity.
type Download struct {
	DownloadUrl *string `json:"download_url,omitempty"`
	Status *string `json:"status,omitempty"`
}

// DownloadLoadMatch mirrors the download fields as an all-optional match
// filter (Go analog of Partial<Download>).
type DownloadLoadMatch struct {
	DownloadUrl *string `json:"download_url,omitempty"`
	Status *string `json:"status,omitempty"`
}

// Fun is the typed data model for the fun entity.
type Fun struct {
	Fact *string `json:"fact,omitempty"`
	Status *string `json:"status,omitempty"`
}

// FunLoadMatch mirrors the fun fields as an all-optional match
// filter (Go analog of Partial<Fun>).
type FunLoadMatch struct {
	Fact *string `json:"fact,omitempty"`
	Status *string `json:"status,omitempty"`
}

// Game is the typed data model for the game entity.
type Game struct {
	Data *[]any `json:"data,omitempty"`
	Status *string `json:"status,omitempty"`
}

// GameListMatch mirrors the game fields as an all-optional match
// filter (Go analog of Partial<Game>).
type GameListMatch struct {
	Data *[]any `json:"data,omitempty"`
	Status *string `json:"status,omitempty"`
}

// Logo is the typed data model for the logo entity.
type Logo struct {
	LogoUrl *string `json:"logo_url,omitempty"`
	Status *string `json:"status,omitempty"`
}

// LogoLoadMatch mirrors the logo fields as an all-optional match
// filter (Go analog of Partial<Logo>).
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

// ToolLoadMatch mirrors the tool fields as an all-optional match
// filter (Go analog of Partial<Tool>).
type ToolLoadMatch struct {
	AudioUrl *string `json:"audio_url,omitempty"`
	OriginalUrl *string `json:"original_url,omitempty"`
	ShortUrl *string `json:"short_url,omitempty"`
	Status *string `json:"status,omitempty"`
	Url *string `json:"url,omitempty"`
}

// ToolCreateData mirrors the tool fields as an all-optional match
// filter (Go analog of Partial<Tool>).
type ToolCreateData struct {
	AudioUrl *string `json:"audio_url,omitempty"`
	OriginalUrl *string `json:"original_url,omitempty"`
	ShortUrl *string `json:"short_url,omitempty"`
	Status *string `json:"status,omitempty"`
	Url *string `json:"url,omitempty"`
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
