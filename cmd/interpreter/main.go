//go:build js && wasm

package main

import (
	"fmt"
	"syscall/js"
	runtimev1 "github.com/cerbos/cerbos/api/genpb/cerbos/runtime/v1"
	"google.golang.org/protobuf/encoding/protojson"
	enginev1 "github.com/cerbos/cerbos/api/genpb/cerbos/engine/v1"
	"github.com/cerbos/cerbos/internal/engine"
	"github.com/cerbos/cerbos/internal/schema"
	"github.com/cerbos/cerbos/internal/engine/tracer"
	"context"
)

func Authz(this js.Value, args []js.Value) interface{} {
	result := make(map[string]interface{})

	rps := new(runtimev1.RunnablePolicySet)
	policyStr := args[0].String()
	err := protojson.Unmarshal([]byte(policyStr), rps)
	if err != nil {
		fmt.Println("Failed to parse policy")
		result["error"] = fmt.Errorf("coult not parse policy: %w", err)
		return result
	}

	input := new(enginev1.CheckInput)
	inputStr := args[1].String()
	err = protojson.Unmarshal([]byte(inputStr), input)
	if err != nil {
		fmt.Println("Failed to parse check input")
		result["error"] = fmt.Errorf("coult not parse check input: %w", err)
		return result
	}

	eval := engine.NewEvaluator(rps, schema.NewNopManager())
	ctx := context.Background()
	have, err := eval.Evaluate(ctx, tracer.Start(nil), input)
	for k, v := range have.Effects {
		result[k] = v.Effect.String()
	}

	return result
}

func main() {
	js.Global().Set("Authz", js.FuncOf(Authz))

	<-make(chan bool)
}
