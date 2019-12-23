package dotfiles

import (
  "fmt"
  "os"
  "gopkg.in/src-d/go-git.v4"
)

func git_run(){
  repos   := gits()

  for i := 0; i < len(repos); i++ {
    var res string
    res = git_clone(repos[i])

    fmt.Printf(res)
  }
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
