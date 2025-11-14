// Copyright 2021-2025 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package svc

import (
	"context"
	"encoding/base64"
	"fmt"

	svcv1 "github.com/cerbos/cerbos/api/genpb/authzen/authorization/v1"
	effectv1 "github.com/cerbos/cerbos/api/genpb/cerbos/effect/v1"
	enginev1 "github.com/cerbos/cerbos/api/genpb/cerbos/engine/v1"
	requestv1 "github.com/cerbos/cerbos/api/genpb/cerbos/request/v1"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/structpb"
)

var _ svcv1.AuthorizationServiceServer = (*AuthzenAuthorizationService)(nil)

// AuthzenAuthorizationService implements the policy checking service.
type AuthzenAuthorizationService struct {
	svc *CerbosService
	*svcv1.UnimplementedAuthorizationServiceServer
}

func NewAuthzenAuthorizationService(svc *CerbosService) *AuthzenAuthorizationService {
	return &AuthzenAuthorizationService{
		svc:                                     svc,
		UnimplementedAuthorizationServiceServer: &svcv1.UnimplementedAuthorizationServiceServer{}}
}

// AccessEvaluation implements authorizationv1.AuthorizationServiceServer.
func (aas *AuthzenAuthorizationService) AccessEvaluation(ctx context.Context, r *svcv1.AccessEvaluationRequest) (*svcv1.AccessEvaluationResponse, error) {
	req, err := toCheckResourcesRequest(r)
	if err != nil {
		return nil, err
	}
	resp, err := aas.svc.CheckResources(ctx, req)
	if err != nil {
		return nil, err
	}
	respAsValue, err := recodeToValue(resp)
	if err != nil {
		return nil, err
	}
	return &svcv1.AccessEvaluationResponse{
		Decision: resp.Results[0].Actions[req.Resources[0].Actions[0]] == effectv1.Effect_EFFECT_ALLOW,
		Context: &svcv1.AccessEvaluationResponse_Context{
			Id: resp.RequestId,
			ReasonUser: &svcv1.AccessEvaluationResponse_Context_Reason{
				Properties: map[string]*structpb.Value{cerbosProp("response"): respAsValue},
			},
		},
	}, nil
}
func cerbosProp(s string) string {
	return "cerbos." + s
}
func lookup[T any](m map[string]*T, k string) *T {
	if v, ok := m[cerbosProp(k)]; ok {
		return v
	}

	return nil
}
func lookupOrEmptyString(m map[string]*structpb.Value, k string) string {
	if v := lookup(m, k); v != nil {
		return v.GetStringValue()
	}
	return ""
}
func toCheckResourcesRequest(req *svcv1.AccessEvaluationRequest) (*requestv1.CheckResourcesRequest, error) {
	auxData, err := extractAuxData(req.GetContext())
	if err != nil {
		return nil, err
	}
	return &requestv1.CheckResourcesRequest{
		RequestId:   lookupOrEmptyString(req.GetContext(), "requestId"),
		IncludeMeta: true,
		Principal:   toPrincipal(req.Subject),
		AuxData:     auxData,
		Resources: []*requestv1.CheckResourcesRequest_ResourceEntry{{
			Actions:  []string{req.Action.GetName()},
			Resource: toResource(req.Resource),
		}},
	}, nil
}

func toResource(res *svcv1.AccessEvaluationRequest_Resource) *enginev1.Resource {
	props := res.Properties
	return &enginev1.Resource{
		Kind:          res.Type,
		PolicyVersion: lookupOrEmptyString(props, "policyVersion"),
		Attr:          props,
		Scope:         lookupOrEmptyString(props, "scope"),
		Id:            res.Id,
	}
}

func toPrincipal(subj *svcv1.AccessEvaluationRequest_Subject) *enginev1.Principal {
	props := subj.Properties
	var roles []string
	for _, v := range lookup(props, "roles").GetListValue().GetValues() {
		if r := v.GetStringValue(); r != "" {
			roles = append(roles, r)
		}
	}
	if len(roles) == 0 {
		roles = []string{subj.Type}
	}
	return &enginev1.Principal{
		Id:            subj.Id,
		PolicyVersion: lookupOrEmptyString(props, "policyVersion"),
		Roles:         roles,
		Attr:          props,
		Scope:         lookupOrEmptyString(props, "scope"),
	}
}

// TODO: consider using protobuf reflection instead
// See:
// https://pkg.go.dev/google.golang.org/protobuf@v1.36.10/reflect/protorange
// https://github.com/cerbos/protoc-gen-go-hashpb/blob/main/internal/generator/generator.go
func recode(from, to proto.Message) error {
	data, err := protojson.Marshal(from)
	if err != nil {
		return err
	}
	return protojson.Unmarshal(data, to)
}

func recodeToValue(from proto.Message) (*structpb.Value, error) {
	return messageToValue(from.ProtoReflect())
}

func messageToValue(msg protoreflect.Message) (*structpb.Value, error) {
	fields := make(map[string]*structpb.Value)
	var rangeErr error

	msg.Range(func(fd protoreflect.FieldDescriptor, v protoreflect.Value) bool {
		fieldValue, err := valueToStructValue(fd, v)
		if err != nil {
			rangeErr = err
			return false
		}
		fields[string(fd.Name())] = fieldValue
		return true
	})

	if rangeErr != nil {
		return nil, rangeErr
	}

	return structpb.NewStructValue(&structpb.Struct{
		Fields: fields,
	}), nil
}

func valueToStructValue(fd protoreflect.FieldDescriptor, v protoreflect.Value) (*structpb.Value, error) {
	switch fd.Kind() {
	case protoreflect.BoolKind:
		return structpb.NewBoolValue(v.Bool()), nil
	case protoreflect.Int32Kind, protoreflect.Int64Kind, protoreflect.Sint32Kind, protoreflect.Sint64Kind, protoreflect.Sfixed32Kind, protoreflect.Sfixed64Kind:
		return structpb.NewNumberValue(float64(v.Int())), nil
	case protoreflect.Uint32Kind, protoreflect.Uint64Kind, protoreflect.Fixed32Kind, protoreflect.Fixed64Kind:
		return structpb.NewNumberValue(float64(v.Uint())), nil
	case protoreflect.FloatKind, protoreflect.DoubleKind:
		return structpb.NewNumberValue(v.Float()), nil
	case protoreflect.StringKind:
		return structpb.NewStringValue(v.String()), nil
	case protoreflect.BytesKind:
		return structpb.NewStringValue(base64.StdEncoding.EncodeToString(v.Bytes())), nil
	case protoreflect.MessageKind:
		if fd.IsList() {
			list := v.List()
			values := make([]*structpb.Value, list.Len())
			for i := 0; i < list.Len(); i++ {
				itemValue, err := messageToValue(list.Get(i).Message())
				if err != nil {
					return nil, err
				}
				values[i] = itemValue
			}
			return structpb.NewListValue(&structpb.ListValue{Values: values}), nil
		} else if fd.IsMap() {
			mapValue := v.Map()
			fields := make(map[string]*structpb.Value)
			mapValue.Range(func(mk protoreflect.MapKey, mv protoreflect.Value) bool {
				keyStr := mk.String()
				valueStruct, err := valueToStructValue(fd.MapValue(), mv)
				if err != nil {
					return false
				}
				fields[keyStr] = valueStruct
				return true
			})
			return structpb.NewStructValue(&structpb.Struct{Fields: fields}), nil
		} else {
			return messageToValue(v.Message())
		}
	case protoreflect.EnumKind:
		enumDesc := fd.Enum()
		enumValue := enumDesc.Values().ByNumber(v.Enum())
		return structpb.NewStringValue(string(enumValue.Name())), nil
	default:
		return structpb.NewNullValue(), nil
	}
}
func extractAuxData(m map[string]*structpb.Value) (*requestv1.AuxData, error) {
	var auxData *structpb.Value
	cAuxData := new(requestv1.AuxData)
	var ok bool
	if auxData, ok = m["auxData"]; !ok {
		return nil, nil
	}
	err := recode(auxData, cAuxData)
	if err != nil {
		return nil, fmt.Errorf("can't extract auxData: %w", err)
	}

	return cAuxData, nil
}
