package main

import (
	"context"
	"fmt"
	"os"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"tempdir_tempfile": resourceTempfile(),
		},

		Schema: map[string]*schema.Schema{
			"path": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Location of temporary directory.",
				DefaultFunc: schema.EnvDefaultFunc("TMPDIR", "/tmp"),
			},
			"pattern": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Pattern used for files created.",
				Default:     "demo-*",
				ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
					if len(val.(string)) >= 16 {
						errs = append(errs, fmt.Errorf("%q must be less than 16 characters", key))
					}
					return
				},
			},
		},

		ConfigureContextFunc: providerConfigure,
	}
}

type Tempdir struct {
	Path    string
	Pattern string
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	dir_path := d.Get("path").(string)
	pattern := d.Get("pattern").(string)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	if _, err := os.Stat(dir_path); os.IsNotExist(err) {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Temporary directory must exist.",
			Detail:   fmt.Sprintf("Temporary directory %s must exist.", dir_path),
		})
	}

	if diags != nil {
		return nil, diags
	}

	client := Tempdir{
		Path:    dir_path,
		Pattern: pattern,
	}

	return client, nil
}
