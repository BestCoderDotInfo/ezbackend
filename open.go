package main

import (
  "fmt"
  "os/exec"
  "bytes"
  "strings"
)

func main() {
  fmt.Println("Open http://127.0.0.1:4000 via browser... \n")

  cmd := exec.Command("bash", "./bin/open");
  cmd.Stdin = strings.NewReader("");
  var out bytes.Buffer;
  cmd.Stdout = &out;
  cmd.Run();
}
