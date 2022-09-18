package main

import (
    "go-web-template/modules/engine"
    "go-web-template/modules/repository"
)

func main() {
    ginEngine := engine.InitGinManager().GetGinEngine()

    ginEngine.Run(":8081")

    defer repository.CloseMySQL()
}
