package main

import (
  "fmt"
  "os"
  "path/filepath"

  "gopkg.in/src-d/go-git.v4"
)

type Gits struct {
  url string
  dir string
}

func  git_clone(repo Gits) string{
  var err error
  fmt.Printf(repo.dir + "\n")

  if _, err := os.Stat(repo.dir); err == nil {
    return "Directory already exists\n"
  }

  _, err = git.PlainClone(repo.dir, false, &git.CloneOptions{
    URL: repo.url,
  })

  if err != nil {
    fmt.Printf("error -> " + repo.url + "\n")
    os.Exit(1)
  }
  return "clone repository\n"
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


func main() {
  var repos []Gits
  var rcfiles []string

  repos = append(repos, Gits{"https://github.com/anyenv/anyenv",      filepath.Join(os.Getenv("HOME"), ".anyenv") })
  repos = append(repos, Gits{"https://github.com/mollifier/anyframe", filepath.Join(os.Getenv("HOME"), ".temp/anyframe")})
  rcfiles = append(rcfiles, "vimrc")
  rcfiles = append(rcfiles, "zshrc")
  rcfiles = append(rcfiles, "bashrc")
  rcfiles = append(rcfiles, "aliasrc")
  rcfiles = append(rcfiles, "tmux.conf")
  rcfiles = append(rcfiles, "hyper.js")

  for i := 0; i < len(repos); i++ {
    var res string
    res = git_clone(repos[i])

    fmt.Printf(res)
  }

  for i := 0; i < len(rcfiles); i++ {
    symlink(rcfiles[i])
  }
}
