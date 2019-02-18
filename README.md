# Pseudoglobals

A project that takes care of pseudoglobals that can be passed to http request handlers. It was built together with [renra/go-logger](https://github.com/renra/go-logger) and [renra/go-helm-config](https://github.com/renra/go-helm-config) (TODO: add links) but it can be used with anything that implements the same interface. Right now it provides access to `Config`, `Logger` and `Clients`.

## Usage

```go
package main

import (
  "fmt"
  "github.com/renra/go-pseudoglobals/pseudoglobals"
)

type ConfigInstance struct {
}

func (ci *ConfigInstance) GetString(key string) string {
  return key;
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
  config := ConfigInstance{}

  clients := map[string]interface{} {
    "postgres": "fake postgres client",
  }

  globals := pseudoglobals.New(&config, LoggerImplementation{}, clients)

  defer func() {
    if r := recover(); r != nil {
      globals.LogErrorWithTrace(fmt.Sprintf("%s", r), "I could be a stack trace")
    }
  }()

  globals.Log(globals.Config().GetString("some_config"))

  panic("Can you do this?")
}
```

`Clients` is a generic type that you can use to store any kind of client in - postgres, redis, elasticsearch, http, etc. The pain is that you have to type-assert it when you read it back from config. So I would recommend to create a wrapper class that will provide that functionality automatically.
