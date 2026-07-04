# frozen_string_literal: true

# Typed models for the Abhi SDK.
#
# GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
# params (op.<name>.points[].args.params[]). Member types come from the
# canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
# @voxgig/apidef VALID_CANON). Ruby types are unenforced; these YARD
# annotations document the shapes. Do not edit by hand.

# Anime entity data model.
#
# @!attribute [rw] data
#   @return [Hash, nil]
#
# @!attribute [rw] status
#   @return [String, nil]
Anime = Struct.new(
  :data,
  :status,
  keyword_init: true
)

# Match filter for Anime#load (any subset of Anime fields).
#
# @!attribute [rw] data
#   @return [Hash, nil]
#
# @!attribute [rw] status
#   @return [String, nil]
AnimeLoadMatch = Struct.new(
  :data,
  :status,
  keyword_init: true
)

# Download entity data model.
#
# @!attribute [rw] download_url
#   @return [String, nil]
#
# @!attribute [rw] status
#   @return [String, nil]
Download = Struct.new(
  :download_url,
  :status,
  keyword_init: true
)

# Match filter for Download#load (any subset of Download fields).
#
# @!attribute [rw] download_url
#   @return [String, nil]
#
# @!attribute [rw] status
#   @return [String, nil]
DownloadLoadMatch = Struct.new(
  :download_url,
  :status,
  keyword_init: true
)

# Fun entity data model.
#
# @!attribute [rw] fact
#   @return [String, nil]
#
# @!attribute [rw] status
#   @return [String, nil]
Fun = Struct.new(
  :fact,
  :status,
  keyword_init: true
)

# Match filter for Fun#load (any subset of Fun fields).
#
# @!attribute [rw] fact
#   @return [String, nil]
#
# @!attribute [rw] status
#   @return [String, nil]
FunLoadMatch = Struct.new(
  :fact,
  :status,
  keyword_init: true
)

# Game entity data model.
#
# @!attribute [rw] data
#   @return [Array, nil]
#
# @!attribute [rw] status
#   @return [String, nil]
Game = Struct.new(
  :data,
  :status,
  keyword_init: true
)

# Match filter for Game#list (any subset of Game fields).
#
# @!attribute [rw] data
#   @return [Array, nil]
#
# @!attribute [rw] status
#   @return [String, nil]
GameListMatch = Struct.new(
  :data,
  :status,
  keyword_init: true
)

# Logo entity data model.
#
# @!attribute [rw] logo_url
#   @return [String, nil]
#
# @!attribute [rw] status
#   @return [String, nil]
Logo = Struct.new(
  :logo_url,
  :status,
  keyword_init: true
)

# Match filter for Logo#load (any subset of Logo fields).
#
# @!attribute [rw] logo_url
#   @return [String, nil]
#
# @!attribute [rw] status
#   @return [String, nil]
LogoLoadMatch = Struct.new(
  :logo_url,
  :status,
  keyword_init: true
)

# Tool entity data model.
#
# @!attribute [rw] audio_url
#   @return [String, nil]
#
# @!attribute [rw] original_url
#   @return [String, nil]
#
# @!attribute [rw] short_url
#   @return [String, nil]
#
# @!attribute [rw] status
#   @return [String, nil]
#
# @!attribute [rw] url
#   @return [String]
Tool = Struct.new(
  :audio_url,
  :original_url,
  :short_url,
  :status,
  :url,
  keyword_init: true
)

# Match filter for Tool#load (any subset of Tool fields).
#
# @!attribute [rw] audio_url
#   @return [String, nil]
#
# @!attribute [rw] original_url
#   @return [String, nil]
#
# @!attribute [rw] short_url
#   @return [String, nil]
#
# @!attribute [rw] status
#   @return [String, nil]
#
# @!attribute [rw] url
#   @return [String, nil]
ToolLoadMatch = Struct.new(
  :audio_url,
  :original_url,
  :short_url,
  :status,
  :url,
  keyword_init: true
)

# Match filter for Tool#create (any subset of Tool fields).
#
# @!attribute [rw] audio_url
#   @return [String, nil]
#
# @!attribute [rw] original_url
#   @return [String, nil]
#
# @!attribute [rw] short_url
#   @return [String, nil]
#
# @!attribute [rw] status
#   @return [String, nil]
#
# @!attribute [rw] url
#   @return [String, nil]
ToolCreateData = Struct.new(
  :audio_url,
  :original_url,
  :short_url,
  :status,
  :url,
  keyword_init: true
)

