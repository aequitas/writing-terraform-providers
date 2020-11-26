terraform {
  required_providers {
    tempdir = {
      source = "custom.example.com/aequitas/tempdir"
    }
  }
}

variable "content" {
  default = "foo_bar"
}

resource "tempdir_tempfile" "example" {
  content = var.content
}

output "absolute_path" {
  value = tempdir_tempfile.example.path
}
