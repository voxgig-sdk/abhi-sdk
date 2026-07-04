<?php
declare(strict_types=1);

// Typed models for the Abhi SDK.
//
// GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
// params (op.<name>.points[].args.params[]). Field/param types come from the
// canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
// @voxgig/apidef VALID_CANON). Do not edit by hand.
//
// These are documentation-grade value objects (PHP 8 typed properties),
// registered on the composer classmap autoload. The SDK boundary exchanges
// assoc-arrays; these classes name the shapes for tooling and typed callers.

/** Anime entity data model. */
class Anime
{
    public ?array $data = null;
    public ?string $status = null;
}

/** Match filter for Anime#load (any subset of Anime fields). */
class AnimeLoadMatch
{
    public ?array $data = null;
    public ?string $status = null;
}

/** Download entity data model. */
class Download
{
    public ?string $download_url = null;
    public ?string $status = null;
}

/** Match filter for Download#load (any subset of Download fields). */
class DownloadLoadMatch
{
    public ?string $download_url = null;
    public ?string $status = null;
}

/** Fun entity data model. */
class Fun
{
    public ?string $fact = null;
    public ?string $status = null;
}

/** Match filter for Fun#load (any subset of Fun fields). */
class FunLoadMatch
{
    public ?string $fact = null;
    public ?string $status = null;
}

/** Game entity data model. */
class Game
{
    public ?array $data = null;
    public ?string $status = null;
}

/** Match filter for Game#list (any subset of Game fields). */
class GameListMatch
{
    public ?array $data = null;
    public ?string $status = null;
}

/** Logo entity data model. */
class Logo
{
    public ?string $logo_url = null;
    public ?string $status = null;
}

/** Match filter for Logo#load (any subset of Logo fields). */
class LogoLoadMatch
{
    public ?string $logo_url = null;
    public ?string $status = null;
}

/** Tool entity data model. */
class Tool
{
    public ?string $audio_url = null;
    public ?string $original_url = null;
    public ?string $short_url = null;
    public ?string $status = null;
    public string $url;
}

/** Match filter for Tool#load (any subset of Tool fields). */
class ToolLoadMatch
{
    public ?string $audio_url = null;
    public ?string $original_url = null;
    public ?string $short_url = null;
    public ?string $status = null;
    public ?string $url = null;
}

/** Match filter for Tool#create (any subset of Tool fields). */
class ToolCreateData
{
    public ?string $audio_url = null;
    public ?string $original_url = null;
    public ?string $short_url = null;
    public ?string $status = null;
    public ?string $url = null;
}

