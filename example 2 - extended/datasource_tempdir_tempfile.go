package main

import (
	"fmt"
	"io/ioutil"
	"path"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceTempfile() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceTempfileRead,

		Schema: map[string]*schema.Schema{
			"number": {
				Type:        schema.TypeString,
				Description: "Random number of temporary file.",
				Optional:    true,
			},
			"match_content": {
				Type:        schema.TypeString,
				Description: "Regex to match content.",
				Optional:    true,
			},

			"content": {
				Type:        schema.TypeString,
				Description: "Content of the file.",
				Computed:    true,
			},
			"path": {
				Type:        schema.TypeString,
				Description: "Absolute path of the file.",
				Computed:    true,
			},
		},
	}
}

func dataSourceTempfileRead(d *schema.ResourceData, m interface{}) error {
	number := d.Get("number").(string)
	id := fmt.Sprintf("demo-%s", number)

	temporary_directory := m.(string)

	file_path := path.Join(temporary_directory, id)

	var content string
	file_bytes, err := ioutil.ReadFile(file_path)
	if err != nil {
		return fmt.Errorf("could not find file with number %s: %s", number, err)
	}

	// else if match_content, for file in tmpdir, if file_content match match_content, content = file_content

	content = string(file_bytes)

	d.SetId(id)

	d.Set("content", content)
	d.Set("path", file_path)

	return nil
}
