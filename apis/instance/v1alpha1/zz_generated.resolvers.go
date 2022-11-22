/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	v1alpha1 "github.com/linode/provider-linode/apis/image/v1alpha1"
	v1alpha11 "github.com/linode/provider-linode/apis/rdns/v1alpha1"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this Disk.
func (mg *Disk) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Image),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ImageRef,
		Selector:     mg.Spec.ForProvider.ImageSelector,
		To: reference.To{
			List:    &v1alpha1.ImageList{},
			Managed: &v1alpha1.Image{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Image")
	}
	mg.Spec.ForProvider.Image = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ImageRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this IP.
func (mg *IP) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Rdns),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.RdnsRef,
		Selector:     mg.Spec.ForProvider.RdnsSelector,
		To: reference.To{
			List:    &v1alpha11.RDNSList{},
			Managed: &v1alpha11.RDNS{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Rdns")
	}
	mg.Spec.ForProvider.Rdns = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RdnsRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Instance.
func (mg *Instance) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Image),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ImageRef,
		Selector:     mg.Spec.ForProvider.ImageSelector,
		To: reference.To{
			List:    &v1alpha1.ImageList{},
			Managed: &v1alpha1.Image{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Image")
	}
	mg.Spec.ForProvider.Image = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ImageRef = rsp.ResolvedReference

	return nil
}
