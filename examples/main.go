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

type FakeClient struct {
}

func (fc *FakeClient) DoSomething() {
  fmt.Println("Fake client doing something")
}

type PG struct {
  globals *pseudoglobals.Pseudoglobals
}

func (pg *PG) Config() pseudoglobals.ConfigInstance {
  return pg.globals.Config()
}

func (pg *PG) Logger() pseudoglobals.LoggerInstance {
  return pg.globals.Logger()
}

func (pg *PG) Log(msg string) {
  pg.globals.Log(msg)
}

func (pg *PG) LogErrorWithTrace(msg string, trace string) {
  pg.globals.LogErrorWithTrace(msg, trace)
}

func (pg *PG) Clients() map[string]interface{} {
  return pg.globals.Clients()
}

func (pg *PG) PostgresClient() *FakeClient {
  return pg.globals.Clients()["postgres"].(*FakeClient)
}

func main() {
  config := ConfigInstance{}
  globals := &PG{pseudoglobals.New(&config, LoggerImplementation{}, map[string]interface{}{
    "postgres": &FakeClient{},
  })}

  defer func() {
    if r := recover(); r != nil {
      globals.LogErrorWithTrace(fmt.Sprintf("%s", r), "I could be a stack trace")
    }
  }()

  globals.PostgresClient().DoSomething()

  globals.Log(globals.Config().GetString("some_config"))

  panic("Can you do this?")
}
