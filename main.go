package gerst

import (
  "fmt"
  "os/exec"
  "strings"
)

func EscapeSingleQuotes(str string) string {
  return strings.ReplaceAll(str, "'", "'\\''")
}

func ExecShellCommand(shellCommand string) string {
  out, err := exec.Command("bash", "-c", shellCommand).Output()
  if err != nil {
    return fmt.Sprintf("failed to execute shell command: %s", shellCommand)
  }
  return string(out)
}

func LogString(str string) {
  fmt.Printf("\n\n\n%v\n\n\n", str)
}

func ManipulateJson(json string, operation string) string {
  return ExecShellCommand(fmt.Sprintf("echo '%s' | jq '%s'", EscapeSingleQuotes(json), operation))
}
