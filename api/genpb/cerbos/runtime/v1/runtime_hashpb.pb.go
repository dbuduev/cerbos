// Code generated by protoc-gen-go-hashpb. Do not edit.
// protoc-gen-go-hashpb v0.3.4
// Source: cerbos/runtime/v1/runtime.proto

package runtimev1

import (
	hash "hash"
)

// HashPB computes a hash of the message using the given hash function
// The ignore set must contain fully-qualified field names (pkg.msg.field) that should be ignored from the hash
func (m *RunnablePolicySet) HashPB(hasher hash.Hash, ignore map[string]struct{}) {
	if m != nil {
		cerbos_runtime_v1_RunnablePolicySet_hashpb_sum(m, hasher, ignore)
	}
}

// HashPB computes a hash of the message using the given hash function
// The ignore set must contain fully-qualified field names (pkg.msg.field) that should be ignored from the hash
func (m *RunnableRolePolicySet) HashPB(hasher hash.Hash, ignore map[string]struct{}) {
	if m != nil {
		cerbos_runtime_v1_RunnableRolePolicySet_hashpb_sum(m, hasher, ignore)
	}
}

// HashPB computes a hash of the message using the given hash function
// The ignore set must contain fully-qualified field names (pkg.msg.field) that should be ignored from the hash
func (m *RunnableRolePolicySet_Metadata) HashPB(hasher hash.Hash, ignore map[string]struct{}) {
	if m != nil {
		cerbos_runtime_v1_RunnableRolePolicySet_Metadata_hashpb_sum(m, hasher, ignore)
	}
}

// HashPB computes a hash of the message using the given hash function
// The ignore set must contain fully-qualified field names (pkg.msg.field) that should be ignored from the hash
func (m *RunnableRolePolicySet_AllowActions) HashPB(hasher hash.Hash, ignore map[string]struct{}) {
	if m != nil {
		cerbos_runtime_v1_RunnableRolePolicySet_AllowActions_hashpb_sum(m, hasher, ignore)
	}
}

// HashPB computes a hash of the message using the given hash function
// The ignore set must contain fully-qualified field names (pkg.msg.field) that should be ignored from the hash
func (m *RunnableResourcePolicySet) HashPB(hasher hash.Hash, ignore map[string]struct{}) {
	if m != nil {
		cerbos_runtime_v1_RunnableResourcePolicySet_hashpb_sum(m, hasher, ignore)
	}
}

// HashPB computes a hash of the message using the given hash function
// The ignore set must contain fully-qualified field names (pkg.msg.field) that should be ignored from the hash
func (m *RunnableResourcePolicySet_Metadata) HashPB(hasher hash.Hash, ignore map[string]struct{}) {
	if m != nil {
		cerbos_runtime_v1_RunnableResourcePolicySet_Metadata_hashpb_sum(m, hasher, ignore)
	}
}

// HashPB computes a hash of the message using the given hash function
// The ignore set must contain fully-qualified field names (pkg.msg.field) that should be ignored from the hash
func (m *RunnableResourcePolicySet_Policy) HashPB(hasher hash.Hash, ignore map[string]struct{}) {
	if m != nil {
		cerbos_runtime_v1_RunnableResourcePolicySet_Policy_hashpb_sum(m, hasher, ignore)
	}
}

// HashPB computes a hash of the message using the given hash function
// The ignore set must contain fully-qualified field names (pkg.msg.field) that should be ignored from the hash
func (m *RunnableResourcePolicySet_Policy_Rule) HashPB(hasher hash.Hash, ignore map[string]struct{}) {
	if m != nil {
		cerbos_runtime_v1_RunnableResourcePolicySet_Policy_Rule_hashpb_sum(m, hasher, ignore)
	}
}

// HashPB computes a hash of the message using the given hash function
// The ignore set must contain fully-qualified field names (pkg.msg.field) that should be ignored from the hash
func (m *RunnableDerivedRole) HashPB(hasher hash.Hash, ignore map[string]struct{}) {
	if m != nil {
		cerbos_runtime_v1_RunnableDerivedRole_hashpb_sum(m, hasher, ignore)
	}
}

// HashPB computes a hash of the message using the given hash function
// The ignore set must contain fully-qualified field names (pkg.msg.field) that should be ignored from the hash
func (m *RunnableDerivedRolesSet) HashPB(hasher hash.Hash, ignore map[string]struct{}) {
	if m != nil {
		cerbos_runtime_v1_RunnableDerivedRolesSet_hashpb_sum(m, hasher, ignore)
	}
}

// HashPB computes a hash of the message using the given hash function
// The ignore set must contain fully-qualified field names (pkg.msg.field) that should be ignored from the hash
func (m *RunnableDerivedRolesSet_Metadata) HashPB(hasher hash.Hash, ignore map[string]struct{}) {
	if m != nil {
		cerbos_runtime_v1_RunnableDerivedRolesSet_Metadata_hashpb_sum(m, hasher, ignore)
	}
}

// HashPB computes a hash of the message using the given hash function
// The ignore set must contain fully-qualified field names (pkg.msg.field) that should be ignored from the hash
func (m *RunnableVariablesSet) HashPB(hasher hash.Hash, ignore map[string]struct{}) {
	if m != nil {
		cerbos_runtime_v1_RunnableVariablesSet_hashpb_sum(m, hasher, ignore)
	}
}

// HashPB computes a hash of the message using the given hash function
// The ignore set must contain fully-qualified field names (pkg.msg.field) that should be ignored from the hash
func (m *RunnableVariablesSet_Metadata) HashPB(hasher hash.Hash, ignore map[string]struct{}) {
	if m != nil {
		cerbos_runtime_v1_RunnableVariablesSet_Metadata_hashpb_sum(m, hasher, ignore)
	}
}

// HashPB computes a hash of the message using the given hash function
// The ignore set must contain fully-qualified field names (pkg.msg.field) that should be ignored from the hash
func (m *RunnablePrincipalPolicySet) HashPB(hasher hash.Hash, ignore map[string]struct{}) {
	if m != nil {
		cerbos_runtime_v1_RunnablePrincipalPolicySet_hashpb_sum(m, hasher, ignore)
	}
}

// HashPB computes a hash of the message using the given hash function
// The ignore set must contain fully-qualified field names (pkg.msg.field) that should be ignored from the hash
func (m *RunnablePrincipalPolicySet_Metadata) HashPB(hasher hash.Hash, ignore map[string]struct{}) {
	if m != nil {
		cerbos_runtime_v1_RunnablePrincipalPolicySet_Metadata_hashpb_sum(m, hasher, ignore)
	}
}

// HashPB computes a hash of the message using the given hash function
// The ignore set must contain fully-qualified field names (pkg.msg.field) that should be ignored from the hash
func (m *RunnablePrincipalPolicySet_Policy) HashPB(hasher hash.Hash, ignore map[string]struct{}) {
	if m != nil {
		cerbos_runtime_v1_RunnablePrincipalPolicySet_Policy_hashpb_sum(m, hasher, ignore)
	}
}

// HashPB computes a hash of the message using the given hash function
// The ignore set must contain fully-qualified field names (pkg.msg.field) that should be ignored from the hash
func (m *RunnablePrincipalPolicySet_Policy_ActionRule) HashPB(hasher hash.Hash, ignore map[string]struct{}) {
	if m != nil {
		cerbos_runtime_v1_RunnablePrincipalPolicySet_Policy_ActionRule_hashpb_sum(m, hasher, ignore)
	}
}

// HashPB computes a hash of the message using the given hash function
// The ignore set must contain fully-qualified field names (pkg.msg.field) that should be ignored from the hash
func (m *RunnablePrincipalPolicySet_Policy_ResourceRules) HashPB(hasher hash.Hash, ignore map[string]struct{}) {
	if m != nil {
		cerbos_runtime_v1_RunnablePrincipalPolicySet_Policy_ResourceRules_hashpb_sum(m, hasher, ignore)
	}
}

// HashPB computes a hash of the message using the given hash function
// The ignore set must contain fully-qualified field names (pkg.msg.field) that should be ignored from the hash
func (m *Expr) HashPB(hasher hash.Hash, ignore map[string]struct{}) {
	if m != nil {
		cerbos_runtime_v1_Expr_hashpb_sum(m, hasher, ignore)
	}
}

// HashPB computes a hash of the message using the given hash function
// The ignore set must contain fully-qualified field names (pkg.msg.field) that should be ignored from the hash
func (m *Output) HashPB(hasher hash.Hash, ignore map[string]struct{}) {
	if m != nil {
		cerbos_runtime_v1_Output_hashpb_sum(m, hasher, ignore)
	}
}

// HashPB computes a hash of the message using the given hash function
// The ignore set must contain fully-qualified field names (pkg.msg.field) that should be ignored from the hash
func (m *Output_When) HashPB(hasher hash.Hash, ignore map[string]struct{}) {
	if m != nil {
		cerbos_runtime_v1_Output_When_hashpb_sum(m, hasher, ignore)
	}
}

// HashPB computes a hash of the message using the given hash function
// The ignore set must contain fully-qualified field names (pkg.msg.field) that should be ignored from the hash
func (m *Variable) HashPB(hasher hash.Hash, ignore map[string]struct{}) {
	if m != nil {
		cerbos_runtime_v1_Variable_hashpb_sum(m, hasher, ignore)
	}
}

// HashPB computes a hash of the message using the given hash function
// The ignore set must contain fully-qualified field names (pkg.msg.field) that should be ignored from the hash
func (m *Condition) HashPB(hasher hash.Hash, ignore map[string]struct{}) {
	if m != nil {
		cerbos_runtime_v1_Condition_hashpb_sum(m, hasher, ignore)
	}
}

// HashPB computes a hash of the message using the given hash function
// The ignore set must contain fully-qualified field names (pkg.msg.field) that should be ignored from the hash
func (m *Condition_ExprList) HashPB(hasher hash.Hash, ignore map[string]struct{}) {
	if m != nil {
		cerbos_runtime_v1_Condition_ExprList_hashpb_sum(m, hasher, ignore)
	}
}

// HashPB computes a hash of the message using the given hash function
// The ignore set must contain fully-qualified field names (pkg.msg.field) that should be ignored from the hash
func (m *CompileErrors) HashPB(hasher hash.Hash, ignore map[string]struct{}) {
	if m != nil {
		cerbos_runtime_v1_CompileErrors_hashpb_sum(m, hasher, ignore)
	}
}

// HashPB computes a hash of the message using the given hash function
// The ignore set must contain fully-qualified field names (pkg.msg.field) that should be ignored from the hash
func (m *CompileErrors_Err) HashPB(hasher hash.Hash, ignore map[string]struct{}) {
	if m != nil {
		cerbos_runtime_v1_CompileErrors_Err_hashpb_sum(m, hasher, ignore)
	}
}

// HashPB computes a hash of the message using the given hash function
// The ignore set must contain fully-qualified field names (pkg.msg.field) that should be ignored from the hash
func (m *IndexBuildErrors) HashPB(hasher hash.Hash, ignore map[string]struct{}) {
	if m != nil {
		cerbos_runtime_v1_IndexBuildErrors_hashpb_sum(m, hasher, ignore)
	}
}

// HashPB computes a hash of the message using the given hash function
// The ignore set must contain fully-qualified field names (pkg.msg.field) that should be ignored from the hash
func (m *IndexBuildErrors_DuplicateDef) HashPB(hasher hash.Hash, ignore map[string]struct{}) {
	if m != nil {
		cerbos_runtime_v1_IndexBuildErrors_DuplicateDef_hashpb_sum(m, hasher, ignore)
	}
}

// HashPB computes a hash of the message using the given hash function
// The ignore set must contain fully-qualified field names (pkg.msg.field) that should be ignored from the hash
func (m *IndexBuildErrors_MissingImport) HashPB(hasher hash.Hash, ignore map[string]struct{}) {
	if m != nil {
		cerbos_runtime_v1_IndexBuildErrors_MissingImport_hashpb_sum(m, hasher, ignore)
	}
}

// HashPB computes a hash of the message using the given hash function
// The ignore set must contain fully-qualified field names (pkg.msg.field) that should be ignored from the hash
func (m *IndexBuildErrors_MissingScope) HashPB(hasher hash.Hash, ignore map[string]struct{}) {
	if m != nil {
		cerbos_runtime_v1_IndexBuildErrors_MissingScope_hashpb_sum(m, hasher, ignore)
	}
}

// HashPB computes a hash of the message using the given hash function
// The ignore set must contain fully-qualified field names (pkg.msg.field) that should be ignored from the hash
func (m *IndexBuildErrors_ScopePermissionsConflicts) HashPB(hasher hash.Hash, ignore map[string]struct{}) {
	if m != nil {
		cerbos_runtime_v1_IndexBuildErrors_ScopePermissionsConflicts_hashpb_sum(m, hasher, ignore)
	}
}

// HashPB computes a hash of the message using the given hash function
// The ignore set must contain fully-qualified field names (pkg.msg.field) that should be ignored from the hash
func (m *IndexBuildErrors_LoadFailure) HashPB(hasher hash.Hash, ignore map[string]struct{}) {
	if m != nil {
		cerbos_runtime_v1_IndexBuildErrors_LoadFailure_hashpb_sum(m, hasher, ignore)
	}
}

// HashPB computes a hash of the message using the given hash function
// The ignore set must contain fully-qualified field names (pkg.msg.field) that should be ignored from the hash
func (m *IndexBuildErrors_Disabled) HashPB(hasher hash.Hash, ignore map[string]struct{}) {
	if m != nil {
		cerbos_runtime_v1_IndexBuildErrors_Disabled_hashpb_sum(m, hasher, ignore)
	}
}

// HashPB computes a hash of the message using the given hash function
// The ignore set must contain fully-qualified field names (pkg.msg.field) that should be ignored from the hash
func (m *Errors) HashPB(hasher hash.Hash, ignore map[string]struct{}) {
	if m != nil {
		cerbos_runtime_v1_Errors_hashpb_sum(m, hasher, ignore)
	}
}
