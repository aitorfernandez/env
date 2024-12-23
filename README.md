# env

A lightweight Go wrapper around Redis to manage environment variables.

## Problem

While environment variables are a common practice during development, relying on them in production can lead to several issues:

- **Inconsistent Practices**: Environment variables differ across setups, making debugging and deployment harder.
- **Messy Config Files**: Hardcoding constants or using multiple config files can quickly become unmanageable.
- **Environment Drift**: Configuration changes between environments, but code does not.

## Solution

Centralize all configuration in a single source of truth: **Redis**. This approach ensures consistency, simplifies management, and supports dynamic updates without requiring code changes.

## Usage

Here's how you can use the `env` package:

```go
package main

import (
  "github.com/aitorfernandez/env"
)

func main() {
  // Get the value of a key.
  key, err := env.Get("key")
  if err != nil {
    // handle error
  }

  // Get the value of a field from a hash.
  secret, err := env.Hget("jwt", "secret")
  if err != nil {
    // handle error
  }

  // Get all members of a set.
  routes, err := env.Smembers("routes")
  if err != nil {
    // handle error
  }

  // Fetch a value with a guaranteed return or log.Fatal on error.
  addr := env.MustHget("account", "addr")
}
```
