// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

// Package apis contains Kubernetes API for the provider.
package apis

import (
	"k8s.io/apimachinery/pkg/runtime"

	v1alpha1 "github.com/linode/provider-linode/apis/database/v1alpha1"
	v1alpha1domain "github.com/linode/provider-linode/apis/domain/v1alpha1"
	v1alpha1firewall "github.com/linode/provider-linode/apis/firewall/v1alpha1"
	v1alpha1image "github.com/linode/provider-linode/apis/image/v1alpha1"
	v1alpha1instance "github.com/linode/provider-linode/apis/instance/v1alpha1"
	v1alpha1ipv6 "github.com/linode/provider-linode/apis/ipv6/v1alpha1"
	v1alpha1lke "github.com/linode/provider-linode/apis/lke/v1alpha1"
	v1alpha1nodebalancer "github.com/linode/provider-linode/apis/nodebalancer/v1alpha1"
	v1alpha1objectstorage "github.com/linode/provider-linode/apis/objectstorage/v1alpha1"
	v1alpha1rdns "github.com/linode/provider-linode/apis/rdns/v1alpha1"
	v1alpha1sshkey "github.com/linode/provider-linode/apis/sshkey/v1alpha1"
	v1alpha1stackscript "github.com/linode/provider-linode/apis/stackscript/v1alpha1"
	v1alpha1token "github.com/linode/provider-linode/apis/token/v1alpha1"
	v1alpha1user "github.com/linode/provider-linode/apis/user/v1alpha1"
	v1alpha1apis "github.com/linode/provider-linode/apis/v1alpha1"
	v1beta1 "github.com/linode/provider-linode/apis/v1beta1"
	v1alpha1volume "github.com/linode/provider-linode/apis/volume/v1alpha1"
)

func init() {
	// Register the types with the Scheme so the components can map objects to GroupVersionKinds and back
	AddToSchemes = append(AddToSchemes,
		v1alpha1.SchemeBuilder.AddToScheme,
		v1alpha1domain.SchemeBuilder.AddToScheme,
		v1alpha1firewall.SchemeBuilder.AddToScheme,
		v1alpha1image.SchemeBuilder.AddToScheme,
		v1alpha1instance.SchemeBuilder.AddToScheme,
		v1alpha1ipv6.SchemeBuilder.AddToScheme,
		v1alpha1lke.SchemeBuilder.AddToScheme,
		v1alpha1nodebalancer.SchemeBuilder.AddToScheme,
		v1alpha1objectstorage.SchemeBuilder.AddToScheme,
		v1alpha1rdns.SchemeBuilder.AddToScheme,
		v1alpha1sshkey.SchemeBuilder.AddToScheme,
		v1alpha1stackscript.SchemeBuilder.AddToScheme,
		v1alpha1token.SchemeBuilder.AddToScheme,
		v1alpha1user.SchemeBuilder.AddToScheme,
		v1alpha1apis.SchemeBuilder.AddToScheme,
		v1beta1.SchemeBuilder.AddToScheme,
		v1alpha1volume.SchemeBuilder.AddToScheme,
	)
}

// AddToSchemes may be used to add all resources defined in the project to a Scheme
var AddToSchemes runtime.SchemeBuilder

// AddToScheme adds all Resources to the Scheme
func AddToScheme(s *runtime.Scheme) error {
	return AddToSchemes.AddToScheme(s)
}
