package main

import (
  "fmt"
  "app/pseudoglobals"
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
