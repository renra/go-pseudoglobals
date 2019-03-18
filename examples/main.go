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
