package containerinstance

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// ContainerGroupNetworkProtocol enumerates the values for container group network protocol.
type ContainerGroupNetworkProtocol string

const (
	// TCP ...
	TCP ContainerGroupNetworkProtocol = "TCP"
	// UDP ...
	UDP ContainerGroupNetworkProtocol = "UDP"
)

// PossibleContainerGroupNetworkProtocolValues returns an array of possible values for the ContainerGroupNetworkProtocol const type.
func PossibleContainerGroupNetworkProtocolValues() []ContainerGroupNetworkProtocol {
	return []ContainerGroupNetworkProtocol{TCP, UDP}
}

// ContainerRestartPolicy enumerates the values for container restart policy.
type ContainerRestartPolicy string

const (
	// Always ...
	Always ContainerRestartPolicy = "always"
)

// PossibleContainerRestartPolicyValues returns an array of possible values for the ContainerRestartPolicy const type.
func PossibleContainerRestartPolicyValues() []ContainerRestartPolicy {
	return []ContainerRestartPolicy{Always}
}

// OperatingSystemTypes enumerates the values for operating system types.
type OperatingSystemTypes string

const (
	// Linux ...
	Linux OperatingSystemTypes = "Linux"
	// Windows ...
	Windows OperatingSystemTypes = "Windows"
)

// PossibleOperatingSystemTypesValues returns an array of possible values for the OperatingSystemTypes const type.
func PossibleOperatingSystemTypesValues() []OperatingSystemTypes {
	return []OperatingSystemTypes{Linux, Windows}
}