package main

import (
    "flag"
    "fmt"
    "os"
    "os/exec"
    "log"
    "bytes"
    "strings"
    "path"
)

func main() {
    dir, _ := os.Getwd()
    currentDir := strings.Replace(dir, " ", "\\ ", -1)
    var rootDir = ""
    var filepath = ""
    i := strings.Index(currentDir, "bestcoder.info")
    j := strings.Index(currentDir, "go-cli")
    if i > -1 && j > -1 {
      rootDir = path.Dir(currentDir)
    } else if i > -1 && j <= -1 {
      rootDir = currentDir
    } else {
      fmt.Println("The path not found!")
      os.Exit(1)
    }
    taskStr := flag.String("task", "", "Task required! Please choose one: setup|start|open|push|update")
    flag.Parse()
    switch task := *taskStr; task {
    case "setup":
      fmt.Println("Setup starting...\n")
      filepath = path.Join(rootDir, "/bin/setup")
      cmd := exec.Command("bash", filepath);
      cmd.Stdin = strings.NewReader("");
      var out bytes.Buffer;
      cmd.Stdout = &out;
      cmd.Run();
    case "start":
      fmt.Println("Your application will run at http://127.0.0.1:4000...\n")
      filepath = path.Join(rootDir, "/bin/start")
      cmd := exec.Command("bash", filepath);
      cmd.Stdin = strings.NewReader("");
      var out bytes.Buffer;
      cmd.Stdout = &out;
      err := cmd.Run();
      if err != nil {
              log.Fatal(err);
      }
      fmt.Printf("Output: \n",out.String());
    case "open":
      fmt.Println("Open http://127.0.0.1:4000 via browser... \n")
      filepath = path.Join(rootDir, "/bin/open")
      cmd := exec.Command("bash", filepath);
      cmd.Stdin = strings.NewReader("");
      var out bytes.Buffer;
      cmd.Stdout = &out;
      cmd.Run();
    case "push":
      fmt.Println("All your changes will push with new remote branch\n")
      filepath = path.Join(rootDir, "/bin/push")
      cmd := exec.Command("bash", filepath);
      cmd.Stdin = strings.NewReader("");
      var out bytes.Buffer;
      cmd.Stdout = &out;
      err := cmd.Run();
      if err != nil {
              log.Fatal(err);
      }
      fmt.Printf("Output: \n",out.String());
    case "update":
      fmt.Println("Current branch is master \n")
      fmt.Println("Your source code has been updated...\n");
      fmt.Println("The local branch merged have been deleted...\n");
      filepath = path.Join(rootDir, "/bin/update")
      cmd := exec.Command("bash", filepath);
      cmd.Stdin = strings.NewReader("");
      var out bytes.Buffer;
      cmd.Stdout = &out;
      err := cmd.Run();
      if err != nil {
              log.Fatal(err);
      }
      fmt.Printf("Output: \n",out.String());
    default:
      flag.PrintDefaults()
      os.Exit(1)
    }

}
