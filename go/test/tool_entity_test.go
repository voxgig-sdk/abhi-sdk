package sdktest

import (
	"encoding/json"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
	"time"

	sdk "github.com/voxgig-sdk/abhi-sdk"
	"github.com/voxgig-sdk/abhi-sdk/core"

	vs "github.com/voxgig/struct"
)

func TestToolEntity(t *testing.T) {
	t.Run("instance", func(t *testing.T) {
		testsdk := sdk.TestSDK(nil, nil)
		ent := testsdk.Tool(nil)
		if ent == nil {
			t.Fatal("expected non-nil ToolEntity")
		}
	})

	t.Run("basic", func(t *testing.T) {
		setup := toolBasicSetup(nil)
		// Per-op sdk-test-control.json skip — basic test exercises a flow
		// with multiple ops; skipping any op skips the whole flow.
		_mode := "unit"
		if setup.live {
			_mode = "live"
		}
		for _, _op := range []string{"create", "load"} {
			if _shouldSkip, _reason := isControlSkipped("entityOp", "tool." + _op, _mode); _shouldSkip {
				if _reason == "" {
					_reason = "skipped via sdk-test-control.json"
				}
				t.Skip(_reason)
				return
			}
		}
		// The basic flow consumes synthetic IDs from the fixture. In live mode
		// without an *_ENTID env override, those IDs hit the live API and 4xx.
		if setup.syntheticOnly {
			t.Skip("live entity test uses synthetic IDs from fixture — set ABHI_TEST_TOOL_ENTID JSON to run live")
			return
		}
		client := setup.client

		// CREATE
		toolRef01Ent := client.Tool(nil)
		toolRef01Data := core.ToMapAny(vs.GetProp(
			vs.GetPath([]any{"new", "tool"}, setup.data), "tool_ref01"))

		toolRef01DataResult, err := toolRef01Ent.Create(toolRef01Data, nil)
		if err != nil {
			t.Fatalf("create failed: %v", err)
		}
		toolRef01Data = core.ToMapAny(toolRef01DataResult)
		if toolRef01Data == nil {
			t.Fatal("expected create result to be a map")
		}

		// LOAD
		toolRef01MatchDt0 := map[string]any{}
		toolRef01DataDt0Loaded, err := toolRef01Ent.Load(toolRef01MatchDt0, nil)
		if err != nil {
			t.Fatalf("load failed: %v", err)
		}
		if toolRef01DataDt0Loaded == nil {
			t.Fatal("expected load result to be non-nil")
		}

	})
}

func toolBasicSetup(extra map[string]any) *entityTestSetup {
	loadEnvLocal()

	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)

	entityDataFile := filepath.Join(dir, "..", "..", ".sdk", "test", "entity", "tool", "ToolTestData.json")

	entityDataSource, err := os.ReadFile(entityDataFile)
	if err != nil {
		panic("failed to read tool test data: " + err.Error())
	}

	var entityData map[string]any
	if err := json.Unmarshal(entityDataSource, &entityData); err != nil {
		panic("failed to parse tool test data: " + err.Error())
	}

	options := map[string]any{}
	options["entity"] = entityData["existing"]

	client := sdk.TestSDK(options, extra)

	// Generate idmap via transform, matching TS pattern.
	idmap := vs.Transform(
		[]any{"tool01", "tool02", "tool03"},
		map[string]any{
			"`$PACK`": []any{"", map[string]any{
				"`$KEY`": "`$COPY`",
				"`$VAL`": []any{"`$FORMAT`", "upper", "`$COPY`"},
			}},
		},
	)

	// Detect ENTID env override before envOverride consumes it. When live
	// mode is on without a real override, the basic test runs against synthetic
	// IDs from the fixture and 4xx's. Surface this so the test can skip.
	entidEnvRaw := os.Getenv("ABHI_TEST_TOOL_ENTID")
	idmapOverridden := entidEnvRaw != "" && strings.HasPrefix(strings.TrimSpace(entidEnvRaw), "{")

	env := envOverride(map[string]any{
		"ABHI_TEST_TOOL_ENTID": idmap,
		"ABHI_TEST_LIVE":      "FALSE",
		"ABHI_TEST_EXPLAIN":   "FALSE",
		"ABHI_APIKEY":         "NONE",
	})

	idmapResolved := core.ToMapAny(env["ABHI_TEST_TOOL_ENTID"])
	if idmapResolved == nil {
		idmapResolved = core.ToMapAny(idmap)
	}

	if env["ABHI_TEST_LIVE"] == "TRUE" {
		mergedOpts := vs.Merge([]any{
			map[string]any{
				"apikey": env["ABHI_APIKEY"],
			},
			extra,
		})
		client = sdk.NewAbhiSDK(core.ToMapAny(mergedOpts))
	}

	live := env["ABHI_TEST_LIVE"] == "TRUE"
	return &entityTestSetup{
		client:        client,
		data:          entityData,
		idmap:         idmapResolved,
		env:           env,
		explain:       env["ABHI_TEST_EXPLAIN"] == "TRUE",
		live:          live,
		syntheticOnly: live && !idmapOverridden,
		now:           time.Now().UnixMilli(),
	}
}
