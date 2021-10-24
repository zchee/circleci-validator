// Copyright 2021 The circleci-validator Authors
// SPDX-License-Identifier: BSD-3-Clause

package ccivalidator

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
)

// AddSSHKeysCommandSchema json schema for the AddSSHKeys command.
type AddSSHKeysCommandSchema struct {
	AddSshKeys *AddSSHKeysParameters `json:"add_ssh_keys"`
}

func (r *AddSSHKeysCommandSchema) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// "AddSshKeys" field is required
	if r.AddSshKeys == nil {
		return nil, errors.New("add_ssh_keys is a required field")
	}
	// Marshal the "add_ssh_keys" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"add_ssh_keys\": ")
	if tmp, err := json.Marshal(r.AddSshKeys); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (r *AddSSHKeysCommandSchema) UnmarshalJSON(b []byte) error {
	add_ssh_keysReceived := false
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "add_ssh_keys":
			if err := json.Unmarshal([]byte(v), &r.AddSshKeys); err != nil {
				return err
			}
			add_ssh_keysReceived = true
		}
	}
	// check if add_ssh_keys (a required property) was received
	if !add_ssh_keysReceived {
		return errors.New("\"add_ssh_keys\" is required but was not present")
	}
	return nil
}

// AttachCommandSchema json schema for the attach command.
type AttachCommandSchema struct {
	AttachWorkspace *AttachParameters `json:"attach_workspace"`
}

func (r *AttachCommandSchema) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// "AttachWorkspace" field is required
	if r.AttachWorkspace == nil {
		return nil, errors.New("attach_workspace is a required field")
	}
	// Marshal the "attach_workspace" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"attach_workspace\": ")
	if tmp, err := json.Marshal(r.AttachWorkspace); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (r *AttachCommandSchema) UnmarshalJSON(b []byte) error {
	attach_workspaceReceived := false
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "attach_workspace":
			if err := json.Unmarshal([]byte(v), &r.AttachWorkspace); err != nil {
				return err
			}
			attach_workspaceReceived = true
		}
	}
	// check if attach_workspace (a required property) was received
	if !attach_workspaceReceived {
		return errors.New("\"attach_workspace\" is required but was not present")
	}
	return nil
}

// CheckoutCommandSchema
type CheckoutCommandSchema struct {
	Checkout *CheckoutParameter `json:"checkout"`
}

func (r *CheckoutCommandSchema) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// "Checkout" field is required
	if r.Checkout == nil {
		return nil, errors.New("checkout is a required field")
	}
	// Marshal the "checkout" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"checkout\": ")
	if tmp, err := json.Marshal(r.Checkout); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (r *CheckoutCommandSchema) UnmarshalJSON(b []byte) error {
	checkoutReceived := false
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "checkout":
			if err := json.Unmarshal([]byte(v), &r.Checkout); err != nil {
				return err
			}
			checkoutReceived = true
		}
	}
	// check if checkout (a required property) was received
	if !checkoutReceived {
		return errors.New("\"checkout\" is required but was not present")
	}
	return nil
}

// CommandSchema
type CommandSchema struct{}

// CircleCIConfigSchema json schema for the circleci config.
type CircleCIConfigSchema struct {
	Commands  *CommandSchema     `json:"commands,omitempty"`
	Jobs      *JobSchema         `json:"jobs"`
	Orbs      []*ConfigOrbImport `json:"orbs,omitempty"`
	Setup     bool               `json:"setup"`
	Version   float64            `json:"version"`
	Workflows *WorkflowSchema    `json:"workflows"`
}

func (r *CircleCIConfigSchema) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// Marshal the "commands" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"commands\": ")
	if tmp, err := json.Marshal(r.Commands); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "Jobs" field is required
	if r.Jobs == nil {
		return nil, errors.New("jobs is a required field")
	}
	// Marshal the "jobs" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"jobs\": ")
	if tmp, err := json.Marshal(r.Jobs); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "orbs" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"orbs\": ")
	if tmp, err := json.Marshal(r.Orbs); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "Setup" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "setup" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"setup\": ")
	if tmp, err := json.Marshal(r.Setup); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "Version" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "version" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"version\": ")
	if tmp, err := json.Marshal(r.Version); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "Workflows" field is required
	if r.Workflows == nil {
		return nil, errors.New("workflows is a required field")
	}
	// Marshal the "workflows" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"workflows\": ")
	if tmp, err := json.Marshal(r.Workflows); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (r *CircleCIConfigSchema) UnmarshalJSON(b []byte) error {
	jobsReceived := false
	setupReceived := false
	versionReceived := false
	workflowsReceived := false
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "commands":
			if err := json.Unmarshal([]byte(v), &r.Commands); err != nil {
				return err
			}
		case "jobs":
			if err := json.Unmarshal([]byte(v), &r.Jobs); err != nil {
				return err
			}
			jobsReceived = true
		case "orbs":
			if err := json.Unmarshal([]byte(v), &r.Orbs); err != nil {
				return err
			}
		case "setup":
			if err := json.Unmarshal([]byte(v), &r.Setup); err != nil {
				return err
			}
			setupReceived = true
		case "version":
			if err := json.Unmarshal([]byte(v), &r.Version); err != nil {
				return err
			}
			versionReceived = true
		case "workflows":
			if err := json.Unmarshal([]byte(v), &r.Workflows); err != nil {
				return err
			}
			workflowsReceived = true
		}
	}
	// check if jobs (a required property) was received
	if !jobsReceived {
		return errors.New("\"jobs\" is required but was not present")
	}
	// check if setup (a required property) was received
	if !setupReceived {
		return errors.New("\"setup\" is required but was not present")
	}
	// check if version (a required property) was received
	if !versionReceived {
		return errors.New("\"version\" is required but was not present")
	}
	// check if workflows (a required property) was received
	if !workflowsReceived {
		return errors.New("\"workflows\" is required but was not present")
	}
	return nil
}

// DockerExecutorSchema a json representation of the docker executor schema to be converted to yaml.
type DockerExecutorSchema struct {
	Docker        []*DockerImageSchema `json:"docker"`
	ResourceClass string               `json:"resource_class"`
}

func (r *DockerExecutorSchema) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// "Docker" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "docker" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"docker\": ")
	if tmp, err := json.Marshal(r.Docker); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "ResourceClass" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "resource_class" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"resource_class\": ")
	if tmp, err := json.Marshal(r.ResourceClass); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (r *DockerExecutorSchema) UnmarshalJSON(b []byte) error {
	dockerReceived := false
	resource_classReceived := false
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "docker":
			if err := json.Unmarshal([]byte(v), &r.Docker); err != nil {
				return err
			}
			dockerReceived = true
		case "resource_class":
			if err := json.Unmarshal([]byte(v), &r.ResourceClass); err != nil {
				return err
			}
			resource_classReceived = true
		}
	}
	// check if docker (a required property) was received
	if !dockerReceived {
		return errors.New("\"docker\" is required but was not present")
	}
	// check if resource_class (a required property) was received
	if !resource_classReceived {
		return errors.New("\"resource_class\" is required but was not present")
	}
	return nil
}

// DockerImageSchema
type DockerImageSchema struct {
	Auth        *DockerAuth       `json:"auth,omitempty"`
	AwsAuth     *DockerAuthAWS    `json:"aws_auth,omitempty"`
	Command     []string          `json:"command,omitempty"`
	Entrypoint  []string          `json:"entrypoint,omitempty"`
	Environment map[string]string `json:"environment,omitempty"`
	Image       string            `json:"image"`
	Name        string            `json:"name,omitempty"`
	User        string            `json:"user,omitempty"`
}

func (r *DockerImageSchema) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// Marshal the "auth" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"auth\": ")
	if tmp, err := json.Marshal(r.Auth); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "aws_auth" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"aws_auth\": ")
	if tmp, err := json.Marshal(r.AwsAuth); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "command" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"command\": ")
	if tmp, err := json.Marshal(r.Command); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "entrypoint" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"entrypoint\": ")
	if tmp, err := json.Marshal(r.Entrypoint); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "environment" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"environment\": ")
	if tmp, err := json.Marshal(r.Environment); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "Image" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "image" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"image\": ")
	if tmp, err := json.Marshal(r.Image); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "name" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"name\": ")
	if tmp, err := json.Marshal(r.Name); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "user" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"user\": ")
	if tmp, err := json.Marshal(r.User); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (r *DockerImageSchema) UnmarshalJSON(b []byte) error {
	imageReceived := false
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "auth":
			if err := json.Unmarshal([]byte(v), &r.Auth); err != nil {
				return err
			}
		case "aws_auth":
			if err := json.Unmarshal([]byte(v), &r.AwsAuth); err != nil {
				return err
			}
		case "command":
			if err := json.Unmarshal([]byte(v), &r.Command); err != nil {
				return err
			}
		case "entrypoint":
			if err := json.Unmarshal([]byte(v), &r.Entrypoint); err != nil {
				return err
			}
		case "environment":
			if err := json.Unmarshal([]byte(v), &r.Environment); err != nil {
				return err
			}
		case "image":
			if err := json.Unmarshal([]byte(v), &r.Image); err != nil {
				return err
			}
			imageReceived = true
		case "name":
			if err := json.Unmarshal([]byte(v), &r.Name); err != nil {
				return err
			}
		case "user":
			if err := json.Unmarshal([]byte(v), &r.User); err != nil {
				return err
			}
		}
	}
	// check if image (a required property) was received
	if !imageReceived {
		return errors.New("\"image\" is required but was not present")
	}
	return nil
}

// JobSchema
type JobSchema struct {
	AdditionalProperties map[string]interface{} `json:"-,omitempty"`
}

func (r *JobSchema) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// Marshal any additional Properties
	for k, v := range r.AdditionalProperties {
		if comma {
			buf.WriteString(",")
		}
		buf.WriteString(fmt.Sprintf("\"%s\":", k))
		if tmp, err := json.Marshal(v); err != nil {
			return nil, err
		} else {
			buf.Write(tmp)
		}
		comma = true
	}

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (r *JobSchema) UnmarshalJSON(b []byte) error {
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		default:
			// an additional "interface{}" value
			var additionalValue interface{}
			if err := json.Unmarshal([]byte(v), &additionalValue); err != nil {
				return err // invalid additionalProperty
			}
			if r.AdditionalProperties == nil {
				r.AdditionalProperties = make(map[string]interface{})
			}
			r.AdditionalProperties[k] = additionalValue
		}
	}
	return nil
}

// JobStepsSchema
type JobStepsSchema struct {
	Steps []interface{} `json:"steps"`
}

func (r *JobStepsSchema) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// "Steps" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "steps" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"steps\": ")
	if tmp, err := json.Marshal(r.Steps); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (r *JobStepsSchema) UnmarshalJSON(b []byte) error {
	stepsReceived := false
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "steps":
			if err := json.Unmarshal([]byte(v), &r.Steps); err != nil {
				return err
			}
			stepsReceived = true
		}
	}
	// check if steps (a required property) was received
	if !stepsReceived {
		return errors.New("\"steps\" is required but was not present")
	}
	return nil
}

// MacOSExecutorSchema a json representation of the macOS executor schema to be converted to yaml.
type MacOSExecutorSchema struct {
	Macos         *Macos `json:"macos"`
	ResourceClass string `json:"resource_class"`
}

func (r *MacOSExecutorSchema) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// "Macos" field is required
	if r.Macos == nil {
		return nil, errors.New("macos is a required field")
	}
	// Marshal the "macos" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"macos\": ")
	if tmp, err := json.Marshal(r.Macos); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "ResourceClass" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "resource_class" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"resource_class\": ")
	if tmp, err := json.Marshal(r.ResourceClass); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (r *MacOSExecutorSchema) UnmarshalJSON(b []byte) error {
	macosReceived := false
	resource_classReceived := false
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "macos":
			if err := json.Unmarshal([]byte(v), &r.Macos); err != nil {
				return err
			}
			macosReceived = true
		case "resource_class":
			if err := json.Unmarshal([]byte(v), &r.ResourceClass); err != nil {
				return err
			}
			resource_classReceived = true
		}
	}
	// check if macos (a required property) was received
	if !macosReceived {
		return errors.New("\"macos\" is required but was not present")
	}
	// check if resource_class (a required property) was received
	if !resource_classReceived {
		return errors.New("\"resource_class\" is required but was not present")
	}
	return nil
}

// MachineExecutorSchema
type MachineExecutorSchema struct {
	Machine       *Machine `json:"machine"`
	ResourceClass string   `json:"resource_class"`
}

func (r *MachineExecutorSchema) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// "Machine" field is required
	if r.Machine == nil {
		return nil, errors.New("machine is a required field")
	}
	// Marshal the "machine" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"machine\": ")
	if tmp, err := json.Marshal(r.Machine); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "ResourceClass" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "resource_class" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"resource_class\": ")
	if tmp, err := json.Marshal(r.ResourceClass); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (r *MachineExecutorSchema) UnmarshalJSON(b []byte) error {
	machineReceived := false
	resource_classReceived := false
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "machine":
			if err := json.Unmarshal([]byte(v), &r.Machine); err != nil {
				return err
			}
			machineReceived = true
		case "resource_class":
			if err := json.Unmarshal([]byte(v), &r.ResourceClass); err != nil {
				return err
			}
			resource_classReceived = true
		}
	}
	// check if machine (a required property) was received
	if !machineReceived {
		return errors.New("\"machine\" is required but was not present")
	}
	// check if resource_class (a required property) was received
	if !resource_classReceived {
		return errors.New("\"resource_class\" is required but was not present")
	}
	return nil
}

// PersistCommandSchema json schema for the persist command.
type PersistCommandSchema struct {
	PersistToWorkspace *PersistParameters `json:"persist_to_workspace"`
}

func (r *PersistCommandSchema) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// "PersistToWorkspace" field is required
	if r.PersistToWorkspace == nil {
		return nil, errors.New("persist_to_workspace is a required field")
	}
	// Marshal the "persist_to_workspace" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"persist_to_workspace\": ")
	if tmp, err := json.Marshal(r.PersistToWorkspace); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (r *PersistCommandSchema) UnmarshalJSON(b []byte) error {
	persist_to_workspaceReceived := false
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "persist_to_workspace":
			if err := json.Unmarshal([]byte(v), &r.PersistToWorkspace); err != nil {
				return err
			}
			persist_to_workspaceReceived = true
		}
	}
	// check if persist_to_workspace (a required property) was received
	if !persist_to_workspaceReceived {
		return errors.New("\"persist_to_workspace\" is required but was not present")
	}
	return nil
}

// PipelineParameterSchema
type PipelineParameterSchema struct {
	AdditionalProperties map[string]*PipelineParameterSchemaItem `json:"-,omitempty"`
}

func (r *PipelineParameterSchema) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// Marshal any additional Properties
	for k, v := range r.AdditionalProperties {
		if comma {
			buf.WriteString(",")
		}
		buf.WriteString(fmt.Sprintf("\"%s\":", k))
		if tmp, err := json.Marshal(v); err != nil {
			return nil, err
		} else {
			buf.Write(tmp)
		}
		comma = true
	}

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (r *PipelineParameterSchema) UnmarshalJSON(b []byte) error {
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		default:
			// an additional "*PipelineParameterSchemaItem" value
			var additionalValue *PipelineParameterSchemaItem
			if err := json.Unmarshal([]byte(v), &additionalValue); err != nil {
				return err // invalid additionalProperty
			}
			if r.AdditionalProperties == nil {
				r.AdditionalProperties = make(map[string]*PipelineParameterSchemaItem)
			}
			r.AdditionalProperties[k] = additionalValue
		}
	}
	return nil
}

// PipelineParameterSchemaItem
type PipelineParameterSchemaItem struct {
	Default       interface{} `json:"default"`
	Enum          []string    `json:"enum,omitempty"`
	ParameterType string      `json:"parameterType"`
}

func (r *PipelineParameterSchemaItem) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// "Default" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "default" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"default\": ")
	if tmp, err := json.Marshal(r.Default); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "enum" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"enum\": ")
	if tmp, err := json.Marshal(r.Enum); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "ParameterType" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "parameterType" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"parameterType\": ")
	if tmp, err := json.Marshal(r.ParameterType); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (r *PipelineParameterSchemaItem) UnmarshalJSON(b []byte) error {
	defaultReceived := false
	parameterTypeReceived := false
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "default":
			if err := json.Unmarshal([]byte(v), &r.Default); err != nil {
				return err
			}
			defaultReceived = true
		case "enum":
			if err := json.Unmarshal([]byte(v), &r.Enum); err != nil {
				return err
			}
		case "parameterType":
			if err := json.Unmarshal([]byte(v), &r.ParameterType); err != nil {
				return err
			}
			parameterTypeReceived = true
		}
	}
	// check if default (a required property) was received
	if !defaultReceived {
		return errors.New("\"default\" is required but was not present")
	}
	// check if parameterType (a required property) was received
	if !parameterTypeReceived {
		return errors.New("\"parameterType\" is required but was not present")
	}
	return nil
}

// RestoreCacheCommandSchema json schema for the restorecache command.
type RestoreCacheCommandSchema struct {
	RestoreCache *RestoreCacheParameters `json:"restore_cache"`
}

func (r *RestoreCacheCommandSchema) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// "RestoreCache" field is required
	if r.RestoreCache == nil {
		return nil, errors.New("restore_cache is a required field")
	}
	// Marshal the "restore_cache" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"restore_cache\": ")
	if tmp, err := json.Marshal(r.RestoreCache); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (r *RestoreCacheCommandSchema) UnmarshalJSON(b []byte) error {
	restore_cacheReceived := false
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "restore_cache":
			if err := json.Unmarshal([]byte(v), &r.RestoreCache); err != nil {
				return err
			}
			restore_cacheReceived = true
		}
	}
	// check if restore_cache (a required property) was received
	if !restore_cacheReceived {
		return errors.New("\"restore_cache\" is required but was not present")
	}
	return nil
}

// RunCommandSchema json schema for the run command.
type RunCommandSchema struct {
	Run *RunParameters `json:"run"`
}

func (r *RunCommandSchema) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// "Run" field is required
	if r.Run == nil {
		return nil, errors.New("run is a required field")
	}
	// Marshal the "run" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"run\": ")
	if tmp, err := json.Marshal(r.Run); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (r *RunCommandSchema) UnmarshalJSON(b []byte) error {
	runReceived := false
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "run":
			if err := json.Unmarshal([]byte(v), &r.Run); err != nil {
				return err
			}
			runReceived = true
		}
	}
	// check if run (a required property) was received
	if !runReceived {
		return errors.New("\"run\" is required but was not present")
	}
	return nil
}

// SaveCacheCommandSchema json schema for the savecache command.
type SaveCacheCommandSchema struct {
	SaveCache *SaveCacheParameters `json:"save_cache"`
}

func (r *SaveCacheCommandSchema) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// "SaveCache" field is required
	if r.SaveCache == nil {
		return nil, errors.New("save_cache is a required field")
	}
	// Marshal the "save_cache" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"save_cache\": ")
	if tmp, err := json.Marshal(r.SaveCache); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (r *SaveCacheCommandSchema) UnmarshalJSON(b []byte) error {
	save_cacheReceived := false
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "save_cache":
			if err := json.Unmarshal([]byte(v), &r.SaveCache); err != nil {
				return err
			}
			save_cacheReceived = true
		}
	}
	// check if save_cache (a required property) was received
	if !save_cacheReceived {
		return errors.New("\"save_cache\" is required but was not present")
	}
	return nil
}

// SetupRemoteDockerCommandSchema json schema for the setupremotedocker command.
type SetupRemoteDockerCommandSchema struct {
	SetupRemoteDocker *SetupRemoteDockerParameters `json:"setup_remote_docker"`
}

func (r *SetupRemoteDockerCommandSchema) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// "SetupRemoteDocker" field is required
	if r.SetupRemoteDocker == nil {
		return nil, errors.New("setup_remote_docker is a required field")
	}
	// Marshal the "setup_remote_docker" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"setup_remote_docker\": ")
	if tmp, err := json.Marshal(r.SetupRemoteDocker); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (r *SetupRemoteDockerCommandSchema) UnmarshalJSON(b []byte) error {
	setup_remote_dockerReceived := false
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "setup_remote_docker":
			if err := json.Unmarshal([]byte(v), &r.SetupRemoteDocker); err != nil {
				return err
			}
			setup_remote_dockerReceived = true
		}
	}
	// check if setup_remote_docker (a required property) was received
	if !setup_remote_dockerReceived {
		return errors.New("\"setup_remote_docker\" is required but was not present")
	}
	return nil
}

// StoreArtifactsCommandSchema json schema for the storeartifacts command.
type StoreArtifactsCommandSchema struct {
	StoreArtifacts *StoreArtifactsParameters `json:"store_artifacts"`
}

func (r *StoreArtifactsCommandSchema) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// "StoreArtifacts" field is required
	if r.StoreArtifacts == nil {
		return nil, errors.New("store_artifacts is a required field")
	}
	// Marshal the "store_artifacts" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"store_artifacts\": ")
	if tmp, err := json.Marshal(r.StoreArtifacts); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (r *StoreArtifactsCommandSchema) UnmarshalJSON(b []byte) error {
	store_artifactsReceived := false
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "store_artifacts":
			if err := json.Unmarshal([]byte(v), &r.StoreArtifacts); err != nil {
				return err
			}
			store_artifactsReceived = true
		}
	}
	// check if store_artifacts (a required property) was received
	if !store_artifactsReceived {
		return errors.New("\"store_artifacts\" is required but was not present")
	}
	return nil
}

// StoreTestResultsCommandSchema json schema for the storetestresults command.
type StoreTestResultsCommandSchema struct {
	StoreTestResults *StoreTestResultsParameters `json:"store_test_results"`
}

func (r *StoreTestResultsCommandSchema) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// "StoreTestResults" field is required
	if r.StoreTestResults == nil {
		return nil, errors.New("store_test_results is a required field")
	}
	// Marshal the "store_test_results" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"store_test_results\": ")
	if tmp, err := json.Marshal(r.StoreTestResults); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (r *StoreTestResultsCommandSchema) UnmarshalJSON(b []byte) error {
	store_test_resultsReceived := false
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "store_test_results":
			if err := json.Unmarshal([]byte(v), &r.StoreTestResults); err != nil {
				return err
			}
			store_test_resultsReceived = true
		}
	}
	// check if store_test_results (a required property) was received
	if !store_test_resultsReceived {
		return errors.New("\"store_test_results\" is required but was not present")
	}
	return nil
}

// WindowsExecutorSchema a json representation of the Windows executor schema to be converted to yaml.
type WindowsExecutorSchema struct {
	Machine       *Machine `json:"machine"`
	ResourceClass string   `json:"resource_class"`
	Shell         string   `json:"shell"`
}

func (r *WindowsExecutorSchema) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// "Machine" field is required
	if r.Machine == nil {
		return nil, errors.New("machine is a required field")
	}
	// Marshal the "machine" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"machine\": ")
	if tmp, err := json.Marshal(r.Machine); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "ResourceClass" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "resource_class" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"resource_class\": ")
	if tmp, err := json.Marshal(r.ResourceClass); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "Shell" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "shell" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"shell\": ")
	if tmp, err := json.Marshal(r.Shell); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (r *WindowsExecutorSchema) UnmarshalJSON(b []byte) error {
	machineReceived := false
	resource_classReceived := false
	shellReceived := false
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "machine":
			if err := json.Unmarshal([]byte(v), &r.Machine); err != nil {
				return err
			}
			machineReceived = true
		case "resource_class":
			if err := json.Unmarshal([]byte(v), &r.ResourceClass); err != nil {
				return err
			}
			resource_classReceived = true
		case "shell":
			if err := json.Unmarshal([]byte(v), &r.Shell); err != nil {
				return err
			}
			shellReceived = true
		}
	}
	// check if machine (a required property) was received
	if !machineReceived {
		return errors.New("\"machine\" is required but was not present")
	}
	// check if resource_class (a required property) was received
	if !resource_classReceived {
		return errors.New("\"resource_class\" is required but was not present")
	}
	// check if shell (a required property) was received
	if !shellReceived {
		return errors.New("\"shell\" is required but was not present")
	}
	return nil
}

// WorkflowFilterSchema
type WorkflowFilterSchema struct {
	// A map defining rules for execution on specific branches
	Branches *Branches `json:"branches,omitempty"`

	// A map defining rules for execution on specific tags
	Tags *Tags `json:"tags,omitempty"`
}

// WorkflowJobSchema
type WorkflowJobSchema struct {
	AdditionalProperties map[string]*WorkflowJobSchemaItem `json:"-,omitempty"`
}

func (r *WorkflowJobSchema) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// Marshal any additional Properties
	for k, v := range r.AdditionalProperties {
		if comma {
			buf.WriteString(",")
		}
		buf.WriteString(fmt.Sprintf("\"%s\":", k))
		if tmp, err := json.Marshal(v); err != nil {
			return nil, err
		} else {
			buf.Write(tmp)
		}
		comma = true
	}

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (r *WorkflowJobSchema) UnmarshalJSON(b []byte) error {
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		default:
			// an additional "*WorkflowJobSchemaItem" value
			var additionalValue *WorkflowJobSchemaItem
			if err := json.Unmarshal([]byte(v), &additionalValue); err != nil {
				return err // invalid additionalProperty
			}
			if r.AdditionalProperties == nil {
				r.AdditionalProperties = make(map[string]*WorkflowJobSchemaItem)
			}
			r.AdditionalProperties[k] = additionalValue
		}
	}
	return nil
}

// WorkflowJobSchemaItem
type WorkflowJobSchemaItem struct {
	Context  []string              `json:"context,omitempty"`
	Filters  *WorkflowFilterSchema `json:"filters,omitempty"`
	JobType  string                `json:"jobType,omitempty"`
	Matrix   *WorkflowMatrixSchema `json:"matrix,omitempty"`
	Requires []string              `json:"requires,omitempty"`
}

// WorkflowMatrixSchema a map of parameter names to every value the job should be called with.
type WorkflowMatrixSchema struct {
	Parameters map[string][]string `json:"parameters"`
}

func (r *WorkflowMatrixSchema) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// "Parameters" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "parameters" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"parameters\": ")
	if tmp, err := json.Marshal(r.Parameters); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (r *WorkflowMatrixSchema) UnmarshalJSON(b []byte) error {
	parametersReceived := false
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "parameters":
			if err := json.Unmarshal([]byte(v), &r.Parameters); err != nil {
				return err
			}
			parametersReceived = true
		}
	}
	// check if parameters (a required property) was received
	if !parametersReceived {
		return errors.New("\"parameters\" is required but was not present")
	}
	return nil
}

// WorkflowSchema
type WorkflowSchema struct {
	AdditionalProperties map[string]*WorkflowSchemaItem `json:"-,omitempty"`
}

func (r *WorkflowSchema) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// Marshal any additional Properties
	for k, v := range r.AdditionalProperties {
		if comma {
			buf.WriteString(",")
		}
		buf.WriteString(fmt.Sprintf("\"%s\":", k))
		if tmp, err := json.Marshal(v); err != nil {
			return nil, err
		} else {
			buf.Write(tmp)
		}
		comma = true
	}

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (r *WorkflowSchema) UnmarshalJSON(b []byte) error {
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		default:
			// an additional "*WorkflowSchemaItem" value
			var additionalValue *WorkflowSchemaItem
			if err := json.Unmarshal([]byte(v), &additionalValue); err != nil {
				return err // invalid additionalProperty
			}
			if r.AdditionalProperties == nil {
				r.AdditionalProperties = make(map[string]*WorkflowSchemaItem)
			}
			r.AdditionalProperties[k] = additionalValue
		}
	}
	return nil
}

// WorkflowSchemaItem
type WorkflowSchemaItem struct {
	Jobs []*WorkflowJobSchema `json:"jobs"`
}

func (r *WorkflowSchemaItem) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// "Jobs" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "jobs" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"jobs\": ")
	if tmp, err := json.Marshal(r.Jobs); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (r *WorkflowSchemaItem) UnmarshalJSON(b []byte) error {
	jobsReceived := false
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "jobs":
			if err := json.Unmarshal([]byte(v), &r.Jobs); err != nil {
				return err
			}
			jobsReceived = true
		}
	}
	// check if jobs (a required property) was received
	if !jobsReceived {
		return errors.New("\"jobs\" is required but was not present")
	}
	return nil
}
