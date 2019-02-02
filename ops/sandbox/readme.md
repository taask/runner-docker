## Taask Docker Sandbox

This image builds upon dind (Docker-in-Docker) and incorporates [Google's gVisor](https://github.com/google/gvisor) to provide the tools needed to safely run untrusted Docker images as Taask tasks.

`runner-docker` will run tasks as normal Docker containers (using the hosts's daemon currently) under most circumstances. To run a task as a gVisor sandboxed container, include the `io.taask.container.config/sandbox:true` annotation, or run `runner-docker` with the `TAASK_RUNNER_DOCKER_SANDBOX=true` environment variable to enforce the sandbox on all tasks.