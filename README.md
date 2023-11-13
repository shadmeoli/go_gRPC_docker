---
theme: /home/meoli/.config/slides/theme.json
author: Shadrack Meoli
---

# Go + gRPC + Docker

###### My idea of a microservice

> I am looking to making a microservice in GO and RPC for the microservices to communicate to each other and docker to manage the microservice on deployment

---

# Intro

#### What is go?

Go is an expressive, concise, clean, and efficient language.

---

#### What is go?

Go is an expressive, concise, clean, and efficient language.

Famous for it concurrency mechanisms and ease to write.

##### Quick example

```go

package main

import "fmt"

func helloWorld() {
    fmt.printf("++ GO - gRPC - Docker")
}

```

_classic print to terminal example_

---

#### What is go?

Go is an expressive, concise, clean, and efficient language.

Famous for it concurrency mechanisms and ease to write.

Types : its novel type system enables flexible and modular program construction.

```go
package main

import { 'regexp' }

type User struct {
    name            string
    email           string
    phoneNumber     int
    password        string
}


// user mail validator
func mailValidator(mail: string) IsValid() bool {
    emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
    return regexp.MatchString(emailRegex, string(e))
}
```

---

#### What is go?

Go is an expressive, concise, clean, and efficient language.

Famous for it concurrency mechanisms and ease to write.

Types : its novel type system enables flexible and modular program construction.

Go is compiled to machine code yet has the convenience of garbage collection and the power of run-time reflection.

---

#### What is go?

Go is an expressive, concise, clean, and efficient language.

Famous for it concurrency mechanisms and ease to write.

Types : its novel type system enables flexible and modular program construction.

Go is compiled to machine code yet has the convenience of garbage collection and the power of run-time reflection.

> It's a fast, statically typed, compiled language that feels like a dynamically typed, interpreted language.

---

# Project setup and file structure

I will be making a basic API in go to showcase the same.
I learning go as doing this project so pardon my project layout and file placements

> file structure

```sh
.
├── Dockerfile
├── go.mod
├── go.sum
├── pkg
├── public
├── README.md
├── server.go
├── src
├── tmp
│   └── main
└── views
    └── index.html

5 directories, 7 files
```

Here is how the app is structured

---

# Source code

Repo: [github]("https://github.com/shadmeoli/go_gRPC")

####
