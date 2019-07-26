package test

import (
	"os"
	// "src/modules/logger"
	"time"

	"testing"

	"../network_module_helpers"

	test_helper "../terraform-module-test-lib"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/oracle/oci-go-sdk/core"
	"github.com/stretchr/testify/assert"
)

// An example of how to test the simple Terraform module in examples/terraform-basic-example using Terratest.

var vcnName string
var expectedCidr string
var expectedVcnLabel string
var vcnId string
var expectedSubnets int

func TestSimple(t *testing.T) {
	t.Parallel()

	terraformDir := "../../sdf-tf-core-subnet/examples/simple"
	terraformOptions := configureTerraformOptions(t, terraformDir)

	var vars Inputs
	err := test_helper.GetConfig("inputs_config.json", &vars)
	if err != nil {
		// logger.Logf(t, err.Error())
		t.Fail()
	}

	vcnName = vars.Vcn_display_name
	expectedCidr = vars.Vcn_cidr
	expectedVcnLabel = vars.Vcn_label
	expectedSubnets = vars.Vcn_expectedSubnets

	defer terraform.Destroy(t, terraformOptions)
	terraform.Init(t, terraformOptions)
	terraform.Apply(t, terraformOptions)
	time.Sleep(60)
	startTests(t, terraformOptions)

}

func startTests(t *testing.T, terraformOptions *terraform.Options) {

	//The below is validation by Go SDK
	context, client := network_module_helpers.CreateVNCClient()
	var id string
	id = os.Getenv("TF_VAR_compartment_id")

	// Using Helpers:
	vcns := network_module_helpers.ListVCN(context, client, id)
	//fmt.Println(vcns)
	testVcn(t, vcns)

	subnets := network_module_helpers.ListSubnets(context, client, id, vcnId)
	//fmt.Println(subnets)
	testSubnet(t, subnets)
}

func testVcn(t *testing.T, vcns []core.Vcn) {

	vcnCidr := map[string]string{}
	vcnLabel := map[string]string{}
	for _, v := range vcns {
		if *v.DisplayName == vcnName {
			if v.LifecycleState == "AVAILABLE" {
				vcnCidr[*v.DisplayName] = *v.CidrBlock
				vcnLabel[*v.DisplayName] = *v.DnsLabel
				vcnId = *v.Id
				//fmt.Println("vcnID   : ", vcnId)
			}
		}
	}

	assert.Equal(t, 1, len(vcnCidr), "Length of vcn cidr should be 1")
	assert.Equal(t, expectedCidr, vcnCidr[vcnName], "vcn cidr should be equal to "+expectedCidr)
	assert.Equal(t, expectedVcnLabel, vcnLabel[vcnName], "vcn label should be equal to "+expectedVcnLabel)
	assert.NotEmpty(t, vcnId, "vcn ociId should not be null")

}

func testSubnet(t *testing.T, subnets []core.Subnet) {
	assert.Equal(t, expectedSubnets, len(subnets), "Length of vcn subnets should be ", expectedSubnets)
}

func configureTerraformOptions(t *testing.T, terraformDir string) *terraform.Options {

	terraformOptions := &terraform.Options{
		TerraformDir: "../../sdf-tf-core-subnet/examples/simple",
		Vars: map[string]interface{}{
			"default_compartment_id": os.Getenv("TF_VAR_compartment_id"),
			"tenancy_id":             os.Getenv("TF_VAR_tenancy_id"),
			"user_id":                os.Getenv("TF_VAR_user_id"),
			"region":                 os.Getenv("TF_VAR_region"),
		},
		Upgrade: true,
	}
	return terraformOptions
}
