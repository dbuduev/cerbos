package main

import (
	"testing"
	_ "embed"
	"google.golang.org/protobuf/encoding/protojson"
	runtimev1 "github.com/cerbos/cerbos/api/genpb/cerbos/runtime/v1"
	"github.com/stretchr/testify/require"
	"github.com/cerbos/cerbos/internal/engine"
	"github.com/cerbos/cerbos/internal/schema"
	"context"
	"github.com/cerbos/cerbos/internal/engine/tracer"
	"encoding/json"
	privatev1 "github.com/cerbos/cerbos/api/genpb/cerbos/private/v1"
)

//go:embed testdata/rps.yaml
var rpsBytes []byte

//go:embed testdata/input.json
var inputBytes []byte

func TestInterpreter(t *testing.T) {
	is := require.New(t)
	rps := new(runtimev1.RunnablePolicySet)
	err := protojson.Unmarshal(rpsBytes, rps)
	is.NoError(err)
	tc := new(privatev1.EngineTestCase)
	err = protojson.Unmarshal(inputBytes, tc)
	err = json.Unmarshal(inputBytes, tc)

	eval := engine.NewEvaluator(rps, schema.NewNopManager())
	ctx := context.Background()
	for i, input := range tc.Inputs {
		have, err := eval.Evaluate(ctx, tracer.Start(nil), input)
		is.NoError(err)
		is.Empty(have.ValidationErrors)

		want := tc.WantOutputs[i]
		is.Equal(len(want.Actions), len(have.Effects))
		for k, v := range have.Effects {
			is.Equal(want.Actions[k].Effect.String(), v.Effect.String())
		}
		is.Equal(len(want.EffectiveDerivedRoles), len(have.EffectiveDerivedRoles), "Number of effective derived roles")
		for _, role := range want.EffectiveDerivedRoles {
			is.Contains(have.EffectiveDerivedRoles, role)
		}
	}
}
