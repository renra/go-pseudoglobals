package pseudoglobals

const SEVERITY_INFO = 1
const SEVERITY_ERROR = 0

type ConfigInstance interface {
  GetString(string) string
}

type configImplementation interface {
  Load() ConfigInstance
}

type LoggerInstance interface {
  LogWithSeverity(map[string]string, int)
}

type loggerImplementation interface {
  New(string, int, map[int]string) LoggerInstance
}

type Pseudoglobals struct {
  Config ConfigInstance
  Logger LoggerInstance
}

func New(c configImplementation, l loggerImplementation) (* Pseudoglobals) {
  config := c.Load()

  return &Pseudoglobals{
    Config: config,
    Logger: l.New(
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
  p.Logger.LogWithSeverity(map[string]string{"message": msg}, SEVERITY_INFO)
}

func (p * Pseudoglobals) LogErrorWithTrace(msg string, trace string) {
  p.Logger.LogWithSeverity(map[string]string{"msg": msg, "trace": trace}, SEVERITY_ERROR)
}

