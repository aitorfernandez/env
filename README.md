# env

A Golang small wrapper around Redis to manage environment variables.

## Problem

Environment variables although a common practice during development it's not a healthy practice to use in Production.

Config files with constants variables in the app lead to messy files, config variables changes between environments, code does not.

## Solution

Move everything into a single source of true, Redis.

## Usage

    package main

    import (
      "github.com/aitorfernandez/env"
    )

    func main() {
      // Get the value of key.
      key, err := env.Get("key")

      // Get the value of a hash field.
      secret, err := env.Hget("jwt", "secret")

      // Get all the members in a set.
      routes, err := env.Smembers("routes")

      // With wrap around log.Fatal.
      addr := env.MustHget("account", "addr")
    }
