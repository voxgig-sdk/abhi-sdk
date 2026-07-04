# Typed models for the Abhi SDK.
#
# GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
# params (op.<name>.points[].args.params[]). Field/param types come from the
# canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
# @voxgig/apidef VALID_CANON). Do not edit by hand.

from __future__ import annotations

from dataclasses import dataclass
from typing import Optional, Any


@dataclass
class Anime:
    data: Optional[dict] = None
    status: Optional[str] = None


@dataclass
class AnimeLoadMatch:
    data: Optional[dict] = None
    status: Optional[str] = None


@dataclass
class Download:
    download_url: Optional[str] = None
    status: Optional[str] = None


@dataclass
class DownloadLoadMatch:
    download_url: Optional[str] = None
    status: Optional[str] = None


@dataclass
class Fun:
    fact: Optional[str] = None
    status: Optional[str] = None


@dataclass
class FunLoadMatch:
    fact: Optional[str] = None
    status: Optional[str] = None


@dataclass
class Game:
    data: Optional[list] = None
    status: Optional[str] = None


@dataclass
class GameListMatch:
    data: Optional[list] = None
    status: Optional[str] = None


@dataclass
class Logo:
    logo_url: Optional[str] = None
    status: Optional[str] = None


@dataclass
class LogoLoadMatch:
    logo_url: Optional[str] = None
    status: Optional[str] = None


@dataclass
class Tool:
    url: str
    audio_url: Optional[str] = None
    original_url: Optional[str] = None
    short_url: Optional[str] = None
    status: Optional[str] = None


@dataclass
class ToolLoadMatch:
    audio_url: Optional[str] = None
    original_url: Optional[str] = None
    short_url: Optional[str] = None
    status: Optional[str] = None
    url: Optional[str] = None


@dataclass
class ToolCreateData:
    audio_url: Optional[str] = None
    original_url: Optional[str] = None
    short_url: Optional[str] = None
    status: Optional[str] = None
    url: Optional[str] = None

