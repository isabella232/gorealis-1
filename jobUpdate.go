/**
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package realis

import (
	"github.com/paypal/gorealis/gen-go/apache/aurora"
)

// Structure to collect all information required to create job update
type JobUpdate struct {
	task    *AuroraTask
	request *aurora.JobUpdateRequest
}

// Create a default JobUpdate object with an empty task and no fields filled in.
func NewJobUpdate(task *AuroraTask) *JobUpdate {
	newTask := NewTask()

	req := aurora.JobUpdateRequest{}
	req.TaskConfig = newTask.TaskConfig()
	req.Settings = NewUpdateSettings()

	return &JobUpdate{task: newTask, request: &req}
}

func JobUpdateFormTask(task *AuroraTask) *JobUpdate {
	// Perform a deep copy to avoid unexpected behavior
	newTask := task.Clone()

	req := aurora.JobUpdateRequest{}
	req.TaskConfig = newTask.TaskConfig()
	req.Settings = NewUpdateSettings()

	return &JobUpdate{task: newTask, request: &req}
}

// Set instance count the job will have after the update.
func (j *JobUpdate) InstanceCount(inst int32) *JobUpdate {
	j.request.InstanceCount = inst
	return j
}

// Max number of instances being updated at any given moment.
func (j *JobUpdate) BatchSize(size int32) *JobUpdate {
	j.request.Settings.UpdateGroupSize = size
	return j
}

// Minimum number of seconds a shard must remain in RUNNING state before considered a success.
func (j *JobUpdate) WatchTime(ms int32) *JobUpdate {
	j.request.Settings.MinWaitInInstanceRunningMs = ms
	return j
}

// Wait for all instances in a group to be done before moving on.
func (j *JobUpdate) WaitForBatchCompletion(batchWait bool) *JobUpdate {
	j.request.Settings.WaitForBatchCompletion = batchWait
	return j
}

// Max number of instance failures to tolerate before marking instance as FAILED.
func (j *JobUpdate) MaxPerInstanceFailures(inst int32) *JobUpdate {
	j.request.Settings.MaxPerInstanceFailures = inst
	return j
}

// Max number of FAILED instances to tolerate before terminating the update.
func (j *JobUpdate) MaxFailedInstances(inst int32) *JobUpdate {
	j.request.Settings.MaxFailedInstances = inst
	return j
}

// When False, prevents auto rollback of a failed update.
func (j *JobUpdate) RollbackOnFail(rollback bool) *JobUpdate {
	j.request.Settings.RollbackOnFailure = rollback
	return j
}

func NewUpdateSettings() *aurora.JobUpdateSettings {

	us := aurora.JobUpdateSettings{}
	// Mirrors defaults set by Pystachio
	us.UpdateOnlyTheseInstances = []*aurora.Range{}
	us.UpdateGroupSize = 1
	us.WaitForBatchCompletion = false
	us.MinWaitInInstanceRunningMs = 45000
	us.MaxPerInstanceFailures = 0
	us.MaxFailedInstances = 0
	us.RollbackOnFailure = true

	return &us
}

/*
   AuroraTask specific API, see task.go for further documentation.
   These functions are provided for the convenience of chaining API calls.
*/

func (j *JobUpdate) ExecutorName(name string) *JobUpdate {
	j.task.ExecutorName(name)
	return j
}

func (j *JobUpdate) ExecutorData(data string) *JobUpdate {
	j.task.ExecutorData(data)
	return j
}

func (j *JobUpdate) CPU(cpus float64) *JobUpdate {
	j.task.CPU(cpus)
	return j
}

func (j *JobUpdate) RAM(ram int64) *JobUpdate {
	j.task.RAM(ram)
	return j
}

func (j *JobUpdate) Disk(disk int64) *JobUpdate {
	j.task.Disk(disk)
	return j
}

func (j *JobUpdate) Tier(tier string) *JobUpdate {
	j.task.Tier(tier)
	return j
}

func (j *JobUpdate) MaxFailure(maxFail int32) *JobUpdate {
	j.task.MaxFailure(maxFail)
	return j
}

func (j *JobUpdate) IsService(isService bool) *JobUpdate {
	j.task.IsService(isService)
	return j
}

func (j *JobUpdate) TaskConfig() *aurora.TaskConfig {
	return j.task.TaskConfig()
}

func (j *JobUpdate) AddURIs(extract bool, cache bool, values ...string) *JobUpdate {
	j.task.AddURIs(extract, cache, values...)
	return j
}

func (j *JobUpdate) AddLabel(key string, value string) *JobUpdate {
	j.task.AddLabel(key, value)
	return j
}

func (j *JobUpdate) AddNamedPorts(names ...string) *JobUpdate {
	j.task.AddNamedPorts(names...)
	return j
}

func (j *JobUpdate) AddPorts(num int) *JobUpdate {
	j.task.AddPorts(num)
	return j
}
func (j *JobUpdate) AddValueConstraint(name string, negated bool, values ...string) *JobUpdate {
	j.task.AddValueConstraint(name, negated, values...)
	return j
}

func (j *JobUpdate) AddLimitConstraint(name string, limit int32) *JobUpdate {
	j.task.AddLimitConstraint(name, limit)
	return j
}

func (j *JobUpdate) AddDedicatedConstraint(role, name string) *JobUpdate {
	j.task.AddDedicatedConstraint(role, name)
	return j
}

func (j *JobUpdate) Container(container Container) *JobUpdate {
	j.task.Container(container)
	return j
}