package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "os"
    "os/exec"
)

type Meta struct {
    Name        string `json:"name"`
    Version     string `json:"version"`
    Description string `json:"description"`
    Author      string `json:"author"`
    License     string `json:"license"`
    Install     string `json:"install"`
    Uninstall   string `json:"uninstall"`
}

func loadMeta(path string) (*Meta, error) {
    content, err := ioutil.ReadFile(path)
    if err != nil {
        return nil, err
    }
    var meta Meta
    if err := json.Unmarshal(content, &meta); err != nil {
        return nil, err
    }
    return &meta, nil
}

func installPackage(pkgPath string) {
    meta, err := loadMeta(pkgPath + "/meta.json")
    if err != nil {
        fmt.Println("Error loading meta.json:", err)
        return
    }

    fmt.Println("Installing", meta.Name, "version", meta.Version)
    cmd := exec.Command("/bin/sh", pkgPath+"/"+meta.Install)
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    if err := cmd.Run(); err != nil {
        fmt.Println("Installation failed:", err)
    } else {
        fmt.Println("Installation complete.")
    }
}

func uninstallPackage(pkgPath string) {
    meta, err := loadMeta(pkgPath + "/meta.json")
    if err != nil {
        fmt.Println("Error loading meta.json:", err)
        return
    }

    fmt.Println("Uninstalling", meta.Name)
    cmd := exec.Command("/bin/sh", pkgPath+"/"+meta.Uninstall)
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    if err := cmd.Run(); err != nil {
        fmt.Println("Uninstallation failed:", err)
    } else {
        fmt.Println("Uninstallation complete.")
    }
}

func main() {
    if len(os.Args) != 3 {
        fmt.Println("Usage: shpkg [install|uninstall] [package-path]")
        return
    }

    command := os.Args[1]
    pkgPath := os.Args[2]

    if command == "install" {
        installPackage(pkgPath)
    } else if command == "uninstall" {
        uninstallPackage(pkgPath)
    } else {
        fmt.Println("Invalid command. Use install or uninstall.")
    }
}
