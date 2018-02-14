package main

import (
  "net/http"
  // "fmt"
  "github.com/gin-gonic/gin"
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB

type Todo struct {
  gorm.Model
  Title string
  Completed string
}

type TodoJSON struct {
  Title string `json:"title" binding:"required"`
  Completed string `json:"completed" binding:"required"`
}

func init() {
  // Open the DB connection
  var err error
  db, err = gorm.Open("sqlite3", "database.db")

  if err != nil {
    panic("Error connecting to database")
  }
  defer db.Close()

  // Migrate the schema
  db.AutoMigrate(&Todo{})
}

func main() {
  router := gin.Default()

  v1 := router.Group("/api/v1/todo")
  {
    v1.GET("/ping", pong)
    v1.POST("/", createTodo)
    // v1.GET("/", getTodos)
    // v1.GET("/:id", getTodo)
    // v1.PUT("/:id", updateTodo)
    // v1.DELETE("/:id", deleteTodo)
  }

  router.Run(":8080")
}

func pong(c *gin.Context) {
  db.Create(&Todo{Title: "Test", Completed: "true"})

  c.JSON(http.StatusOK, gin.H{
    "message": "pong",
  })
}

func createTodo(c *gin.Context) {
  var json TodoJSON

  if err := c.ShouldBindJSON(&json); err == nil {
    db.Create(Todo{Title: json.Title, Completed: json.Completed})

    c.JSON(http.StatusOK, gin.H{"status": "Todo created"})
  } else {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
  }
}
