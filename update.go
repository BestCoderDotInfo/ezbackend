package main

import (
  "fmt"
  "os/exec"
  "log"
  "bytes"
  "strings"
)

func main() {
  fmt.Println("Current branch is master \n")
  fmt.Println("Your source code has been updated...\n");
  fmt.Println("The local branch merged have been deleted...\n");

  cmd := exec.Command("bash", "./bin/update");
  cmd.Stdin = strings.NewReader("");
  var out bytes.Buffer;
  cmd.Stdout = &out;
  err := cmd.Run();
  if err != nil {
          log.Fatal(err);
  }
  fmt.Printf("Output: \n",out.String());
}
