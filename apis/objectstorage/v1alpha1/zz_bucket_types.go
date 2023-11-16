// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

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

type BucketInitParameters struct {

	// The Access Control Level of the bucket using a canned ACL string. See all ACL strings in the Linode API v4 documentation.
	// The Access Control Level of the bucket using a canned ACL string.
	ACL *string `json:"acl,omitempty" tf:"acl,omitempty"`

	// The cert used by this Object Storage Bucket.
	Cert []CertInitParameters `json:"cert,omitempty" tf:"cert,omitempty"`

	// The cluster of the Linode Object Storage Bucket.
	// The cluster of the Linode Object Storage Bucket.
	Cluster *string `json:"cluster,omitempty" tf:"cluster,omitempty"`

	// If true, the bucket will have CORS enabled for all origins.
	// If true, the bucket will be created with CORS enabled for all origins.
	CorsEnabled *bool `json:"corsEnabled,omitempty" tf:"cors_enabled,omitempty"`

	// The label of the Linode Object Storage Bucket.
	// The label of the Linode Object Storage Bucket.
	Label *string `json:"label,omitempty" tf:"label,omitempty"`

	// Lifecycle rules to be applied to the bucket.
	LifecycleRule []LifecycleRuleInitParameters `json:"lifecycleRule,omitempty" tf:"lifecycle_rule,omitempty"`

	// Whether to enable versioning. Once you version-enable a bucket, it can never return to an unversioned state. You can, however, suspend versioning on that bucket. (Requires access_key and secret_key)
	// Whether to enable versioning.
	Versioning *bool `json:"versioning,omitempty" tf:"versioning,omitempty"`
}

type BucketObservation struct {

	// The Access Control Level of the bucket using a canned ACL string. See all ACL strings in the Linode API v4 documentation.
	// The Access Control Level of the bucket using a canned ACL string.
	ACL *string `json:"acl,omitempty" tf:"acl,omitempty"`

	// The access key to authenticate with.
	// The S3 access key to use for this resource. (Required for lifecycle_rule and versioning)
	AccessKey *string `json:"accessKey,omitempty" tf:"access_key,omitempty"`

	// The cert used by this Object Storage Bucket.
	Cert []CertParameters `json:"cert,omitempty" tf:"cert,omitempty"`

	// The cluster of the Linode Object Storage Bucket.
	// The cluster of the Linode Object Storage Bucket.
	Cluster *string `json:"cluster,omitempty" tf:"cluster,omitempty"`

	// If true, the bucket will have CORS enabled for all origins.
	// If true, the bucket will be created with CORS enabled for all origins.
	CorsEnabled *bool `json:"corsEnabled,omitempty" tf:"cors_enabled,omitempty"`

	// The hostname where this bucket can be accessed. This hostname can be accessed through a browser if the bucket is made public.
	Hostname *string `json:"hostname,omitempty" tf:"hostname,omitempty"`

	// The unique identifier for the rule.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The label of the Linode Object Storage Bucket.
	// The label of the Linode Object Storage Bucket.
	Label *string `json:"label,omitempty" tf:"label,omitempty"`

	// Lifecycle rules to be applied to the bucket.
	LifecycleRule []LifecycleRuleObservation `json:"lifecycleRule,omitempty" tf:"lifecycle_rule,omitempty"`

	// The secret key to authenticate with.
	// The S3 secret key to use for this resource. (Required for lifecycle_rule and versioning)
	SecretKey *string `json:"secretKey,omitempty" tf:"secret_key,omitempty"`

	// Whether to enable versioning. Once you version-enable a bucket, it can never return to an unversioned state. You can, however, suspend versioning on that bucket. (Requires access_key and secret_key)
	// Whether to enable versioning.
	Versioning *bool `json:"versioning,omitempty" tf:"versioning,omitempty"`
}

type BucketParameters struct {

	// The Access Control Level of the bucket using a canned ACL string. See all ACL strings in the Linode API v4 documentation.
	// The Access Control Level of the bucket using a canned ACL string.
	// +kubebuilder:validation:Optional
	ACL *string `json:"acl,omitempty" tf:"acl,omitempty"`

	// The access key to authenticate with.
	// The S3 access key to use for this resource. (Required for lifecycle_rule and versioning)
	// +crossplane:generate:reference:type=Key
	// +crossplane:generate:reference:refFieldName=AccessKeyRef
	// +crossplane:generate:reference:selectorFieldName=AccessKeySelector
	// +kubebuilder:validation:Optional
	AccessKey *string `json:"accessKey,omitempty" tf:"access_key,omitempty"`

	// Reference to a Key to populate accessKey.
	// +kubebuilder:validation:Optional
	AccessKeyRef *v1.Reference `json:"accessKeyRef,omitempty" tf:"-"`

	// Selector for a Key to populate accessKey.
	// +kubebuilder:validation:Optional
	AccessKeySelector *v1.Selector `json:"accessKeySelector,omitempty" tf:"-"`

	// The cert used by this Object Storage Bucket.
	// +kubebuilder:validation:Optional
	Cert []CertParameters `json:"cert,omitempty" tf:"cert,omitempty"`

	// The cluster of the Linode Object Storage Bucket.
	// The cluster of the Linode Object Storage Bucket.
	// +kubebuilder:validation:Optional
	Cluster *string `json:"cluster,omitempty" tf:"cluster,omitempty"`

	// If true, the bucket will have CORS enabled for all origins.
	// If true, the bucket will be created with CORS enabled for all origins.
	// +kubebuilder:validation:Optional
	CorsEnabled *bool `json:"corsEnabled,omitempty" tf:"cors_enabled,omitempty"`

	// The label of the Linode Object Storage Bucket.
	// The label of the Linode Object Storage Bucket.
	// +kubebuilder:validation:Optional
	Label *string `json:"label,omitempty" tf:"label,omitempty"`

	// Lifecycle rules to be applied to the bucket.
	// +kubebuilder:validation:Optional
	LifecycleRule []LifecycleRuleParameters `json:"lifecycleRule,omitempty" tf:"lifecycle_rule,omitempty"`

	// The secret key to authenticate with.
	// The S3 secret key to use for this resource. (Required for lifecycle_rule and versioning)
	// +crossplane:generate:reference:type=Key
	// +crossplane:generate:reference:refFieldName=SecretKeyRef
	// +crossplane:generate:reference:selectorFieldName=SecretKeySelector
	// +kubebuilder:validation:Optional
	SecretKey *string `json:"secretKey,omitempty" tf:"secret_key,omitempty"`

	// Reference to a Key to populate secretKey.
	// +kubebuilder:validation:Optional
	SecretKeyRef *v1.Reference `json:"secretKeyRef,omitempty" tf:"-"`

	// Selector for a Key to populate secretKey.
	// +kubebuilder:validation:Optional
	SecretKeySelector *v1.Selector `json:"secretKeySelector,omitempty" tf:"-"`

	// Whether to enable versioning. Once you version-enable a bucket, it can never return to an unversioned state. You can, however, suspend versioning on that bucket. (Requires access_key and secret_key)
	// Whether to enable versioning.
	// +kubebuilder:validation:Optional
	Versioning *bool `json:"versioning,omitempty" tf:"versioning,omitempty"`
}

type CertInitParameters struct {
}

type CertObservation struct {
}

type CertParameters struct {

	// The Base64 encoded and PEM formatted SSL certificate.
	// The Base64 encoded and PEM formatted SSL certificate.
	// +kubebuilder:validation:Required
	CertificateSecretRef v1.SecretKeySelector `json:"certificateSecretRef" tf:"-"`

	// The private key associated with the TLS/SSL certificate.
	// The private key associated with the TLS/SSL certificate.
	// +kubebuilder:validation:Required
	PrivateKeySecretRef v1.SecretKeySelector `json:"privateKeySecretRef" tf:"-"`
}

type ExpirationInitParameters struct {

	// Specifies the date after which you want the corresponding action to take effect.
	// Specifies the date after which you want the corresponding action to take effect.
	Date *string `json:"date,omitempty" tf:"date,omitempty"`

	// Specifies the number of days after object creation when the specific rule action takes effect.
	// Specifies the number of days after object creation when the specific rule action takes effect.
	Days *int64 `json:"days,omitempty" tf:"days,omitempty"`

	// On a versioned bucket (versioning-enabled or versioning-suspended bucket), you can add this element in the lifecycle configuration to direct Linode Object Storage to delete expired object delete markers. This cannot be specified with Days or Date in a Lifecycle Expiration Policy.
	// Directs Linode Object Storage to remove expired deleted markers.
	ExpiredObjectDeleteMarker *bool `json:"expiredObjectDeleteMarker,omitempty" tf:"expired_object_delete_marker,omitempty"`
}

type ExpirationObservation struct {

	// Specifies the date after which you want the corresponding action to take effect.
	// Specifies the date after which you want the corresponding action to take effect.
	Date *string `json:"date,omitempty" tf:"date,omitempty"`

	// Specifies the number of days after object creation when the specific rule action takes effect.
	// Specifies the number of days after object creation when the specific rule action takes effect.
	Days *int64 `json:"days,omitempty" tf:"days,omitempty"`

	// On a versioned bucket (versioning-enabled or versioning-suspended bucket), you can add this element in the lifecycle configuration to direct Linode Object Storage to delete expired object delete markers. This cannot be specified with Days or Date in a Lifecycle Expiration Policy.
	// Directs Linode Object Storage to remove expired deleted markers.
	ExpiredObjectDeleteMarker *bool `json:"expiredObjectDeleteMarker,omitempty" tf:"expired_object_delete_marker,omitempty"`
}

type ExpirationParameters struct {

	// Specifies the date after which you want the corresponding action to take effect.
	// Specifies the date after which you want the corresponding action to take effect.
	// +kubebuilder:validation:Optional
	Date *string `json:"date,omitempty" tf:"date,omitempty"`

	// Specifies the number of days after object creation when the specific rule action takes effect.
	// Specifies the number of days after object creation when the specific rule action takes effect.
	// +kubebuilder:validation:Optional
	Days *int64 `json:"days,omitempty" tf:"days,omitempty"`

	// On a versioned bucket (versioning-enabled or versioning-suspended bucket), you can add this element in the lifecycle configuration to direct Linode Object Storage to delete expired object delete markers. This cannot be specified with Days or Date in a Lifecycle Expiration Policy.
	// Directs Linode Object Storage to remove expired deleted markers.
	// +kubebuilder:validation:Optional
	ExpiredObjectDeleteMarker *bool `json:"expiredObjectDeleteMarker,omitempty" tf:"expired_object_delete_marker,omitempty"`
}

type LifecycleRuleInitParameters struct {

	// Specifies the number of days after initiating a multipart upload when the multipart upload must be completed.
	// Specifies the number of days after initiating a multipart upload when the multipart upload must be completed.
	AbortIncompleteMultipartUploadDays *int64 `json:"abortIncompleteMultipartUploadDays,omitempty" tf:"abort_incomplete_multipart_upload_days,omitempty"`

	// Specifies whether the lifecycle rule is active.
	// Specifies whether the lifecycle rule is active.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Specifies a period in the object's expire.
	Expiration []ExpirationInitParameters `json:"expiration,omitempty" tf:"expiration,omitempty"`

	// The unique identifier for the rule.
	// The unique identifier for the rule.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies when non-current object versions expire.
	NoncurrentVersionExpiration []NoncurrentVersionExpirationInitParameters `json:"noncurrentVersionExpiration,omitempty" tf:"noncurrent_version_expiration,omitempty"`

	// The object key prefix identifying one or more objects to which the rule applies.
	// The object key prefix identifying one or more objects to which the rule applies.
	Prefix *string `json:"prefix,omitempty" tf:"prefix,omitempty"`
}

type LifecycleRuleObservation struct {

	// Specifies the number of days after initiating a multipart upload when the multipart upload must be completed.
	// Specifies the number of days after initiating a multipart upload when the multipart upload must be completed.
	AbortIncompleteMultipartUploadDays *int64 `json:"abortIncompleteMultipartUploadDays,omitempty" tf:"abort_incomplete_multipart_upload_days,omitempty"`

	// Specifies whether the lifecycle rule is active.
	// Specifies whether the lifecycle rule is active.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Specifies a period in the object's expire.
	Expiration []ExpirationObservation `json:"expiration,omitempty" tf:"expiration,omitempty"`

	// The unique identifier for the rule.
	// The unique identifier for the rule.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies when non-current object versions expire.
	NoncurrentVersionExpiration []NoncurrentVersionExpirationObservation `json:"noncurrentVersionExpiration,omitempty" tf:"noncurrent_version_expiration,omitempty"`

	// The object key prefix identifying one or more objects to which the rule applies.
	// The object key prefix identifying one or more objects to which the rule applies.
	Prefix *string `json:"prefix,omitempty" tf:"prefix,omitempty"`
}

type LifecycleRuleParameters struct {

	// Specifies the number of days after initiating a multipart upload when the multipart upload must be completed.
	// Specifies the number of days after initiating a multipart upload when the multipart upload must be completed.
	// +kubebuilder:validation:Optional
	AbortIncompleteMultipartUploadDays *int64 `json:"abortIncompleteMultipartUploadDays,omitempty" tf:"abort_incomplete_multipart_upload_days,omitempty"`

	// Specifies whether the lifecycle rule is active.
	// Specifies whether the lifecycle rule is active.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled" tf:"enabled,omitempty"`

	// Specifies a period in the object's expire.
	// +kubebuilder:validation:Optional
	Expiration []ExpirationParameters `json:"expiration,omitempty" tf:"expiration,omitempty"`

	// The unique identifier for the rule.
	// The unique identifier for the rule.
	// +kubebuilder:validation:Optional
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies when non-current object versions expire.
	// +kubebuilder:validation:Optional
	NoncurrentVersionExpiration []NoncurrentVersionExpirationParameters `json:"noncurrentVersionExpiration,omitempty" tf:"noncurrent_version_expiration,omitempty"`

	// The object key prefix identifying one or more objects to which the rule applies.
	// The object key prefix identifying one or more objects to which the rule applies.
	// +kubebuilder:validation:Optional
	Prefix *string `json:"prefix,omitempty" tf:"prefix,omitempty"`
}

type NoncurrentVersionExpirationInitParameters struct {

	// Specifies the number of days after object creation when the specific rule action takes effect.
	// Specifies the number of days non-current object versions expire.
	Days *int64 `json:"days,omitempty" tf:"days,omitempty"`
}

type NoncurrentVersionExpirationObservation struct {

	// Specifies the number of days after object creation when the specific rule action takes effect.
	// Specifies the number of days non-current object versions expire.
	Days *int64 `json:"days,omitempty" tf:"days,omitempty"`
}

type NoncurrentVersionExpirationParameters struct {

	// Specifies the number of days after object creation when the specific rule action takes effect.
	// Specifies the number of days non-current object versions expire.
	// +kubebuilder:validation:Optional
	Days *int64 `json:"days" tf:"days,omitempty"`
}

// BucketSpec defines the desired state of Bucket
type BucketSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     BucketParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider BucketInitParameters `json:"initProvider,omitempty"`
}

// BucketStatus defines the observed state of Bucket.
type BucketStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        BucketObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Bucket is the Schema for the Buckets API. Manages a Linode Object Storage Bucket.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,linode}
type Bucket struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.cluster) || (has(self.initProvider) && has(self.initProvider.cluster))",message="spec.forProvider.cluster is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.label) || (has(self.initProvider) && has(self.initProvider.label))",message="spec.forProvider.label is a required parameter"
	Spec   BucketSpec   `json:"spec"`
	Status BucketStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BucketList contains a list of Buckets
type BucketList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Bucket `json:"items"`
}

// Repository type metadata.
var (
	Bucket_Kind             = "Bucket"
	Bucket_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Bucket_Kind}.String()
	Bucket_KindAPIVersion   = Bucket_Kind + "." + CRDGroupVersion.String()
	Bucket_GroupVersionKind = CRDGroupVersion.WithKind(Bucket_Kind)
)

func init() {
	SchemeBuilder.Register(&Bucket{}, &BucketList{})
}
