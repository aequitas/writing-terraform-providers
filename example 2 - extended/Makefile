.PHONY: build install init apply destroy test mrproper

registry = custom.example.com
namespace = aequitas
name = tempdir
version = 0.0.1

os = darwin
arch = amd64

provider_name = terraform-provider-${name}
dist_name = ${provider_name}_v${version}

tmpdir ?= ${TMPDIR}

# build provider binary
build: ${provider_name}
${provider_name}: $(shell find . -name "*.go")
	-goimports -w .
	gofmt -w .
	go build .

# copy provider binary to location expected by Terraform
# https://www.terraform.io/upgrade-guides/0-13.html#new-filesystem-layout-for-local-copies-of-providers
install: terraform.d/plugins/${registry}/${namespace}/tempdir/${version}/${os}_${arch}/${dist_name}
terraform.d/plugins/${registry}/${namespace}/tempdir/${version}/${os}_${arch}/${dist_name}: ${provider_name}
	mkdir -p ${@D}
	cp $< $@

# initialize Terraform for current project
init: .terraform/plugins/selections.json
.terraform/plugins/selections.json: install
	terraform init

plan: init
	terraform plan

apply: init
	echo yes | terraform apply

destroy: init
	echo yes | terraform destroy

import: init
	echo foo_bar > ${tmpdir}/test-421337.tmp
	terraform import tempdir_tempfile.example test-421337.tmp

test: build
	go test -v github.com/aequitas/terraform-provider-tempdir

test_acc: build
	TF_ACC=1 go test -v github.com/aequitas/terraform-provider-tempdir

# cleanup provisioned resources and Terraform state
clean:
	rm -rf terraform.tfstate terraform.tfstate.backup ${tmpdir}/test-*.tmp

# prestine condition, cleanup and also remove all build artifacts and Terraform cache
mrproper: clean
	rm -fr terraform.d .terraform ${provider_name}