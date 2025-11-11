// Copyright 2021-2025 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package svc

import (
	"context"
	svcv1 "github.com/cerbos/cerbos/api/genpb/authzen/authorization/v1"
	"github.com/cerbos/cerbos/internal/auxdata"
	"github.com/cerbos/cerbos/internal/engine"
)

var _ svcv1.AuthorizationServiceServer = (*AuthzenAuthorizationService)(nil)

// AuthzenAuthorizationService implements the policy checking service.
type AuthzenAuthorizationService struct {
	eng     *engine.Engine
	auxData *auxdata.AuxData
	svcv1.UnimplementedAuthorizationServiceServer
	reqLimits RequestLimits
}

func NewAuthzenAuthorizationService(eng *engine.Engine, auxData *auxdata.AuxData, reqLimits RequestLimits) *CerbosService {
	return &CerbosService{
		eng:                              eng,
		auxData:                          auxData,
		reqLimits:                        reqLimits,
		UnimplementedCerbosServiceServer: &svcv1.UnimplementedAuthorizationServiceServer
	}
}

// AccessEvaluation implements authorizationv1.AuthorizationServiceServer.
func (a *AuthzenAuthorizationService) AccessEvaluation(context.Context, *svcv1.AccessEvaluationRequest) (*svcv1.AccessEvaluationResponse, error) {
	panic("unimplemented")
}
