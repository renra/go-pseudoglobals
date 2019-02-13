package pseudoglobals

const SEVERITY_INFO = 1
const SEVERITY_ERROR = 0

type ConfigInstance interface {
  Get(string) interface{}
  GetString(string) string
}

type LoggerInstance interface {
  LogWithSeverity(map[string]string, int)
}

type LoggerImplementation interface {
  New(string, int, map[int]string) LoggerInstance
}

type Pseudoglobals struct {
  config ConfigInstance
  logger LoggerInstance
}

func (g *Pseudoglobals) Config() ConfigInstance {
  return g.config
}

func (g Pseudoglobals) Logger() LoggerInstance {
  return g.logger
}

func New(config ConfigInstance, l LoggerImplementation) (* Pseudoglobals) {
  return &Pseudoglobals{
    config: config,
    logger: l.New(
      config.GetString("service"),
      SEVERITY_INFO,
      map[int]string {
        SEVERITY_INFO: "INFO",
        SEVERITY_ERROR: "ERROR",
      },
    ),
  }
}

func (p * Pseudoglobals) Log(msg string) {
  p.Logger().LogWithSeverity(map[string]string{"message": msg}, SEVERITY_INFO)
}

func (p * Pseudoglobals) LogErrorWithTrace(msg string, trace string) {
  p.Logger().LogWithSeverity(map[string]string{"msg": msg, "trace": trace}, SEVERITY_ERROR)
}
