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
  fmt.Printf("Can't update latest source code from remote repo. Something errors \n",out.String());
}
