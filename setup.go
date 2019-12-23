package main

import (
  "fmt"
  dotfiles "./pkg"
)

func main() {
  dotfiles.cli()
  dotfiles.git_run()
  dotfiles.symlink_run()
}
