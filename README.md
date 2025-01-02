# OpenGL Exercises with Go-GL

> This repository contains chapter examples and solutions to OpenGL exercises from [LearnOpenGL.com](https://learnopengl.com/), implemented in Go using the go-gl bindings.
Managed by [Merlindorin](https://github.com/merlindorin), this project aims to demonstrate foundational OpenGL concepts using the Go programming language.

## Table of Contents

<!-- TOC -->
* [OpenGL Exercises with Go-GL](#opengl-exercises-with-go-gl)
  * [Table of Contents](#table-of-contents)
  * [Summary](#summary)
  * [Prerequisites](#prerequisites)
  * [Installation](#installation)
    * [Clone the Repository](#clone-the-repository)
  * [Usage](#usage)
  * [Development](#development)
    * [Repository Structure](#repository-structure)
    * [Development with Taskfile](#development-with-taskfile)
    * [Installing Tools with Aqua](#installing-tools-with-aqua)
  * [Contributing](#contributing)
  * [License](#license)
  * [References](#references)
<!-- TOC -->

## Summary

Explore fundamental OpenGL concepts and practices through structured exercises. Using Go with go-gl, this repository provides a practical learning path parallel to the tutorials on LearnOpenGL.com.

## Prerequisites

To effectively use this repository, you'll need:

- [Go](https://golang.org/dl/): Ensure the latest version is installed for compatibility.
- [Git](https://git-scm.com): For version control and managing repository copies.
- [Task](https://taskfile.dev/): Utilized for script and automation task management.
- [Aqua](https://aquaproj.github.io): Manages CLI tool versions consistently used in the project.

## Installation

### Clone the Repository

1. Clone this repository to your local machine:
   ```bash
   git clone https://github.com/merlindorin/opengl-go-exercises.git
   cd opengl-go-exercises
   ```

## Usage

After cloning the repository, navigate to any exercise directory within `exercises/` and run the main Go file. Here is an example command for running the first exercise:

```bash
cd exercices/1.1.hello_window
go run main.go
```

Each subdirectory represents a separate exercise from [LearnOpenGL.com](https://learnopengl.com/), implemented using Go.

## Development

To contribute or modify the exercises, follow these steps:

1. Ensure all dependencies are installed.
2. Navigate to the desired exercise directory.
3. Develop and test your changes.

### Repository Structure

- `exercices`: Contains subdirectories for each exercise, numbered according to their sequence in the tutorial.
    - Each exercise includes `.go` files for primary logic and `.glsl` files for shaders.
- `utils`: Holds utility Go files for shader and program management, aiding in common tasks such as error checking and shader compilation.
- `Taskfile.yaml`: Configuration for automating tasks like building and testing.
- `aqua.yaml`: Specifies CLI tools managed by Aqua for development.

### Development with Taskfile

Use Taskfile to automate repetitive tasks. It helps standardize tasks like building the exercise binaries and running tests. Review the `Taskfile.yaml` to explore available tasks.

### Installing Tools with Aqua

1. Ensure Aqua is installed by following the [Aqua installation guide](https://aquaproj.github.io/docs/install).
2. Use the `aqua.yaml` file provided to detail required tool versions.
3. Run Aqua to install these tools:
   ```bash
   aqua i
   ```

   This command will guide Aqua in installing the specified tools, ensuring all contributors use consistent tool versions.

## Contributing

Interested in contributing?

- Fork the repository.
- Create a new branch: `git checkout -b feature/new-feature`.
- Commit your changes: `git commit -m 'Add new feature'`.
- Push your branch: `git push origin feature/new-feature`.
- Open a pull request for discussion and review.

## License

This project is licensed under the MIT License. See the [LICENSE.md](./LICENSE.md) file for details.

## References

- [LearnOpenGL.com](https://learnopengl.com/) - A comprehensive resource for learning OpenGL.
- [go-gl](https://github.com/go-gl) - Go bindings for OpenGL.