// Code generated by protoc-gen-go-hashpb. Do not edit.
// protoc-gen-go-hashpb v0.1.0

package telemetryv1

import (
	protowire "google.golang.org/protobuf/encoding/protowire"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	hash "hash"
	math "math"
	sort "sort"
)

func cerbos_telemetry_v1_Event_ApiActivity_hashpb_sum(m *Event_ApiActivity, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.telemetry.v1.Event.ApiActivity.version"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.Version))

	}
	if _, ok := ignore["cerbos.telemetry.v1.Event.ApiActivity.uptime"]; !ok {
		if m.Uptime != nil {
			google_protobuf_Duration_hashpb_sum(m.Uptime, hasher, ignore)
		}

	}
	if _, ok := ignore["cerbos.telemetry.v1.Event.ApiActivity.method_calls"]; !ok {
		if len(m.MethodCalls) > 0 {
			for _, v := range m.MethodCalls {
				if v != nil {
					cerbos_telemetry_v1_Event_CountStat_hashpb_sum(v, hasher, ignore)
				}

			}
		}
	}
	if _, ok := ignore["cerbos.telemetry.v1.Event.ApiActivity.user_agents"]; !ok {
		if len(m.UserAgents) > 0 {
			for _, v := range m.UserAgents {
				if v != nil {
					cerbos_telemetry_v1_Event_CountStat_hashpb_sum(v, hasher, ignore)
				}

			}
		}
	}
}

func cerbos_telemetry_v1_Event_CountStat_hashpb_sum(m *Event_CountStat, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.telemetry.v1.Event.CountStat.key"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.Key))

	}
	if _, ok := ignore["cerbos.telemetry.v1.Event.CountStat.count"]; !ok {
		_, _ = hasher.Write(protowire.AppendVarint(nil, m.Count))

	}
}

func cerbos_telemetry_v1_Event_hashpb_sum(m *Event, hasher hash.Hash, ignore map[string]struct{}) {
	if m.Data != nil {
		if _, ok := ignore["cerbos.telemetry.v1.Event.data"]; !ok {
			switch t := m.Data.(type) {
			case *Event_ApiActivity_:
				if t.ApiActivity != nil {
					cerbos_telemetry_v1_Event_ApiActivity_hashpb_sum(t.ApiActivity, hasher, ignore)
				}

			}
		}
	}
}

func cerbos_telemetry_v1_ServerLaunch_Cerbos_hashpb_sum(m *ServerLaunch_Cerbos, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.telemetry.v1.ServerLaunch.Cerbos.version"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.Version))

	}
	if _, ok := ignore["cerbos.telemetry.v1.ServerLaunch.Cerbos.commit"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.Commit))

	}
	if _, ok := ignore["cerbos.telemetry.v1.ServerLaunch.Cerbos.build_date"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.BuildDate))

	}
	if _, ok := ignore["cerbos.telemetry.v1.ServerLaunch.Cerbos.module_version"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.ModuleVersion))

	}
	if _, ok := ignore["cerbos.telemetry.v1.ServerLaunch.Cerbos.module_checksum"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.ModuleChecksum))

	}
}

func cerbos_telemetry_v1_ServerLaunch_Features_AdminApi_hashpb_sum(m *ServerLaunch_Features_AdminApi, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.telemetry.v1.ServerLaunch.Features.AdminApi.enabled"]; !ok {
		_, _ = hasher.Write(protowire.AppendVarint(nil, protowire.EncodeBool(m.Enabled)))

	}
}

func cerbos_telemetry_v1_ServerLaunch_Features_Audit_hashpb_sum(m *ServerLaunch_Features_Audit, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.telemetry.v1.ServerLaunch.Features.Audit.enabled"]; !ok {
		_, _ = hasher.Write(protowire.AppendVarint(nil, protowire.EncodeBool(m.Enabled)))

	}
	if _, ok := ignore["cerbos.telemetry.v1.ServerLaunch.Features.Audit.backend"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.Backend))

	}
}

func cerbos_telemetry_v1_ServerLaunch_Features_Schema_hashpb_sum(m *ServerLaunch_Features_Schema, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.telemetry.v1.ServerLaunch.Features.Schema.enforcement"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.Enforcement))

	}
}

func cerbos_telemetry_v1_ServerLaunch_Features_Storage_Blob_hashpb_sum(m *ServerLaunch_Features_Storage_Blob, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.telemetry.v1.ServerLaunch.Features.Storage.Blob.provider"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.Provider))

	}
	if _, ok := ignore["cerbos.telemetry.v1.ServerLaunch.Features.Storage.Blob.poll_interval"]; !ok {
		if m.PollInterval != nil {
			google_protobuf_Duration_hashpb_sum(m.PollInterval, hasher, ignore)
		}

	}
}

func cerbos_telemetry_v1_ServerLaunch_Features_Storage_Disk_hashpb_sum(m *ServerLaunch_Features_Storage_Disk, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.telemetry.v1.ServerLaunch.Features.Storage.Disk.watch"]; !ok {
		_, _ = hasher.Write(protowire.AppendVarint(nil, protowire.EncodeBool(m.Watch)))

	}
}

func cerbos_telemetry_v1_ServerLaunch_Features_Storage_Git_hashpb_sum(m *ServerLaunch_Features_Storage_Git, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.telemetry.v1.ServerLaunch.Features.Storage.Git.protocol"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.Protocol))

	}
	if _, ok := ignore["cerbos.telemetry.v1.ServerLaunch.Features.Storage.Git.auth"]; !ok {
		_, _ = hasher.Write(protowire.AppendVarint(nil, protowire.EncodeBool(m.Auth)))

	}
	if _, ok := ignore["cerbos.telemetry.v1.ServerLaunch.Features.Storage.Git.poll_interval"]; !ok {
		if m.PollInterval != nil {
			google_protobuf_Duration_hashpb_sum(m.PollInterval, hasher, ignore)
		}

	}
}

func cerbos_telemetry_v1_ServerLaunch_Features_Storage_hashpb_sum(m *ServerLaunch_Features_Storage, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.telemetry.v1.ServerLaunch.Features.Storage.driver"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.Driver))

	}
	if m.Store != nil {
		if _, ok := ignore["cerbos.telemetry.v1.ServerLaunch.Features.Storage.store"]; !ok {
			switch t := m.Store.(type) {
			case *ServerLaunch_Features_Storage_Disk_:
				if t.Disk != nil {
					cerbos_telemetry_v1_ServerLaunch_Features_Storage_Disk_hashpb_sum(t.Disk, hasher, ignore)
				}

			case *ServerLaunch_Features_Storage_Git_:
				if t.Git != nil {
					cerbos_telemetry_v1_ServerLaunch_Features_Storage_Git_hashpb_sum(t.Git, hasher, ignore)
				}

			case *ServerLaunch_Features_Storage_Blob_:
				if t.Blob != nil {
					cerbos_telemetry_v1_ServerLaunch_Features_Storage_Blob_hashpb_sum(t.Blob, hasher, ignore)
				}

			}
		}
	}
}

func cerbos_telemetry_v1_ServerLaunch_Features_hashpb_sum(m *ServerLaunch_Features, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.telemetry.v1.ServerLaunch.Features.audit"]; !ok {
		if m.Audit != nil {
			cerbos_telemetry_v1_ServerLaunch_Features_Audit_hashpb_sum(m.Audit, hasher, ignore)
		}

	}
	if _, ok := ignore["cerbos.telemetry.v1.ServerLaunch.Features.schema"]; !ok {
		if m.Schema != nil {
			cerbos_telemetry_v1_ServerLaunch_Features_Schema_hashpb_sum(m.Schema, hasher, ignore)
		}

	}
	if _, ok := ignore["cerbos.telemetry.v1.ServerLaunch.Features.admin_api"]; !ok {
		if m.AdminApi != nil {
			cerbos_telemetry_v1_ServerLaunch_Features_AdminApi_hashpb_sum(m.AdminApi, hasher, ignore)
		}

	}
	if _, ok := ignore["cerbos.telemetry.v1.ServerLaunch.Features.storage"]; !ok {
		if m.Storage != nil {
			cerbos_telemetry_v1_ServerLaunch_Features_Storage_hashpb_sum(m.Storage, hasher, ignore)
		}

	}
}

func cerbos_telemetry_v1_ServerLaunch_Source_hashpb_sum(m *ServerLaunch_Source, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.telemetry.v1.ServerLaunch.Source.cerbos"]; !ok {
		if m.Cerbos != nil {
			cerbos_telemetry_v1_ServerLaunch_Cerbos_hashpb_sum(m.Cerbos, hasher, ignore)
		}

	}
	if _, ok := ignore["cerbos.telemetry.v1.ServerLaunch.Source.os"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.Os))

	}
	if _, ok := ignore["cerbos.telemetry.v1.ServerLaunch.Source.arch"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.Arch))

	}
	if _, ok := ignore["cerbos.telemetry.v1.ServerLaunch.Source.num_cpus"]; !ok {
		_, _ = hasher.Write(protowire.AppendVarint(nil, uint64(m.NumCpus)))

	}
}

func cerbos_telemetry_v1_ServerLaunch_Stats_Policy_hashpb_sum(m *ServerLaunch_Stats_Policy, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.telemetry.v1.ServerLaunch.Stats.Policy.count"]; !ok {
		if len(m.Count) > 0 {
			keys := make([]string, len(m.Count))
			i := 0
			for k := range m.Count {
				keys[i] = k
				i++
			}

			sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })

			for _, k := range keys {
				_, _ = hasher.Write(protowire.AppendVarint(nil, uint64(m.Count[k])))

			}
		}
	}
	if _, ok := ignore["cerbos.telemetry.v1.ServerLaunch.Stats.Policy.avg_rule_count"]; !ok {
		if len(m.AvgRuleCount) > 0 {
			keys := make([]string, len(m.AvgRuleCount))
			i := 0
			for k := range m.AvgRuleCount {
				keys[i] = k
				i++
			}

			sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })

			for _, k := range keys {
				_, _ = hasher.Write(protowire.AppendFixed64(nil, math.Float64bits(m.AvgRuleCount[k])))

			}
		}
	}
	if _, ok := ignore["cerbos.telemetry.v1.ServerLaunch.Stats.Policy.avg_condition_count"]; !ok {
		if len(m.AvgConditionCount) > 0 {
			keys := make([]string, len(m.AvgConditionCount))
			i := 0
			for k := range m.AvgConditionCount {
				keys[i] = k
				i++
			}

			sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })

			for _, k := range keys {
				_, _ = hasher.Write(protowire.AppendFixed64(nil, math.Float64bits(m.AvgConditionCount[k])))

			}
		}
	}
}

func cerbos_telemetry_v1_ServerLaunch_Stats_Schema_hashpb_sum(m *ServerLaunch_Stats_Schema, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.telemetry.v1.ServerLaunch.Stats.Schema.count"]; !ok {
		_, _ = hasher.Write(protowire.AppendVarint(nil, uint64(m.Count)))

	}
}

func cerbos_telemetry_v1_ServerLaunch_Stats_hashpb_sum(m *ServerLaunch_Stats, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.telemetry.v1.ServerLaunch.Stats.policy"]; !ok {
		if m.Policy != nil {
			cerbos_telemetry_v1_ServerLaunch_Stats_Policy_hashpb_sum(m.Policy, hasher, ignore)
		}

	}
	if _, ok := ignore["cerbos.telemetry.v1.ServerLaunch.Stats.schema"]; !ok {
		if m.Schema != nil {
			cerbos_telemetry_v1_ServerLaunch_Stats_Schema_hashpb_sum(m.Schema, hasher, ignore)
		}

	}
}

func cerbos_telemetry_v1_ServerLaunch_hashpb_sum(m *ServerLaunch, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.telemetry.v1.ServerLaunch.version"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.Version))

	}
	if _, ok := ignore["cerbos.telemetry.v1.ServerLaunch.source"]; !ok {
		if m.Source != nil {
			cerbos_telemetry_v1_ServerLaunch_Source_hashpb_sum(m.Source, hasher, ignore)
		}

	}
	if _, ok := ignore["cerbos.telemetry.v1.ServerLaunch.features"]; !ok {
		if m.Features != nil {
			cerbos_telemetry_v1_ServerLaunch_Features_hashpb_sum(m.Features, hasher, ignore)
		}

	}
	if _, ok := ignore["cerbos.telemetry.v1.ServerLaunch.stats"]; !ok {
		if m.Stats != nil {
			cerbos_telemetry_v1_ServerLaunch_Stats_hashpb_sum(m.Stats, hasher, ignore)
		}

	}
}

func cerbos_telemetry_v1_ServerStop_hashpb_sum(m *ServerStop, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.telemetry.v1.ServerStop.version"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.Version))

	}
	if _, ok := ignore["cerbos.telemetry.v1.ServerStop.uptime"]; !ok {
		if m.Uptime != nil {
			google_protobuf_Duration_hashpb_sum(m.Uptime, hasher, ignore)
		}

	}
	if _, ok := ignore["cerbos.telemetry.v1.ServerStop.requests_total"]; !ok {
		_, _ = hasher.Write(protowire.AppendVarint(nil, m.RequestsTotal))

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