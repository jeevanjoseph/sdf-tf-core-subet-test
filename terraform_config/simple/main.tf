// Copyright (c) 2018, Oracle and/or its affiliates. All rights reserved.

module "oci_subnets" {
  source           = "../../../sdf-tf-core-subnet/"
  #source          = "oracle-terraform-modules/default-vcn/oci"
  
  default_compartment_id  = var.default_compartment_id
  # vcn_id = data.terraform_remote_state.network.outputs.vcn.id
  vcn_id                  = oci_core_vcn.this.id
  vcn_cidr                = oci_core_vcn.this.cidr_block
  
  subnets = {
    test1 = {
      compartment_id    = null
      dynamic_cidr      = false
      cidr              = "192.168.0.0/30"
      cidr_len          = null
      cidr_num          = null
      enable_dns        = true
      dns_label         = "test1"
      private           = true
      ad                = null
      dhcp_options_id   = null
      route_table_id    = null
      security_list_ids = null
    },
    test2 = {
      compartment_id    = null
      dynamic_cidr      = false
      cidr              = "192.168.0.4/30"
      cidr_len          = null
      cidr_num          = null
      enable_dns        = true
      dns_label         = "test2"
      private           = true
      ad                = 0
      dhcp_options_id   = null
      route_table_id    = null
      security_list_ids = null
    },
    test3 = {
      compartment_id    = null
      dynamic_cidr      = false
      cidr              = "192.168.0.8/30"
      cidr_len          = null
      cidr_num          = null
      enable_dns        = false
      dns_label         = null
      private           = true
      ad                = null
      dhcp_options_id   = null
      route_table_id    = null
      security_list_ids = null
    },
    test4 = {
      compartment_id    = null
      dynamic_cidr      = false
      cidr              = "192.168.0.12/30"
      cidr_len          = null
      cidr_num          = null
      enable_dns        = false
      dns_label         = "test4"
      private           = true
      ad                = null
      dhcp_options_id   = null
      route_table_id    = null
      security_list_ids = null
    },
    test5 = {
      compartment_id    = null
      dynamic_cidr      = false
      cidr              = "192.168.0.16/30"
      cidr_len          = null
      cidr_num          = null
      enable_dns        = true
      dns_label         = null
      private           = false
      ad                = null
      dhcp_options_id   = null
      route_table_id    = null
      security_list_ids = null
    }
  }
}

resource "oci_core_vcn" "this" {
  dns_label      = "terravcn"
  cidr_block     = "192.168.0.0/16"
  compartment_id = var.default_compartment_id
  display_name   = "terravcn"
}
