package main

import (
	"regexp"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"testing"
)

func TestAccResourceTempdirTempfile(t *testing.T) {
	testConfig := `
  provider "tempdir" {
    path = "/tmp"
    pattern = "test-*.tmp" 
  }
  
	resource "tempdir_tempfile" "test" {
    content = "pannenkoek met stroop"
	}`

	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr("tempdir_tempfile.test", "path", regexp.MustCompile("/tmp/test-[0-9]+\\.tmp")),
				),
			},
		},
	})
}

func TestAccResourceTempdirTempfileUpdate(t *testing.T) {
	testConfig := `
	resource "tempdir_tempfile" "test" {
    content = "pannenkoek met stroop"
	}`

	testConfig2 := `
	resource "tempdir_tempfile" "test" {
    content = "pannenkoek met stroop en poedersuiker"
	}`

	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("tempdir_tempfile.test", "content", "pannenkoek met stroop"),
				),
			},
			{
				Config: testConfig2,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("tempdir_tempfile.test", "content", "pannenkoek met stroop en poedersuiker"),
				),
			},
		},
	})
}
