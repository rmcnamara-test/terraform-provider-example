package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() *schema.Provider {
			return &schema.Provider{
				DataSourcesMap: map[string]*schema.Resource{
					"example": dataSourceHelloWorld(),
				},
			}
		},
	})
}

func dataSourceHelloWorld() *schema.Resource {
	return &schema.Resource{
		Read: func(d *schema.ResourceData, m interface{}) error {
			return nil
		},
	}
}
