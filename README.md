# Pseudoglobals

A project that takes care of pseudoglobals that can be passed to http request handlers. It was built together with [renra/go-logger](https://github.com/renra/go-logger) and [renra/go-helm-config](https://github.com/renra/go-helm-config) (TODO: add links) but it can be used with anything that implements the same interface. Right now it provides access to `Config`, `Logger` and `Clients`.

## Usage

```go
package main

import (
  "fmt"
  "app/pseudoglobals"
  "github.com/renra/go-errtrace/errtrace"
)

type ConfigInstance struct {
}

func (ci *ConfigInstance) Get(key string) (interface{}, *errtrace.Error) {
  return key, nil;
}

func (ci *ConfigInstance) GetP(key string) interface{} {
  return key;
}

func (ci *ConfigInstance) GetString(key string) (string, *errtrace.Error) {
  return key, nil;
}

func (ci *ConfigInstance) GetStringP(key string) string {
  return key;
}

func (ci *ConfigInstance) GetInt(key string) (int, *errtrace.Error) {
  return 4, nil;
}

func (ci *ConfigInstance) GetIntP(key string) int {
  return 4;
}

func (ci *ConfigInstance) GetFloat(key string) (float64, *errtrace.Error) {
  return 3.14, nil;
}

func (ci *ConfigInstance) GetFloatP(key string) float64 {
  return 3.14;
}

func (ci *ConfigInstance) GetBool(key string) (bool, *errtrace.Error) {
  return true, nil;
}

func (ci *ConfigInstance) GetBoolP(key string) bool {
  return true;
}

type LoggerInstance struct {
}

func (lInst *LoggerInstance) LogWithSeverity(data map[string]string, severity int) {
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
      globals.LogErrorWithTrace(errtrace.Wrap(r))
    }
  }()

  globals.Log(globals.Config().GetStringP("some_config"))

  panic(errtrace.New("Can you do this?"))
}
```

`Clients` is a generic type that you can use to store any kind of client in - postgres, redis, elasticsearch, http, etc. The pain is that you have to type-assert it when you read it back from config. So I would recommend to create a wrapper class that will provide that functionality automatically.
