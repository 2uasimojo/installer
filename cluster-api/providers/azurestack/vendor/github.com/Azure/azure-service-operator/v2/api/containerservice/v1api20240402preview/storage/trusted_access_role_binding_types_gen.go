// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package storage

import (
	v20231001s "github.com/Azure/azure-service-operator/v2/api/containerservice/v1api20231001/storage"
	v20231102ps "github.com/Azure/azure-service-operator/v2/api/containerservice/v1api20231102preview/storage"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/configmaps"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/core"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/secrets"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/conversion"
)

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Storage version of v1api20240402preview.TrustedAccessRoleBinding
// Generator information:
// - Generated from: /containerservice/resource-manager/Microsoft.ContainerService/aks/preview/2024-04-02-preview/managedClusters.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedClusters/{resourceName}/trustedAccessRoleBindings/{trustedAccessRoleBindingName}
type TrustedAccessRoleBinding struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TrustedAccessRoleBinding_Spec   `json:"spec,omitempty"`
	Status            TrustedAccessRoleBinding_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &TrustedAccessRoleBinding{}

// GetConditions returns the conditions of the resource
func (binding *TrustedAccessRoleBinding) GetConditions() conditions.Conditions {
	return binding.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (binding *TrustedAccessRoleBinding) SetConditions(conditions conditions.Conditions) {
	binding.Status.Conditions = conditions
}

var _ conversion.Convertible = &TrustedAccessRoleBinding{}

// ConvertFrom populates our TrustedAccessRoleBinding from the provided hub TrustedAccessRoleBinding
func (binding *TrustedAccessRoleBinding) ConvertFrom(hub conversion.Hub) error {
	// intermediate variable for conversion
	var source v20231001s.TrustedAccessRoleBinding

	err := source.ConvertFrom(hub)
	if err != nil {
		return errors.Wrap(err, "converting from hub to source")
	}

	err = binding.AssignProperties_From_TrustedAccessRoleBinding(&source)
	if err != nil {
		return errors.Wrap(err, "converting from source to binding")
	}

	return nil
}

// ConvertTo populates the provided hub TrustedAccessRoleBinding from our TrustedAccessRoleBinding
func (binding *TrustedAccessRoleBinding) ConvertTo(hub conversion.Hub) error {
	// intermediate variable for conversion
	var destination v20231001s.TrustedAccessRoleBinding
	err := binding.AssignProperties_To_TrustedAccessRoleBinding(&destination)
	if err != nil {
		return errors.Wrap(err, "converting to destination from binding")
	}
	err = destination.ConvertTo(hub)
	if err != nil {
		return errors.Wrap(err, "converting from destination to hub")
	}

	return nil
}

var _ configmaps.Exporter = &TrustedAccessRoleBinding{}

// ConfigMapDestinationExpressions returns the Spec.OperatorSpec.ConfigMapExpressions property
func (binding *TrustedAccessRoleBinding) ConfigMapDestinationExpressions() []*core.DestinationExpression {
	if binding.Spec.OperatorSpec == nil {
		return nil
	}
	return binding.Spec.OperatorSpec.ConfigMapExpressions
}

var _ secrets.Exporter = &TrustedAccessRoleBinding{}

// SecretDestinationExpressions returns the Spec.OperatorSpec.SecretExpressions property
func (binding *TrustedAccessRoleBinding) SecretDestinationExpressions() []*core.DestinationExpression {
	if binding.Spec.OperatorSpec == nil {
		return nil
	}
	return binding.Spec.OperatorSpec.SecretExpressions
}

var _ genruntime.KubernetesResource = &TrustedAccessRoleBinding{}

// AzureName returns the Azure name of the resource
func (binding *TrustedAccessRoleBinding) AzureName() string {
	return binding.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2024-04-02-preview"
func (binding TrustedAccessRoleBinding) GetAPIVersion() string {
	return "2024-04-02-preview"
}

// GetResourceScope returns the scope of the resource
func (binding *TrustedAccessRoleBinding) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (binding *TrustedAccessRoleBinding) GetSpec() genruntime.ConvertibleSpec {
	return &binding.Spec
}

// GetStatus returns the status of this resource
func (binding *TrustedAccessRoleBinding) GetStatus() genruntime.ConvertibleStatus {
	return &binding.Status
}

// GetSupportedOperations returns the operations supported by the resource
func (binding *TrustedAccessRoleBinding) GetSupportedOperations() []genruntime.ResourceOperation {
	return []genruntime.ResourceOperation{
		genruntime.ResourceOperationDelete,
		genruntime.ResourceOperationGet,
		genruntime.ResourceOperationPut,
	}
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.ContainerService/managedClusters/trustedAccessRoleBindings"
func (binding *TrustedAccessRoleBinding) GetType() string {
	return "Microsoft.ContainerService/managedClusters/trustedAccessRoleBindings"
}

// NewEmptyStatus returns a new empty (blank) status
func (binding *TrustedAccessRoleBinding) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &TrustedAccessRoleBinding_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (binding *TrustedAccessRoleBinding) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(binding.Spec)
	return binding.Spec.Owner.AsResourceReference(group, kind)
}

// SetStatus sets the status of this resource
func (binding *TrustedAccessRoleBinding) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*TrustedAccessRoleBinding_STATUS); ok {
		binding.Status = *st
		return nil
	}

	// Convert status to required version
	var st TrustedAccessRoleBinding_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	binding.Status = st
	return nil
}

// AssignProperties_From_TrustedAccessRoleBinding populates our TrustedAccessRoleBinding from the provided source TrustedAccessRoleBinding
func (binding *TrustedAccessRoleBinding) AssignProperties_From_TrustedAccessRoleBinding(source *v20231001s.TrustedAccessRoleBinding) error {

	// ObjectMeta
	binding.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec TrustedAccessRoleBinding_Spec
	err := spec.AssignProperties_From_TrustedAccessRoleBinding_Spec(&source.Spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_From_TrustedAccessRoleBinding_Spec() to populate field Spec")
	}
	binding.Spec = spec

	// Status
	var status TrustedAccessRoleBinding_STATUS
	err = status.AssignProperties_From_TrustedAccessRoleBinding_STATUS(&source.Status)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_From_TrustedAccessRoleBinding_STATUS() to populate field Status")
	}
	binding.Status = status

	// Invoke the augmentConversionForTrustedAccessRoleBinding interface (if implemented) to customize the conversion
	var bindingAsAny any = binding
	if augmentedBinding, ok := bindingAsAny.(augmentConversionForTrustedAccessRoleBinding); ok {
		err := augmentedBinding.AssignPropertiesFrom(source)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_TrustedAccessRoleBinding populates the provided destination TrustedAccessRoleBinding from our TrustedAccessRoleBinding
func (binding *TrustedAccessRoleBinding) AssignProperties_To_TrustedAccessRoleBinding(destination *v20231001s.TrustedAccessRoleBinding) error {

	// ObjectMeta
	destination.ObjectMeta = *binding.ObjectMeta.DeepCopy()

	// Spec
	var spec v20231001s.TrustedAccessRoleBinding_Spec
	err := binding.Spec.AssignProperties_To_TrustedAccessRoleBinding_Spec(&spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_To_TrustedAccessRoleBinding_Spec() to populate field Spec")
	}
	destination.Spec = spec

	// Status
	var status v20231001s.TrustedAccessRoleBinding_STATUS
	err = binding.Status.AssignProperties_To_TrustedAccessRoleBinding_STATUS(&status)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_To_TrustedAccessRoleBinding_STATUS() to populate field Status")
	}
	destination.Status = status

	// Invoke the augmentConversionForTrustedAccessRoleBinding interface (if implemented) to customize the conversion
	var bindingAsAny any = binding
	if augmentedBinding, ok := bindingAsAny.(augmentConversionForTrustedAccessRoleBinding); ok {
		err := augmentedBinding.AssignPropertiesTo(destination)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (binding *TrustedAccessRoleBinding) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: binding.Spec.OriginalVersion,
		Kind:    "TrustedAccessRoleBinding",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1api20240402preview.TrustedAccessRoleBinding
// Generator information:
// - Generated from: /containerservice/resource-manager/Microsoft.ContainerService/aks/preview/2024-04-02-preview/managedClusters.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedClusters/{resourceName}/trustedAccessRoleBindings/{trustedAccessRoleBindingName}
type TrustedAccessRoleBindingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TrustedAccessRoleBinding `json:"items"`
}

type augmentConversionForTrustedAccessRoleBinding interface {
	AssignPropertiesFrom(src *v20231001s.TrustedAccessRoleBinding) error
	AssignPropertiesTo(dst *v20231001s.TrustedAccessRoleBinding) error
}

// Storage version of v1api20240402preview.TrustedAccessRoleBinding_Spec
type TrustedAccessRoleBinding_Spec struct {
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName       string                                `json:"azureName,omitempty"`
	OperatorSpec    *TrustedAccessRoleBindingOperatorSpec `json:"operatorSpec,omitempty"`
	OriginalVersion string                                `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a containerservice.azure.com/ManagedCluster resource
	Owner       *genruntime.KnownResourceReference `group:"containerservice.azure.com" json:"owner,omitempty" kind:"ManagedCluster"`
	PropertyBag genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	Roles       []string                           `json:"roles,omitempty"`

	// +kubebuilder:validation:Required
	// SourceResourceReference: The ARM resource ID of source resource that trusted access is configured for.
	SourceResourceReference *genruntime.ResourceReference `armReference:"SourceResourceId" json:"sourceResourceReference,omitempty"`
}

var _ genruntime.ConvertibleSpec = &TrustedAccessRoleBinding_Spec{}

// ConvertSpecFrom populates our TrustedAccessRoleBinding_Spec from the provided source
func (binding *TrustedAccessRoleBinding_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*v20231001s.TrustedAccessRoleBinding_Spec)
	if ok {
		// Populate our instance from source
		return binding.AssignProperties_From_TrustedAccessRoleBinding_Spec(src)
	}

	// Convert to an intermediate form
	src = &v20231001s.TrustedAccessRoleBinding_Spec{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = binding.AssignProperties_From_TrustedAccessRoleBinding_Spec(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our TrustedAccessRoleBinding_Spec
func (binding *TrustedAccessRoleBinding_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*v20231001s.TrustedAccessRoleBinding_Spec)
	if ok {
		// Populate destination from our instance
		return binding.AssignProperties_To_TrustedAccessRoleBinding_Spec(dst)
	}

	// Convert to an intermediate form
	dst = &v20231001s.TrustedAccessRoleBinding_Spec{}
	err := binding.AssignProperties_To_TrustedAccessRoleBinding_Spec(dst)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecTo()")
	}

	// Update dst from our instance
	err = dst.ConvertSpecTo(destination)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecTo()")
	}

	return nil
}

// AssignProperties_From_TrustedAccessRoleBinding_Spec populates our TrustedAccessRoleBinding_Spec from the provided source TrustedAccessRoleBinding_Spec
func (binding *TrustedAccessRoleBinding_Spec) AssignProperties_From_TrustedAccessRoleBinding_Spec(source *v20231001s.TrustedAccessRoleBinding_Spec) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// AzureName
	binding.AzureName = source.AzureName

	// OperatorSpec
	if source.OperatorSpec != nil {
		var operatorSpec TrustedAccessRoleBindingOperatorSpec
		err := operatorSpec.AssignProperties_From_TrustedAccessRoleBindingOperatorSpec(source.OperatorSpec)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_From_TrustedAccessRoleBindingOperatorSpec() to populate field OperatorSpec")
		}
		binding.OperatorSpec = &operatorSpec
	} else {
		binding.OperatorSpec = nil
	}

	// OriginalVersion
	binding.OriginalVersion = source.OriginalVersion

	// Owner
	if source.Owner != nil {
		owner := source.Owner.Copy()
		binding.Owner = &owner
	} else {
		binding.Owner = nil
	}

	// Roles
	binding.Roles = genruntime.CloneSliceOfString(source.Roles)

	// SourceResourceReference
	if source.SourceResourceReference != nil {
		sourceResourceReference := source.SourceResourceReference.Copy()
		binding.SourceResourceReference = &sourceResourceReference
	} else {
		binding.SourceResourceReference = nil
	}

	// Update the property bag
	if len(propertyBag) > 0 {
		binding.PropertyBag = propertyBag
	} else {
		binding.PropertyBag = nil
	}

	// Invoke the augmentConversionForTrustedAccessRoleBinding_Spec interface (if implemented) to customize the conversion
	var bindingAsAny any = binding
	if augmentedBinding, ok := bindingAsAny.(augmentConversionForTrustedAccessRoleBinding_Spec); ok {
		err := augmentedBinding.AssignPropertiesFrom(source)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_TrustedAccessRoleBinding_Spec populates the provided destination TrustedAccessRoleBinding_Spec from our TrustedAccessRoleBinding_Spec
func (binding *TrustedAccessRoleBinding_Spec) AssignProperties_To_TrustedAccessRoleBinding_Spec(destination *v20231001s.TrustedAccessRoleBinding_Spec) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(binding.PropertyBag)

	// AzureName
	destination.AzureName = binding.AzureName

	// OperatorSpec
	if binding.OperatorSpec != nil {
		var operatorSpec v20231001s.TrustedAccessRoleBindingOperatorSpec
		err := binding.OperatorSpec.AssignProperties_To_TrustedAccessRoleBindingOperatorSpec(&operatorSpec)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_To_TrustedAccessRoleBindingOperatorSpec() to populate field OperatorSpec")
		}
		destination.OperatorSpec = &operatorSpec
	} else {
		destination.OperatorSpec = nil
	}

	// OriginalVersion
	destination.OriginalVersion = binding.OriginalVersion

	// Owner
	if binding.Owner != nil {
		owner := binding.Owner.Copy()
		destination.Owner = &owner
	} else {
		destination.Owner = nil
	}

	// Roles
	destination.Roles = genruntime.CloneSliceOfString(binding.Roles)

	// SourceResourceReference
	if binding.SourceResourceReference != nil {
		sourceResourceReference := binding.SourceResourceReference.Copy()
		destination.SourceResourceReference = &sourceResourceReference
	} else {
		destination.SourceResourceReference = nil
	}

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// Invoke the augmentConversionForTrustedAccessRoleBinding_Spec interface (if implemented) to customize the conversion
	var bindingAsAny any = binding
	if augmentedBinding, ok := bindingAsAny.(augmentConversionForTrustedAccessRoleBinding_Spec); ok {
		err := augmentedBinding.AssignPropertiesTo(destination)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

// Storage version of v1api20240402preview.TrustedAccessRoleBinding_STATUS
type TrustedAccessRoleBinding_STATUS struct {
	Conditions        []conditions.Condition `json:"conditions,omitempty"`
	Id                *string                `json:"id,omitempty"`
	Name              *string                `json:"name,omitempty"`
	PropertyBag       genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	ProvisioningState *string                `json:"provisioningState,omitempty"`
	Roles             []string               `json:"roles,omitempty"`
	SourceResourceId  *string                `json:"sourceResourceId,omitempty"`
	SystemData        *SystemData_STATUS     `json:"systemData,omitempty"`
	Type              *string                `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &TrustedAccessRoleBinding_STATUS{}

// ConvertStatusFrom populates our TrustedAccessRoleBinding_STATUS from the provided source
func (binding *TrustedAccessRoleBinding_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	src, ok := source.(*v20231001s.TrustedAccessRoleBinding_STATUS)
	if ok {
		// Populate our instance from source
		return binding.AssignProperties_From_TrustedAccessRoleBinding_STATUS(src)
	}

	// Convert to an intermediate form
	src = &v20231001s.TrustedAccessRoleBinding_STATUS{}
	err := src.ConvertStatusFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusFrom()")
	}

	// Update our instance from src
	err = binding.AssignProperties_From_TrustedAccessRoleBinding_STATUS(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusFrom()")
	}

	return nil
}

// ConvertStatusTo populates the provided destination from our TrustedAccessRoleBinding_STATUS
func (binding *TrustedAccessRoleBinding_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	dst, ok := destination.(*v20231001s.TrustedAccessRoleBinding_STATUS)
	if ok {
		// Populate destination from our instance
		return binding.AssignProperties_To_TrustedAccessRoleBinding_STATUS(dst)
	}

	// Convert to an intermediate form
	dst = &v20231001s.TrustedAccessRoleBinding_STATUS{}
	err := binding.AssignProperties_To_TrustedAccessRoleBinding_STATUS(dst)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusTo()")
	}

	// Update dst from our instance
	err = dst.ConvertStatusTo(destination)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusTo()")
	}

	return nil
}

// AssignProperties_From_TrustedAccessRoleBinding_STATUS populates our TrustedAccessRoleBinding_STATUS from the provided source TrustedAccessRoleBinding_STATUS
func (binding *TrustedAccessRoleBinding_STATUS) AssignProperties_From_TrustedAccessRoleBinding_STATUS(source *v20231001s.TrustedAccessRoleBinding_STATUS) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// Conditions
	binding.Conditions = genruntime.CloneSliceOfCondition(source.Conditions)

	// Id
	binding.Id = genruntime.ClonePointerToString(source.Id)

	// Name
	binding.Name = genruntime.ClonePointerToString(source.Name)

	// ProvisioningState
	binding.ProvisioningState = genruntime.ClonePointerToString(source.ProvisioningState)

	// Roles
	binding.Roles = genruntime.CloneSliceOfString(source.Roles)

	// SourceResourceId
	binding.SourceResourceId = genruntime.ClonePointerToString(source.SourceResourceId)

	// SystemData
	if source.SystemData != nil {
		var systemDataSTATUSStash v20231102ps.SystemData_STATUS
		err := systemDataSTATUSStash.AssignProperties_From_SystemData_STATUS(source.SystemData)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_From_SystemData_STATUS() to populate field SystemData_STATUSStash from SystemData")
		}
		var systemDatum SystemData_STATUS
		err = systemDatum.AssignProperties_From_SystemData_STATUS(&systemDataSTATUSStash)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_From_SystemData_STATUS() to populate field SystemData from SystemData_STATUSStash")
		}
		binding.SystemData = &systemDatum
	} else {
		binding.SystemData = nil
	}

	// Type
	binding.Type = genruntime.ClonePointerToString(source.Type)

	// Update the property bag
	if len(propertyBag) > 0 {
		binding.PropertyBag = propertyBag
	} else {
		binding.PropertyBag = nil
	}

	// Invoke the augmentConversionForTrustedAccessRoleBinding_STATUS interface (if implemented) to customize the conversion
	var bindingAsAny any = binding
	if augmentedBinding, ok := bindingAsAny.(augmentConversionForTrustedAccessRoleBinding_STATUS); ok {
		err := augmentedBinding.AssignPropertiesFrom(source)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_TrustedAccessRoleBinding_STATUS populates the provided destination TrustedAccessRoleBinding_STATUS from our TrustedAccessRoleBinding_STATUS
func (binding *TrustedAccessRoleBinding_STATUS) AssignProperties_To_TrustedAccessRoleBinding_STATUS(destination *v20231001s.TrustedAccessRoleBinding_STATUS) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(binding.PropertyBag)

	// Conditions
	destination.Conditions = genruntime.CloneSliceOfCondition(binding.Conditions)

	// Id
	destination.Id = genruntime.ClonePointerToString(binding.Id)

	// Name
	destination.Name = genruntime.ClonePointerToString(binding.Name)

	// ProvisioningState
	destination.ProvisioningState = genruntime.ClonePointerToString(binding.ProvisioningState)

	// Roles
	destination.Roles = genruntime.CloneSliceOfString(binding.Roles)

	// SourceResourceId
	destination.SourceResourceId = genruntime.ClonePointerToString(binding.SourceResourceId)

	// SystemData
	if binding.SystemData != nil {
		var systemDataSTATUSStash v20231102ps.SystemData_STATUS
		err := binding.SystemData.AssignProperties_To_SystemData_STATUS(&systemDataSTATUSStash)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_To_SystemData_STATUS() to populate field SystemData_STATUSStash from SystemData")
		}
		var systemDatum v20231001s.SystemData_STATUS
		err = systemDataSTATUSStash.AssignProperties_To_SystemData_STATUS(&systemDatum)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_To_SystemData_STATUS() to populate field SystemData from SystemData_STATUSStash")
		}
		destination.SystemData = &systemDatum
	} else {
		destination.SystemData = nil
	}

	// Type
	destination.Type = genruntime.ClonePointerToString(binding.Type)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// Invoke the augmentConversionForTrustedAccessRoleBinding_STATUS interface (if implemented) to customize the conversion
	var bindingAsAny any = binding
	if augmentedBinding, ok := bindingAsAny.(augmentConversionForTrustedAccessRoleBinding_STATUS); ok {
		err := augmentedBinding.AssignPropertiesTo(destination)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

type augmentConversionForTrustedAccessRoleBinding_Spec interface {
	AssignPropertiesFrom(src *v20231001s.TrustedAccessRoleBinding_Spec) error
	AssignPropertiesTo(dst *v20231001s.TrustedAccessRoleBinding_Spec) error
}

type augmentConversionForTrustedAccessRoleBinding_STATUS interface {
	AssignPropertiesFrom(src *v20231001s.TrustedAccessRoleBinding_STATUS) error
	AssignPropertiesTo(dst *v20231001s.TrustedAccessRoleBinding_STATUS) error
}

// Storage version of v1api20240402preview.TrustedAccessRoleBindingOperatorSpec
// Details for configuring operator behavior. Fields in this struct are interpreted by the operator directly rather than being passed to Azure
type TrustedAccessRoleBindingOperatorSpec struct {
	ConfigMapExpressions []*core.DestinationExpression `json:"configMapExpressions,omitempty"`
	PropertyBag          genruntime.PropertyBag        `json:"$propertyBag,omitempty"`
	SecretExpressions    []*core.DestinationExpression `json:"secretExpressions,omitempty"`
}

// AssignProperties_From_TrustedAccessRoleBindingOperatorSpec populates our TrustedAccessRoleBindingOperatorSpec from the provided source TrustedAccessRoleBindingOperatorSpec
func (operator *TrustedAccessRoleBindingOperatorSpec) AssignProperties_From_TrustedAccessRoleBindingOperatorSpec(source *v20231001s.TrustedAccessRoleBindingOperatorSpec) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// ConfigMapExpressions
	if source.ConfigMapExpressions != nil {
		configMapExpressionList := make([]*core.DestinationExpression, len(source.ConfigMapExpressions))
		for configMapExpressionIndex, configMapExpressionItem := range source.ConfigMapExpressions {
			// Shadow the loop variable to avoid aliasing
			configMapExpressionItem := configMapExpressionItem
			if configMapExpressionItem != nil {
				configMapExpression := *configMapExpressionItem.DeepCopy()
				configMapExpressionList[configMapExpressionIndex] = &configMapExpression
			} else {
				configMapExpressionList[configMapExpressionIndex] = nil
			}
		}
		operator.ConfigMapExpressions = configMapExpressionList
	} else {
		operator.ConfigMapExpressions = nil
	}

	// SecretExpressions
	if source.SecretExpressions != nil {
		secretExpressionList := make([]*core.DestinationExpression, len(source.SecretExpressions))
		for secretExpressionIndex, secretExpressionItem := range source.SecretExpressions {
			// Shadow the loop variable to avoid aliasing
			secretExpressionItem := secretExpressionItem
			if secretExpressionItem != nil {
				secretExpression := *secretExpressionItem.DeepCopy()
				secretExpressionList[secretExpressionIndex] = &secretExpression
			} else {
				secretExpressionList[secretExpressionIndex] = nil
			}
		}
		operator.SecretExpressions = secretExpressionList
	} else {
		operator.SecretExpressions = nil
	}

	// Update the property bag
	if len(propertyBag) > 0 {
		operator.PropertyBag = propertyBag
	} else {
		operator.PropertyBag = nil
	}

	// Invoke the augmentConversionForTrustedAccessRoleBindingOperatorSpec interface (if implemented) to customize the conversion
	var operatorAsAny any = operator
	if augmentedOperator, ok := operatorAsAny.(augmentConversionForTrustedAccessRoleBindingOperatorSpec); ok {
		err := augmentedOperator.AssignPropertiesFrom(source)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_TrustedAccessRoleBindingOperatorSpec populates the provided destination TrustedAccessRoleBindingOperatorSpec from our TrustedAccessRoleBindingOperatorSpec
func (operator *TrustedAccessRoleBindingOperatorSpec) AssignProperties_To_TrustedAccessRoleBindingOperatorSpec(destination *v20231001s.TrustedAccessRoleBindingOperatorSpec) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(operator.PropertyBag)

	// ConfigMapExpressions
	if operator.ConfigMapExpressions != nil {
		configMapExpressionList := make([]*core.DestinationExpression, len(operator.ConfigMapExpressions))
		for configMapExpressionIndex, configMapExpressionItem := range operator.ConfigMapExpressions {
			// Shadow the loop variable to avoid aliasing
			configMapExpressionItem := configMapExpressionItem
			if configMapExpressionItem != nil {
				configMapExpression := *configMapExpressionItem.DeepCopy()
				configMapExpressionList[configMapExpressionIndex] = &configMapExpression
			} else {
				configMapExpressionList[configMapExpressionIndex] = nil
			}
		}
		destination.ConfigMapExpressions = configMapExpressionList
	} else {
		destination.ConfigMapExpressions = nil
	}

	// SecretExpressions
	if operator.SecretExpressions != nil {
		secretExpressionList := make([]*core.DestinationExpression, len(operator.SecretExpressions))
		for secretExpressionIndex, secretExpressionItem := range operator.SecretExpressions {
			// Shadow the loop variable to avoid aliasing
			secretExpressionItem := secretExpressionItem
			if secretExpressionItem != nil {
				secretExpression := *secretExpressionItem.DeepCopy()
				secretExpressionList[secretExpressionIndex] = &secretExpression
			} else {
				secretExpressionList[secretExpressionIndex] = nil
			}
		}
		destination.SecretExpressions = secretExpressionList
	} else {
		destination.SecretExpressions = nil
	}

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// Invoke the augmentConversionForTrustedAccessRoleBindingOperatorSpec interface (if implemented) to customize the conversion
	var operatorAsAny any = operator
	if augmentedOperator, ok := operatorAsAny.(augmentConversionForTrustedAccessRoleBindingOperatorSpec); ok {
		err := augmentedOperator.AssignPropertiesTo(destination)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

type augmentConversionForTrustedAccessRoleBindingOperatorSpec interface {
	AssignPropertiesFrom(src *v20231001s.TrustedAccessRoleBindingOperatorSpec) error
	AssignPropertiesTo(dst *v20231001s.TrustedAccessRoleBindingOperatorSpec) error
}

func init() {
	SchemeBuilder.Register(&TrustedAccessRoleBinding{}, &TrustedAccessRoleBindingList{})
}
