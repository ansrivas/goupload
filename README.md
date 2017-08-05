# goupload
This is an experimental project to understand Golang's `Chi` mux with `pongo2` templates.
---

[![Build Status](https://travis-ci.org/ansrivas/goupload.svg?branch=master)](https://travis-ci.org/ansrivas/goupload)


# Usage:
---
1. Run help:

    ```bash
    $ make help
    help:           Show available options with this Makefile
    test:           Run all the tests
    clean:          Clean the application
    build:          Build the application
    app_help:       Display flags accepted by the application
    test_run:       Run the application in a test mode
    remove:         Remove running docker containers and remove the volumes
    dock_run_fg:    Run docker containers, foreground.
    dock_run_bg:    Run docker containers, background.
    build_docker:   Build docker containers

    ```

2. Run tests using:

    ```
    make test
    ```

3. Execute docker-app:

    ```
    make dock_run_fg
    ```

4. Navigate to `localhost:8080` and login using `admin@gmail.com:admin`
