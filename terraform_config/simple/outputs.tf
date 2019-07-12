// Copyright (c) 2018, Oracle and/or its affiliates. All rights reserved.

output "subnets" {
  description = "The Subnet(s) that have been created as a part of this module."
  value       = module.oci_subnets.subnets
}
