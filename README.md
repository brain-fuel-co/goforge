# GoForge

## Inspired by

- Mix: a build tool for Elixir that creates, compiles, tests, manages dependencies, and much more.
- Polylith: a framework for architecting composable projects in a variety of languages
- Leiningen: a project automation framework for Clojure
- Cobra: a Go framework for making CLI apps
- Maven and Gradle: frameworks for making Java apps

## Why use GoForge?

- Straightforward way to make composable code in Go.
- See what's going on in a project at a glance.
- Straightforward commands to make and extend a project.
- Straightforward commands to make and version artifacts.
- Plays nicely with Git & GitHub.
- Plays nicely with Cobra
- Plays nicely with Viper

## How to use GoForge?

```bash
goforge new project <git user account>/<project name>
# OR
goforge new project <website>/<project name>
# OR
goforge new project <website>/<path/to/project name>
```

```bash
goforge new component <component name>
```

```bash
goforge new base <base name>
```

```bash
goforge new app <app name>
# OR
goforge new app --cobra <app name>
```

## Commands

### Show

#### Components
##### All
##### For a particular app
##### For a particular base

#### Bases
##### All
##### For a particular app
##### Using a particular component

#### Apps
##### All
##### Using a particular base
##### Using a particular component

#### Deps
##### All
##### For a particular component
##### For a particular base
##### For a particular app

### New

With each new entity, it should be added to Git.

#### Project
#### Component
#### Base
#### App

### Build
#### App
#### All Apps

### Clean
#### All
#### App

### Test
#### Component
#### All Components
#### Base
#### All Bases
#### App
#### All Apps

### Tag

### Bump
#### Major
#### Minor
#### Patch

### Release

Build, Test, Tag

#### Major
#### Minor
#### Patch

## TODO

- Make GoForge compatible with GitHub Actions
