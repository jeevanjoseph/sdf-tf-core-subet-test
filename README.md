# sdf-tf-core-subnet-test

Testing for sdf-tf-core-subnet module

### Set the ENV variables:
```
export TF_VAR_tenancy_id=ocid1.tenancy.xxx
export TF_VAR_compartment_id=ocid1.compartment.xxx
export TF_VAR_user_id=ocid1.user.xxx
export TF_VAR_fingerprint=a3:03:xx:xx
export TF_VAR_private_key_path=/Users/xx/.oci/oci_api_key.pem
export TF_VAR_region=us-phoenix-1
```

```
$ mkdir test; cd test
$ git clone git@orahub.oraclecorp.com:ateam/sdf-tf-core-subnet.git
$ git clone git@orahub.oraclecorp.com:ateam/sdf-tf-core-subnet-test.git
$ cd sdf-tf-core-subnet-test
$ cd Simple_test
$ go test -run TestSimple
```