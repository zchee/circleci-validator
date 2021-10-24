// Copyright 2021 The circleci-validator Authors
// SPDX-License-Identifier: BSD-3-Clause

package ccivalidator

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
)

// AddSSHKeys the AddSSHKeys command is a special step that adds SSH keys from a project’s settings to a container. Also configures SSH to use these keys.
type AddSSHKeys struct {
	Name       string                `json:"name"`
	Parameters *AddSSHKeysParameters `json:"parameters"`
}

func (r *AddSSHKeys) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// "Name" field is required
	// only required object types supported for marshal checking (for now)
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
	// "Parameters" field is required
	if r.Parameters == nil {
		return nil, errors.New("parameters is a required field")
	}
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

func (r *AddSSHKeys) UnmarshalJSON(b []byte) error {
	nameReceived := false
	parametersReceived := false
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "name":
			if err := json.Unmarshal([]byte(v), &r.Name); err != nil {
				return err
			}
			nameReceived = true
		case "parameters":
			if err := json.Unmarshal([]byte(v), &r.Parameters); err != nil {
				return err
			}
			parametersReceived = true
		}
	}
	// check if name (a required property) was received
	if !nameReceived {
		return errors.New("\"name\" is required but was not present")
	}
	// check if parameters (a required property) was received
	if !parametersReceived {
		return errors.New("\"parameters\" is required but was not present")
	}
	return nil
}

// AddSSHKeysParameters command parameters for the AddSSHKeys command.
type AddSSHKeysParameters struct {
	// List of fingerprints corresponding to the keys to be added.
	Fingerprints []string `json:"fingerprints"`

	// Title of the step to be shown in the CircleCI UI (default: full command)
	Name string `json:"name,omitempty"`
}

func (r *AddSSHKeysParameters) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// "Fingerprints" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "fingerprints" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"fingerprints\": ")
	if tmp, err := json.Marshal(r.Fingerprints); err != nil {
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

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (r *AddSSHKeysParameters) UnmarshalJSON(b []byte) error {
	fingerprintsReceived := false
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "fingerprints":
			if err := json.Unmarshal([]byte(v), &r.Fingerprints); err != nil {
				return err
			}
			fingerprintsReceived = true
		case "name":
			if err := json.Unmarshal([]byte(v), &r.Name); err != nil {
				return err
			}
		}
	}
	// check if fingerprints (a required property) was received
	if !fingerprintsReceived {
		return errors.New("\"fingerprints\" is required but was not present")
	}
	return nil
}

// Attach special step used to attach the workflow’s workspace to the current container.
//
// The full contents of the workspace are downloaded and copied into the directory the workspace is being attached at.
type Attach struct {
	Name       string            `json:"name"`
	Parameters *AttachParameters `json:"parameters"`
}

func (r *Attach) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// "Name" field is required
	// only required object types supported for marshal checking (for now)
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
	// "Parameters" field is required
	if r.Parameters == nil {
		return nil, errors.New("parameters is a required field")
	}
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

func (r *Attach) UnmarshalJSON(b []byte) error {
	nameReceived := false
	parametersReceived := false
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "name":
			if err := json.Unmarshal([]byte(v), &r.Name); err != nil {
				return err
			}
			nameReceived = true
		case "parameters":
			if err := json.Unmarshal([]byte(v), &r.Parameters); err != nil {
				return err
			}
			parametersReceived = true
		}
	}
	// check if name (a required property) was received
	if !nameReceived {
		return errors.New("\"name\" is required but was not present")
	}
	// check if parameters (a required property) was received
	if !parametersReceived {
		return errors.New("\"parameters\" is required but was not present")
	}
	return nil
}

// AttachParameters command parameters for the attach command.
type AttachParameters struct {
	// Directory to attach the workspace to.
	At string `json:"at"`

	// Title of the step to be shown in the CircleCI UI (default: full command)
	Name string `json:"name,omitempty"`
}

func (r *AttachParameters) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// "At" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "at" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"at\": ")
	if tmp, err := json.Marshal(r.At); err != nil {
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

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (r *AttachParameters) UnmarshalJSON(b []byte) error {
	atReceived := false
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "at":
			if err := json.Unmarshal([]byte(v), &r.At); err != nil {
				return err
			}
			atReceived = true
		case "name":
			if err := json.Unmarshal([]byte(v), &r.Name); err != nil {
				return err
			}
		}
	}
	// check if at (a required property) was received
	if !atReceived {
		return errors.New("\"at\" is required but was not present")
	}
	return nil
}

// Branches a map defining rules for execution on specific branches.
type Branches struct {
	// Either a single branch specifier, or a list of branch specifiers
	Ignore []string `json:"ignore,omitempty"`

	// Either a single branch specifier, or a list of branch specifiers
	Only []string `json:"only,omitempty"`
}

// Checkout A special step used to check out source code to the configured path.
// (defaults to the working_directory).
type Checkout struct {
	Name       string             `json:"name"`
	Parameters *CheckoutParameter `json:"parameters,omitempty"`
}

func (r *Checkout) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// "Name" field is required
	// only required object types supported for marshal checking (for now)
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

func (r *Checkout) UnmarshalJSON(b []byte) error {
	nameReceived := false
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "name":
			if err := json.Unmarshal([]byte(v), &r.Name); err != nil {
				return err
			}
			nameReceived = true
		case "parameters":
			if err := json.Unmarshal([]byte(v), &r.Parameters); err != nil {
				return err
			}
		}
	}
	// check if name (a required property) was received
	if !nameReceived {
		return errors.New("\"name\" is required but was not present")
	}
	return nil
}

// CheckoutParameter command parameters for the checkout command.
type CheckoutParameter struct {
	// Title of the step to be shown in the CircleCI UI (default: full command)
	Name string `json:"name,omitempty"`

	// Checkout directory.
	// Will be interpreted relative to the working_directory of the job.
	Path string `json:"path,omitempty"`
}

// CircleCIConfigObject CircleCI configuration object
type CircleCIConfigObject struct {
	Commands  []interface{} `json:"commands,omitempty"`
	Jobs      []*Job        `json:"jobs,omitempty"`
	Version   float64       `json:"version"`
	Workflows []*Workflow   `json:"workflows,omitempty"`
}

func (r *CircleCIConfigObject) MarshalJSON() ([]byte, error) {
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

func (r *CircleCIConfigObject) UnmarshalJSON(b []byte) error {
	versionReceived := false
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
		case "version":
			if err := json.Unmarshal([]byte(v), &r.Version); err != nil {
				return err
			}
			versionReceived = true
		case "workflows":
			if err := json.Unmarshal([]byte(v), &r.Workflows); err != nil {
				return err
			}
		}
	}
	// check if version (a required property) was received
	if !versionReceived {
		return errors.New("\"version\" is required but was not present")
	}
	return nil
}

// CommandParameters parameter definitions for the command.
type CommandParameters struct {
	AdditionalProperties map[string]interface{} `json:"-,omitempty"`

	// Title of the step to be shown in the CircleCI UI (default: full command)
	Name string `json:"name,omitempty"`
}

func (r *CommandParameters) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
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

func (r *CommandParameters) UnmarshalJSON(b []byte) error {
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "name":
			if err := json.Unmarshal([]byte(v), &r.Name); err != nil {
				return err
			}
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

// ConfigOrbImport orb import object.
type ConfigOrbImport struct {
	OrbAlias  string `json:"orbAlias"`
	OrbImport string `json:"orbImport"`
}

func (r *ConfigOrbImport) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// "OrbAlias" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "orbAlias" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"orbAlias\": ")
	if tmp, err := json.Marshal(r.OrbAlias); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "OrbImport" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "orbImport" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"orbImport\": ")
	if tmp, err := json.Marshal(r.OrbImport); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (r *ConfigOrbImport) UnmarshalJSON(b []byte) error {
	orbAliasReceived := false
	orbImportReceived := false
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "orbAlias":
			if err := json.Unmarshal([]byte(v), &r.OrbAlias); err != nil {
				return err
			}
			orbAliasReceived = true
		case "orbImport":
			if err := json.Unmarshal([]byte(v), &r.OrbImport); err != nil {
				return err
			}
			orbImportReceived = true
		}
	}
	// check if orbAlias (a required property) was received
	if !orbAliasReceived {
		return errors.New("\"orbAlias\" is required but was not present")
	}
	// check if orbImport (a required property) was received
	if !orbImportReceived {
		return errors.New("\"orbImport\" is required but was not present")
	}
	return nil
}

// DockerAuth authentication for registries using standard `docker login` credentials.
type DockerAuth struct {
	// Specify an environment variable (e.g. $DOCKER_PASSWORD)
	Password string `json:"password"`

	Username string `json:"username"`
}

func (r *DockerAuth) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// "Password" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "password" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"password\": ")
	if tmp, err := json.Marshal(r.Password); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "Username" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "username" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"username\": ")
	if tmp, err := json.Marshal(r.Username); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (r *DockerAuth) UnmarshalJSON(b []byte) error {
	passwordReceived := false
	usernameReceived := false
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "password":
			if err := json.Unmarshal([]byte(v), &r.Password); err != nil {
				return err
			}
			passwordReceived = true
		case "username":
			if err := json.Unmarshal([]byte(v), &r.Username); err != nil {
				return err
			}
			usernameReceived = true
		}
	}
	// check if password (a required property) was received
	if !passwordReceived {
		return errors.New("\"password\" is required but was not present")
	}
	// check if username (a required property) was received
	if !usernameReceived {
		return errors.New("\"username\" is required but was not present")
	}
	return nil
}

// DockerAuthAWS authentication for AWS Elastic Container Registry (ECR).
type DockerAuthAWS struct {
	AwsAccessKeyId string `json:"aws_access_key_id"`

	// Specify an environment variable (e.g. $ECR_AWS_SECRET_ACCESS_KEY)
	AwsSecretAccessKey string `json:"aws_secret_access_key"`
}

func (r *DockerAuthAWS) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// "AwsAccessKeyId" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "aws_access_key_id" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"aws_access_key_id\": ")
	if tmp, err := json.Marshal(r.AwsAccessKeyId); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "AwsSecretAccessKey" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "aws_secret_access_key" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"aws_secret_access_key\": ")
	if tmp, err := json.Marshal(r.AwsSecretAccessKey); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (r *DockerAuthAWS) UnmarshalJSON(b []byte) error {
	aws_access_key_idReceived := false
	aws_secret_access_keyReceived := false
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "aws_access_key_id":
			if err := json.Unmarshal([]byte(v), &r.AwsAccessKeyId); err != nil {
				return err
			}
			aws_access_key_idReceived = true
		case "aws_secret_access_key":
			if err := json.Unmarshal([]byte(v), &r.AwsSecretAccessKey); err != nil {
				return err
			}
			aws_secret_access_keyReceived = true
		}
	}
	// check if aws_access_key_id (a required property) was received
	if !aws_access_key_idReceived {
		return errors.New("\"aws_access_key_id\" is required but was not present")
	}
	// check if aws_secret_access_key (a required property) was received
	if !aws_secret_access_keyReceived {
		return errors.New("\"aws_secret_access_key\" is required but was not present")
	}
	return nil
}

// DockerExecutor a docker based CircleCI executor.
type DockerExecutor struct {
	Description string `json:"description,omitempty"`

	// The name of a custom Docker image to use.
	Image *DockerImage `json:"image"`

	// Instantiate a Docker executor.
	ResourceClass string `json:"resourceClass"`

	// Add additional Docker images which will be accessible from the primary container.
	// This is typically used for adding a database as a service container.
	ServiceImages []*DockerImage `json:"serviceImages"`
}

func (r *DockerExecutor) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// Marshal the "description" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"description\": ")
	if tmp, err := json.Marshal(r.Description); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "Image" field is required
	if r.Image == nil {
		return nil, errors.New("image is a required field")
	}
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
	// "ResourceClass" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "resourceClass" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"resourceClass\": ")
	if tmp, err := json.Marshal(r.ResourceClass); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "ServiceImages" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "serviceImages" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"serviceImages\": ")
	if tmp, err := json.Marshal(r.ServiceImages); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (r *DockerExecutor) UnmarshalJSON(b []byte) error {
	imageReceived := false
	resourceClassReceived := false
	serviceImagesReceived := false
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "description":
			if err := json.Unmarshal([]byte(v), &r.Description); err != nil {
				return err
			}
		case "image":
			if err := json.Unmarshal([]byte(v), &r.Image); err != nil {
				return err
			}
			imageReceived = true
		case "resourceClass":
			if err := json.Unmarshal([]byte(v), &r.ResourceClass); err != nil {
				return err
			}
			resourceClassReceived = true
		case "serviceImages":
			if err := json.Unmarshal([]byte(v), &r.ServiceImages); err != nil {
				return err
			}
			serviceImagesReceived = true
		}
	}
	// check if image (a required property) was received
	if !imageReceived {
		return errors.New("\"image\" is required but was not present")
	}
	// check if resourceClass (a required property) was received
	if !resourceClassReceived {
		return errors.New("\"resourceClass\" is required but was not present")
	}
	// check if serviceImages (a required property) was received
	if !serviceImagesReceived {
		return errors.New("\"serviceImages\" is required but was not present")
	}
	return nil
}

// DockerImage
type DockerImage struct {
	Auth        *DockerAuth       `json:"auth,omitempty"`
	AwsAuth     *DockerAuthAWS    `json:"aws_auth,omitempty"`
	Command     []string          `json:"command,omitempty"`
	Entrypoint  []string          `json:"entrypoint,omitempty"`
	Environment map[string]string `json:"environment,omitempty"`
	Image       string            `json:"image"`
	Name        string            `json:"name,omitempty"`
	User        string            `json:"user,omitempty"`
}

func (r *DockerImage) MarshalJSON() ([]byte, error) {
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

func (r *DockerImage) UnmarshalJSON(b []byte) error {
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

// DockerImageMap
type DockerImageMap struct {
	Image string `json:"image"`
}

func (r *DockerImageMap) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
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

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (r *DockerImageMap) UnmarshalJSON(b []byte) error {
	imageReceived := false
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "image":
			if err := json.Unmarshal([]byte(v), &r.Image); err != nil {
				return err
			}
			imageReceived = true
		}
	}
	// check if image (a required property) was received
	if !imageReceived {
		return errors.New("\"image\" is required but was not present")
	}
	return nil
}

// Git
type Git struct {
	// The long (40-character) git SHA of the build prior to the one being built.
	BaseRevision string `json:"base_revision"`

	// The name of the git branch that was pushed to trigger the pipeline.
	Branch  string `json:"branch"`
	IsLocal bool   `json:"_isLocal"`

	// The long (40-character) git SHA that is being built.
	Revision string `json:"revision"`

	// The name of the git tag that was pushed to trigger the pipeline. If the pipeline was not triggered by a tag, then this is the empty string.
	Tag string `json:"tag"`
}

func (r *Git) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// "BaseRevision" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "base_revision" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"base_revision\": ")
	if tmp, err := json.Marshal(r.BaseRevision); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "Branch" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "branch" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"branch\": ")
	if tmp, err := json.Marshal(r.Branch); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "IsLocal" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "_isLocal" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"_isLocal\": ")
	if tmp, err := json.Marshal(r.IsLocal); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "Revision" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "revision" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"revision\": ")
	if tmp, err := json.Marshal(r.Revision); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "Tag" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "tag" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"tag\": ")
	if tmp, err := json.Marshal(r.Tag); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (r *Git) UnmarshalJSON(b []byte) error {
	base_revisionReceived := false
	branchReceived := false
	_isLocalReceived := false
	revisionReceived := false
	tagReceived := false
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "base_revision":
			if err := json.Unmarshal([]byte(v), &r.BaseRevision); err != nil {
				return err
			}
			base_revisionReceived = true
		case "branch":
			if err := json.Unmarshal([]byte(v), &r.Branch); err != nil {
				return err
			}
			branchReceived = true
		case "_isLocal":
			if err := json.Unmarshal([]byte(v), &r.IsLocal); err != nil {
				return err
			}
			_isLocalReceived = true
		case "revision":
			if err := json.Unmarshal([]byte(v), &r.Revision); err != nil {
				return err
			}
			revisionReceived = true
		case "tag":
			if err := json.Unmarshal([]byte(v), &r.Tag); err != nil {
				return err
			}
			tagReceived = true
		}
	}
	// check if base_revision (a required property) was received
	if !base_revisionReceived {
		return errors.New("\"base_revision\" is required but was not present")
	}
	// check if branch (a required property) was received
	if !branchReceived {
		return errors.New("\"branch\" is required but was not present")
	}
	// check if _isLocal (a required property) was received
	if !_isLocalReceived {
		return errors.New("\"_isLocal\" is required but was not present")
	}
	// check if revision (a required property) was received
	if !revisionReceived {
		return errors.New("\"revision\" is required but was not present")
	}
	// check if tag (a required property) was received
	if !tagReceived {
		return errors.New("\"tag\" is required but was not present")
	}
	return nil
}

// Job jobs define a collection of steps to be run within a given executor, and are orchestrated using workflows.
type Job struct {
	// the reusable executor to use for this job. The Executor must have already been instantiated and added to the config.
	Executor interface{} `json:"executor"`

	// the name of the current Job.
	Name string `json:"name"`

	// a list of Commands to execute within the job in the order which they were added.
	Steps []interface{} `json:"steps"`
}

func (r *Job) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// "Executor" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "executor" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"executor\": ")
	if tmp, err := json.Marshal(r.Executor); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "Name" field is required
	// only required object types supported for marshal checking (for now)
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

func (r *Job) UnmarshalJSON(b []byte) error {
	executorReceived := false
	nameReceived := false
	stepsReceived := false
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "executor":
			if err := json.Unmarshal([]byte(v), &r.Executor); err != nil {
				return err
			}
			executorReceived = true
		case "name":
			if err := json.Unmarshal([]byte(v), &r.Name); err != nil {
				return err
			}
			nameReceived = true
		case "steps":
			if err := json.Unmarshal([]byte(v), &r.Steps); err != nil {
				return err
			}
			stepsReceived = true
		}
	}
	// check if executor (a required property) was received
	if !executorReceived {
		return errors.New("\"executor\" is required but was not present")
	}
	// check if name (a required property) was received
	if !nameReceived {
		return errors.New("\"name\" is required but was not present")
	}
	// check if steps (a required property) was received
	if !stepsReceived {
		return errors.New("\"steps\" is required but was not present")
	}
	return nil
}

// MacOSExecutor a macOS virtual machine with configurable Xcode version.
type MacOSExecutor struct {
	Description string `json:"description,omitempty"`

	ResourceClass string `json:"resourceClass"`

	// Select an xcode version
	Xcode string `json:"xcode"`
}

func (r *MacOSExecutor) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// Marshal the "description" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"description\": ")
	if tmp, err := json.Marshal(r.Description); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "ResourceClass" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "resourceClass" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"resourceClass\": ")
	if tmp, err := json.Marshal(r.ResourceClass); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "Xcode" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "xcode" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"xcode\": ")
	if tmp, err := json.Marshal(r.Xcode); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (r *MacOSExecutor) UnmarshalJSON(b []byte) error {
	resourceClassReceived := false
	xcodeReceived := false
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "description":
			if err := json.Unmarshal([]byte(v), &r.Description); err != nil {
				return err
			}
		case "resourceClass":
			if err := json.Unmarshal([]byte(v), &r.ResourceClass); err != nil {
				return err
			}
			resourceClassReceived = true
		case "xcode":
			if err := json.Unmarshal([]byte(v), &r.Xcode); err != nil {
				return err
			}
			xcodeReceived = true
		}
	}
	// check if resourceClass (a required property) was received
	if !resourceClassReceived {
		return errors.New("\"resourceClass\" is required but was not present")
	}
	// check if xcode (a required property) was received
	if !xcodeReceived {
		return errors.New("\"xcode\" is required but was not present")
	}
	return nil
}

// Machine
type Machine struct {
	Image string `json:"image"`
}

func (r *Machine) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
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

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (r *Machine) UnmarshalJSON(b []byte) error {
	imageReceived := false
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "image":
			if err := json.Unmarshal([]byte(v), &r.Image); err != nil {
				return err
			}
			imageReceived = true
		}
	}
	// check if image (a required property) was received
	if !imageReceived {
		return errors.New("\"image\" is required but was not present")
	}
	return nil
}

// MachineExecutor the linux virtual machine executor.
type MachineExecutor struct {
	Description string `json:"description,omitempty"`

	// Select one of the Ubuntu Linux VM Images provided by CircleCI.
	Image         string `json:"image"`
	ResourceClass string `json:"resourceClass"`
}

func (r *MachineExecutor) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// Marshal the "description" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"description\": ")
	if tmp, err := json.Marshal(r.Description); err != nil {
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
	// "ResourceClass" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "resourceClass" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"resourceClass\": ")
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

func (r *MachineExecutor) UnmarshalJSON(b []byte) error {
	imageReceived := false
	resourceClassReceived := false
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "description":
			if err := json.Unmarshal([]byte(v), &r.Description); err != nil {
				return err
			}
		case "image":
			if err := json.Unmarshal([]byte(v), &r.Image); err != nil {
				return err
			}
			imageReceived = true
		case "resourceClass":
			if err := json.Unmarshal([]byte(v), &r.ResourceClass); err != nil {
				return err
			}
			resourceClassReceived = true
		}
	}
	// check if image (a required property) was received
	if !imageReceived {
		return errors.New("\"image\" is required but was not present")
	}
	// check if resourceClass (a required property) was received
	if !resourceClassReceived {
		return errors.New("\"resourceClass\" is required but was not present")
	}
	return nil
}

// Macos
type Macos struct {
	Xcode string `json:"xcode"`
}

func (r *Macos) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// "Xcode" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "xcode" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"xcode\": ")
	if tmp, err := json.Marshal(r.Xcode); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (r *Macos) UnmarshalJSON(b []byte) error {
	xcodeReceived := false
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "xcode":
			if err := json.Unmarshal([]byte(v), &r.Xcode); err != nil {
				return err
			}
			xcodeReceived = true
		}
	}
	// check if xcode (a required property) was received
	if !xcodeReceived {
		return errors.New("\"xcode\" is required but was not present")
	}
	return nil
}

// Persist special step used to persist the workflow’s workspace to the current container.
//
// The full contents of the workspace are downloaded and copied into the directory the workspace is being Persisted at.
type Persist struct {
	Name       string             `json:"name"`
	Parameters *PersistParameters `json:"parameters"`
}

func (r *Persist) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// "Name" field is required
	// only required object types supported for marshal checking (for now)
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
	// "Parameters" field is required
	if r.Parameters == nil {
		return nil, errors.New("parameters is a required field")
	}
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

func (r *Persist) UnmarshalJSON(b []byte) error {
	nameReceived := false
	parametersReceived := false
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "name":
			if err := json.Unmarshal([]byte(v), &r.Name); err != nil {
				return err
			}
			nameReceived = true
		case "parameters":
			if err := json.Unmarshal([]byte(v), &r.Parameters); err != nil {
				return err
			}
			parametersReceived = true
		}
	}
	// check if name (a required property) was received
	if !nameReceived {
		return errors.New("\"name\" is required but was not present")
	}
	// check if parameters (a required property) was received
	if !parametersReceived {
		return errors.New("\"parameters\" is required but was not present")
	}
	return nil
}

// PersistParameters command parameters for the persist command.
type PersistParameters struct {
	// Title of the step to be shown in the CircleCI UI (default: full command)
	Name string `json:"name,omitempty"`

	// Glob identifying file(s), or a non-glob path to a directory to add to the shared workspace. Interpreted as relative to the workspace root. Must not be the workspace root itself.
	Paths []string `json:"paths"`

	// Either an absolute path or a path relative to `working_directory`
	Root string `json:"root"`
}

func (r *PersistParameters) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
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
	// "Paths" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "paths" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"paths\": ")
	if tmp, err := json.Marshal(r.Paths); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "Root" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "root" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"root\": ")
	if tmp, err := json.Marshal(r.Root); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (r *PersistParameters) UnmarshalJSON(b []byte) error {
	pathsReceived := false
	rootReceived := false
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "name":
			if err := json.Unmarshal([]byte(v), &r.Name); err != nil {
				return err
			}
		case "paths":
			if err := json.Unmarshal([]byte(v), &r.Paths); err != nil {
				return err
			}
			pathsReceived = true
		case "root":
			if err := json.Unmarshal([]byte(v), &r.Root); err != nil {
				return err
			}
			rootReceived = true
		}
	}
	// check if paths (a required property) was received
	if !pathsReceived {
		return errors.New("\"paths\" is required but was not present")
	}
	// check if root (a required property) was received
	if !rootReceived {
		return errors.New("\"root\" is required but was not present")
	}
	return nil
}

// Pipeline access pipeline variables from within CircleCI Cloud.
type Pipeline struct {
	// A globally unique id representing for the pipeline
	Id string `json:"id"`

	// Pipeline parameter values are passed at the config level on CircleCI. These values will not be present on a local system.
	IsLocal bool `json:"_isLocal"`

	// A project unique integer id for the pipeline
	Number int `json:"number"`

	// Array of user defined parameters
	Parameters []*PipelineParameter `json:"parameters"`
}

func (r *Pipeline) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// "Id" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "id" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"id\": ")
	if tmp, err := json.Marshal(r.Id); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "IsLocal" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "_isLocal" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"_isLocal\": ")
	if tmp, err := json.Marshal(r.IsLocal); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "Number" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "number" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"number\": ")
	if tmp, err := json.Marshal(r.Number); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
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

func (r *Pipeline) UnmarshalJSON(b []byte) error {
	idReceived := false
	_isLocalReceived := false
	numberReceived := false
	parametersReceived := false
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "id":
			if err := json.Unmarshal([]byte(v), &r.Id); err != nil {
				return err
			}
			idReceived = true
		case "_isLocal":
			if err := json.Unmarshal([]byte(v), &r.IsLocal); err != nil {
				return err
			}
			_isLocalReceived = true
		case "number":
			if err := json.Unmarshal([]byte(v), &r.Number); err != nil {
				return err
			}
			numberReceived = true
		case "parameters":
			if err := json.Unmarshal([]byte(v), &r.Parameters); err != nil {
				return err
			}
			parametersReceived = true
		}
	}
	// check if id (a required property) was received
	if !idReceived {
		return errors.New("\"id\" is required but was not present")
	}
	// check if _isLocal (a required property) was received
	if !_isLocalReceived {
		return errors.New("\"_isLocal\" is required but was not present")
	}
	// check if number (a required property) was received
	if !numberReceived {
		return errors.New("\"number\" is required but was not present")
	}
	// check if parameters (a required property) was received
	if !parametersReceived {
		return errors.New("\"parameters\" is required but was not present")
	}
	return nil
}

// PipelineParameter a pipeline parameter.
type PipelineParameter struct {
	DefaultValue  interface{}        `json:"defaultValue"`
	EnumValues    []string           `json:"enumValues"`
	Name          string             `json:"name"`
	ParameterType string             `json:"parameterType"`
	Value         *PipelineParameter `json:"value"`
}

func (r *PipelineParameter) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// "DefaultValue" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "defaultValue" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"defaultValue\": ")
	if tmp, err := json.Marshal(r.DefaultValue); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "EnumValues" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "enumValues" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"enumValues\": ")
	if tmp, err := json.Marshal(r.EnumValues); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "Name" field is required
	// only required object types supported for marshal checking (for now)
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
	// "Value" field is required
	if r.Value == nil {
		return nil, errors.New("value is a required field")
	}
	// Marshal the "value" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"value\": ")
	if tmp, err := json.Marshal(r.Value); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (r *PipelineParameter) UnmarshalJSON(b []byte) error {
	defaultValueReceived := false
	enumValuesReceived := false
	nameReceived := false
	parameterTypeReceived := false
	valueReceived := false
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "defaultValue":
			if err := json.Unmarshal([]byte(v), &r.DefaultValue); err != nil {
				return err
			}
			defaultValueReceived = true
		case "enumValues":
			if err := json.Unmarshal([]byte(v), &r.EnumValues); err != nil {
				return err
			}
			enumValuesReceived = true
		case "name":
			if err := json.Unmarshal([]byte(v), &r.Name); err != nil {
				return err
			}
			nameReceived = true
		case "parameterType":
			if err := json.Unmarshal([]byte(v), &r.ParameterType); err != nil {
				return err
			}
			parameterTypeReceived = true
		case "value":
			if err := json.Unmarshal([]byte(v), &r.Value); err != nil {
				return err
			}
			valueReceived = true
		}
	}
	// check if defaultValue (a required property) was received
	if !defaultValueReceived {
		return errors.New("\"defaultValue\" is required but was not present")
	}
	// check if enumValues (a required property) was received
	if !enumValuesReceived {
		return errors.New("\"enumValues\" is required but was not present")
	}
	// check if name (a required property) was received
	if !nameReceived {
		return errors.New("\"name\" is required but was not present")
	}
	// check if parameterType (a required property) was received
	if !parameterTypeReceived {
		return errors.New("\"parameterType\" is required but was not present")
	}
	// check if value (a required property) was received
	if !valueReceived {
		return errors.New("\"value\" is required but was not present")
	}
	return nil
}

// Project pipeline project level information.
type Project struct {
	// The URL where the current project is hosted. E.g. https://github.com/circleci/circleci-docs
	GitUrl  string `json:"git_url"`
	IsLocal bool   `json:"_isLocal"`

	// The lower-case name of the VCS provider, E.g. “github”, “bitbucket”
	Vcs string `json:"vcs"`
}

func (r *Project) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// "GitUrl" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "git_url" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"git_url\": ")
	if tmp, err := json.Marshal(r.GitUrl); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "IsLocal" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "_isLocal" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"_isLocal\": ")
	if tmp, err := json.Marshal(r.IsLocal); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "Vcs" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "vcs" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"vcs\": ")
	if tmp, err := json.Marshal(r.Vcs); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (r *Project) UnmarshalJSON(b []byte) error {
	git_urlReceived := false
	_isLocalReceived := false
	vcsReceived := false
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "git_url":
			if err := json.Unmarshal([]byte(v), &r.GitUrl); err != nil {
				return err
			}
			git_urlReceived = true
		case "_isLocal":
			if err := json.Unmarshal([]byte(v), &r.IsLocal); err != nil {
				return err
			}
			_isLocalReceived = true
		case "vcs":
			if err := json.Unmarshal([]byte(v), &r.Vcs); err != nil {
				return err
			}
			vcsReceived = true
		}
	}
	// check if git_url (a required property) was received
	if !git_urlReceived {
		return errors.New("\"git_url\" is required but was not present")
	}
	// check if _isLocal (a required property) was received
	if !_isLocalReceived {
		return errors.New("\"_isLocal\" is required but was not present")
	}
	// check if vcs (a required property) was received
	if !vcsReceived {
		return errors.New("\"vcs\" is required but was not present")
	}
	return nil
}

// Restore restores a previously saved cache based on a key..cache needs to have been saved first for this key using save_cache step.
//
// Learn more in the caching documentation.
type Restore struct {
	Name       string                  `json:"name"`
	Parameters *RestoreCacheParameters `json:"parameters"`
}

func (r *Restore) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// "Name" field is required
	// only required object types supported for marshal checking (for now)
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
	// "Parameters" field is required
	if r.Parameters == nil {
		return nil, errors.New("parameters is a required field")
	}
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

func (r *Restore) UnmarshalJSON(b []byte) error {
	nameReceived := false
	parametersReceived := false
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "name":
			if err := json.Unmarshal([]byte(v), &r.Name); err != nil {
				return err
			}
			nameReceived = true
		case "parameters":
			if err := json.Unmarshal([]byte(v), &r.Parameters); err != nil {
				return err
			}
			parametersReceived = true
		}
	}
	// check if name (a required property) was received
	if !nameReceived {
		return errors.New("\"name\" is required but was not present")
	}
	// check if parameters (a required property) was received
	if !parametersReceived {
		return errors.New("\"parameters\" is required but was not present")
	}
	return nil
}

// RestoreCacheParameters command parameters for the restorecache command.
type RestoreCacheParameters struct {
	// List of cache keys to lookup for a cache to restore. Only first existing key will be restored.
	Keys []string `json:"keys"`

	// Title of the step to be shown in the CircleCI UI (default: full command)
	Name string `json:"name,omitempty"`
}

func (r *RestoreCacheParameters) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// "Keys" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "keys" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"keys\": ")
	if tmp, err := json.Marshal(r.Keys); err != nil {
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

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (r *RestoreCacheParameters) UnmarshalJSON(b []byte) error {
	keysReceived := false
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "keys":
			if err := json.Unmarshal([]byte(v), &r.Keys); err != nil {
				return err
			}
			keysReceived = true
		case "name":
			if err := json.Unmarshal([]byte(v), &r.Name); err != nil {
				return err
			}
		}
	}
	// check if keys (a required property) was received
	if !keysReceived {
		return errors.New("\"keys\" is required but was not present")
	}
	return nil
}

// Run the run command step is used for invoking all command-line programs.
type Run struct {
	Name       string         `json:"name"`
	Parameters *RunParameters `json:"parameters"`
}

func (r *Run) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// "Name" field is required
	// only required object types supported for marshal checking (for now)
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
	// "Parameters" field is required
	if r.Parameters == nil {
		return nil, errors.New("parameters is a required field")
	}
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

func (r *Run) UnmarshalJSON(b []byte) error {
	nameReceived := false
	parametersReceived := false
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "name":
			if err := json.Unmarshal([]byte(v), &r.Name); err != nil {
				return err
			}
			nameReceived = true
		case "parameters":
			if err := json.Unmarshal([]byte(v), &r.Parameters); err != nil {
				return err
			}
			parametersReceived = true
		}
	}
	// check if name (a required property) was received
	if !nameReceived {
		return errors.New("\"name\" is required but was not present")
	}
	// check if parameters (a required property) was received
	if !parametersReceived {
		return errors.New("\"parameters\" is required but was not present")
	}
	return nil
}

// RunParameters command parameters for the run command.
type RunParameters struct {
	// Whether or not this step should run in the background (default: false)
	Background bool `json:"background,omitempty"`

	// Command to run via the shell
	Command string `json:"command"`

	// Additional environmental variables, locally scoped to command
	Environment map[string]string `json:"environment,omitempty"`

	// Title of the step to be shown in the CircleCI UI (default: full command)
	Name string `json:"name,omitempty"`

	// Elapsed time the command can run without output. The string is a decimal with unit suffix, such as “20m”, “1.25h”, “5s” (default: 10 minutes)
	NoOutputTimeout string `json:"no_output_timeout,omitempty"`

	// Shell to use for execution command (default: See Default Shell Options)
	Shell string `json:"shell,omitempty"`

	// Specify when to enable or disable the step. (default: on_success)
	When string `json:"when,omitempty"`

	// In which directory to run this step. Will be interpreted relative to the working_directory of the job). (default: .)
	WorkingDirectory string `json:"working_directory,omitempty"`
}

func (r *RunParameters) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// Marshal the "background" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"background\": ")
	if tmp, err := json.Marshal(r.Background); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "Command" field is required
	// only required object types supported for marshal checking (for now)
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
	// Marshal the "no_output_timeout" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"no_output_timeout\": ")
	if tmp, err := json.Marshal(r.NoOutputTimeout); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
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
	// Marshal the "when" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"when\": ")
	if tmp, err := json.Marshal(r.When); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "working_directory" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"working_directory\": ")
	if tmp, err := json.Marshal(r.WorkingDirectory); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (r *RunParameters) UnmarshalJSON(b []byte) error {
	commandReceived := false
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "background":
			if err := json.Unmarshal([]byte(v), &r.Background); err != nil {
				return err
			}
		case "command":
			if err := json.Unmarshal([]byte(v), &r.Command); err != nil {
				return err
			}
			commandReceived = true
		case "environment":
			if err := json.Unmarshal([]byte(v), &r.Environment); err != nil {
				return err
			}
		case "name":
			if err := json.Unmarshal([]byte(v), &r.Name); err != nil {
				return err
			}
		case "no_output_timeout":
			if err := json.Unmarshal([]byte(v), &r.NoOutputTimeout); err != nil {
				return err
			}
		case "shell":
			if err := json.Unmarshal([]byte(v), &r.Shell); err != nil {
				return err
			}
		case "when":
			if err := json.Unmarshal([]byte(v), &r.When); err != nil {
				return err
			}
		case "working_directory":
			if err := json.Unmarshal([]byte(v), &r.WorkingDirectory); err != nil {
				return err
			}
		}
	}
	// check if command (a required property) was received
	if !commandReceived {
		return errors.New("\"command\" is required but was not present")
	}
	return nil
}

// Save generates and stores a cache of a file or directory of files such as dependencies or source code in our object storage.
//
// Later jobs can restore this cache.
type Save struct {
	Name       string               `json:"name"`
	Parameters *SaveCacheParameters `json:"parameters"`
}

func (r *Save) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// "Name" field is required
	// only required object types supported for marshal checking (for now)
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
	// "Parameters" field is required
	if r.Parameters == nil {
		return nil, errors.New("parameters is a required field")
	}
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

func (r *Save) UnmarshalJSON(b []byte) error {
	nameReceived := false
	parametersReceived := false
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "name":
			if err := json.Unmarshal([]byte(v), &r.Name); err != nil {
				return err
			}
			nameReceived = true
		case "parameters":
			if err := json.Unmarshal([]byte(v), &r.Parameters); err != nil {
				return err
			}
			parametersReceived = true
		}
	}
	// check if name (a required property) was received
	if !nameReceived {
		return errors.New("\"name\" is required but was not present")
	}
	// check if parameters (a required property) was received
	if !parametersReceived {
		return errors.New("\"parameters\" is required but was not present")
	}
	return nil
}

// SaveCacheParameters command parameters for the savecache command.
type SaveCacheParameters struct {
	// Unique identifier for this cache
	Key string `json:"key"`

	// Title of the step to be shown in the CircleCI UI (default: full command)
	Name string `json:"name,omitempty"`

	// List of directories which should be added to the cache
	Paths []string `json:"paths"`

	// Specify when to enable or disable the step.
	When string `json:"when,omitempty"`
}

func (r *SaveCacheParameters) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// "Key" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "key" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"key\": ")
	if tmp, err := json.Marshal(r.Key); err != nil {
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
	// "Paths" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "paths" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"paths\": ")
	if tmp, err := json.Marshal(r.Paths); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "when" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"when\": ")
	if tmp, err := json.Marshal(r.When); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (r *SaveCacheParameters) UnmarshalJSON(b []byte) error {
	keyReceived := false
	pathsReceived := false
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "key":
			if err := json.Unmarshal([]byte(v), &r.Key); err != nil {
				return err
			}
			keyReceived = true
		case "name":
			if err := json.Unmarshal([]byte(v), &r.Name); err != nil {
				return err
			}
		case "paths":
			if err := json.Unmarshal([]byte(v), &r.Paths); err != nil {
				return err
			}
			pathsReceived = true
		case "when":
			if err := json.Unmarshal([]byte(v), &r.When); err != nil {
				return err
			}
		}
	}
	// check if key (a required property) was received
	if !keyReceived {
		return errors.New("\"key\" is required but was not present")
	}
	// check if paths (a required property) was received
	if !pathsReceived {
		return errors.New("\"paths\" is required but was not present")
	}
	return nil
}

// SetupRemoteDocker creates a remote docker environment configured to execute docker commands.
type SetupRemoteDocker struct {
	Name       string                       `json:"name"`
	Parameters *SetupRemoteDockerParameters `json:"parameters"`
}

func (r *SetupRemoteDocker) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// "Name" field is required
	// only required object types supported for marshal checking (for now)
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
	// "Parameters" field is required
	if r.Parameters == nil {
		return nil, errors.New("parameters is a required field")
	}
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

func (r *SetupRemoteDocker) UnmarshalJSON(b []byte) error {
	nameReceived := false
	parametersReceived := false
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "name":
			if err := json.Unmarshal([]byte(v), &r.Name); err != nil {
				return err
			}
			nameReceived = true
		case "parameters":
			if err := json.Unmarshal([]byte(v), &r.Parameters); err != nil {
				return err
			}
			parametersReceived = true
		}
	}
	// check if name (a required property) was received
	if !nameReceived {
		return errors.New("\"name\" is required but was not present")
	}
	// check if parameters (a required property) was received
	if !parametersReceived {
		return errors.New("\"parameters\" is required but was not present")
	}
	return nil
}

// SetupRemoteDockerParameters command parameters for the setupremotedocker command.
type SetupRemoteDockerParameters struct {
	// Title of the step to be shown in the CircleCI UI (default: full command)
	Name string `json:"name,omitempty"`

	// SetupRemoteDocker directory.
	// Will be interpreted relative to the working_directory of the job.
	Version string `json:"version"`
}

func (r *SetupRemoteDockerParameters) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
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

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (r *SetupRemoteDockerParameters) UnmarshalJSON(b []byte) error {
	versionReceived := false
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "name":
			if err := json.Unmarshal([]byte(v), &r.Name); err != nil {
				return err
			}
		case "version":
			if err := json.Unmarshal([]byte(v), &r.Version); err != nil {
				return err
			}
			versionReceived = true
		}
	}
	// check if version (a required property) was received
	if !versionReceived {
		return errors.New("\"version\" is required but was not present")
	}
	return nil
}

// StoreArtifacts a special step used to check out source code to the configured path (defaults to the working_directory).
type StoreArtifacts struct {
	Name       string                    `json:"name"`
	Parameters *StoreArtifactsParameters `json:"parameters"`
}

func (r *StoreArtifacts) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// "Name" field is required
	// only required object types supported for marshal checking (for now)
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
	// "Parameters" field is required
	if r.Parameters == nil {
		return nil, errors.New("parameters is a required field")
	}
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

func (r *StoreArtifacts) UnmarshalJSON(b []byte) error {
	nameReceived := false
	parametersReceived := false
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "name":
			if err := json.Unmarshal([]byte(v), &r.Name); err != nil {
				return err
			}
			nameReceived = true
		case "parameters":
			if err := json.Unmarshal([]byte(v), &r.Parameters); err != nil {
				return err
			}
			parametersReceived = true
		}
	}
	// check if name (a required property) was received
	if !nameReceived {
		return errors.New("\"name\" is required but was not present")
	}
	// check if parameters (a required property) was received
	if !parametersReceived {
		return errors.New("\"parameters\" is required but was not present")
	}
	return nil
}

// StoreArtifactsParameters command parameters for the storeartifacts command.
type StoreArtifactsParameters struct {
	// Prefix added to the artifact paths in the artifacts API (default: the directory of the file specified in path)
	Destination string `json:"destination,omitempty"`

	// Title of the step to be shown in the CircleCI UI (default: full command)
	Name string `json:"name,omitempty"`

	// Directory in the primary container to save as job artifacts
	Path string `json:"path"`
}

func (r *StoreArtifactsParameters) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// Marshal the "destination" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"destination\": ")
	if tmp, err := json.Marshal(r.Destination); err != nil {
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
	// "Path" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "path" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"path\": ")
	if tmp, err := json.Marshal(r.Path); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (r *StoreArtifactsParameters) UnmarshalJSON(b []byte) error {
	pathReceived := false
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "destination":
			if err := json.Unmarshal([]byte(v), &r.Destination); err != nil {
				return err
			}
		case "name":
			if err := json.Unmarshal([]byte(v), &r.Name); err != nil {
				return err
			}
		case "path":
			if err := json.Unmarshal([]byte(v), &r.Path); err != nil {
				return err
			}
			pathReceived = true
		}
	}
	// check if path (a required property) was received
	if !pathReceived {
		return errors.New("\"path\" is required but was not present")
	}
	return nil
}

// StoreTestResults special step used to upload and store test results for a build.
//
// Test results are visible on the CircleCI web application, under each build’s “Test Summary” section.
//
// Storing test results is useful for timing analysis of your test suites.
type StoreTestResults struct {
	Name       string                      `json:"name"`
	Parameters *StoreTestResultsParameters `json:"parameters"`
}

func (r *StoreTestResults) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// "Name" field is required
	// only required object types supported for marshal checking (for now)
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
	// "Parameters" field is required
	if r.Parameters == nil {
		return nil, errors.New("parameters is a required field")
	}
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

func (r *StoreTestResults) UnmarshalJSON(b []byte) error {
	nameReceived := false
	parametersReceived := false
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "name":
			if err := json.Unmarshal([]byte(v), &r.Name); err != nil {
				return err
			}
			nameReceived = true
		case "parameters":
			if err := json.Unmarshal([]byte(v), &r.Parameters); err != nil {
				return err
			}
			parametersReceived = true
		}
	}
	// check if name (a required property) was received
	if !nameReceived {
		return errors.New("\"name\" is required but was not present")
	}
	// check if parameters (a required property) was received
	if !parametersReceived {
		return errors.New("\"parameters\" is required but was not present")
	}
	return nil
}

// StoreTestResultsParameters command parameters for the storetestresults command.
type StoreTestResultsParameters struct {
	// Title of the step to be shown in the CircleCI UI (default: full command)
	Name string `json:"name,omitempty"`

	// Path (absolute, or relative to your working_directory) to directory containing subdirectories of JUnit XML or Cucumber JSON test metadata files
	Path string `json:"path"`
}

func (r *StoreTestResultsParameters) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
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
	// "Path" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "path" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"path\": ")
	if tmp, err := json.Marshal(r.Path); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (r *StoreTestResultsParameters) UnmarshalJSON(b []byte) error {
	pathReceived := false
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "name":
			if err := json.Unmarshal([]byte(v), &r.Name); err != nil {
				return err
			}
		case "path":
			if err := json.Unmarshal([]byte(v), &r.Path); err != nil {
				return err
			}
			pathReceived = true
		}
	}
	// check if path (a required property) was received
	if !pathReceived {
		return errors.New("\"path\" is required but was not present")
	}
	return nil
}

// Tags a map defining rules for execution on specific tags.
type Tags struct {
	// Either a single tag specifier, or a list of tag specifiers
	Ignore []string `json:"ignore,omitempty"`

	// Either a single tag specifier, or a list of tag specifiers
	Only []string `json:"only,omitempty"`
}

// WindowsExecutor a Windows virtual machine (CircleCI Cloud).
type WindowsExecutor struct {
	Description string `json:"description,omitempty"`

	// Select one of the available Windows VM Images provided by CircleCI
	Image         string `json:"image"`
	ResourceClass string `json:"resourceClass"`
}

func (r *WindowsExecutor) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// Marshal the "description" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"description\": ")
	if tmp, err := json.Marshal(r.Description); err != nil {
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
	// "ResourceClass" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "resourceClass" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"resourceClass\": ")
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

func (r *WindowsExecutor) UnmarshalJSON(b []byte) error {
	imageReceived := false
	resourceClassReceived := false
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "description":
			if err := json.Unmarshal([]byte(v), &r.Description); err != nil {
				return err
			}
		case "image":
			if err := json.Unmarshal([]byte(v), &r.Image); err != nil {
				return err
			}
			imageReceived = true
		case "resourceClass":
			if err := json.Unmarshal([]byte(v), &r.ResourceClass); err != nil {
				return err
			}
			resourceClassReceived = true
		}
	}
	// check if image (a required property) was received
	if !imageReceived {
		return errors.New("\"image\" is required but was not present")
	}
	// check if resourceClass (a required property) was received
	if !resourceClassReceived {
		return errors.New("\"resourceClass\" is required but was not present")
	}
	return nil
}

// Workflow a workflow is a set of rules for defining a collection of jobs and their run order.
type Workflow struct {
	// The jobs to execute when this Workflow is triggered.
	Jobs []*WorkflowJob `json:"jobs"`

	// The name of the Workflow.
	Name string `json:"name"`
}

func (r *Workflow) MarshalJSON() ([]byte, error) {
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
	// "Name" field is required
	// only required object types supported for marshal checking (for now)
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

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (r *Workflow) UnmarshalJSON(b []byte) error {
	jobsReceived := false
	nameReceived := false
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
		case "name":
			if err := json.Unmarshal([]byte(v), &r.Name); err != nil {
				return err
			}
			nameReceived = true
		}
	}
	// check if jobs (a required property) was received
	if !jobsReceived {
		return errors.New("\"jobs\" is required but was not present")
	}
	// check if name (a required property) was received
	if !nameReceived {
		return errors.New("\"name\" is required but was not present")
	}
	return nil
}

// WorkflowJob assign parameters and filters to a Job within a Workflow.
//
// Utility class for assigning parameters to a job.
//
// Should only be instantiated for specific use cases.
type WorkflowJob struct {
	Job        *Job                   `json:"job"`
	Parameters *WorkflowJobParameters `json:"parameters,omitempty"`
}

func (r *WorkflowJob) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// "Job" field is required
	if r.Job == nil {
		return nil, errors.New("job is a required field")
	}
	// Marshal the "job" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"job\": ")
	if tmp, err := json.Marshal(r.Job); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
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

func (r *WorkflowJob) UnmarshalJSON(b []byte) error {
	jobReceived := false
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "job":
			if err := json.Unmarshal([]byte(v), &r.Job); err != nil {
				return err
			}
			jobReceived = true
		case "parameters":
			if err := json.Unmarshal([]byte(v), &r.Parameters); err != nil {
				return err
			}
		}
	}
	// check if job (a required property) was received
	if !jobReceived {
		return errors.New("\"job\" is required but was not present")
	}
	return nil
}

// WorkflowJobParameters
type WorkflowJobParameters struct {
	AdditionalProperties map[string]interface{} `json:"-,omitempty"`

	Context []string `json:"context,omitempty"`

	// {@link https://circleci.com/docs/2.0/configuration-reference/#filters} Filter workflow job's execution by branch or git tag.
	Filters *WorkflowFilterSchema `json:"filters,omitempty"`

	// An "approval" type job is a special job which pauses the workflow. This "job" is not defined outside of the workflow, you may enter any potential name for the job name. As long as the parameter of "type" is present and equal to "approval" this job will act as a placeholder that awaits user input to continue.
	JobType string `json:"jobType,omitempty"`

	// {@link https://circleci.com/docs/2.0/configuration-reference/#matrix-requires-version-21} The matrix stanza allows you to run a parameterized job multiple times with different arguments.
	Matrix *WorkflowMatrixSchema `json:"matrix,omitempty"`

	Name string `json:"name,omitempty"`

	// A list of jobs that must succeed for the job to start. Note: When jobs in the current workflow that are listed as dependencies are not executed (due to a filter function for example), their requirement as a dependency for other jobs will be ignored by the requires option. However, if all dependencies of a job are filtered, then that job will not be executed either.
	Requires []string `json:"requires,omitempty"`
}

func (r *WorkflowJobParameters) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// Marshal the "context" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"context\": ")
	if tmp, err := json.Marshal(r.Context); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "filters" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"filters\": ")
	if tmp, err := json.Marshal(r.Filters); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "jobType" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"jobType\": ")
	if tmp, err := json.Marshal(r.JobType); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "matrix" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"matrix\": ")
	if tmp, err := json.Marshal(r.Matrix); err != nil {
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
	// Marshal the "requires" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"requires\": ")
	if tmp, err := json.Marshal(r.Requires); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
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

func (r *WorkflowJobParameters) UnmarshalJSON(b []byte) error {
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "context":
			if err := json.Unmarshal([]byte(v), &r.Context); err != nil {
				return err
			}
		case "filters":
			if err := json.Unmarshal([]byte(v), &r.Filters); err != nil {
				return err
			}
		case "jobType":
			if err := json.Unmarshal([]byte(v), &r.JobType); err != nil {
				return err
			}
		case "matrix":
			if err := json.Unmarshal([]byte(v), &r.Matrix); err != nil {
				return err
			}
		case "name":
			if err := json.Unmarshal([]byte(v), &r.Name); err != nil {
				return err
			}
		case "requires":
			if err := json.Unmarshal([]byte(v), &r.Requires); err != nil {
				return err
			}
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
