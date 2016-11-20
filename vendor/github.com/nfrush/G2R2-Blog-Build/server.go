package main

import (
  "log"
  "os"
  "os/signal"
  "syscall"

  "github.com/graphql-go/handler"
  "github.com/labstack/echo/engine/fasthttp"

  func main() {
    // Creates a GraphQL-go HTTP handler with the defined schema
	  handler := handler.New(&handler.Config{
		    Schema: &schema.Schema,
		    Pretty: true,
	  })
  }
)
