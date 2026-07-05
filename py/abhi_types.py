# Typed models for the Abhi SDK.
#
# GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
# params (op.<name>.points[].args.params[]). Field/param types come from the
# canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
# @voxgig/apidef VALID_CANON). Do not edit by hand.
#
# These are TypedDicts, not dataclasses: the SDK ops return/accept plain dicts
# at runtime, and a TypedDict IS a dict shape, so the types match the runtime.
# Optional (req:false) keys are modelled as TypedDict key-optionality
# (total=False), split into a required base + total=False subclass when a type
# has both required and optional keys.

from __future__ import annotations

from typing import TypedDict, Any


class Anime(TypedDict, total=False):
    data: dict
    status: str


class AnimeLoadMatch(TypedDict, total=False):
    data: dict
    status: str


class Download(TypedDict, total=False):
    download_url: str
    status: str


class DownloadLoadMatch(TypedDict, total=False):
    download_url: str
    status: str


class Fun(TypedDict, total=False):
    fact: str
    status: str


class FunLoadMatch(TypedDict, total=False):
    fact: str
    status: str


class Game(TypedDict, total=False):
    data: list
    status: str


class GameListMatch(TypedDict, total=False):
    data: list
    status: str


class Logo(TypedDict, total=False):
    logo_url: str
    status: str


class LogoLoadMatch(TypedDict, total=False):
    logo_url: str
    status: str


class ToolRequired(TypedDict):
    url: str


class Tool(ToolRequired, total=False):
    audio_url: str
    original_url: str
    short_url: str
    status: str


class ToolLoadMatch(TypedDict, total=False):
    audio_url: str
    original_url: str
    short_url: str
    status: str
    url: str


class ToolCreateDataRequired(TypedDict):
    url: str


class ToolCreateData(ToolCreateDataRequired, total=False):
    audio_url: str
    original_url: str
    short_url: str
    status: str
