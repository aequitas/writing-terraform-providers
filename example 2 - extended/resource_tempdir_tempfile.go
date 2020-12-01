package main

import (
	"context"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceTempfile() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceTempfileCreate,
		ReadContext:   resourceTempfileRead,
		UpdateContext: resourceTempfileUpdate,
		DeleteContext: resourceTempfileDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			// arguments/inputs
			"content": {
				Type:         schema.TypeString,
				Description:  "Content of the file.",
				Optional:     true,
				ForceNew:     false,
				ExactlyOneOf: []string{"content", "empty"},
				// ConflictsWith
				// AtLeastOneOf
				// RequiredWith
			},

			"empty": {
				Type:        schema.TypeBool,
				Description: "Create empty tempfile",
				Optional:    true,
				ForceNew:    false,
			},

			// attributes/outputs
			"path": {
				Type:        schema.TypeString,
				Description: "Absolute path of the file.",
				Computed:    true,
			},
		},
	}
}

func resourceTempfileCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	content := d.Get("content").(string)
	empty := d.Get("empty").(bool)

	client := m.(Tempdir)
	dir_path := client.Path
	pattern := client.Pattern

	file, err := ioutil.TempFile(dir_path, pattern)
	if err != nil {
		return diag.FromErr(err)
	}

	id := filepath.Base(file.Name()) // demo-XXXXXXXX
	file_path := file.Name()

	if empty {
		content = ""
	}

	err = ioutil.WriteFile(file_path, []byte(content), 0644)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(id)

	return resourceTempfileRead(ctx, d, m)
}

func resourceTempfileRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	id := d.Id()

	client := m.(Tempdir)
	dir_path := client.Path

	file_path := path.Join(dir_path, id)

	var content string
	if _, err := os.Stat(file_path); !os.IsNotExist(err) {
		file_bytes, _ := ioutil.ReadFile(file_path)
		content = string(file_bytes)
	} else {
		content = ""
		file_path = ""
	}

	d.Set("content", content)
	d.Set("path", file_path)

	return nil
}

func resourceTempfileUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	id := d.Id()
	content := d.Get("content").(string)
	empty := d.Get("empty").(bool)

	client := m.(Tempdir)
	dir_path := client.Path

	file_path := path.Join(dir_path, id)

	if empty {
		content = ""
	}

	err := ioutil.WriteFile(file_path, []byte(content), 0644)
	if err != nil {
		return diag.FromErr(err)
	}

	return resourceTempfileRead(ctx, d, m)
}

func resourceTempfileDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	id := d.Id()

	client := m.(Tempdir)
	dir_path := client.Path

	file_path := path.Join(dir_path, id)

	os.Remove(file_path)

	return nil
}
