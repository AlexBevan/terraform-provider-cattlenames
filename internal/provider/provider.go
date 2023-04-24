package provider

import (
	"terraform-provider-cattlenames/internal/resource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"cattlenames": resource.ResourceCattleNames(),
		},
	}
}