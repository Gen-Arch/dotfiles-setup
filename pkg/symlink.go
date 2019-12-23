package dotfiles

import (
  "fmt"
  "os"
  "path/filepath"
)

func symlink_run() {
  rcfiles := rcs()
  for i := 0; i < len(rcfiles); i++ {
    symlink(rcfiles[i])
  }
}

func symlink(file string) {
  var symfile  string = filepath.Join(os.Getenv("HOME"), "." + file)
  var realfile string = filepath.Join(os.Getenv("HOME"), "dotfiles", "configs", file)

  if _, err := os.Stat(symfile); err == nil {
    os.Remove(symfile)
    fmt.Printf("remove symlink -> " + symfile + "\n")
  }

  os.Symlink(realfile, symfile)
  fmt.Printf("create symlink -> " + symfile + "\n")
}
