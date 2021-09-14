// +build !ignore_autogenerated

/*


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

// Code generated by controller-gen. DO NOT EDIT.

package v1beta1

import (
	"github.com/application-stacks/runtime-component-operator/common"
	monitoringv1 "github.com/coreos/prometheus-operator/pkg/apis/monitoring/v1"
	routev1 "github.com/openshift/api/route/v1"
	appsv1 "k8s.io/api/apps/v1"
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OperationStatusCondition) DeepCopyInto(out *OperationStatusCondition) {
	*out = *in
	if in.LastTransitionTime != nil {
		in, out := &in.LastTransitionTime, &out.LastTransitionTime
		*out = (*in).DeepCopy()
	}
	in.LastUpdateTime.DeepCopyInto(&out.LastUpdateTime)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OperationStatusCondition.
func (in *OperationStatusCondition) DeepCopy() *OperationStatusCondition {
	if in == nil {
		return nil
	}
	out := new(OperationStatusCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RuntimeComponent) DeepCopyInto(out *RuntimeComponent) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RuntimeComponent.
func (in *RuntimeComponent) DeepCopy() *RuntimeComponent {
	if in == nil {
		return nil
	}
	out := new(RuntimeComponent)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RuntimeComponent) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RuntimeComponentAffinity) DeepCopyInto(out *RuntimeComponentAffinity) {
	*out = *in
	if in.NodeAffinity != nil {
		in, out := &in.NodeAffinity, &out.NodeAffinity
		*out = new(v1.NodeAffinity)
		(*in).DeepCopyInto(*out)
	}
	if in.PodAffinity != nil {
		in, out := &in.PodAffinity, &out.PodAffinity
		*out = new(v1.PodAffinity)
		(*in).DeepCopyInto(*out)
	}
	if in.PodAntiAffinity != nil {
		in, out := &in.PodAntiAffinity, &out.PodAntiAffinity
		*out = new(v1.PodAntiAffinity)
		(*in).DeepCopyInto(*out)
	}
	if in.Architecture != nil {
		in, out := &in.Architecture, &out.Architecture
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.NodeAffinityLabels != nil {
		in, out := &in.NodeAffinityLabels, &out.NodeAffinityLabels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RuntimeComponentAffinity.
func (in *RuntimeComponentAffinity) DeepCopy() *RuntimeComponentAffinity {
	if in == nil {
		return nil
	}
	out := new(RuntimeComponentAffinity)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RuntimeComponentAutoScaling) DeepCopyInto(out *RuntimeComponentAutoScaling) {
	*out = *in
	if in.TargetCPUUtilizationPercentage != nil {
		in, out := &in.TargetCPUUtilizationPercentage, &out.TargetCPUUtilizationPercentage
		*out = new(int32)
		**out = **in
	}
	if in.MinReplicas != nil {
		in, out := &in.MinReplicas, &out.MinReplicas
		*out = new(int32)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RuntimeComponentAutoScaling.
func (in *RuntimeComponentAutoScaling) DeepCopy() *RuntimeComponentAutoScaling {
	if in == nil {
		return nil
	}
	out := new(RuntimeComponentAutoScaling)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RuntimeComponentBindingExpose) DeepCopyInto(out *RuntimeComponentBindingExpose) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RuntimeComponentBindingExpose.
func (in *RuntimeComponentBindingExpose) DeepCopy() *RuntimeComponentBindingExpose {
	if in == nil {
		return nil
	}
	out := new(RuntimeComponentBindingExpose)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RuntimeComponentBindings) DeepCopyInto(out *RuntimeComponentBindings) {
	*out = *in
	if in.AutoDetect != nil {
		in, out := &in.AutoDetect, &out.AutoDetect
		*out = new(bool)
		**out = **in
	}
	if in.Embedded != nil {
		in, out := &in.Embedded, &out.Embedded
		*out = new(runtime.RawExtension)
		(*in).DeepCopyInto(*out)
	}
	if in.Expose != nil {
		in, out := &in.Expose, &out.Expose
		*out = new(RuntimeComponentBindingExpose)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RuntimeComponentBindings.
func (in *RuntimeComponentBindings) DeepCopy() *RuntimeComponentBindings {
	if in == nil {
		return nil
	}
	out := new(RuntimeComponentBindings)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RuntimeComponentDeployment) DeepCopyInto(out *RuntimeComponentDeployment) {
	*out = *in
	if in.UpdateStrategy != nil {
		in, out := &in.UpdateStrategy, &out.UpdateStrategy
		*out = new(appsv1.DeploymentStrategy)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RuntimeComponentDeployment.
func (in *RuntimeComponentDeployment) DeepCopy() *RuntimeComponentDeployment {
	if in == nil {
		return nil
	}
	out := new(RuntimeComponentDeployment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RuntimeComponentList) DeepCopyInto(out *RuntimeComponentList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RuntimeComponent, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RuntimeComponentList.
func (in *RuntimeComponentList) DeepCopy() *RuntimeComponentList {
	if in == nil {
		return nil
	}
	out := new(RuntimeComponentList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RuntimeComponentList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RuntimeComponentMonitoring) DeepCopyInto(out *RuntimeComponentMonitoring) {
	*out = *in
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Endpoints != nil {
		in, out := &in.Endpoints, &out.Endpoints
		*out = make([]monitoringv1.Endpoint, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RuntimeComponentMonitoring.
func (in *RuntimeComponentMonitoring) DeepCopy() *RuntimeComponentMonitoring {
	if in == nil {
		return nil
	}
	out := new(RuntimeComponentMonitoring)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RuntimeComponentPersistentVolumeClaim) DeepCopyInto(out *RuntimeComponentPersistentVolumeClaim) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.Spec != nil {
		in, out := &in.Spec, &out.Spec
		*out = new(v1.PersistentVolumeClaimSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(v1.PersistentVolumeClaimStatus)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RuntimeComponentPersistentVolumeClaim.
func (in *RuntimeComponentPersistentVolumeClaim) DeepCopy() *RuntimeComponentPersistentVolumeClaim {
	if in == nil {
		return nil
	}
	out := new(RuntimeComponentPersistentVolumeClaim)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RuntimeComponentRoute) DeepCopyInto(out *RuntimeComponentRoute) {
	*out = *in
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Termination != nil {
		in, out := &in.Termination, &out.Termination
		*out = new(routev1.TLSTerminationType)
		**out = **in
	}
	if in.InsecureEdgeTerminationPolicy != nil {
		in, out := &in.InsecureEdgeTerminationPolicy, &out.InsecureEdgeTerminationPolicy
		*out = new(routev1.InsecureEdgeTerminationPolicyType)
		**out = **in
	}
	if in.CertificateSecretRef != nil {
		in, out := &in.CertificateSecretRef, &out.CertificateSecretRef
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RuntimeComponentRoute.
func (in *RuntimeComponentRoute) DeepCopy() *RuntimeComponentRoute {
	if in == nil {
		return nil
	}
	out := new(RuntimeComponentRoute)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RuntimeComponentService) DeepCopyInto(out *RuntimeComponentService) {
	*out = *in
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(v1.ServiceType)
		**out = **in
	}
	if in.TargetPort != nil {
		in, out := &in.TargetPort, &out.TargetPort
		*out = new(int32)
		**out = **in
	}
	if in.NodePort != nil {
		in, out := &in.NodePort, &out.NodePort
		*out = new(int32)
		**out = **in
	}
	if in.Ports != nil {
		in, out := &in.Ports, &out.Ports
		*out = make([]v1.ServicePort, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Consumes != nil {
		in, out := &in.Consumes, &out.Consumes
		*out = make([]ServiceBindingConsumes, len(*in))
		copy(*out, *in)
	}
	if in.Provides != nil {
		in, out := &in.Provides, &out.Provides
		*out = new(ServiceBindingProvides)
		(*in).DeepCopyInto(*out)
	}
	if in.CertificateSecretRef != nil {
		in, out := &in.CertificateSecretRef, &out.CertificateSecretRef
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RuntimeComponentService.
func (in *RuntimeComponentService) DeepCopy() *RuntimeComponentService {
	if in == nil {
		return nil
	}
	out := new(RuntimeComponentService)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RuntimeComponentSpec) DeepCopyInto(out *RuntimeComponentSpec) {
	*out = *in
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int32)
		**out = **in
	}
	if in.Autoscaling != nil {
		in, out := &in.Autoscaling, &out.Autoscaling
		*out = new(RuntimeComponentAutoScaling)
		(*in).DeepCopyInto(*out)
	}
	if in.PullPolicy != nil {
		in, out := &in.PullPolicy, &out.PullPolicy
		*out = new(v1.PullPolicy)
		**out = **in
	}
	if in.PullSecret != nil {
		in, out := &in.PullSecret, &out.PullSecret
		*out = new(string)
		**out = **in
	}
	if in.Volumes != nil {
		in, out := &in.Volumes, &out.Volumes
		*out = make([]v1.Volume, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.VolumeMounts != nil {
		in, out := &in.VolumeMounts, &out.VolumeMounts
		*out = make([]v1.VolumeMount, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ResourceConstraints != nil {
		in, out := &in.ResourceConstraints, &out.ResourceConstraints
		*out = new(v1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
	if in.ReadinessProbe != nil {
		in, out := &in.ReadinessProbe, &out.ReadinessProbe
		*out = new(v1.Probe)
		(*in).DeepCopyInto(*out)
	}
	if in.LivenessProbe != nil {
		in, out := &in.LivenessProbe, &out.LivenessProbe
		*out = new(v1.Probe)
		(*in).DeepCopyInto(*out)
	}
	if in.StartupProbe != nil {
		in, out := &in.StartupProbe, &out.StartupProbe
		*out = new(v1.Probe)
		(*in).DeepCopyInto(*out)
	}
	if in.Service != nil {
		in, out := &in.Service, &out.Service
		*out = new(RuntimeComponentService)
		(*in).DeepCopyInto(*out)
	}
	if in.Expose != nil {
		in, out := &in.Expose, &out.Expose
		*out = new(bool)
		**out = **in
	}
	if in.Deployment != nil {
		in, out := &in.Deployment, &out.Deployment
		*out = new(RuntimeComponentDeployment)
		(*in).DeepCopyInto(*out)
	}
	if in.StatefulSet != nil {
		in, out := &in.StatefulSet, &out.StatefulSet
		*out = new(RuntimeComponentStatefulSet)
		(*in).DeepCopyInto(*out)
	}
	if in.EnvFrom != nil {
		in, out := &in.EnvFrom, &out.EnvFrom
		*out = make([]v1.EnvFromSource, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make([]v1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ServiceAccountName != nil {
		in, out := &in.ServiceAccountName, &out.ServiceAccountName
		*out = new(string)
		**out = **in
	}
	if in.Architecture != nil {
		in, out := &in.Architecture, &out.Architecture
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Storage != nil {
		in, out := &in.Storage, &out.Storage
		*out = new(RuntimeComponentStorage)
		(*in).DeepCopyInto(*out)
	}
	if in.CreateKnativeService != nil {
		in, out := &in.CreateKnativeService, &out.CreateKnativeService
		*out = new(bool)
		**out = **in
	}
	if in.Monitoring != nil {
		in, out := &in.Monitoring, &out.Monitoring
		*out = new(RuntimeComponentMonitoring)
		(*in).DeepCopyInto(*out)
	}
	if in.InitContainers != nil {
		in, out := &in.InitContainers, &out.InitContainers
		*out = make([]v1.Container, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SidecarContainers != nil {
		in, out := &in.SidecarContainers, &out.SidecarContainers
		*out = make([]v1.Container, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Route != nil {
		in, out := &in.Route, &out.Route
		*out = new(RuntimeComponentRoute)
		(*in).DeepCopyInto(*out)
	}
	if in.Bindings != nil {
		in, out := &in.Bindings, &out.Bindings
		*out = new(RuntimeComponentBindings)
		(*in).DeepCopyInto(*out)
	}
	if in.Affinity != nil {
		in, out := &in.Affinity, &out.Affinity
		*out = new(RuntimeComponentAffinity)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RuntimeComponentSpec.
func (in *RuntimeComponentSpec) DeepCopy() *RuntimeComponentSpec {
	if in == nil {
		return nil
	}
	out := new(RuntimeComponentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RuntimeComponentStatefulSet) DeepCopyInto(out *RuntimeComponentStatefulSet) {
	*out = *in
	if in.UpdateStrategy != nil {
		in, out := &in.UpdateStrategy, &out.UpdateStrategy
		*out = new(appsv1.StatefulSetUpdateStrategy)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RuntimeComponentStatefulSet.
func (in *RuntimeComponentStatefulSet) DeepCopy() *RuntimeComponentStatefulSet {
	if in == nil {
		return nil
	}
	out := new(RuntimeComponentStatefulSet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RuntimeComponentStatus) DeepCopyInto(out *RuntimeComponentStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]StatusCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ConsumedServices != nil {
		in, out := &in.ConsumedServices, &out.ConsumedServices
		*out = make(common.ConsumedServices, len(*in))
		for key, val := range *in {
			var outVal []string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make([]string, len(*in))
				copy(*out, *in)
			}
			(*out)[key] = outVal
		}
	}
	if in.ResolvedBindings != nil {
		in, out := &in.ResolvedBindings, &out.ResolvedBindings
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Binding != nil {
		in, out := &in.Binding, &out.Binding
		*out = new(v1.LocalObjectReference)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RuntimeComponentStatus.
func (in *RuntimeComponentStatus) DeepCopy() *RuntimeComponentStatus {
	if in == nil {
		return nil
	}
	out := new(RuntimeComponentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RuntimeComponentStorage) DeepCopyInto(out *RuntimeComponentStorage) {
	*out = *in
	if in.VolumeClaimTemplate != nil {
		in, out := &in.VolumeClaimTemplate, &out.VolumeClaimTemplate
		*out = new(RuntimeComponentPersistentVolumeClaim)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RuntimeComponentStorage.
func (in *RuntimeComponentStorage) DeepCopy() *RuntimeComponentStorage {
	if in == nil {
		return nil
	}
	out := new(RuntimeComponentStorage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RuntimeOperation) DeepCopyInto(out *RuntimeOperation) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RuntimeOperation.
func (in *RuntimeOperation) DeepCopy() *RuntimeOperation {
	if in == nil {
		return nil
	}
	out := new(RuntimeOperation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RuntimeOperation) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RuntimeOperationList) DeepCopyInto(out *RuntimeOperationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RuntimeOperation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RuntimeOperationList.
func (in *RuntimeOperationList) DeepCopy() *RuntimeOperationList {
	if in == nil {
		return nil
	}
	out := new(RuntimeOperationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RuntimeOperationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RuntimeOperationSpec) DeepCopyInto(out *RuntimeOperationSpec) {
	*out = *in
	if in.Command != nil {
		in, out := &in.Command, &out.Command
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RuntimeOperationSpec.
func (in *RuntimeOperationSpec) DeepCopy() *RuntimeOperationSpec {
	if in == nil {
		return nil
	}
	out := new(RuntimeOperationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RuntimeOperationStatus) DeepCopyInto(out *RuntimeOperationStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]OperationStatusCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RuntimeOperationStatus.
func (in *RuntimeOperationStatus) DeepCopy() *RuntimeOperationStatus {
	if in == nil {
		return nil
	}
	out := new(RuntimeOperationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceBindingAuth) DeepCopyInto(out *ServiceBindingAuth) {
	*out = *in
	in.Username.DeepCopyInto(&out.Username)
	in.Password.DeepCopyInto(&out.Password)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceBindingAuth.
func (in *ServiceBindingAuth) DeepCopy() *ServiceBindingAuth {
	if in == nil {
		return nil
	}
	out := new(ServiceBindingAuth)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceBindingConsumes) DeepCopyInto(out *ServiceBindingConsumes) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceBindingConsumes.
func (in *ServiceBindingConsumes) DeepCopy() *ServiceBindingConsumes {
	if in == nil {
		return nil
	}
	out := new(ServiceBindingConsumes)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceBindingProvides) DeepCopyInto(out *ServiceBindingProvides) {
	*out = *in
	if in.Auth != nil {
		in, out := &in.Auth, &out.Auth
		*out = new(ServiceBindingAuth)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceBindingProvides.
func (in *ServiceBindingProvides) DeepCopy() *ServiceBindingProvides {
	if in == nil {
		return nil
	}
	out := new(ServiceBindingProvides)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StatusCondition) DeepCopyInto(out *StatusCondition) {
	*out = *in
	if in.LastTransitionTime != nil {
		in, out := &in.LastTransitionTime, &out.LastTransitionTime
		*out = (*in).DeepCopy()
	}
	in.LastUpdateTime.DeepCopyInto(&out.LastUpdateTime)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StatusCondition.
func (in *StatusCondition) DeepCopy() *StatusCondition {
	if in == nil {
		return nil
	}
	out := new(StatusCondition)
	in.DeepCopyInto(out)
	return out
}
