package dotfiles

import (
  "os"
  flags "github.com/jessevdk/go-flags"
)

type Options struct {
  shell string `short:"s" long:"shell" description:"selectt defaurt shell"`
}


func cli() Options {
  var opts Options

  _, err := flags.Parse(&opts)
  if err != nil {
    os.Exit(1)
  }

  return opts
}
