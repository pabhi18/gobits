# Dockwatch - Docker Container Resource Watcher

Dockwatch is a lightweight CLI tool for monitoring Docker containers' **CPU**, **memory** usage, and **uptime** in real-time.

## Features

- List running Docker containers with their **CPU** usage, **memory** usage, and **uptime**.
- Simple command-line interface using Cobra.

## Installation

### Prerequisites

- Make sure you have **Go** installed on your machine.
- Docker must be installed and running.

### Build and Install

1. Clone this repository:

    ```bash
    git clone https://github.com/pabhi18/dockwatch.git
    cd dockwatch
    ```

2. Build the project:

    ```bash
    go build -o dockwatch
    ```

3. Move the binary to a system path (optional):

    ```bash
    sudo mv dockwatch /usr/local/bin/
    ```

   This will allow you to run `dockwatch` globally from anywhere.

## Usage

### To List Docker Containers

Run the following command to list the containers along with their **CPU**, **memory**, and **uptime**:

```bash
dockwatch list
