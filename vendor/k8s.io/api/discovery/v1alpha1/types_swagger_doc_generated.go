/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

// This file contains a collection of methods that can be used from go-restful to
// generate Swagger API documentation for its models. Please read this PR for more
// information on the implementation: https://github.com/emicklei/go-restful/pull/215
//
// TODOs are ignored from the parser (e.g. TODO(andronat):... || TODO:...) if and only if
// they are on one line! For multiple line or blocks that you want to ignore use ---.
// Any context after a --- is ignored.
//
// Those methods can be generated by using hack/update-generated-swagger-docs.sh

// AUTO-GENERATED FUNCTIONS START HERE. DO NOT EDIT.
var map_Endpoint = map[string]string{
	"":           "Endpoint represents a single logical \"backend\" implementing a service.",
	"addresses":  "addresses of this endpoint. The contents of this field are interpreted according to the corresponding EndpointSlice addressType field. This allows for cases like dual-stack networking where both IPv4 and IPv6 addresses would be included with the IP addressType. Consumers (e.g. kube-proxy) must handle different types of addresses in the context of their own capabilities. This must contain at least one address but no more than 100.",
	"conditions": "conditions contains information about the current status of the endpoint.",
	"hostname":   "hostname of this endpoint. This field may be used by consumers of endpoints to distinguish endpoints from each other (e.g. in DNS names). Multiple endpoints which use the same hostname should be considered fungible (e.g. multiple A values in DNS). Must pass DNS Label (RFC 1123) validation.",
	"targetRef":  "targetRef is a reference to a Kubernetes object that represents this endpoint.",
	"topology":   "topology contains arbitrary topology information associated with the endpoint. These key/value pairs must conform with the label format. https://kubernetes.io/docs/concepts/overview/working-with-objects/labels Topology may include a maximum of 16 key/value pairs. This includes, but is not limited to the following well known keys: * kubernetes.io/hostname: the value indicates the hostname of the node\n  where the endpoint is located. This should match the corresponding\n  node label.\n* topology.kubernetes.io/zone: the value indicates the zone where the\n  endpoint is located. This should match the corresponding node label.\n* topology.kubernetes.io/region: the value indicates the region where the\n  endpoint is located. This should match the corresponding node label.",
}

func (Endpoint) SwaggerDoc() map[string]string {
	return map_Endpoint
}

var map_EndpointConditions = map[string]string{
	"":      "EndpointConditions represents the current condition of an endpoint.",
	"ready": "ready indicates that this endpoint is prepared to receive traffic, according to whatever system is managing the endpoint. A nil value indicates an unknown state. In most cases consumers should interpret this unknown state as ready.",
}

func (EndpointConditions) SwaggerDoc() map[string]string {
	return map_EndpointConditions
}

var map_EndpointPort = map[string]string{
	"":         "EndpointPort represents a Port used by an EndpointSlice",
	"name":     "The name of this port. All ports in an EndpointSlice must have a unique name. If the EndpointSlice is dervied from a Kubernetes service, this corresponds to the Service.ports[].name. Name must either be an empty string or pass DNS_LABEL validation: * must be no more than 63 characters long. * must consist of lower case alphanumeric characters or '-'. * must start and end with an alphanumeric character. Default is empty string.",
	"protocol": "The IP protocol for this port. Must be UDP, TCP, or SCTP. Default is TCP.",
	"port":     "The port number of the endpoint. If this is not specified, ports are not restricted and must be interpreted in the context of the specific consumer.",
}

func (EndpointPort) SwaggerDoc() map[string]string {
	return map_EndpointPort
}

var map_EndpointSlice = map[string]string{
	"":            "EndpointSlice represents a subset of the endpoints that implement a service. For a given service there may be multiple EndpointSlice objects, selected by labels, which must be joined to produce the full set of endpoints.",
	"metadata":    "Standard object's metadata.",
	"addressType": "addressType specifies the type of address carried by this EndpointSlice. All addresses in this slice must be the same type. The following address types are currently supported: * IP:   Represents an IP Address. This can include both IPv4 and IPv6\n        addresses.\n* FQDN: Represents a Fully Qualified Domain Name. Default is IP",
	"endpoints":   "endpoints is a list of unique endpoints in this slice. Each slice may include a maximum of 1000 endpoints.",
	"ports":       "ports specifies the list of network ports exposed by each endpoint in this slice. Each port must have a unique name. When ports is empty, it indicates that there are no defined ports. When a port is defined with a nil port value, it indicates \"all ports\". Each slice may include a maximum of 100 ports.",
}

func (EndpointSlice) SwaggerDoc() map[string]string {
	return map_EndpointSlice
}

var map_EndpointSliceList = map[string]string{
	"":         "EndpointSliceList represents a list of endpoint slices",
	"metadata": "Standard list metadata.",
	"items":    "List of endpoint slices",
}

func (EndpointSliceList) SwaggerDoc() map[string]string {
	return map_EndpointSliceList
}

// AUTO-GENERATED FUNCTIONS END HERE
