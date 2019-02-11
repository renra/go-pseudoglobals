# Pseudoglobals

A project that takes care of pseudoglobals that can be passed to http request handlers. It was built with together with renra/go-logger and renra/go-config (TODO: add links) but it can be used with anything that implements the same interface. Right now it provides access to `Config` and `Logger`.

```go
package main

import (
  "fmt"
  "github.com/renra/go-pseudoglobals/pseudoglobals"
)

type ConfigInstance struct {
}

func (cInst *ConfigInstance) GetString(key string) string {
  return key;
}

type ConfigImplementation struct {
}

func (ci ConfigImplementation) Load() pseudoglobals.ConfigInstance {
  return &ConfigInstance{}
}

type LoggerInstance struct {
}

func (lInst LoggerInstance) LogWithSeverity(data map[string]string, severity int) {
  fmt.Println(fmt.Sprintf("%d: %s", severity, data));
}

type LoggerImplementation struct {
}

func (li LoggerImplementation) New(label string, thresholdSeverity int, severities map[int]string) pseudoglobals.LoggerInstance {
  return &LoggerInstance{}
}


func main() {
  globals := pseudoglobals.New(ConfigImplementation{}, LoggerImplementation{})

  defer func() {
    if r := recover(); r != nil {
      globals.LogErrorWithTrace(fmt.Sprintf("%s", r), "I could be a stack trace")
    }
  }()

  globals.Log(globals.Config.GetString("some_config"))

  panic("Can you do this?")
}

```
