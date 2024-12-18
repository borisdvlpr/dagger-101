# dagger-101

This repo gives you a quick intro to Dagger, covering the basics and showing how to develop and run a Dagger pipeline for a Rust application, using Go. The content comes from [Dagger's Documentation](https://docs.dagger.io) and Lasse Lund's [*Software Engineering for Data Scientists*](https://lasselundstenjensen.github.io/itu-bds-sdse) course contents at the IT University of Copenhagen.

## Intro

Dagger is a tool that lets you define and run pipelines as code. It works on top of Docker, so your pipelines run in containers. This means you can run them the same way locally or on your preferred CI platform, cutting out the waiting time for CI pipeline executions (woohoo, no last minute surprises!). Plus, Dagger has a caching system that stores everything, making each pipeline run much faster.

## Ecosystem

Dagger is made up of several components that work together to manage the delivery of apps to the cloud:

- **Dagger CLI**: The main tool for interacting with Dagger, usable in the terminal or CI scripts, with a real-time terminal UI for visualization.
  
- **Dagger Functions**: Reusable code blocks that perform tasks like pulling images or copying files, running in containers and supporting multiple languages like Python, Go, and TypeScript.

- **Dagger SDKs**: SDKs for writing Dagger Functions in Go, TypeScript, and Python, with community support for other languages.

- **Dagger API**: A unified GraphQL-based interface for managing pipelines, usually accessed through the CLI or SDKs.

- **Dagger Cloud**: A control plane for managing pipelines with features like visualization, logs, and performance insights.

- **The Daggerverse**: A free service to easily find and use publicly available Dagger Functions.

## How does it work?

According to [older versions of the Go SDK Documentation](https://archive.docs.dagger.io/0.9/sdk/go/):

![Go SDK Architecture Reference](./assets/go-sdk-architecture-reference.png)

1. Your Go program imports the Dagger Go library.
2. Using the Go library, your program opens a new session to a Dagger Engine: either by connecting to an existing engine, or by provisioning one on-the-fly.
3. Using the Go library, your program prepares API requests describing pipelines to run, then sends them to the engine. The wire protocol used to communicate with the engine is private and not yet documented, but this will change in the future. For now, the Go library is the only documented API available to your program.
4. When the engine receives an API request, it computes a Directed Acyclic Graph (DAG) of low-level operations required to compute the result, and starts processing operations concurrently.
5. When all operations in the pipeline have been resolved, the engine sends the pipeline result back to your program.
6. Your program may use the pipeline's result as input to new pipelines.

