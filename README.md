# StackJanitor Terraform Provider

==================

<img src="https://cdn.rawgit.com/hashicorp/terraform-website/master/content/source/assets/images/logo-hashicorp.svg" width="600px">

## What is it?

This marks the terraform'd infrastructure as being a candidate for cleanup. 
It extends upon our existing work with StackJanitor (CloudFormation) to clean
up unused infra in AWS.

## Prerequisites

* Must have StackJanitor (version > x.x.x) installed in AWS Account
* S3 backend

## How it works

Setup this provider with the minimum config below

```bash
# For example, restrict template version in 0.1.x
provider "stack_janitor" {
  version = "~> 0.1"
}
```

Then have this configuration block:
```bash
janitor "my_stack" {
  location = "s3://terraform/project/subdirectory" # If not specified, uses the S3 backend of current TF config.
  clean_after = '7' # Number of days before we clean up
  retry_count = '3' # Number of retries before we give up cleaning
  providers = ['aws'] # You might have custom providers --> List them all here.
}
```

After 7 days, StackJanitor will grab the terraform state, initialise the
providers listed (via `terraform init`) and run
`terraform destroy -auto-approve`.

## Use cases

We deploy infrastructure from PR branches. The original use case for
StackJanitor applies here. We need a way to clean up our unused
terraform stacks.

## Failure cases

Say you deploy a static website in S3. TF can't delete the bucket if it's
not empty so the delete will fail. Manually delete all the S3 content
and rerun the delete.

Maintainers
-----------

@edify42

Requirements
------------

-	[Terraform](https://www.terraform.io/downloads.html) 0.10.x
-	[Go](https://golang.org/doc/install) 1.11 (to build the provider plugin)

Building The Provider
---------------------

Clone repository to: `$GOPATH/src/github.com/terraform-providers/terraform-provider-template`

```sh
$ mkdir -p $GOPATH/src/github.com/terraform-providers; cd $GOPATH/src/github.com/terraform-providers
$ git clone git@github.com:terraform-providers/terraform-provider-template
```

Enter the provider directory and build the provider

```sh
$ cd $GOPATH/src/github.com/terraform-providers/terraform-provider-template
$ make build
```

Using the provider
----------------------
## Fill in for each provider

Developing the Provider
---------------------------

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed on your machine (version 1.11+ is *required*). You'll also need to correctly setup a [GOPATH](http://golang.org/doc/code.html#GOPATH), as well as adding `$GOPATH/bin` to your `$PATH`.

To compile the provider, run `make build`. This will build the provider and put the provider binary in the `$GOPATH/bin` directory.

```sh
$ make build
...
$ $GOPATH/bin/terraform-provider-template
...
```

In order to test the provider, you can simply run `make test`.

```sh
$ make test
```

In order to run the full suite of Acceptance tests, run `make testacc`.

*Note:* Acceptance tests create real resources, and often cost money to run.

```sh
$ make testacc
```
