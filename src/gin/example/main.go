package main

    import (
        _ "GoLearningExample/src/gin/example/api/database"
        "GoLearningExample/src/gin/example/api/router"
        orm "GoLearningExample/src/gin/example/api/database"
    )

    func main() {
        defer orm.Eloquent.Close()
        router := router.InitRouter()
        router.Run(":8000")
    }