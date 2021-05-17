# Students' Automobile Club, KGEC

[![Contributors](https://img.shields.io/github/contributors/dsckgec/sac-kgec-web.svg)](https://github.com/dsckgec/sac-kgec-web/graphs/contributors) [![Forks](https://img.shields.io/github/forks/dsckgec/sac-kgec-web.svg)](https://github.com/dsckgec/sac-kgec-web/network/members) [![Issues](https://img.shields.io/github/issues/dsckgec/sac-kgec-web.svg)](https://github.com/dsckgec/sac-kgec-web/issues) [![Pull Request](https://img.shields.io/github/issues-pr-closed-raw/dsckgec/sac-kgec-web)](https://github.com/dsckgec/sac-kgec-web/pulls)


A basic GitHub repository template for initializing open source projects on a single click.

## Contents

1. [Description](#description)
1. [Project structure](#project-structure)
1. [Project roadmap](#project-roadmap)
1. [Getting started](#getting-started)
1. [Live demo](#live-demo)
1. [Built with](#built-with)
1. [Contributing](#contributing)
1. [Authors](#authors)
1. [License](#license)
1. [Acknowledgments](#acknowledgments)

## Description

This project hosts the revamped version of SAC KGEC's official website


## Project structure

```
/
  ├── .github/            github related files like PR templates, contribution guidelines
  ├── client/             client side code            
      ├── css/            stores all the css files
      ├── img/            directory for images used in frontend
      ├── js/             directory for frontend javascript           
      ├── blogs.json      stores individual blogs as json
      ├── *.html          all HTML files go here
  ├── server/             server side code
      ├── blogs.go        extracts blogs from the blogs.json from client side
      ├── controller.go   handler functions for routes go here
      ├── routes.go       endpoint definitions go here
      ├── server.go       initiates the routes and the gin-gonic server
  ├── .gitignore          stores files and directories to be ignored in commits
  ├── go.mod              stores definitions of go packages and modules used
  ├── go.sum              stores definitions of go packages and modules used
  ├── LICENSE             the open source license
  ├── main.go             entry point to the server
  └── readme.md           project readme!
```

## Project roadmap
```
To be updated after today's meet
```

## Getting started
Everyone is welcomed to contribute to our project. Mentioning in bold, you do not need to know the tech stack and tools beforehand to be a part of our project. This is a learn-and-build projects where the contributors build alongside learning the various concepts and technologies involved.
Below are a few prerequisites and installation guides:

### Prerequisites
- Go - [download](https://golang.org/dl/)
- A web browser

### Installing

A step by step series of examples that tell you how to get a development env running.

1. Fork and clone the repository followed by opening the project in your text editor (with a terminal)
2. create a .env file and add this line `PORT=5000`
3. That was all for setting up the project. Whenever you need to run the project, type go run main.go in the terminal


## Live demo


## Built with

- Golang (with Gin-Gonic)
- HTML, CSS, JS

## Contributing

Please read [contributing.md](./.github/contributing.md) for details on our code of conduct, and the process for submitting pull requests to us.

## Authors

<a href="https://github.com/DSCKGEC/sac-kgec-web/graphs/contributors">
  <img src="https://contrib.rocks/image?repo=DSCKGEC/sac-kgec-web" />
</a>

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- 
