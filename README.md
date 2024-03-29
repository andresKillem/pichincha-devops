## DevOps Assessment
* Golang microservice, dockerized and deployed on ECS with Kubernets cluster
* For the CI I used Shippable

## Results

![alt text](https://github.com/andresKillem/pichincha-devops/blob/master/Images/1.png)
![alt text](https://github.com/andresKillem/pichincha-devops/blob/master/Images/2.png)
![alt text](https://github.com/andresKillem/pichincha-devops/blob/master/Images/3.png)
![alt text](https://github.com/andresKillem/pichincha-devops/blob/master/Images/4.png)
![alt text](https://github.com/andresKillem/pichincha-devops/blob/master/Images/5.png)
![alt text](https://github.com/andresKillem/pichincha-devops/blob/master/Images/6.png)

# Requirements
* go 1.9 or higher.

# Installation & Usage

* First of all we have to ensure that our project is in the correct path, follow option #1 in this [dep tutorial](https://golang.github.io/dep/docs/new-project.html) to know where your project should be.
* Run `$ make deps` to install all the dependencies.

## Using Dep

* `dep ensure` ensures that all the go dependencies are installed (this is run in the Makefile).
* If you are going from scratch and don't need the examples provided in this template you will need to remove `Gopkg.lock` and `Gopkg.toml` and initialize dep with `$ dep init`
* After running `$ dep init` a _vendor_ folder will be created and you can _add your dependencies_ with the following `$ dep ensure -add github.com/foo/bar github.com/baz/quux` if you want to know more about dep just check the tutorial above.
* To understand how are you going to be using dep please check [this](https://golang.github.io/dep/docs/daily-dep.html#using-dep-ensure) out

## Creating an API through [gin-gonic](https://github.com/gin-gonic/gin)

* When installing the dependencies for this project you will get gin-gonic for API definition.
* To run the main for the API application you have to `$ make run` or `$ go run cmd/api/main.go`, this will run an HTTP server in port 8080
* For testing all the define endpoint you can try out these different CURL commands:
    * Create: `$ curl -H "Content-type: application/json" -d '{"i_am": "1", "title": "Some Todo Title", "the_rest": "description", "when_finish": "2018-12-06T14:26:40.623Z"}' "http://localhost:8080/todo"`
    * Read: `$ curl -X GET "http://localhost:8080/todo/1"`
    * Update: `$ curl -X PUT -H "Content-type: application/json" -d '{"i_am": "1", "title": "Some Todo Title", "the_rest": "description", "when_finish": "2018-12-06T14:26:40.623Z"}' "http://localhost:8080/todo"`
    * Delete: `$ curl -X DELETE "http://localhost:8080/todo/1"`

## How to make build of a main.go file and run it.
* `$ make build` or `$ go build -o ./tmp/web-server ./cmd/api/main.go`
* The command above will create a file called `web-server` in folder _tmp_, that file is an executable with the main in _./cmd/api/main.go_
* To run your executable you have to:
    * Make it executable: `chmod +x ./tmp/web-server`
    * Run it: `./tmp/web-server`
* You can build whatever you want (it doesn't have to be a web-server), for example there is another main with which you can follow the same steps _./cmd/cli/main.go_

## Hot reload for the Web Server
* For Mac/Windows users you will have to modify [install_dep.sh](./scripts/install_dep.sh) and change the OS variable to match your Operating System.
* We have to install [Air](https://github.com/cosmtrek/air) (if you ran `$ make deps` command this should be already installed and you can skip this):
    * `curl -fLo ~/.air \
           https://raw.githubusercontent.com/cosmtrek/air/master/bin/linux/air`
    * `chmod +x ~/.air`
* To run the project with hot reload: `$ make run-dev` or `~/.air -c .air.config`

## How to run the tests
* After you install everything with dep you should be able to do `$ make test` or `$ go test -cover -v ./...` this will run all the test files in the project.
* Test should be in the same folder of the file they are testing and the file name of the test must have the suffix `_test`, if you see the example in _test_ folder you will get it right away.

## TL;DR How to run/build
* Build:  `$ make build` or `$ go build -o <destination_of_executable_relative_to_root> <path_of_main_file_relative_to_root>`
* Run:
    * Without executable: `$ make run` or `$ go run <path_of_main_file_relative_to_root>`
    * With executable:
        * Make the file executable: `$ chmod +x <path_to_executable_relative_to_root>
        * Run it: `$ <path_to_executable_relative_to_root>

# Standard Go Project Layout

This is a basic layout for Go application projects. It's not an official standard defined by the core Go dev team; however, it is a set of common historical and emerging project layout patterns in the Go ecosystem. Some of these patterns are more popular than others. It also has a number of small enhancements along with several supporting directories common to any large enough real world application.

If you are trying to learn Go or if you are building a PoC or a toy project for yourself this project layout is an overkill. Start with something really simple (a single `main.go` file is more than enough). As your project grows keep in mind that it'll be important to make sure your code is well structured otherwise you'll end up with a messy code with lots of hidden dependencies and global state. When you have more people working on the project you'll need even more structure. That's when it's important to introduce a common way to manage packages/libraries. When you have an open source project or when you know other projects import the code from your project repository that's when it's important to have private packages and code. Clone the repository, keep what you need and delete everything else! Just because it's there it doesn't mean you have to use it all. None of these patterns are used in every single project. Even the `vendor` pattern is not universal.

This project layout is intentionally generic and it doesn't try to impose a specific Go package structure.

This is a community effort. Open an issue if you see a new pattern or if you think one of the existing patterns needs to be updated.

If you need help with naming, formatting and style start by running [`gofmt`](https://golang.org/cmd/gofmt/) and [`golint`](https://github.com/golang/lint). Also make sure to read these Go code style guidelines and recommendations:
* https://talks.golang.org/2014/names.slide
* https://golang.org/doc/effective_go.html#names
* https://blog.golang.org/package-names
* https://github.com/golang/go/wiki/CodeReviewComments

See [`Go Project Layout`](https://medium.com/golang-learn/go-project-layout-e5213cdcfaa2) for additional background information.

## Go Directories

### `/todo`

This directory varies according to the name of the application, in this example it is _todo_ because that's what this example is about. You will have all
your business related code under this directory.

### `/cmd`

Main applications for this project.

The directory name for each application should match the name of the executable you want to have (e.g., `/cmd/myapp`).

Don't put a lot of code in the application directory. If you think the code can be imported and used in other projects, then it should live in the `/pkg` directory. If the code is not reusable or if you don't want others to reuse it, put that code in the `/internal` directory. You'll be surprised what others will do, so be explicit about your intentions!

It's common to have a small `main` function that imports and invokes the code from the `/internal` and `/pkg` directories and nothing else.

See the [`/cmd`](cmd/README.md) directory for examples.

### `/vendor`

Application dependencies (managed manually or by your favorite dependency management tool like [`dep`](https://github.com/golang/dep)).

Don't commit your application dependencies if you are building a library.

## Service Application Directories

### `/api`

OpenAPI/Swagger specs, JSON schema files, protocol definition files.

## Common Application Directories

### `/configs`

Configuration file templates or default configs.

Put your `confd` or `consul-template` template files here.

### `/scripts`

Scripts to perform various build, install, analysis, etc operations.

These scripts keep the root level Makefile small and simple (e.g., `https://github.com/hashicorp/terraform/blob/master/Makefile`).

See the [`/scripts`](scripts/README.md) directory for examples.

