# BaseGo

Base API project using Go.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.
See deployment for notes on how to deploy the project on a live system.

### Prerequisites

What things you need to install the software and how to install them:

1. Install [Go](https://golang.org/) and [set up `GOPATH`](https://github.com/golang/go/wiki/SettingGOPATH).

    ```bash
    # for macOS, required https://brew.sh/
    $ brew install go

    # create empty directory for GOPATH
    $ mkdir $HOME/go && cd $HOME/go && mkdir bin pkg src && cd -

    # set up GOPATH, the example using https://ohmyz.sh/
    # change ~/.zshrc to ~/.bashrc if you use default bash
    $ echo "export GOPATH=$HOME/go
            export PATH=$PATH:$GOPATH/bin
        " >> ~/.zshrc
    $ source ~/.zshrc
    ```
    For more details, visit [https://golang.org/doc/install](https://golang.org/doc/install).

2. Install [dep](https://github.com/golang/dep) as dependency management tool.

    On MacOS you can install or upgrade to the latest released version with Homebrew:
    
    ```bash
    $ brew install dep
    $ brew upgrade dep
    ```

    On Debian platforms you can install or upgrade to the latest version with apt-get:

    ```bash
    $ sudo apt-get install go-dep
    ```

    On other platforms you can use the `install.sh` script:

    ```bash
    $ curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
    ```

    It will install into your `$GOPATH/bin` directory by default or any other directory you specify using the `INSTALL_DIRECTORY` environment variable.

    If your platform is not supported, you'll need to build it manually.

    If you're interested in hacking on `dep`, you can install via `go get`:

    ```bash
    $ go get -u github.com/golang/dep/cmd/dep
    ```

### Clone

1. Create project directory.

    ```bash
    $ mkdir -p $GOPATH/src/github.com/jonylim/
    ```

2. Change directory to project directory, then `git clone` the project.

    ```bash
    $ cd $GOPATH/src/github.com/jonylim/ && git clone https://github.com/jonylim/basego.git
    ```

3. Change directory to project path, then install the dependencies.

    ```bash
    $ cd $GOPATH/src/github.com/jonylim/basego/ && dep ensure
    ```

### Running the code

From the project's root directory, run the command.

```bash
$ go run *.go
```

### Build

From the project's root directory, run the following commands.

```bash
$ go build -o basego-api cmd/cstd/*.go
```

## Deployment

The application is run using `systemd` services.
For the very first time, configure the systemd syslog by running the script in `first-setup.sh`.

For quick build & deployment, run the following command from the project's root directory.

```bash
$ make clean build deploy
```

## Built with

* [Go](https://golang.org/) - The programming language
* [Dep](https://github.com/golang/dep) - Dependency Management
* [HttpRouter](https://github.com/julienschmidt/httprouter) - HTTP request router
* [Redigo](https://github.com/gomodule/redigo) - Go client for the Redis database
* [pq](https://github.com/lib/pq) - Pure Go Postgres driver for database/sql
* [Imaging](github.com/disintegration/imaging) - Basic image processing functions (resize, rotate, crop, etc.)

## Contributing

Please read [our wiki](https://github.com/jonylim/basego/wikis/home) for details on our code of conduct, and the process for submitting pull requests to us.
- [Project Layout](https://github.com/jonylim/basego/wikis/Contributors/01.-Project-Layout)
- [Effective Go](https://golang.org/doc/effective_go.html)
- [Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments) 

## Versioning

We use [SemVer](http://semver.org/) for versioning.
For the versions available, see the [tags on this repository](https://github.com/jonylim/basego/tags). 

## Learn More

Please check [our wiki](https://github.com/jonylim/basego/wikis/home).