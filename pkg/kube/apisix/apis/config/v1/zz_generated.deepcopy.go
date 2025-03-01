// +build !ignore_autogenerated

// Licensed to the Apache Software Foundation (ASF) under one or more
// contributor license agreements.  See the NOTICE file distributed with
// this work for additional information regarding copyright ownership.
// The ASF licenses this file to You under the Apache License, Version 2.0
// (the "License"); you may not use this file except in compliance with
// the License.  You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ActiveHealthCheck) DeepCopyInto(out *ActiveHealthCheck) {
	*out = *in
	if in.StrictTLS != nil {
		in, out := &in.StrictTLS, &out.StrictTLS
		*out = new(bool)
		**out = **in
	}
	if in.RequestHeaders != nil {
		in, out := &in.RequestHeaders, &out.RequestHeaders
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Healthy != nil {
		in, out := &in.Healthy, &out.Healthy
		*out = new(ActiveHealthCheckHealthy)
		(*in).DeepCopyInto(*out)
	}
	if in.Unhealthy != nil {
		in, out := &in.Unhealthy, &out.Unhealthy
		*out = new(ActiveHealthCheckUnhealthy)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ActiveHealthCheck.
func (in *ActiveHealthCheck) DeepCopy() *ActiveHealthCheck {
	if in == nil {
		return nil
	}
	out := new(ActiveHealthCheck)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ActiveHealthCheckHealthy) DeepCopyInto(out *ActiveHealthCheckHealthy) {
	*out = *in
	in.PassiveHealthCheckHealthy.DeepCopyInto(&out.PassiveHealthCheckHealthy)
	out.Interval = in.Interval
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ActiveHealthCheckHealthy.
func (in *ActiveHealthCheckHealthy) DeepCopy() *ActiveHealthCheckHealthy {
	if in == nil {
		return nil
	}
	out := new(ActiveHealthCheckHealthy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ActiveHealthCheckUnhealthy) DeepCopyInto(out *ActiveHealthCheckUnhealthy) {
	*out = *in
	in.PassiveHealthCheckUnhealthy.DeepCopyInto(&out.PassiveHealthCheckUnhealthy)
	out.Interval = in.Interval
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ActiveHealthCheckUnhealthy.
func (in *ActiveHealthCheckUnhealthy) DeepCopy() *ActiveHealthCheckUnhealthy {
	if in == nil {
		return nil
	}
	out := new(ActiveHealthCheckUnhealthy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApisixMutualTlsClientConfig) DeepCopyInto(out *ApisixMutualTlsClientConfig) {
	*out = *in
	out.CASecret = in.CASecret
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApisixMutualTlsClientConfig.
func (in *ApisixMutualTlsClientConfig) DeepCopy() *ApisixMutualTlsClientConfig {
	if in == nil {
		return nil
	}
	out := new(ApisixMutualTlsClientConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApisixRoute) DeepCopyInto(out *ApisixRoute) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.Spec != nil {
		in, out := &in.Spec, &out.Spec
		*out = new(ApisixRouteSpec)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApisixRoute.
func (in *ApisixRoute) DeepCopy() *ApisixRoute {
	if in == nil {
		return nil
	}
	out := new(ApisixRoute)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ApisixRoute) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApisixRouteList) DeepCopyInto(out *ApisixRouteList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ApisixRoute, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApisixRouteList.
func (in *ApisixRouteList) DeepCopy() *ApisixRouteList {
	if in == nil {
		return nil
	}
	out := new(ApisixRouteList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ApisixRouteList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApisixRouteSpec) DeepCopyInto(out *ApisixRouteSpec) {
	*out = *in
	if in.Rules != nil {
		in, out := &in.Rules, &out.Rules
		*out = make([]Rule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApisixRouteSpec.
func (in *ApisixRouteSpec) DeepCopy() *ApisixRouteSpec {
	if in == nil {
		return nil
	}
	out := new(ApisixRouteSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApisixSecret) DeepCopyInto(out *ApisixSecret) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApisixSecret.
func (in *ApisixSecret) DeepCopy() *ApisixSecret {
	if in == nil {
		return nil
	}
	out := new(ApisixSecret)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApisixTls) DeepCopyInto(out *ApisixTls) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.Spec != nil {
		in, out := &in.Spec, &out.Spec
		*out = new(ApisixTlsSpec)
		(*in).DeepCopyInto(*out)
	}
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApisixTls.
func (in *ApisixTls) DeepCopy() *ApisixTls {
	if in == nil {
		return nil
	}
	out := new(ApisixTls)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ApisixTls) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApisixTlsList) DeepCopyInto(out *ApisixTlsList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ApisixTls, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApisixTlsList.
func (in *ApisixTlsList) DeepCopy() *ApisixTlsList {
	if in == nil {
		return nil
	}
	out := new(ApisixTlsList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ApisixTlsList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApisixTlsSpec) DeepCopyInto(out *ApisixTlsSpec) {
	*out = *in
	if in.Hosts != nil {
		in, out := &in.Hosts, &out.Hosts
		*out = make([]HostType, len(*in))
		copy(*out, *in)
	}
	out.Secret = in.Secret
	if in.Client != nil {
		in, out := &in.Client, &out.Client
		*out = new(ApisixMutualTlsClientConfig)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApisixTlsSpec.
func (in *ApisixTlsSpec) DeepCopy() *ApisixTlsSpec {
	if in == nil {
		return nil
	}
	out := new(ApisixTlsSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApisixUpstream) DeepCopyInto(out *ApisixUpstream) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.Spec != nil {
		in, out := &in.Spec, &out.Spec
		*out = new(ApisixUpstreamSpec)
		(*in).DeepCopyInto(*out)
	}
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApisixUpstream.
func (in *ApisixUpstream) DeepCopy() *ApisixUpstream {
	if in == nil {
		return nil
	}
	out := new(ApisixUpstream)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ApisixUpstream) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApisixUpstreamConfig) DeepCopyInto(out *ApisixUpstreamConfig) {
	*out = *in
	if in.LoadBalancer != nil {
		in, out := &in.LoadBalancer, &out.LoadBalancer
		*out = new(LoadBalancer)
		**out = **in
	}
	if in.Retries != nil {
		in, out := &in.Retries, &out.Retries
		*out = new(int)
		**out = **in
	}
	if in.Timeout != nil {
		in, out := &in.Timeout, &out.Timeout
		*out = new(UpstreamTimeout)
		**out = **in
	}
	if in.HealthCheck != nil {
		in, out := &in.HealthCheck, &out.HealthCheck
		*out = new(HealthCheck)
		(*in).DeepCopyInto(*out)
	}
	if in.Subsets != nil {
		in, out := &in.Subsets, &out.Subsets
		*out = make([]ApisixUpstreamSubset, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApisixUpstreamConfig.
func (in *ApisixUpstreamConfig) DeepCopy() *ApisixUpstreamConfig {
	if in == nil {
		return nil
	}
	out := new(ApisixUpstreamConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApisixUpstreamList) DeepCopyInto(out *ApisixUpstreamList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ApisixUpstream, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApisixUpstreamList.
func (in *ApisixUpstreamList) DeepCopy() *ApisixUpstreamList {
	if in == nil {
		return nil
	}
	out := new(ApisixUpstreamList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ApisixUpstreamList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApisixUpstreamSpec) DeepCopyInto(out *ApisixUpstreamSpec) {
	*out = *in
	in.ApisixUpstreamConfig.DeepCopyInto(&out.ApisixUpstreamConfig)
	if in.PortLevelSettings != nil {
		in, out := &in.PortLevelSettings, &out.PortLevelSettings
		*out = make([]PortLevelSettings, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApisixUpstreamSpec.
func (in *ApisixUpstreamSpec) DeepCopy() *ApisixUpstreamSpec {
	if in == nil {
		return nil
	}
	out := new(ApisixUpstreamSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApisixUpstreamSubset) DeepCopyInto(out *ApisixUpstreamSubset) {
	*out = *in
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApisixUpstreamSubset.
func (in *ApisixUpstreamSubset) DeepCopy() *ApisixUpstreamSubset {
	if in == nil {
		return nil
	}
	out := new(ApisixUpstreamSubset)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Backend) DeepCopyInto(out *Backend) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Backend.
func (in *Backend) DeepCopy() *Backend {
	if in == nil {
		return nil
	}
	out := new(Backend)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HealthCheck) DeepCopyInto(out *HealthCheck) {
	*out = *in
	if in.Active != nil {
		in, out := &in.Active, &out.Active
		*out = new(ActiveHealthCheck)
		(*in).DeepCopyInto(*out)
	}
	if in.Passive != nil {
		in, out := &in.Passive, &out.Passive
		*out = new(PassiveHealthCheck)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HealthCheck.
func (in *HealthCheck) DeepCopy() *HealthCheck {
	if in == nil {
		return nil
	}
	out := new(HealthCheck)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Http) DeepCopyInto(out *Http) {
	*out = *in
	if in.Paths != nil {
		in, out := &in.Paths, &out.Paths
		*out = make([]Path, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Http.
func (in *Http) DeepCopy() *Http {
	if in == nil {
		return nil
	}
	out := new(Http)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoadBalancer) DeepCopyInto(out *LoadBalancer) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoadBalancer.
func (in *LoadBalancer) DeepCopy() *LoadBalancer {
	if in == nil {
		return nil
	}
	out := new(LoadBalancer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PassiveHealthCheck) DeepCopyInto(out *PassiveHealthCheck) {
	*out = *in
	if in.Healthy != nil {
		in, out := &in.Healthy, &out.Healthy
		*out = new(PassiveHealthCheckHealthy)
		(*in).DeepCopyInto(*out)
	}
	if in.Unhealthy != nil {
		in, out := &in.Unhealthy, &out.Unhealthy
		*out = new(PassiveHealthCheckUnhealthy)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PassiveHealthCheck.
func (in *PassiveHealthCheck) DeepCopy() *PassiveHealthCheck {
	if in == nil {
		return nil
	}
	out := new(PassiveHealthCheck)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PassiveHealthCheckHealthy) DeepCopyInto(out *PassiveHealthCheckHealthy) {
	*out = *in
	if in.HTTPCodes != nil {
		in, out := &in.HTTPCodes, &out.HTTPCodes
		*out = make([]int, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PassiveHealthCheckHealthy.
func (in *PassiveHealthCheckHealthy) DeepCopy() *PassiveHealthCheckHealthy {
	if in == nil {
		return nil
	}
	out := new(PassiveHealthCheckHealthy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PassiveHealthCheckUnhealthy) DeepCopyInto(out *PassiveHealthCheckUnhealthy) {
	*out = *in
	if in.HTTPCodes != nil {
		in, out := &in.HTTPCodes, &out.HTTPCodes
		*out = make([]int, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PassiveHealthCheckUnhealthy.
func (in *PassiveHealthCheckUnhealthy) DeepCopy() *PassiveHealthCheckUnhealthy {
	if in == nil {
		return nil
	}
	out := new(PassiveHealthCheckUnhealthy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Path) DeepCopyInto(out *Path) {
	*out = *in
	out.Backend = in.Backend
	if in.Timeout != nil {
		in, out := &in.Timeout, &out.Timeout
		*out = new(UpstreamTimeout)
		**out = **in
	}
	if in.Plugins != nil {
		in, out := &in.Plugins, &out.Plugins
		*out = make([]Plugin, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Path.
func (in *Path) DeepCopy() *Path {
	if in == nil {
		return nil
	}
	out := new(Path)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Plugin) DeepCopyInto(out *Plugin) {
	*out = *in
	in.Config.DeepCopyInto(&out.Config)
	in.ConfigSet.DeepCopyInto(&out.ConfigSet)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Plugin.
func (in *Plugin) DeepCopy() *Plugin {
	if in == nil {
		return nil
	}
	out := new(Plugin)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PortLevelSettings) DeepCopyInto(out *PortLevelSettings) {
	*out = *in
	in.ApisixUpstreamConfig.DeepCopyInto(&out.ApisixUpstreamConfig)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PortLevelSettings.
func (in *PortLevelSettings) DeepCopy() *PortLevelSettings {
	if in == nil {
		return nil
	}
	out := new(PortLevelSettings)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Rule) DeepCopyInto(out *Rule) {
	*out = *in
	in.Http.DeepCopyInto(&out.Http)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Rule.
func (in *Rule) DeepCopy() *Rule {
	if in == nil {
		return nil
	}
	out := new(Rule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UpstreamTimeout) DeepCopyInto(out *UpstreamTimeout) {
	*out = *in
	out.Connect = in.Connect
	out.Send = in.Send
	out.Read = in.Read
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UpstreamTimeout.
func (in *UpstreamTimeout) DeepCopy() *UpstreamTimeout {
	if in == nil {
		return nil
	}
	out := new(UpstreamTimeout)
	in.DeepCopyInto(out)
	return out
}
