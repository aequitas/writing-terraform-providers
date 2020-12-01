terraform {
  required_providers {
    tempdir = {
      source = "custom.example.com/aequitas/tempdir"
    }
  }
}

provider "tempdir" {
  pattern = "test-*.tmp"
}

resource "tempdir_tempfile" "example" {
  content = "foo_bar"
}

output "absolute_path" {
  value = tempdir_tempfile.example.path
}
