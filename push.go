package main

import (
  "fmt"
  "os/exec"
  "log"
  "bytes"
  "strings"
)

func main() {
  fmt.Println("All your changes will push with new remote branch\n")
  cmd := exec.Command("bash", "./bin/push");
  cmd.Stdin = strings.NewReader("");
  var out bytes.Buffer;
  cmd.Stdout = &out;
  err := cmd.Run();
  if err != nil {
          log.Fatal(err);
  }
  fmt.Printf("Output: \n",out.String());
}
