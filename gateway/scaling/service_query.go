// License: OpenFaaS Community Edition (CE) EULA
// Copyright (c) 2017,2019-2024 OpenFaaS Author(s)

// Copyright (c) OpenFaaS Author(s). All rights reserved.

package scaling

// ServiceQuery provides interface for replica querying/setting
type ServiceQuery interface {
	GetReplicas(service, namespace string) (response ServiceQueryResponse, err error)
	SetReplicas(service, namespace string, count uint64) error
}

// ServiceQueryResponse response from querying a function status
type ServiceQueryResponse struct {
	Replicas          uint64
	MaxReplicas       uint64
	MinReplicas       uint64
	ScalingFactor     uint64
	AvailableReplicas uint64
	Annotations       *map[string]string
}
