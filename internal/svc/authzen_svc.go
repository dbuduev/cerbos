// Copyright 2021-2025 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package svc

import (
	"context"
	"fmt"

	svcv1 "github.com/cerbos/cerbos/api/genpb/authzen/authorization/v1"
	enginev1 "github.com/cerbos/cerbos/api/genpb/cerbos/engine/v1"
	requestv1 "github.com/cerbos/cerbos/api/genpb/cerbos/request/v1"
	responsev1 "github.com/cerbos/cerbos/api/genpb/cerbos/response/v1"
	"github.com/cerbos/cerbos/internal/auxdata"
	"github.com/cerbos/cerbos/internal/engine"
	"google.golang.org/protobuf/encoding/protojson"
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
func (aas *AuthzenAuthorizationService) AccessEvaluation(ctx context.Context, req *svcv1.AccessEvaluationRequest) (*svcv1.AccessEvaluationResponse, error) {
	// _log := logging.ReqScopeLog(ctx)
	cReq, err := toCheckResourcesRequest(req)
	if err != nil {
		return nil, err
	}
	cResp, err := aas.svc.CheckResources(ctx, cReq)
	if err != nil {
		return nil, err
	}
	resp, err := toAccessEvaluationResponse(ctx, cResp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func toAccessEvaluationResponse(ctx context.Context, cResp *responsev1.CheckResourcesResponse) (*svcv1.AccessEvaluationResponse, error) {
	panic("unimplemented")
}
func lookup[T any](m map[string]*T, k string) *T {
	if v, ok := m[k]; ok {
		return v
	}

	return nil
}
func lookupOrDefault[T any](m map[string]T, k string, d T) T {
	if v, ok := m[k]; ok {
		return v
	}

	return d
}
func lookupOrEmptyString(m map[string]*structpb.Value, k string) string {
	if v := lookup(m, k); v != nil {
		return v.GetStringValue()
	}
	return ""
}
func toCheckResourcesRequest(req *svcv1.AccessEvaluationRequest) (*requestv1.CheckResourcesRequest, error) {
	c := req.GetContext()
	return &requestv1.CheckResourcesRequest{
		RequestId:   lookupOrEmptyString(c, "requestId"),
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

func extractAuxData(m map[string]*structpb.Value) (*requestv1.AuxData, error) {
	var data []byte
	var auxData *structpb.Value
	var ok bool
	if auxData, ok = m["auxData"]; !ok {
		return nil, nil
	}
	data, err := protojson.Marshal(auxData)
	if err != nil {
		return nil, fmt.Errorf("can't marshal context auxData: %w", err)
	}

	cAuxData := new(requestv1.AuxData)
	if err = protojson.Unmarshal(data, cAuxData); err != nil {
		return nil, fmt.Errorf("can't unmarshal AuxData: %w", err)
	}

	return cAuxData, nil
}
