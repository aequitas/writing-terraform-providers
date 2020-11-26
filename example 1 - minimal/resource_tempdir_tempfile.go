package main

import (
	"io/ioutil"
	"os"
	"path"
	"path/filepath"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceTempfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceTempfileCreate,
		Read:   resourceTempfileRead,
		Delete: resourceTempfileDelete,

		Schema: map[string]*schema.Schema{
			// arguments/inputs
			"content": {
				Type:        schema.TypeString,
				Description: "Content of the file.",
				Required:    true,
				ForceNew:    true,
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

func resourceTempfileCreate(d *schema.ResourceData, m interface{}) error {
	content := d.Get("content").(string)

	file, _ := ioutil.TempFile("/tmp", "demo-*")
	ioutil.WriteFile(file.Name(), []byte(content), 0644)

	id := filepath.Base(file.Name()) // demo-XXXXXXXX
	file_path := file.Name()         // /tmp/demo-XXXXXXXX

	d.SetId(id)
	d.Set("path", file_path)

	return nil
}

func resourceTempfileRead(d *schema.ResourceData, m interface{}) error {
	id := d.Id()

	file_path := path.Join("/tmp", id) // /tmp/demo-XXXXXXXX

	file_bytes, _ := ioutil.ReadFile(file_path)
	content := string(file_bytes)

	d.Set("content", content)
	d.Set("path", file_path)

	return nil
}

func resourceTempfileDelete(d *schema.ResourceData, m interface{}) error {
	id := d.Id()

	file_path := path.Join("/tmp", id)

	os.Remove(file_path)

	return nil
}
