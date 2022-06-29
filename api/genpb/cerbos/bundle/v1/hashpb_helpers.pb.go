// Code generated by protoc-gen-go-hashpb. Do not edit.
// protoc-gen-go-hashpb v0.1.0

package bundlev1

import (
	protowire "google.golang.org/protobuf/encoding/protowire"
	hash "hash"
	sort "sort"
)

func cerbos_bundle_v1_Manifest_hashpb_sum(m *Manifest, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.bundle.v1.Manifest.version"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.Version))

	}
	if _, ok := ignore["cerbos.bundle.v1.Manifest.policy_index"]; !ok {
		if len(m.PolicyIndex) > 0 {
			keys := make([]string, len(m.PolicyIndex))
			i := 0
			for k := range m.PolicyIndex {
				keys[i] = k
				i++
			}

			sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })

			for _, k := range keys {
				_, _ = hasher.Write(protowire.AppendString(nil, m.PolicyIndex[k]))

			}
		}
	}
}
