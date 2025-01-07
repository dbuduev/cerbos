// Code generated by protoc-gen-go-hashpb. Do not edit.
// protoc-gen-go-hashpb v0.3.4

package sourcev1

import (
	v1 "github.com/cerbos/cerbos/api/genpb/cerbos/policy/v1"
	protowire "google.golang.org/protobuf/encoding/protowire"
	structpb "google.golang.org/protobuf/types/known/structpb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	hash "hash"
	math "math"
	sort "sort"
)

func cerbos_policy_v1_Condition_hashpb_sum(m *v1.Condition, hasher hash.Hash, ignore map[string]struct{}) {
	if m.Condition != nil {
		if _, ok := ignore["cerbos.policy.v1.Condition.condition"]; !ok {
			switch t := m.Condition.(type) {
			case *v1.Condition_Match:
				if t.Match != nil {
					cerbos_policy_v1_Match_hashpb_sum(t.Match, hasher, ignore)
				}

			case *v1.Condition_Script:
				_, _ = hasher.Write(protowire.AppendString(nil, t.Script))

			}
		}
	}
}

func cerbos_policy_v1_Constants_hashpb_sum(m *v1.Constants, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.policy.v1.Constants.import"]; !ok {
		if len(m.Import) > 0 {
			for _, v := range m.Import {
				_, _ = hasher.Write(protowire.AppendString(nil, v))

			}
		}
	}
	if _, ok := ignore["cerbos.policy.v1.Constants.local"]; !ok {
		if len(m.Local) > 0 {
			keys := make([]string, len(m.Local))
			i := 0
			for k := range m.Local {
				keys[i] = k
				i++
			}

			sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })

			for _, k := range keys {
				if m.Local[k] != nil {
					google_protobuf_Value_hashpb_sum(m.Local[k], hasher, ignore)
				}

			}
		}
	}
}

func cerbos_policy_v1_DerivedRoles_hashpb_sum(m *v1.DerivedRoles, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.policy.v1.DerivedRoles.name"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.GetName()))

	}
	if _, ok := ignore["cerbos.policy.v1.DerivedRoles.definitions"]; !ok {
		if len(m.Definitions) > 0 {
			for _, v := range m.Definitions {
				if v != nil {
					cerbos_policy_v1_RoleDef_hashpb_sum(v, hasher, ignore)
				}

			}
		}
	}
	if _, ok := ignore["cerbos.policy.v1.DerivedRoles.variables"]; !ok {
		if m.GetVariables() != nil {
			cerbos_policy_v1_Variables_hashpb_sum(m.GetVariables(), hasher, ignore)
		}

	}
	if _, ok := ignore["cerbos.policy.v1.DerivedRoles.constants"]; !ok {
		if m.GetConstants() != nil {
			cerbos_policy_v1_Constants_hashpb_sum(m.GetConstants(), hasher, ignore)
		}

	}
}

func cerbos_policy_v1_ExportConstants_hashpb_sum(m *v1.ExportConstants, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.policy.v1.ExportConstants.name"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.GetName()))

	}
	if _, ok := ignore["cerbos.policy.v1.ExportConstants.definitions"]; !ok {
		if len(m.Definitions) > 0 {
			keys := make([]string, len(m.Definitions))
			i := 0
			for k := range m.Definitions {
				keys[i] = k
				i++
			}

			sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })

			for _, k := range keys {
				if m.Definitions[k] != nil {
					google_protobuf_Value_hashpb_sum(m.Definitions[k], hasher, ignore)
				}

			}
		}
	}
}

func cerbos_policy_v1_ExportVariables_hashpb_sum(m *v1.ExportVariables, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.policy.v1.ExportVariables.name"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.GetName()))

	}
	if _, ok := ignore["cerbos.policy.v1.ExportVariables.definitions"]; !ok {
		if len(m.Definitions) > 0 {
			keys := make([]string, len(m.Definitions))
			i := 0
			for k := range m.Definitions {
				keys[i] = k
				i++
			}

			sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })

			for _, k := range keys {
				_, _ = hasher.Write(protowire.AppendString(nil, m.Definitions[k]))

			}
		}
	}
}

func cerbos_policy_v1_Match_ExprList_hashpb_sum(m *v1.Match_ExprList, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.policy.v1.Match.ExprList.of"]; !ok {
		if len(m.Of) > 0 {
			for _, v := range m.Of {
				if v != nil {
					cerbos_policy_v1_Match_hashpb_sum(v, hasher, ignore)
				}

			}
		}
	}
}

func cerbos_policy_v1_Match_hashpb_sum(m *v1.Match, hasher hash.Hash, ignore map[string]struct{}) {
	if m.Op != nil {
		if _, ok := ignore["cerbos.policy.v1.Match.op"]; !ok {
			switch t := m.Op.(type) {
			case *v1.Match_All:
				if t.All != nil {
					cerbos_policy_v1_Match_ExprList_hashpb_sum(t.All, hasher, ignore)
				}

			case *v1.Match_Any:
				if t.Any != nil {
					cerbos_policy_v1_Match_ExprList_hashpb_sum(t.Any, hasher, ignore)
				}

			case *v1.Match_None:
				if t.None != nil {
					cerbos_policy_v1_Match_ExprList_hashpb_sum(t.None, hasher, ignore)
				}

			case *v1.Match_Expr:
				_, _ = hasher.Write(protowire.AppendString(nil, t.Expr))

			}
		}
	}
}

func cerbos_policy_v1_Metadata_hashpb_sum(m *v1.Metadata, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.policy.v1.Metadata.source_file"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.GetSourceFile()))

	}
	if _, ok := ignore["cerbos.policy.v1.Metadata.annotations"]; !ok {
		if len(m.Annotations) > 0 {
			keys := make([]string, len(m.Annotations))
			i := 0
			for k := range m.Annotations {
				keys[i] = k
				i++
			}

			sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })

			for _, k := range keys {
				_, _ = hasher.Write(protowire.AppendString(nil, m.Annotations[k]))

			}
		}
	}
	if _, ok := ignore["cerbos.policy.v1.Metadata.hash"]; !ok {
		if m.GetHash() != nil {
			google_protobuf_UInt64Value_hashpb_sum(m.GetHash(), hasher, ignore)
		}

	}
	if _, ok := ignore["cerbos.policy.v1.Metadata.store_identifer"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.GetStoreIdentifer()))

	}
	if _, ok := ignore["cerbos.policy.v1.Metadata.store_identifier"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.GetStoreIdentifier()))

	}
	if _, ok := ignore["cerbos.policy.v1.Metadata.source_attributes"]; !ok {
		if m.GetSourceAttributes() != nil {
			cerbos_policy_v1_SourceAttributes_hashpb_sum(m.GetSourceAttributes(), hasher, ignore)
		}

	}
}

func cerbos_policy_v1_Output_When_hashpb_sum(m *v1.Output_When, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.policy.v1.Output.When.rule_activated"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.GetRuleActivated()))

	}
	if _, ok := ignore["cerbos.policy.v1.Output.When.condition_not_met"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.GetConditionNotMet()))

	}
}

func cerbos_policy_v1_Output_hashpb_sum(m *v1.Output, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.policy.v1.Output.expr"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.GetExpr()))

	}
	if _, ok := ignore["cerbos.policy.v1.Output.when"]; !ok {
		if m.GetWhen() != nil {
			cerbos_policy_v1_Output_When_hashpb_sum(m.GetWhen(), hasher, ignore)
		}

	}
}

func cerbos_policy_v1_Policy_hashpb_sum(m *v1.Policy, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.policy.v1.Policy.api_version"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.GetApiVersion()))

	}
	if _, ok := ignore["cerbos.policy.v1.Policy.disabled"]; !ok {
		_, _ = hasher.Write(protowire.AppendVarint(nil, protowire.EncodeBool(m.GetDisabled())))

	}
	if _, ok := ignore["cerbos.policy.v1.Policy.description"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.GetDescription()))

	}
	if _, ok := ignore["cerbos.policy.v1.Policy.metadata"]; !ok {
		if m.GetMetadata() != nil {
			cerbos_policy_v1_Metadata_hashpb_sum(m.GetMetadata(), hasher, ignore)
		}

	}
	if m.PolicyType != nil {
		if _, ok := ignore["cerbos.policy.v1.Policy.policy_type"]; !ok {
			switch t := m.PolicyType.(type) {
			case *v1.Policy_ResourcePolicy:
				if t.ResourcePolicy != nil {
					cerbos_policy_v1_ResourcePolicy_hashpb_sum(t.ResourcePolicy, hasher, ignore)
				}

			case *v1.Policy_PrincipalPolicy:
				if t.PrincipalPolicy != nil {
					cerbos_policy_v1_PrincipalPolicy_hashpb_sum(t.PrincipalPolicy, hasher, ignore)
				}

			case *v1.Policy_DerivedRoles:
				if t.DerivedRoles != nil {
					cerbos_policy_v1_DerivedRoles_hashpb_sum(t.DerivedRoles, hasher, ignore)
				}

			case *v1.Policy_ExportVariables:
				if t.ExportVariables != nil {
					cerbos_policy_v1_ExportVariables_hashpb_sum(t.ExportVariables, hasher, ignore)
				}

			case *v1.Policy_RolePolicy:
				if t.RolePolicy != nil {
					cerbos_policy_v1_RolePolicy_hashpb_sum(t.RolePolicy, hasher, ignore)
				}

			case *v1.Policy_ExportConstants:
				if t.ExportConstants != nil {
					cerbos_policy_v1_ExportConstants_hashpb_sum(t.ExportConstants, hasher, ignore)
				}

			}
		}
	}
	if _, ok := ignore["cerbos.policy.v1.Policy.variables"]; !ok {
		if len(m.Variables) > 0 {
			keys := make([]string, len(m.Variables))
			i := 0
			for k := range m.Variables {
				keys[i] = k
				i++
			}

			sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })

			for _, k := range keys {
				_, _ = hasher.Write(protowire.AppendString(nil, m.Variables[k]))

			}
		}
	}
	if _, ok := ignore["cerbos.policy.v1.Policy.json_schema"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.GetJsonSchema()))

	}
}

func cerbos_policy_v1_PrincipalPolicy_hashpb_sum(m *v1.PrincipalPolicy, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.policy.v1.PrincipalPolicy.principal"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.GetPrincipal()))

	}
	if _, ok := ignore["cerbos.policy.v1.PrincipalPolicy.version"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.GetVersion()))

	}
	if _, ok := ignore["cerbos.policy.v1.PrincipalPolicy.rules"]; !ok {
		if len(m.Rules) > 0 {
			for _, v := range m.Rules {
				if v != nil {
					cerbos_policy_v1_PrincipalRule_hashpb_sum(v, hasher, ignore)
				}

			}
		}
	}
	if _, ok := ignore["cerbos.policy.v1.PrincipalPolicy.scope"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.GetScope()))

	}
	if _, ok := ignore["cerbos.policy.v1.PrincipalPolicy.variables"]; !ok {
		if m.GetVariables() != nil {
			cerbos_policy_v1_Variables_hashpb_sum(m.GetVariables(), hasher, ignore)
		}

	}
	if _, ok := ignore["cerbos.policy.v1.PrincipalPolicy.scope_permissions"]; !ok {
		_, _ = hasher.Write(protowire.AppendVarint(nil, uint64(m.GetScopePermissions())))

	}
	if _, ok := ignore["cerbos.policy.v1.PrincipalPolicy.constants"]; !ok {
		if m.GetConstants() != nil {
			cerbos_policy_v1_Constants_hashpb_sum(m.GetConstants(), hasher, ignore)
		}

	}
}

func cerbos_policy_v1_PrincipalRule_Action_hashpb_sum(m *v1.PrincipalRule_Action, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.policy.v1.PrincipalRule.Action.action"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.GetAction()))

	}
	if _, ok := ignore["cerbos.policy.v1.PrincipalRule.Action.condition"]; !ok {
		if m.GetCondition() != nil {
			cerbos_policy_v1_Condition_hashpb_sum(m.GetCondition(), hasher, ignore)
		}

	}
	if _, ok := ignore["cerbos.policy.v1.PrincipalRule.Action.effect"]; !ok {
		_, _ = hasher.Write(protowire.AppendVarint(nil, uint64(m.GetEffect())))

	}
	if _, ok := ignore["cerbos.policy.v1.PrincipalRule.Action.name"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.GetName()))

	}
	if _, ok := ignore["cerbos.policy.v1.PrincipalRule.Action.output"]; !ok {
		if m.GetOutput() != nil {
			cerbos_policy_v1_Output_hashpb_sum(m.GetOutput(), hasher, ignore)
		}

	}
}

func cerbos_policy_v1_PrincipalRule_hashpb_sum(m *v1.PrincipalRule, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.policy.v1.PrincipalRule.resource"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.GetResource()))

	}
	if _, ok := ignore["cerbos.policy.v1.PrincipalRule.actions"]; !ok {
		if len(m.Actions) > 0 {
			for _, v := range m.Actions {
				if v != nil {
					cerbos_policy_v1_PrincipalRule_Action_hashpb_sum(v, hasher, ignore)
				}

			}
		}
	}
}

func cerbos_policy_v1_ResourcePolicy_hashpb_sum(m *v1.ResourcePolicy, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.policy.v1.ResourcePolicy.resource"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.GetResource()))

	}
	if _, ok := ignore["cerbos.policy.v1.ResourcePolicy.version"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.GetVersion()))

	}
	if _, ok := ignore["cerbos.policy.v1.ResourcePolicy.import_derived_roles"]; !ok {
		if len(m.ImportDerivedRoles) > 0 {
			for _, v := range m.ImportDerivedRoles {
				_, _ = hasher.Write(protowire.AppendString(nil, v))

			}
		}
	}
	if _, ok := ignore["cerbos.policy.v1.ResourcePolicy.rules"]; !ok {
		if len(m.Rules) > 0 {
			for _, v := range m.Rules {
				if v != nil {
					cerbos_policy_v1_ResourceRule_hashpb_sum(v, hasher, ignore)
				}

			}
		}
	}
	if _, ok := ignore["cerbos.policy.v1.ResourcePolicy.scope"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.GetScope()))

	}
	if _, ok := ignore["cerbos.policy.v1.ResourcePolicy.schemas"]; !ok {
		if m.GetSchemas() != nil {
			cerbos_policy_v1_Schemas_hashpb_sum(m.GetSchemas(), hasher, ignore)
		}

	}
	if _, ok := ignore["cerbos.policy.v1.ResourcePolicy.variables"]; !ok {
		if m.GetVariables() != nil {
			cerbos_policy_v1_Variables_hashpb_sum(m.GetVariables(), hasher, ignore)
		}

	}
	if _, ok := ignore["cerbos.policy.v1.ResourcePolicy.scope_permissions"]; !ok {
		_, _ = hasher.Write(protowire.AppendVarint(nil, uint64(m.GetScopePermissions())))

	}
	if _, ok := ignore["cerbos.policy.v1.ResourcePolicy.constants"]; !ok {
		if m.GetConstants() != nil {
			cerbos_policy_v1_Constants_hashpb_sum(m.GetConstants(), hasher, ignore)
		}

	}
}

func cerbos_policy_v1_ResourceRule_hashpb_sum(m *v1.ResourceRule, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.policy.v1.ResourceRule.actions"]; !ok {
		if len(m.Actions) > 0 {
			for _, v := range m.Actions {
				_, _ = hasher.Write(protowire.AppendString(nil, v))

			}
		}
	}
	if _, ok := ignore["cerbos.policy.v1.ResourceRule.derived_roles"]; !ok {
		if len(m.DerivedRoles) > 0 {
			for _, v := range m.DerivedRoles {
				_, _ = hasher.Write(protowire.AppendString(nil, v))

			}
		}
	}
	if _, ok := ignore["cerbos.policy.v1.ResourceRule.roles"]; !ok {
		if len(m.Roles) > 0 {
			for _, v := range m.Roles {
				_, _ = hasher.Write(protowire.AppendString(nil, v))

			}
		}
	}
	if _, ok := ignore["cerbos.policy.v1.ResourceRule.condition"]; !ok {
		if m.GetCondition() != nil {
			cerbos_policy_v1_Condition_hashpb_sum(m.GetCondition(), hasher, ignore)
		}

	}
	if _, ok := ignore["cerbos.policy.v1.ResourceRule.effect"]; !ok {
		_, _ = hasher.Write(protowire.AppendVarint(nil, uint64(m.GetEffect())))

	}
	if _, ok := ignore["cerbos.policy.v1.ResourceRule.name"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.GetName()))

	}
	if _, ok := ignore["cerbos.policy.v1.ResourceRule.output"]; !ok {
		if m.GetOutput() != nil {
			cerbos_policy_v1_Output_hashpb_sum(m.GetOutput(), hasher, ignore)
		}

	}
}

func cerbos_policy_v1_RoleDef_hashpb_sum(m *v1.RoleDef, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.policy.v1.RoleDef.name"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.GetName()))

	}
	if _, ok := ignore["cerbos.policy.v1.RoleDef.parent_roles"]; !ok {
		if len(m.ParentRoles) > 0 {
			for _, v := range m.ParentRoles {
				_, _ = hasher.Write(protowire.AppendString(nil, v))

			}
		}
	}
	if _, ok := ignore["cerbos.policy.v1.RoleDef.condition"]; !ok {
		if m.GetCondition() != nil {
			cerbos_policy_v1_Condition_hashpb_sum(m.GetCondition(), hasher, ignore)
		}

	}
}

func cerbos_policy_v1_RolePolicy_hashpb_sum(m *v1.RolePolicy, hasher hash.Hash, ignore map[string]struct{}) {
	if m.PolicyType != nil {
		if _, ok := ignore["cerbos.policy.v1.RolePolicy.policy_type"]; !ok {
			switch t := m.PolicyType.(type) {
			case *v1.RolePolicy_Role:
				_, _ = hasher.Write(protowire.AppendString(nil, t.Role))

			}
		}
	}
	if _, ok := ignore["cerbos.policy.v1.RolePolicy.scope"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.GetScope()))

	}
	if _, ok := ignore["cerbos.policy.v1.RolePolicy.rules"]; !ok {
		if len(m.Rules) > 0 {
			for _, v := range m.Rules {
				if v != nil {
					cerbos_policy_v1_RoleRule_hashpb_sum(v, hasher, ignore)
				}

			}
		}
	}
	if _, ok := ignore["cerbos.policy.v1.RolePolicy.scope_permissions"]; !ok {
		_, _ = hasher.Write(protowire.AppendVarint(nil, uint64(m.GetScopePermissions())))

	}
	if _, ok := ignore["cerbos.policy.v1.RolePolicy.parent_roles"]; !ok {
		if len(m.ParentRoles) > 0 {
			for _, v := range m.ParentRoles {
				_, _ = hasher.Write(protowire.AppendString(nil, v))

			}
		}
	}
}

func cerbos_policy_v1_RoleRule_hashpb_sum(m *v1.RoleRule, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.policy.v1.RoleRule.resource"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.GetResource()))

	}
	if _, ok := ignore["cerbos.policy.v1.RoleRule.allow_actions"]; !ok {
		if len(m.AllowActions) > 0 {
			for _, v := range m.AllowActions {
				_, _ = hasher.Write(protowire.AppendString(nil, v))

			}
		}
	}
}

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
		_, _ = hasher.Write(protowire.AppendString(nil, m.GetRef()))

	}
	if _, ok := ignore["cerbos.policy.v1.Schemas.Schema.ignore_when"]; !ok {
		if m.GetIgnoreWhen() != nil {
			cerbos_policy_v1_Schemas_IgnoreWhen_hashpb_sum(m.GetIgnoreWhen(), hasher, ignore)
		}

	}
}

func cerbos_policy_v1_Schemas_hashpb_sum(m *v1.Schemas, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.policy.v1.Schemas.principal_schema"]; !ok {
		if m.GetPrincipalSchema() != nil {
			cerbos_policy_v1_Schemas_Schema_hashpb_sum(m.GetPrincipalSchema(), hasher, ignore)
		}

	}
	if _, ok := ignore["cerbos.policy.v1.Schemas.resource_schema"]; !ok {
		if m.GetResourceSchema() != nil {
			cerbos_policy_v1_Schemas_Schema_hashpb_sum(m.GetResourceSchema(), hasher, ignore)
		}

	}
}

func cerbos_policy_v1_SourceAttributes_hashpb_sum(m *v1.SourceAttributes, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.policy.v1.SourceAttributes.attributes"]; !ok {
		if len(m.Attributes) > 0 {
			keys := make([]string, len(m.Attributes))
			i := 0
			for k := range m.Attributes {
				keys[i] = k
				i++
			}

			sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })

			for _, k := range keys {
				if m.Attributes[k] != nil {
					google_protobuf_Value_hashpb_sum(m.Attributes[k], hasher, ignore)
				}

			}
		}
	}
}

func cerbos_policy_v1_Variables_hashpb_sum(m *v1.Variables, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.policy.v1.Variables.import"]; !ok {
		if len(m.Import) > 0 {
			for _, v := range m.Import {
				_, _ = hasher.Write(protowire.AppendString(nil, v))

			}
		}
	}
	if _, ok := ignore["cerbos.policy.v1.Variables.local"]; !ok {
		if len(m.Local) > 0 {
			keys := make([]string, len(m.Local))
			i := 0
			for k := range m.Local {
				keys[i] = k
				i++
			}

			sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })

			for _, k := range keys {
				_, _ = hasher.Write(protowire.AppendString(nil, m.Local[k]))

			}
		}
	}
}

func cerbos_source_v1_Error_hashpb_sum(m *Error, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.source.v1.Error.kind"]; !ok {
		_, _ = hasher.Write(protowire.AppendVarint(nil, uint64(m.GetKind())))

	}
	if _, ok := ignore["cerbos.source.v1.Error.position"]; !ok {
		if m.GetPosition() != nil {
			cerbos_source_v1_Position_hashpb_sum(m.GetPosition(), hasher, ignore)
		}

	}
	if _, ok := ignore["cerbos.source.v1.Error.message"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.GetMessage()))

	}
	if _, ok := ignore["cerbos.source.v1.Error.context"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.GetContext()))

	}
}

func cerbos_source_v1_PolicyWrapper_hashpb_sum(m *PolicyWrapper, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.source.v1.PolicyWrapper.id"]; !ok {
		_, _ = hasher.Write(protowire.AppendVarint(nil, m.GetId()))

	}
	if _, ok := ignore["cerbos.source.v1.PolicyWrapper.key"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.GetKey()))

	}
	if _, ok := ignore["cerbos.source.v1.PolicyWrapper.policy"]; !ok {
		if m.GetPolicy() != nil {
			cerbos_policy_v1_Policy_hashpb_sum(m.GetPolicy(), hasher, ignore)
		}

	}
	if _, ok := ignore["cerbos.source.v1.PolicyWrapper.kind"]; !ok {
		_, _ = hasher.Write(protowire.AppendVarint(nil, uint64(m.GetKind())))

	}
	if _, ok := ignore["cerbos.source.v1.PolicyWrapper.name"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.GetName()))

	}
	if _, ok := ignore["cerbos.source.v1.PolicyWrapper.version"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.GetVersion()))

	}
	if _, ok := ignore["cerbos.source.v1.PolicyWrapper.scope"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.GetScope()))

	}
}

func cerbos_source_v1_Position_hashpb_sum(m *Position, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.source.v1.Position.line"]; !ok {
		_, _ = hasher.Write(protowire.AppendVarint(nil, uint64(m.GetLine())))

	}
	if _, ok := ignore["cerbos.source.v1.Position.column"]; !ok {
		_, _ = hasher.Write(protowire.AppendVarint(nil, uint64(m.GetColumn())))

	}
	if _, ok := ignore["cerbos.source.v1.Position.path"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.GetPath()))

	}
}

func cerbos_source_v1_SourceContext_hashpb_sum(m *SourceContext, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.source.v1.SourceContext.start_position"]; !ok {
		if m.GetStartPosition() != nil {
			cerbos_source_v1_StartPosition_hashpb_sum(m.GetStartPosition(), hasher, ignore)
		}

	}
	if _, ok := ignore["cerbos.source.v1.SourceContext.field_positions"]; !ok {
		if len(m.FieldPositions) > 0 {
			keys := make([]string, len(m.FieldPositions))
			i := 0
			for k := range m.FieldPositions {
				keys[i] = k
				i++
			}

			sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })

			for _, k := range keys {
				if m.FieldPositions[k] != nil {
					cerbos_source_v1_Position_hashpb_sum(m.FieldPositions[k], hasher, ignore)
				}

			}
		}
	}
	if _, ok := ignore["cerbos.source.v1.SourceContext.errors"]; !ok {
		if len(m.Errors) > 0 {
			for _, v := range m.Errors {
				if v != nil {
					cerbos_source_v1_Error_hashpb_sum(v, hasher, ignore)
				}

			}
		}
	}
}

func cerbos_source_v1_StartPosition_hashpb_sum(m *StartPosition, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.source.v1.StartPosition.line"]; !ok {
		_, _ = hasher.Write(protowire.AppendVarint(nil, uint64(m.GetLine())))

	}
	if _, ok := ignore["cerbos.source.v1.StartPosition.column"]; !ok {
		_, _ = hasher.Write(protowire.AppendVarint(nil, uint64(m.GetColumn())))

	}
	if _, ok := ignore["cerbos.source.v1.StartPosition.offset"]; !ok {
		_, _ = hasher.Write(protowire.AppendVarint(nil, uint64(m.GetOffset())))

	}
}

func google_protobuf_ListValue_hashpb_sum(m *structpb.ListValue, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["google.protobuf.ListValue.values"]; !ok {
		if len(m.Values) > 0 {
			for _, v := range m.Values {
				if v != nil {
					google_protobuf_Value_hashpb_sum(v, hasher, ignore)
				}

			}
		}
	}
}

func google_protobuf_Struct_hashpb_sum(m *structpb.Struct, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["google.protobuf.Struct.fields"]; !ok {
		if len(m.Fields) > 0 {
			keys := make([]string, len(m.Fields))
			i := 0
			for k := range m.Fields {
				keys[i] = k
				i++
			}

			sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })

			for _, k := range keys {
				if m.Fields[k] != nil {
					google_protobuf_Value_hashpb_sum(m.Fields[k], hasher, ignore)
				}

			}
		}
	}
}

func google_protobuf_UInt64Value_hashpb_sum(m *wrapperspb.UInt64Value, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["google.protobuf.UInt64Value.value"]; !ok {
		_, _ = hasher.Write(protowire.AppendVarint(nil, m.GetValue()))

	}
}

func google_protobuf_Value_hashpb_sum(m *structpb.Value, hasher hash.Hash, ignore map[string]struct{}) {
	if m.Kind != nil {
		if _, ok := ignore["google.protobuf.Value.kind"]; !ok {
			switch t := m.Kind.(type) {
			case *structpb.Value_NullValue:
				_, _ = hasher.Write(protowire.AppendVarint(nil, uint64(t.NullValue)))

			case *structpb.Value_NumberValue:
				_, _ = hasher.Write(protowire.AppendFixed64(nil, math.Float64bits(t.NumberValue)))

			case *structpb.Value_StringValue:
				_, _ = hasher.Write(protowire.AppendString(nil, t.StringValue))

			case *structpb.Value_BoolValue:
				_, _ = hasher.Write(protowire.AppendVarint(nil, protowire.EncodeBool(t.BoolValue)))

			case *structpb.Value_StructValue:
				if t.StructValue != nil {
					google_protobuf_Struct_hashpb_sum(t.StructValue, hasher, ignore)
				}

			case *structpb.Value_ListValue:
				if t.ListValue != nil {
					google_protobuf_ListValue_hashpb_sum(t.ListValue, hasher, ignore)
				}

			}
		}
	}
}
