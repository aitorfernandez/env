# env

A Golang small wrapper around Redis to manage environment variables.

## Problem

Environment variables although a common practice during development it's not a healthy practice to use in Production.

Config files with constants variables in the app lead to messy files, config variables changes between environments, code does not.

## Solution

Move everything into a single source of true, Redis.
