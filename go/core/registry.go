package core

var UtilityRegistrar func(u *Utility)

var NewBaseFeatureFunc func() Feature

var NewTestFeatureFunc func() Feature

var NewAnimeEntityFunc func(client *AbhiSDK, entopts map[string]any) AbhiEntity

var NewDownloadEntityFunc func(client *AbhiSDK, entopts map[string]any) AbhiEntity

var NewFunEntityFunc func(client *AbhiSDK, entopts map[string]any) AbhiEntity

var NewGameEntityFunc func(client *AbhiSDK, entopts map[string]any) AbhiEntity

var NewLogoEntityFunc func(client *AbhiSDK, entopts map[string]any) AbhiEntity

var NewToolEntityFunc func(client *AbhiSDK, entopts map[string]any) AbhiEntity

