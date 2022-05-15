package main
 
 import (
  "fmt"
  "io/ioutil"
  "os/exec"
 )
 
 func main() {
  cmd := exec.Command("/bin/bash", "-c", `env`)
 
  //Create get command output pipeline
  stdout, err := cmd.StdoutPipe()
  if err != nil {
   fmt.Printf("Error:can not obtain stdout pipe for command:%s\n", err)
   return
  }
 
  //Execute command
  if err := cmd.Start(); err != nil {
   fmt.Println("Error:The command is err,", err)
   return
  }
 
  //Read all outputs
  bytes, err := ioutil.ReadAll(stdout)
  if err != nil {
   fmt.Println("ReadAll Stdout:", err.Error())
   return
  }
 
  if err := cmd.Wait(); err != nil {
   fmt.Println("wait:", err.Error())
   return
  }
  fmt.Printf("stdout:\n\n %s", bytes)
 }
