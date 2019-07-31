package network_module_helpers

import (
	"context"
	"fmt"

	"github.com/oracle/oci-go-sdk/common"
	"github.com/oracle/oci-go-sdk/core"
)

func CreateVNCClient() (context.Context, core.VirtualNetworkClient) {
	ctx := context.Background()
	configProvider := common.DefaultConfigProvider()
	fmt.Println(configProvider)
	c, err := core.NewVirtualNetworkClientWithConfigurationProvider(configProvider)
	if err != nil {
		fmt.Println(err)
	}
	return ctx, c
}

func ListVCN(ctx context.Context, c core.VirtualNetworkClient, compartment_id string) []core.Vcn {
	request := core.ListVcnsRequest{
		CompartmentId: common.String(compartment_id),
	}
	r, err := c.ListVcns(ctx, request)
	if err != nil {
		fmt.Println(err)
	}
	return r.Items
}

func ListSubnets(ctx context.Context, c core.VirtualNetworkClient, compartment_id string, vcn_id string) []core.Subnet {
	request := core.ListSubnetsRequest{
		CompartmentId: common.String(compartment_id),
		VcnId:         common.String(vcn_id),
	}
	r, err := c.ListSubnets(ctx, request)
	if err != nil {
		fmt.Println(err)
	}
	return r.Items
}
