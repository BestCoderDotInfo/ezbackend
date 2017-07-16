package main

import (
  "fmt"
  "os/exec"
  "log"
  "bytes"
  "strings"
)

func main() {
  fmt.Println("Your application will run at http://127.0.0.1:4000...\n")
  cmd1 := exec.Command("bash", "./bin/start");
  cmd1.Stdin = strings.NewReader("");
  var out bytes.Buffer;
  cmd1.Stdout = &out;
  err := cmd1.Run();
  if err != nil {
          log.Fatal(err);
  }
  fmt.Printf("Can't start your application. Something errors \n",out.String());
}
