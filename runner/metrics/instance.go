// Copyright 2025 Cloudbase Solutions SRL
//
//    Licensed under the Apache License, Version 2.0 (the "License"); you may
//    not use this file except in compliance with the License. You may obtain
//    a copy of the License at
//
//         http://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
//    WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
//    License for the specific language governing permissions and limitations
//    under the License.

package metrics

import (
	"context"

	"github.com/cloudbase/garm/metrics"
	"github.com/cloudbase/garm/runner"
)

// CollectInstanceMetric collects the metrics for the runner instances
// reflecting the statuses and the pool they belong to.
func CollectInstanceMetric(ctx context.Context, r *runner.Runner) error {
	// reset metrics
	metrics.InstanceStatus.Reset()

	instances, err := r.ListAllInstances(ctx)
	if err != nil {
		return err
	}

	pools, err := r.ListAllPools(ctx)
	if err != nil {
		return err
	}

	type poolInfo struct {
		Name         string
		Type         string
		ProviderName string
	}

	poolNames := make(map[string]poolInfo)
	for _, pool := range pools {
		switch {
		case pool.OrgName != "":
			poolNames[pool.ID] = poolInfo{
				Name: pool.OrgName,
				Type: string(pool.PoolType()),
			}
		case pool.EnterpriseName != "":
			poolNames[pool.ID] = poolInfo{
				Name: pool.EnterpriseName,
				Type: string(pool.PoolType()),
			}
		default:
			poolNames[pool.ID] = poolInfo{
				Name: pool.RepoName,
				Type: string(pool.PoolType()),
			}
		}
	}

	for _, instance := range instances {
		metrics.InstanceStatus.WithLabelValues(
			instance.Name,                           // label: name
			string(instance.Status),                 // label: status
			string(instance.RunnerStatus),           // label: runner_status
			poolNames[instance.PoolID].Name,         // label: pool_owner
			poolNames[instance.PoolID].Type,         // label: pool_type
			instance.PoolID,                         // label: pool_id
			poolNames[instance.PoolID].ProviderName, // label: provider
		).Set(1)
	}
	return nil
}
