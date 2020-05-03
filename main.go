package main

import (
  "fmt"
  "io/ioutil"
  "os"
  "os/exec"
  "strings"
)

func main() {
  if len(os.Args) != 3 {
    fmt.Println("Usage:", os.Args[0], "{VM ID}", "{YOURFILE}")
    return 
  } else {
    fmt.Println("Clearing Custom CloudInit from VM",os.Args[1],"First")
    out, err1 := exec.Command("qm","set",os.Args[1],"--cicustom","").Output()
    if err1 != nil {
      panic(err1)
    }
    fmt.Println("Getting AutoGeneratred CloudInit File For",os.Args[1])
    out, err2 := exec.Command("qm","cloudinit","dump",os.Args[1],"user").Output()
    if err2 != nil {
      panic(err2)
    }
    fmt.Println("Getting Your CloudInit File",os.Args[2])
    dat, err3 := ioutil.ReadFile(os.Args[2])
    if err3 != nil {
      panic(err3)
    }
    fmt.Println("Creating snippets folder locally /var/lib/vz/snippets")
    _ = os.Mkdir("/var/lib/vz/snippets", 0644)
    fmt.Println("Writing Custom CloudInit File /var/lib/vz/snippets/custom-"+os.Args[1])
    abc := string(out)+strings.Replace(string(dat),"#cloud-config","",1) 
    d1 := []byte(string(abc))
    err4 := ioutil.WriteFile("/var/lib/vz/snippets/custom-"+os.Args[1], d1, 0644)
    if err4 != nil {
      panic(err4)
    }
    fmt.Println("Applying Custom CloudInit For VM",os.Args[1])
    out, err5 := exec.Command("qm","set",os.Args[1],"--cicustom","user=local:snippets/custom-"+os.Args[1]).Output()
    if err5 != nil {
      panic(err5)
    } 
    fmt.Println("Completed")
  }
}
