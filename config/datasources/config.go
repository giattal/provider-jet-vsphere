package datasources

import (
	"github.com/crossplane/terrajet/pkg/config"

	"github.com/giattal/provider-jet-vsphere/config/common"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("vsphere_host", func(r *config.Resource) {
		r.Version = common.VersionV1Alpha2
		r.ExternalName = config.IdentifierFromProvider
		r.References["name"] = config.Reference{
			Type: "Name",
		}
		r.References["datacenter_id"] = config.Reference{
			Type:              "DatacenterId",
			RefFieldName:      "DatacenterIdRefs",
			SelectorFieldName: "DatacenterIdSelector",
		}
		r.References["id"] = config.Reference{
			Type: "Id",
		}
		r.References["resource_pool_id"] = config.Reference{
			Type: "ResourcePoolId",
		}
	})
}
