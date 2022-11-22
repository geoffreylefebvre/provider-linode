/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type RDNSObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type RDNSParameters struct {

	// The Public IPv4 or IPv6 address that will receive the PTR record.  A matching A or AAAA record must exist.
	// The public Linode IPv4 or IPv6 address to operate on.
	// +kubebuilder:validation:Required
	Address *string `json:"address" tf:"address,omitempty"`

	// The name of the RDNS address.
	// The reverse DNS assigned to this address. For public IPv4 addresses, this will be set to a default value provided by Linode if not explicitly set.
	// +kubebuilder:validation:Required
	Rdns *string `json:"rdns" tf:"rdns,omitempty"`

	// If true, the RDNS assignment will be retried within the operation timeout period.
	// If true, the RDNS assignment will be retried within the operation timeout period.
	// +kubebuilder:validation:Optional
	WaitForAvailable *bool `json:"waitForAvailable,omitempty" tf:"wait_for_available,omitempty"`
}

// RDNSSpec defines the desired state of RDNS
type RDNSSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RDNSParameters `json:"forProvider"`
}

// RDNSStatus defines the observed state of RDNS.
type RDNSStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RDNSObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// RDNS is the Schema for the RDNSs API. Manages the RDNS / PTR record for the IP Address associated with a Linode Instance.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,linode}
type RDNS struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RDNSSpec   `json:"spec"`
	Status            RDNSStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RDNSList contains a list of RDNSs
type RDNSList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RDNS `json:"items"`
}

// Repository type metadata.
var (
	RDNS_Kind             = "RDNS"
	RDNS_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RDNS_Kind}.String()
	RDNS_KindAPIVersion   = RDNS_Kind + "." + CRDGroupVersion.String()
	RDNS_GroupVersionKind = CRDGroupVersion.WithKind(RDNS_Kind)
)

func init() {
	SchemeBuilder.Register(&RDNS{}, &RDNSList{})
}
