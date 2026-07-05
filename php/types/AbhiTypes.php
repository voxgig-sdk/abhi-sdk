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

/** Request payload for Anime#load. */
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

/** Request payload for Download#load. */
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

/** Request payload for Fun#load. */
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

/** Request payload for Game#list. */
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

/** Request payload for Logo#load. */
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

/** Request payload for Tool#load. */
class ToolLoadMatch
{
    public ?string $audio_url = null;
    public ?string $original_url = null;
    public ?string $short_url = null;
    public ?string $status = null;
    public ?string $url = null;
}

/** Request payload for Tool#create. */
class ToolCreateData
{
    public ?string $audio_url = null;
    public ?string $original_url = null;
    public ?string $short_url = null;
    public ?string $status = null;
    public string $url;
}

