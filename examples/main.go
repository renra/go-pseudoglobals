package main

import (
  "fmt"
  "app/pseudoglobals"
)

type ConfigInstance struct {
}

func (ci *ConfigInstance) Get(key string) interface{} {
  return key;
}

func (ci *ConfigInstance) GetString(key string) string {
  return key;
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
  globals := pseudoglobals.New(&config, LoggerImplementation{})

  defer func() {
    if r := recover(); r != nil {
      globals.LogErrorWithTrace(fmt.Sprintf("%s", r), "I could be a stack trace")
    }
  }()

  globals.Log(globals.Config().GetString("some_config"))

  panic("Can you do this?")
}
