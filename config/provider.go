/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/upbound/upjet/pkg/config"

	"github.com/linode/provider-linode/config/domain"
	"github.com/linode/provider-linode/config/domain_record"
	"github.com/linode/provider-linode/config/firewall"
	"github.com/linode/provider-linode/config/firewall_device"
	"github.com/linode/provider-linode/config/image"
	"github.com/linode/provider-linode/config/instance"
	"github.com/linode/provider-linode/config/instance_config"
	"github.com/linode/provider-linode/config/instance_disk"
	"github.com/linode/provider-linode/config/instance_ip"
	"github.com/linode/provider-linode/config/instance_shared_ips"
	"github.com/linode/provider-linode/config/lke_cluster"
	"github.com/linode/provider-linode/config/nodebalancer"
	"github.com/linode/provider-linode/config/nodebalancer_config"
	"github.com/linode/provider-linode/config/nodebalancer_node"
	"github.com/linode/provider-linode/config/object_storage_bucket"
	"github.com/linode/provider-linode/config/object_storage_key"
	"github.com/linode/provider-linode/config/object_storage_object"
	"github.com/linode/provider-linode/config/stackscript"
)

const (
	resourcePrefix = "linode"
	modulePath     = "github.com/linode/provider-linode"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		domain.Configure,
		domain_record.Configure,
		firewall.Configure,
		firewall_device.Configure,
		image.Configure,
		instance.Configure,
		instance_config.Configure,
		instance_disk.Configure,
		instance_ip.Configure,
		instance_shared_ips.Configure,
		lke_cluster.Configure,
		nodebalancer.Configure,
		nodebalancer_config.Configure,
		nodebalancer_node.Configure,
		object_storage_bucket.Configure,
		object_storage_key.Configure,
		object_storage_object.Configure,
		stackscript.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
