// Code generated by protoc-gen-go-hashpb. Do not edit.
// protoc-gen-go-hashpb v0.1.0

package runtimev1

import (
	v1 "github.com/cerbos/cerbos/api/genpb/cerbos/policy/v1"
	v1alpha1 "google.golang.org/genproto/googleapis/api/expr/v1alpha1"
	protowire "google.golang.org/protobuf/encoding/protowire"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	hash "hash"
	math "math"
	sort "sort"
)

func cerbos_policy_v1_Schemas_IgnoreWhen_hashpb_sum(m *v1.Schemas_IgnoreWhen, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.policy.v1.Schemas.IgnoreWhen.actions"]; !ok {
		if len(m.Actions) > 0 {
			for _, v := range m.Actions {
				_, _ = hasher.Write(protowire.AppendString(nil, v))

			}
		}
	}
}

func cerbos_policy_v1_Schemas_Schema_hashpb_sum(m *v1.Schemas_Schema, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.policy.v1.Schemas.Schema.ref"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.Ref))

	}
	if _, ok := ignore["cerbos.policy.v1.Schemas.Schema.ignore_when"]; !ok {
		if m.IgnoreWhen != nil {
			cerbos_policy_v1_Schemas_IgnoreWhen_hashpb_sum(m.IgnoreWhen, hasher, ignore)
		}

	}
}

func cerbos_policy_v1_Schemas_hashpb_sum(m *v1.Schemas, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.policy.v1.Schemas.principal_schema"]; !ok {
		if m.PrincipalSchema != nil {
			cerbos_policy_v1_Schemas_Schema_hashpb_sum(m.PrincipalSchema, hasher, ignore)
		}

	}
	if _, ok := ignore["cerbos.policy.v1.Schemas.resource_schema"]; !ok {
		if m.ResourceSchema != nil {
			cerbos_policy_v1_Schemas_Schema_hashpb_sum(m.ResourceSchema, hasher, ignore)
		}

	}
}

func cerbos_runtime_v1_CompileErrors_Err_hashpb_sum(m *CompileErrors_Err, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.runtime.v1.CompileErrors.Err.file"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.File))

	}
	if _, ok := ignore["cerbos.runtime.v1.CompileErrors.Err.error"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.Error))

	}
	if _, ok := ignore["cerbos.runtime.v1.CompileErrors.Err.description"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.Description))

	}
}

func cerbos_runtime_v1_CompileErrors_hashpb_sum(m *CompileErrors, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.runtime.v1.CompileErrors.errors"]; !ok {
		if len(m.Errors) > 0 {
			for _, v := range m.Errors {
				if v != nil {
					cerbos_runtime_v1_CompileErrors_Err_hashpb_sum(v, hasher, ignore)
				}

			}
		}
	}
}

func cerbos_runtime_v1_Condition_ExprList_hashpb_sum(m *Condition_ExprList, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.runtime.v1.Condition.ExprList.expr"]; !ok {
		if len(m.Expr) > 0 {
			for _, v := range m.Expr {
				if v != nil {
					cerbos_runtime_v1_Condition_hashpb_sum(v, hasher, ignore)
				}

			}
		}
	}
}

func cerbos_runtime_v1_Condition_hashpb_sum(m *Condition, hasher hash.Hash, ignore map[string]struct{}) {
	if m.Op != nil {
		if _, ok := ignore["cerbos.runtime.v1.Condition.op"]; !ok {
			switch t := m.Op.(type) {
			case *Condition_All:
				if t.All != nil {
					cerbos_runtime_v1_Condition_ExprList_hashpb_sum(t.All, hasher, ignore)
				}

			case *Condition_Any:
				if t.Any != nil {
					cerbos_runtime_v1_Condition_ExprList_hashpb_sum(t.Any, hasher, ignore)
				}

			case *Condition_None:
				if t.None != nil {
					cerbos_runtime_v1_Condition_ExprList_hashpb_sum(t.None, hasher, ignore)
				}

			case *Condition_Expr:
				if t.Expr != nil {
					cerbos_runtime_v1_Expr_hashpb_sum(t.Expr, hasher, ignore)
				}

			}
		}
	}
}

func cerbos_runtime_v1_Errors_hashpb_sum(m *Errors, hasher hash.Hash, ignore map[string]struct{}) {
	if m.Kind != nil {
		if _, ok := ignore["cerbos.runtime.v1.Errors.kind"]; !ok {
			switch t := m.Kind.(type) {
			case *Errors_IndexBuildErrors:
				if t.IndexBuildErrors != nil {
					cerbos_runtime_v1_IndexBuildErrors_hashpb_sum(t.IndexBuildErrors, hasher, ignore)
				}

			case *Errors_CompileErrors:
				if t.CompileErrors != nil {
					cerbos_runtime_v1_CompileErrors_hashpb_sum(t.CompileErrors, hasher, ignore)
				}

			}
		}
	}
}

func cerbos_runtime_v1_Expr_hashpb_sum(m *Expr, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.runtime.v1.Expr.original"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.Original))

	}
	if _, ok := ignore["cerbos.runtime.v1.Expr.checked"]; !ok {
		if m.Checked != nil {
			google_api_expr_v1alpha1_CheckedExpr_hashpb_sum(m.Checked, hasher, ignore)
		}

	}
}

func cerbos_runtime_v1_IndexBuildErrors_DuplicateDef_hashpb_sum(m *IndexBuildErrors_DuplicateDef, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.runtime.v1.IndexBuildErrors.DuplicateDef.file"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.File))

	}
	if _, ok := ignore["cerbos.runtime.v1.IndexBuildErrors.DuplicateDef.other_file"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.OtherFile))

	}
}

func cerbos_runtime_v1_IndexBuildErrors_LoadFailure_hashpb_sum(m *IndexBuildErrors_LoadFailure, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.runtime.v1.IndexBuildErrors.LoadFailure.file"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.File))

	}
	if _, ok := ignore["cerbos.runtime.v1.IndexBuildErrors.LoadFailure.error"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.Error))

	}
}

func cerbos_runtime_v1_IndexBuildErrors_MissingImport_hashpb_sum(m *IndexBuildErrors_MissingImport, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.runtime.v1.IndexBuildErrors.MissingImport.importing_file"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.ImportingFile))

	}
	if _, ok := ignore["cerbos.runtime.v1.IndexBuildErrors.MissingImport.desc"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.Desc))

	}
}

func cerbos_runtime_v1_IndexBuildErrors_hashpb_sum(m *IndexBuildErrors, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.runtime.v1.IndexBuildErrors.disabled"]; !ok {
		if len(m.Disabled) > 0 {
			for _, v := range m.Disabled {
				_, _ = hasher.Write(protowire.AppendString(nil, v))

			}
		}
	}
	if _, ok := ignore["cerbos.runtime.v1.IndexBuildErrors.duplicate_defs"]; !ok {
		if len(m.DuplicateDefs) > 0 {
			for _, v := range m.DuplicateDefs {
				if v != nil {
					cerbos_runtime_v1_IndexBuildErrors_DuplicateDef_hashpb_sum(v, hasher, ignore)
				}

			}
		}
	}
	if _, ok := ignore["cerbos.runtime.v1.IndexBuildErrors.load_failures"]; !ok {
		if len(m.LoadFailures) > 0 {
			for _, v := range m.LoadFailures {
				if v != nil {
					cerbos_runtime_v1_IndexBuildErrors_LoadFailure_hashpb_sum(v, hasher, ignore)
				}

			}
		}
	}
	if _, ok := ignore["cerbos.runtime.v1.IndexBuildErrors.missing_imports"]; !ok {
		if len(m.MissingImports) > 0 {
			for _, v := range m.MissingImports {
				if v != nil {
					cerbos_runtime_v1_IndexBuildErrors_MissingImport_hashpb_sum(v, hasher, ignore)
				}

			}
		}
	}
	if _, ok := ignore["cerbos.runtime.v1.IndexBuildErrors.missing_scopes"]; !ok {
		if len(m.MissingScopes) > 0 {
			for _, v := range m.MissingScopes {
				_, _ = hasher.Write(protowire.AppendString(nil, v))

			}
		}
	}
}

func cerbos_runtime_v1_RunnableDerivedRole_hashpb_sum(m *RunnableDerivedRole, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.runtime.v1.RunnableDerivedRole.name"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.Name))

	}
	if _, ok := ignore["cerbos.runtime.v1.RunnableDerivedRole.parent_roles"]; !ok {
		if len(m.ParentRoles) > 0 {
			keys := make([]string, len(m.ParentRoles))
			i := 0
			for k := range m.ParentRoles {
				keys[i] = k
				i++
			}

			sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })

			for _, k := range keys {
				if m.ParentRoles[k] != nil {
					google_protobuf_Empty_hashpb_sum(m.ParentRoles[k], hasher, ignore)
				}

			}
		}
	}
	if _, ok := ignore["cerbos.runtime.v1.RunnableDerivedRole.variables"]; !ok {
		if len(m.Variables) > 0 {
			keys := make([]string, len(m.Variables))
			i := 0
			for k := range m.Variables {
				keys[i] = k
				i++
			}

			sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })

			for _, k := range keys {
				if m.Variables[k] != nil {
					cerbos_runtime_v1_Expr_hashpb_sum(m.Variables[k], hasher, ignore)
				}

			}
		}
	}
	if _, ok := ignore["cerbos.runtime.v1.RunnableDerivedRole.condition"]; !ok {
		if m.Condition != nil {
			cerbos_runtime_v1_Condition_hashpb_sum(m.Condition, hasher, ignore)
		}

	}
}

func cerbos_runtime_v1_RunnableDerivedRolesSet_Metadata_hashpb_sum(m *RunnableDerivedRolesSet_Metadata, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.runtime.v1.RunnableDerivedRolesSet.Metadata.fqn"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.Fqn))

	}
}

func cerbos_runtime_v1_RunnableDerivedRolesSet_hashpb_sum(m *RunnableDerivedRolesSet, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.runtime.v1.RunnableDerivedRolesSet.meta"]; !ok {
		if m.Meta != nil {
			cerbos_runtime_v1_RunnableDerivedRolesSet_Metadata_hashpb_sum(m.Meta, hasher, ignore)
		}

	}
	if _, ok := ignore["cerbos.runtime.v1.RunnableDerivedRolesSet.derived_roles"]; !ok {
		if len(m.DerivedRoles) > 0 {
			keys := make([]string, len(m.DerivedRoles))
			i := 0
			for k := range m.DerivedRoles {
				keys[i] = k
				i++
			}

			sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })

			for _, k := range keys {
				if m.DerivedRoles[k] != nil {
					cerbos_runtime_v1_RunnableDerivedRole_hashpb_sum(m.DerivedRoles[k], hasher, ignore)
				}

			}
		}
	}
}

func cerbos_runtime_v1_RunnablePolicySet_hashpb_sum(m *RunnablePolicySet, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.runtime.v1.RunnablePolicySet.fqn"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.Fqn))

	}
	if m.PolicySet != nil {
		if _, ok := ignore["cerbos.runtime.v1.RunnablePolicySet.policy_set"]; !ok {
			switch t := m.PolicySet.(type) {
			case *RunnablePolicySet_ResourcePolicy:
				if t.ResourcePolicy != nil {
					cerbos_runtime_v1_RunnableResourcePolicySet_hashpb_sum(t.ResourcePolicy, hasher, ignore)
				}

			case *RunnablePolicySet_PrincipalPolicy:
				if t.PrincipalPolicy != nil {
					cerbos_runtime_v1_RunnablePrincipalPolicySet_hashpb_sum(t.PrincipalPolicy, hasher, ignore)
				}

			case *RunnablePolicySet_DerivedRoles:
				if t.DerivedRoles != nil {
					cerbos_runtime_v1_RunnableDerivedRolesSet_hashpb_sum(t.DerivedRoles, hasher, ignore)
				}

			}
		}
	}
}

func cerbos_runtime_v1_RunnablePrincipalPolicySet_Metadata_hashpb_sum(m *RunnablePrincipalPolicySet_Metadata, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.runtime.v1.RunnablePrincipalPolicySet.Metadata.fqn"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.Fqn))

	}
	if _, ok := ignore["cerbos.runtime.v1.RunnablePrincipalPolicySet.Metadata.principal"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.Principal))

	}
	if _, ok := ignore["cerbos.runtime.v1.RunnablePrincipalPolicySet.Metadata.version"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.Version))

	}
}

func cerbos_runtime_v1_RunnablePrincipalPolicySet_Policy_ActionRule_hashpb_sum(m *RunnablePrincipalPolicySet_Policy_ActionRule, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.runtime.v1.RunnablePrincipalPolicySet.Policy.ActionRule.action"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.Action))

	}
	if _, ok := ignore["cerbos.runtime.v1.RunnablePrincipalPolicySet.Policy.ActionRule.name"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.Name))

	}
	if _, ok := ignore["cerbos.runtime.v1.RunnablePrincipalPolicySet.Policy.ActionRule.condition"]; !ok {
		if m.Condition != nil {
			cerbos_runtime_v1_Condition_hashpb_sum(m.Condition, hasher, ignore)
		}

	}
	if _, ok := ignore["cerbos.runtime.v1.RunnablePrincipalPolicySet.Policy.ActionRule.effect"]; !ok {
		_, _ = hasher.Write(protowire.AppendVarint(nil, uint64(m.Effect)))

	}
}

func cerbos_runtime_v1_RunnablePrincipalPolicySet_Policy_ResourceRules_hashpb_sum(m *RunnablePrincipalPolicySet_Policy_ResourceRules, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.runtime.v1.RunnablePrincipalPolicySet.Policy.ResourceRules.action_rules"]; !ok {
		if len(m.ActionRules) > 0 {
			for _, v := range m.ActionRules {
				if v != nil {
					cerbos_runtime_v1_RunnablePrincipalPolicySet_Policy_ActionRule_hashpb_sum(v, hasher, ignore)
				}

			}
		}
	}
}

func cerbos_runtime_v1_RunnablePrincipalPolicySet_Policy_hashpb_sum(m *RunnablePrincipalPolicySet_Policy, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.runtime.v1.RunnablePrincipalPolicySet.Policy.scope"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.Scope))

	}
	if _, ok := ignore["cerbos.runtime.v1.RunnablePrincipalPolicySet.Policy.variables"]; !ok {
		if len(m.Variables) > 0 {
			keys := make([]string, len(m.Variables))
			i := 0
			for k := range m.Variables {
				keys[i] = k
				i++
			}

			sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })

			for _, k := range keys {
				if m.Variables[k] != nil {
					cerbos_runtime_v1_Expr_hashpb_sum(m.Variables[k], hasher, ignore)
				}

			}
		}
	}
	if _, ok := ignore["cerbos.runtime.v1.RunnablePrincipalPolicySet.Policy.resource_rules"]; !ok {
		if len(m.ResourceRules) > 0 {
			keys := make([]string, len(m.ResourceRules))
			i := 0
			for k := range m.ResourceRules {
				keys[i] = k
				i++
			}

			sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })

			for _, k := range keys {
				if m.ResourceRules[k] != nil {
					cerbos_runtime_v1_RunnablePrincipalPolicySet_Policy_ResourceRules_hashpb_sum(m.ResourceRules[k], hasher, ignore)
				}

			}
		}
	}
}

func cerbos_runtime_v1_RunnablePrincipalPolicySet_hashpb_sum(m *RunnablePrincipalPolicySet, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.runtime.v1.RunnablePrincipalPolicySet.meta"]; !ok {
		if m.Meta != nil {
			cerbos_runtime_v1_RunnablePrincipalPolicySet_Metadata_hashpb_sum(m.Meta, hasher, ignore)
		}

	}
	if _, ok := ignore["cerbos.runtime.v1.RunnablePrincipalPolicySet.policies"]; !ok {
		if len(m.Policies) > 0 {
			for _, v := range m.Policies {
				if v != nil {
					cerbos_runtime_v1_RunnablePrincipalPolicySet_Policy_hashpb_sum(v, hasher, ignore)
				}

			}
		}
	}
}

func cerbos_runtime_v1_RunnableResourcePolicySet_Metadata_hashpb_sum(m *RunnableResourcePolicySet_Metadata, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.runtime.v1.RunnableResourcePolicySet.Metadata.fqn"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.Fqn))

	}
	if _, ok := ignore["cerbos.runtime.v1.RunnableResourcePolicySet.Metadata.resource"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.Resource))

	}
	if _, ok := ignore["cerbos.runtime.v1.RunnableResourcePolicySet.Metadata.version"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.Version))

	}
}

func cerbos_runtime_v1_RunnableResourcePolicySet_Policy_Rule_hashpb_sum(m *RunnableResourcePolicySet_Policy_Rule, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.runtime.v1.RunnableResourcePolicySet.Policy.Rule.name"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.Name))

	}
	if _, ok := ignore["cerbos.runtime.v1.RunnableResourcePolicySet.Policy.Rule.actions"]; !ok {
		if len(m.Actions) > 0 {
			keys := make([]string, len(m.Actions))
			i := 0
			for k := range m.Actions {
				keys[i] = k
				i++
			}

			sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })

			for _, k := range keys {
				if m.Actions[k] != nil {
					google_protobuf_Empty_hashpb_sum(m.Actions[k], hasher, ignore)
				}

			}
		}
	}
	if _, ok := ignore["cerbos.runtime.v1.RunnableResourcePolicySet.Policy.Rule.derived_roles"]; !ok {
		if len(m.DerivedRoles) > 0 {
			keys := make([]string, len(m.DerivedRoles))
			i := 0
			for k := range m.DerivedRoles {
				keys[i] = k
				i++
			}

			sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })

			for _, k := range keys {
				if m.DerivedRoles[k] != nil {
					google_protobuf_Empty_hashpb_sum(m.DerivedRoles[k], hasher, ignore)
				}

			}
		}
	}
	if _, ok := ignore["cerbos.runtime.v1.RunnableResourcePolicySet.Policy.Rule.roles"]; !ok {
		if len(m.Roles) > 0 {
			keys := make([]string, len(m.Roles))
			i := 0
			for k := range m.Roles {
				keys[i] = k
				i++
			}

			sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })

			for _, k := range keys {
				if m.Roles[k] != nil {
					google_protobuf_Empty_hashpb_sum(m.Roles[k], hasher, ignore)
				}

			}
		}
	}
	if _, ok := ignore["cerbos.runtime.v1.RunnableResourcePolicySet.Policy.Rule.condition"]; !ok {
		if m.Condition != nil {
			cerbos_runtime_v1_Condition_hashpb_sum(m.Condition, hasher, ignore)
		}

	}
	if _, ok := ignore["cerbos.runtime.v1.RunnableResourcePolicySet.Policy.Rule.effect"]; !ok {
		_, _ = hasher.Write(protowire.AppendVarint(nil, uint64(m.Effect)))

	}
}

func cerbos_runtime_v1_RunnableResourcePolicySet_Policy_hashpb_sum(m *RunnableResourcePolicySet_Policy, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.runtime.v1.RunnableResourcePolicySet.Policy.scope"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.Scope))

	}
	if _, ok := ignore["cerbos.runtime.v1.RunnableResourcePolicySet.Policy.derived_roles"]; !ok {
		if len(m.DerivedRoles) > 0 {
			keys := make([]string, len(m.DerivedRoles))
			i := 0
			for k := range m.DerivedRoles {
				keys[i] = k
				i++
			}

			sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })

			for _, k := range keys {
				if m.DerivedRoles[k] != nil {
					cerbos_runtime_v1_RunnableDerivedRole_hashpb_sum(m.DerivedRoles[k], hasher, ignore)
				}

			}
		}
	}
	if _, ok := ignore["cerbos.runtime.v1.RunnableResourcePolicySet.Policy.variables"]; !ok {
		if len(m.Variables) > 0 {
			keys := make([]string, len(m.Variables))
			i := 0
			for k := range m.Variables {
				keys[i] = k
				i++
			}

			sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })

			for _, k := range keys {
				if m.Variables[k] != nil {
					cerbos_runtime_v1_Expr_hashpb_sum(m.Variables[k], hasher, ignore)
				}

			}
		}
	}
	if _, ok := ignore["cerbos.runtime.v1.RunnableResourcePolicySet.Policy.rules"]; !ok {
		if len(m.Rules) > 0 {
			for _, v := range m.Rules {
				if v != nil {
					cerbos_runtime_v1_RunnableResourcePolicySet_Policy_Rule_hashpb_sum(v, hasher, ignore)
				}

			}
		}
	}
	if _, ok := ignore["cerbos.runtime.v1.RunnableResourcePolicySet.Policy.schemas"]; !ok {
		if m.Schemas != nil {
			cerbos_policy_v1_Schemas_hashpb_sum(m.Schemas, hasher, ignore)
		}

	}
}

func cerbos_runtime_v1_RunnableResourcePolicySet_hashpb_sum(m *RunnableResourcePolicySet, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.runtime.v1.RunnableResourcePolicySet.meta"]; !ok {
		if m.Meta != nil {
			cerbos_runtime_v1_RunnableResourcePolicySet_Metadata_hashpb_sum(m.Meta, hasher, ignore)
		}

	}
	if _, ok := ignore["cerbos.runtime.v1.RunnableResourcePolicySet.policies"]; !ok {
		if len(m.Policies) > 0 {
			for _, v := range m.Policies {
				if v != nil {
					cerbos_runtime_v1_RunnableResourcePolicySet_Policy_hashpb_sum(v, hasher, ignore)
				}

			}
		}
	}
	if _, ok := ignore["cerbos.runtime.v1.RunnableResourcePolicySet.schemas"]; !ok {
		if m.Schemas != nil {
			cerbos_policy_v1_Schemas_hashpb_sum(m.Schemas, hasher, ignore)
		}

	}
}

func google_api_expr_v1alpha1_CheckedExpr_hashpb_sum(m *v1alpha1.CheckedExpr, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["google.api.expr.v1alpha1.CheckedExpr.reference_map"]; !ok {
		if len(m.ReferenceMap) > 0 {
			keys := make([]int64, len(m.ReferenceMap))
			i := 0
			for k := range m.ReferenceMap {
				keys[i] = k
				i++
			}

			sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })

			for _, k := range keys {
				if m.ReferenceMap[k] != nil {
					google_api_expr_v1alpha1_Reference_hashpb_sum(m.ReferenceMap[k], hasher, ignore)
				}

			}
		}
	}
	if _, ok := ignore["google.api.expr.v1alpha1.CheckedExpr.type_map"]; !ok {
		if len(m.TypeMap) > 0 {
			keys := make([]int64, len(m.TypeMap))
			i := 0
			for k := range m.TypeMap {
				keys[i] = k
				i++
			}

			sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })

			for _, k := range keys {
				if m.TypeMap[k] != nil {
					google_api_expr_v1alpha1_Type_hashpb_sum(m.TypeMap[k], hasher, ignore)
				}

			}
		}
	}
	if _, ok := ignore["google.api.expr.v1alpha1.CheckedExpr.expr"]; !ok {
		if m.Expr != nil {
			google_api_expr_v1alpha1_Expr_hashpb_sum(m.Expr, hasher, ignore)
		}

	}
	if _, ok := ignore["google.api.expr.v1alpha1.CheckedExpr.source_info"]; !ok {
		if m.SourceInfo != nil {
			google_api_expr_v1alpha1_SourceInfo_hashpb_sum(m.SourceInfo, hasher, ignore)
		}

	}
}

func google_api_expr_v1alpha1_Constant_hashpb_sum(m *v1alpha1.Constant, hasher hash.Hash, ignore map[string]struct{}) {
	if m.ConstantKind != nil {
		if _, ok := ignore["google.api.expr.v1alpha1.Constant.constant_kind"]; !ok {
			switch t := m.ConstantKind.(type) {
			case *v1alpha1.Constant_NullValue:
				_, _ = hasher.Write(protowire.AppendVarint(nil, uint64(t.NullValue)))

			case *v1alpha1.Constant_BoolValue:
				_, _ = hasher.Write(protowire.AppendVarint(nil, protowire.EncodeBool(t.BoolValue)))

			case *v1alpha1.Constant_Int64Value:
				_, _ = hasher.Write(protowire.AppendVarint(nil, uint64(t.Int64Value)))

			case *v1alpha1.Constant_Uint64Value:
				_, _ = hasher.Write(protowire.AppendVarint(nil, t.Uint64Value))

			case *v1alpha1.Constant_DoubleValue:
				_, _ = hasher.Write(protowire.AppendFixed64(nil, math.Float64bits(t.DoubleValue)))

			case *v1alpha1.Constant_StringValue:
				_, _ = hasher.Write(protowire.AppendString(nil, t.StringValue))

			case *v1alpha1.Constant_BytesValue:
				_, _ = hasher.Write(protowire.AppendBytes(nil, t.BytesValue))

			case *v1alpha1.Constant_DurationValue:
				if t.DurationValue != nil {
					google_protobuf_Duration_hashpb_sum(t.DurationValue, hasher, ignore)
				}

			case *v1alpha1.Constant_TimestampValue:
				if t.TimestampValue != nil {
					google_protobuf_Timestamp_hashpb_sum(t.TimestampValue, hasher, ignore)
				}

			}
		}
	}
}

func google_api_expr_v1alpha1_Expr_Call_hashpb_sum(m *v1alpha1.Expr_Call, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["google.api.expr.v1alpha1.Expr.Call.target"]; !ok {
		if m.Target != nil {
			google_api_expr_v1alpha1_Expr_hashpb_sum(m.Target, hasher, ignore)
		}

	}
	if _, ok := ignore["google.api.expr.v1alpha1.Expr.Call.function"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.Function))

	}
	if _, ok := ignore["google.api.expr.v1alpha1.Expr.Call.args"]; !ok {
		if len(m.Args) > 0 {
			for _, v := range m.Args {
				if v != nil {
					google_api_expr_v1alpha1_Expr_hashpb_sum(v, hasher, ignore)
				}

			}
		}
	}
}

func google_api_expr_v1alpha1_Expr_Comprehension_hashpb_sum(m *v1alpha1.Expr_Comprehension, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["google.api.expr.v1alpha1.Expr.Comprehension.iter_var"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.IterVar))

	}
	if _, ok := ignore["google.api.expr.v1alpha1.Expr.Comprehension.iter_range"]; !ok {
		if m.IterRange != nil {
			google_api_expr_v1alpha1_Expr_hashpb_sum(m.IterRange, hasher, ignore)
		}

	}
	if _, ok := ignore["google.api.expr.v1alpha1.Expr.Comprehension.accu_var"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.AccuVar))

	}
	if _, ok := ignore["google.api.expr.v1alpha1.Expr.Comprehension.accu_init"]; !ok {
		if m.AccuInit != nil {
			google_api_expr_v1alpha1_Expr_hashpb_sum(m.AccuInit, hasher, ignore)
		}

	}
	if _, ok := ignore["google.api.expr.v1alpha1.Expr.Comprehension.loop_condition"]; !ok {
		if m.LoopCondition != nil {
			google_api_expr_v1alpha1_Expr_hashpb_sum(m.LoopCondition, hasher, ignore)
		}

	}
	if _, ok := ignore["google.api.expr.v1alpha1.Expr.Comprehension.loop_step"]; !ok {
		if m.LoopStep != nil {
			google_api_expr_v1alpha1_Expr_hashpb_sum(m.LoopStep, hasher, ignore)
		}

	}
	if _, ok := ignore["google.api.expr.v1alpha1.Expr.Comprehension.result"]; !ok {
		if m.Result != nil {
			google_api_expr_v1alpha1_Expr_hashpb_sum(m.Result, hasher, ignore)
		}

	}
}

func google_api_expr_v1alpha1_Expr_CreateList_hashpb_sum(m *v1alpha1.Expr_CreateList, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["google.api.expr.v1alpha1.Expr.CreateList.elements"]; !ok {
		if len(m.Elements) > 0 {
			for _, v := range m.Elements {
				if v != nil {
					google_api_expr_v1alpha1_Expr_hashpb_sum(v, hasher, ignore)
				}

			}
		}
	}
}

func google_api_expr_v1alpha1_Expr_CreateStruct_Entry_hashpb_sum(m *v1alpha1.Expr_CreateStruct_Entry, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["google.api.expr.v1alpha1.Expr.CreateStruct.Entry.id"]; !ok {
		_, _ = hasher.Write(protowire.AppendVarint(nil, uint64(m.Id)))

	}
	if m.KeyKind != nil {
		if _, ok := ignore["google.api.expr.v1alpha1.Expr.CreateStruct.Entry.key_kind"]; !ok {
			switch t := m.KeyKind.(type) {
			case *v1alpha1.Expr_CreateStruct_Entry_FieldKey:
				_, _ = hasher.Write(protowire.AppendString(nil, t.FieldKey))

			case *v1alpha1.Expr_CreateStruct_Entry_MapKey:
				if t.MapKey != nil {
					google_api_expr_v1alpha1_Expr_hashpb_sum(t.MapKey, hasher, ignore)
				}

			}
		}
	}
	if _, ok := ignore["google.api.expr.v1alpha1.Expr.CreateStruct.Entry.value"]; !ok {
		if m.Value != nil {
			google_api_expr_v1alpha1_Expr_hashpb_sum(m.Value, hasher, ignore)
		}

	}
}

func google_api_expr_v1alpha1_Expr_CreateStruct_hashpb_sum(m *v1alpha1.Expr_CreateStruct, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["google.api.expr.v1alpha1.Expr.CreateStruct.message_name"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.MessageName))

	}
	if _, ok := ignore["google.api.expr.v1alpha1.Expr.CreateStruct.entries"]; !ok {
		if len(m.Entries) > 0 {
			for _, v := range m.Entries {
				if v != nil {
					google_api_expr_v1alpha1_Expr_CreateStruct_Entry_hashpb_sum(v, hasher, ignore)
				}

			}
		}
	}
}

func google_api_expr_v1alpha1_Expr_Ident_hashpb_sum(m *v1alpha1.Expr_Ident, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["google.api.expr.v1alpha1.Expr.Ident.name"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.Name))

	}
}

func google_api_expr_v1alpha1_Expr_Select_hashpb_sum(m *v1alpha1.Expr_Select, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["google.api.expr.v1alpha1.Expr.Select.operand"]; !ok {
		if m.Operand != nil {
			google_api_expr_v1alpha1_Expr_hashpb_sum(m.Operand, hasher, ignore)
		}

	}
	if _, ok := ignore["google.api.expr.v1alpha1.Expr.Select.field"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.Field))

	}
	if _, ok := ignore["google.api.expr.v1alpha1.Expr.Select.test_only"]; !ok {
		_, _ = hasher.Write(protowire.AppendVarint(nil, protowire.EncodeBool(m.TestOnly)))

	}
}

func google_api_expr_v1alpha1_Expr_hashpb_sum(m *v1alpha1.Expr, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["google.api.expr.v1alpha1.Expr.id"]; !ok {
		_, _ = hasher.Write(protowire.AppendVarint(nil, uint64(m.Id)))

	}
	if m.ExprKind != nil {
		if _, ok := ignore["google.api.expr.v1alpha1.Expr.expr_kind"]; !ok {
			switch t := m.ExprKind.(type) {
			case *v1alpha1.Expr_ConstExpr:
				if t.ConstExpr != nil {
					google_api_expr_v1alpha1_Constant_hashpb_sum(t.ConstExpr, hasher, ignore)
				}

			case *v1alpha1.Expr_IdentExpr:
				if t.IdentExpr != nil {
					google_api_expr_v1alpha1_Expr_Ident_hashpb_sum(t.IdentExpr, hasher, ignore)
				}

			case *v1alpha1.Expr_SelectExpr:
				if t.SelectExpr != nil {
					google_api_expr_v1alpha1_Expr_Select_hashpb_sum(t.SelectExpr, hasher, ignore)
				}

			case *v1alpha1.Expr_CallExpr:
				if t.CallExpr != nil {
					google_api_expr_v1alpha1_Expr_Call_hashpb_sum(t.CallExpr, hasher, ignore)
				}

			case *v1alpha1.Expr_ListExpr:
				if t.ListExpr != nil {
					google_api_expr_v1alpha1_Expr_CreateList_hashpb_sum(t.ListExpr, hasher, ignore)
				}

			case *v1alpha1.Expr_StructExpr:
				if t.StructExpr != nil {
					google_api_expr_v1alpha1_Expr_CreateStruct_hashpb_sum(t.StructExpr, hasher, ignore)
				}

			case *v1alpha1.Expr_ComprehensionExpr:
				if t.ComprehensionExpr != nil {
					google_api_expr_v1alpha1_Expr_Comprehension_hashpb_sum(t.ComprehensionExpr, hasher, ignore)
				}

			}
		}
	}
}

func google_api_expr_v1alpha1_Reference_hashpb_sum(m *v1alpha1.Reference, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["google.api.expr.v1alpha1.Reference.name"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.Name))

	}
	if _, ok := ignore["google.api.expr.v1alpha1.Reference.overload_id"]; !ok {
		if len(m.OverloadId) > 0 {
			for _, v := range m.OverloadId {
				_, _ = hasher.Write(protowire.AppendString(nil, v))

			}
		}
	}
	if _, ok := ignore["google.api.expr.v1alpha1.Reference.value"]; !ok {
		if m.Value != nil {
			google_api_expr_v1alpha1_Constant_hashpb_sum(m.Value, hasher, ignore)
		}

	}
}

func google_api_expr_v1alpha1_SourceInfo_hashpb_sum(m *v1alpha1.SourceInfo, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["google.api.expr.v1alpha1.SourceInfo.syntax_version"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.SyntaxVersion))

	}
	if _, ok := ignore["google.api.expr.v1alpha1.SourceInfo.location"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.Location))

	}
	if _, ok := ignore["google.api.expr.v1alpha1.SourceInfo.line_offsets"]; !ok {
		if len(m.LineOffsets) > 0 {
			for _, v := range m.LineOffsets {
				_, _ = hasher.Write(protowire.AppendVarint(nil, uint64(v)))

			}
		}
	}
	if _, ok := ignore["google.api.expr.v1alpha1.SourceInfo.positions"]; !ok {
		if len(m.Positions) > 0 {
			keys := make([]int64, len(m.Positions))
			i := 0
			for k := range m.Positions {
				keys[i] = k
				i++
			}

			sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })

			for _, k := range keys {
				_, _ = hasher.Write(protowire.AppendVarint(nil, uint64(m.Positions[k])))

			}
		}
	}
	if _, ok := ignore["google.api.expr.v1alpha1.SourceInfo.macro_calls"]; !ok {
		if len(m.MacroCalls) > 0 {
			keys := make([]int64, len(m.MacroCalls))
			i := 0
			for k := range m.MacroCalls {
				keys[i] = k
				i++
			}

			sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })

			for _, k := range keys {
				if m.MacroCalls[k] != nil {
					google_api_expr_v1alpha1_Expr_hashpb_sum(m.MacroCalls[k], hasher, ignore)
				}

			}
		}
	}
}

func google_api_expr_v1alpha1_Type_AbstractType_hashpb_sum(m *v1alpha1.Type_AbstractType, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["google.api.expr.v1alpha1.Type.AbstractType.name"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.Name))

	}
	if _, ok := ignore["google.api.expr.v1alpha1.Type.AbstractType.parameter_types"]; !ok {
		if len(m.ParameterTypes) > 0 {
			for _, v := range m.ParameterTypes {
				if v != nil {
					google_api_expr_v1alpha1_Type_hashpb_sum(v, hasher, ignore)
				}

			}
		}
	}
}

func google_api_expr_v1alpha1_Type_FunctionType_hashpb_sum(m *v1alpha1.Type_FunctionType, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["google.api.expr.v1alpha1.Type.FunctionType.result_type"]; !ok {
		if m.ResultType != nil {
			google_api_expr_v1alpha1_Type_hashpb_sum(m.ResultType, hasher, ignore)
		}

	}
	if _, ok := ignore["google.api.expr.v1alpha1.Type.FunctionType.arg_types"]; !ok {
		if len(m.ArgTypes) > 0 {
			for _, v := range m.ArgTypes {
				if v != nil {
					google_api_expr_v1alpha1_Type_hashpb_sum(v, hasher, ignore)
				}

			}
		}
	}
}

func google_api_expr_v1alpha1_Type_ListType_hashpb_sum(m *v1alpha1.Type_ListType, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["google.api.expr.v1alpha1.Type.ListType.elem_type"]; !ok {
		if m.ElemType != nil {
			google_api_expr_v1alpha1_Type_hashpb_sum(m.ElemType, hasher, ignore)
		}

	}
}

func google_api_expr_v1alpha1_Type_MapType_hashpb_sum(m *v1alpha1.Type_MapType, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["google.api.expr.v1alpha1.Type.MapType.key_type"]; !ok {
		if m.KeyType != nil {
			google_api_expr_v1alpha1_Type_hashpb_sum(m.KeyType, hasher, ignore)
		}

	}
	if _, ok := ignore["google.api.expr.v1alpha1.Type.MapType.value_type"]; !ok {
		if m.ValueType != nil {
			google_api_expr_v1alpha1_Type_hashpb_sum(m.ValueType, hasher, ignore)
		}

	}
}

func google_api_expr_v1alpha1_Type_hashpb_sum(m *v1alpha1.Type, hasher hash.Hash, ignore map[string]struct{}) {
	if m.TypeKind != nil {
		if _, ok := ignore["google.api.expr.v1alpha1.Type.type_kind"]; !ok {
			switch t := m.TypeKind.(type) {
			case *v1alpha1.Type_Dyn:
				if t.Dyn != nil {
					google_protobuf_Empty_hashpb_sum(t.Dyn, hasher, ignore)
				}

			case *v1alpha1.Type_Null:
				_, _ = hasher.Write(protowire.AppendVarint(nil, uint64(t.Null)))

			case *v1alpha1.Type_Primitive:
				_, _ = hasher.Write(protowire.AppendVarint(nil, uint64(t.Primitive)))

			case *v1alpha1.Type_Wrapper:
				_, _ = hasher.Write(protowire.AppendVarint(nil, uint64(t.Wrapper)))

			case *v1alpha1.Type_WellKnown:
				_, _ = hasher.Write(protowire.AppendVarint(nil, uint64(t.WellKnown)))

			case *v1alpha1.Type_ListType_:
				if t.ListType != nil {
					google_api_expr_v1alpha1_Type_ListType_hashpb_sum(t.ListType, hasher, ignore)
				}

			case *v1alpha1.Type_MapType_:
				if t.MapType != nil {
					google_api_expr_v1alpha1_Type_MapType_hashpb_sum(t.MapType, hasher, ignore)
				}

			case *v1alpha1.Type_Function:
				if t.Function != nil {
					google_api_expr_v1alpha1_Type_FunctionType_hashpb_sum(t.Function, hasher, ignore)
				}

			case *v1alpha1.Type_MessageType:
				_, _ = hasher.Write(protowire.AppendString(nil, t.MessageType))

			case *v1alpha1.Type_TypeParam:
				_, _ = hasher.Write(protowire.AppendString(nil, t.TypeParam))

			case *v1alpha1.Type_Type:
				if t.Type != nil {
					google_api_expr_v1alpha1_Type_hashpb_sum(t.Type, hasher, ignore)
				}

			case *v1alpha1.Type_Error:
				if t.Error != nil {
					google_protobuf_Empty_hashpb_sum(t.Error, hasher, ignore)
				}

			case *v1alpha1.Type_AbstractType_:
				if t.AbstractType != nil {
					google_api_expr_v1alpha1_Type_AbstractType_hashpb_sum(t.AbstractType, hasher, ignore)
				}

			}
		}
	}
}

func google_protobuf_Duration_hashpb_sum(m *durationpb.Duration, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["google.protobuf.Duration.seconds"]; !ok {
		_, _ = hasher.Write(protowire.AppendVarint(nil, uint64(m.Seconds)))

	}
	if _, ok := ignore["google.protobuf.Duration.nanos"]; !ok {
		_, _ = hasher.Write(protowire.AppendVarint(nil, uint64(m.Nanos)))

	}
}

func google_protobuf_Empty_hashpb_sum(m *emptypb.Empty, hasher hash.Hash, ignore map[string]struct{}) {
}

func google_protobuf_Timestamp_hashpb_sum(m *timestamppb.Timestamp, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["google.protobuf.Timestamp.seconds"]; !ok {
		_, _ = hasher.Write(protowire.AppendVarint(nil, uint64(m.Seconds)))

	}
	if _, ok := ignore["google.protobuf.Timestamp.nanos"]; !ok {
		_, _ = hasher.Write(protowire.AppendVarint(nil, uint64(m.Nanos)))

	}
}
