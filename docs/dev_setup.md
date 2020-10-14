# Introduction

Welcome to the development setup guide! Follow the instructions below to be up & running in no time.

## Pre-requisites

- Install Golang >= 1.15.2
- Install PostgreSQL >= 9.5

## Checks

- Ensure you are running the postgres service
- Ensure you have set the appropriate password for the postgres user
- Create a new database in postgres called `caros`
- Grant the necessary access to the above database to the user from Step 2

## Installing Dependencies

- When you build the app, golang automatically downloads the necessary dependencies by reading the `go.mod` file.
- Do not delete the checksum file as that is required for dependency management

## Environment Setup

We make use of environment variables to configure variable data like `database password`, `database username` and so on. So, you'll need to set environment variables using one of the following approaches:

- Copy the provided `sample.env` file from project root as `.env` and replace the values accordingly
- Set all the environment variables specified in `sample.env` into your environment

## Running the App

- From the project root, run the command `go run src/app.go`

## Running tests

- From the project root, run the command `go test ./...`.
- This will run all the tests from all the packages
