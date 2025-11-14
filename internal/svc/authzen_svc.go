// Copyright 2021-2025 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package svc

import (
	"context"
	"fmt"

	svcv1 "github.com/cerbos/cerbos/api/genpb/authzen/authorization/v1"
	effectv1 "github.com/cerbos/cerbos/api/genpb/cerbos/effect/v1"
	enginev1 "github.com/cerbos/cerbos/api/genpb/cerbos/engine/v1"
	requestv1 "github.com/cerbos/cerbos/api/genpb/cerbos/request/v1"
	responsev1 "github.com/cerbos/cerbos/api/genpb/cerbos/response/v1"
	"github.com/cerbos/cerbos/internal/auxdata"
	"github.com/cerbos/cerbos/internal/engine"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/structpb"
)

var _ svcv1.AuthorizationServiceServer = (*AuthzenAuthorizationService)(nil)

// AuthzenAuthorizationService implements the policy checking service.
type AuthzenAuthorizationService struct {
	svc *CerbosService
	*svcv1.UnimplementedAuthorizationServiceServer
}

func NewAuthzenAuthorizationService(eng *engine.Engine, auxData *auxdata.AuxData, reqLimits RequestLimits) *AuthzenAuthorizationService {
	return &AuthzenAuthorizationService{
		svc:                                     NewCerbosService(eng, auxData, reqLimits),
		UnimplementedAuthorizationServiceServer: &svcv1.UnimplementedAuthorizationServiceServer{}}
}

// AccessEvaluation implements authorizationv1.AuthorizationServiceServer.
func (aas *AuthzenAuthorizationService) AccessEvaluation(ctx context.Context, r *svcv1.AccessEvaluationRequest) (*svcv1.AccessEvaluationResponse, error) {
	// _log := logging.ReqScopeLog(ctx)
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
				Properties: map[string]*structpb.Value{cerbosProp("response"): respAsValue}
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
	return &requestv1.CheckResourcesRequest{
		RequestId:   lookupOrEmptyString(req.GetContext(),"requestId"),
		IncludeMeta: true,
		Principal:   toPrincipal(req.Subject),
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

// TODO: temp
func recode(from, to proto.Message) error {
	data, err := protojson.Marshal(from)
	if err != nil {
		return err
	}
	return protojson.Unmarshal(data, to)
}

func recodeToValue(from proto.Message) (*structpb.Value, error) {
	v := new(structpb.Value)
	if err := recode(from, v); err != nil {
		return nil, err
	}
	return v, nil
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
