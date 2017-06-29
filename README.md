# Formed

## Command Formed

 - [Getting started](#getting-started)
 - [Introduction](#introduction)
 - [Backend](#backend-cli)
 - [Tests](#tests)


### Getting started

The formed CLI expects a few prerequisites to be installed for it to be
able to build. A valid `go` installation (1.8+) and a valid `$GOPATH`.

The quickest way to get up and running is to do the following:

```
make install
cd dist

./formed query
```

Then go to the following url [localhost:8080](localhost:8080/query) and
start editing the form!

### Introduction

The formed command code is written as one helpful CLI, so executing it is as
helpful as possible to new comers. There are a few defaults that can be useful
when developing with the backend CLI and they can be seen in the help section:

```
./formed query -help
USAGE
  query [flags]

FLAGS
  -api tcp://0.0.0.0:8080      listen address for query API
  -debug false                 debug logging
  -filestore ./data/store.csv  location of where the file store
  -ui.local false              ignores embedded files and goes straight to the filesystem
```

### Backend CLI

The backend CLI is written in `go`, using little dependencies as possible,
because most of the `stdlib` can perform most if not all the tasks required.
Including routing!

The dependency management is managed by `glide` as this seems to be the most
stable and reliable at the time of writing.

The `main.go` file (entry point) can be found in `cmd/formed/main.go`. This
is the entry point to the CLI and the code found with in the folder creates
the dependencies for the server. It also provides a helpful usage API to help
with starting said server.

The `pkg` directory is where all the CLI dependencies (libraries) are. This
is the [current standard practice](https://peter.bourgon.org/go-best-practices-2016/#repository-structure)
in go.

Building the whole of the backend can be done with the following command:

```
make clean build
```

#### Store

The store package is an abstraction over raw files, it provides two simple
methods for reading and writing files. The abstraction is at a quite high level
so that the user of the store doesn't have to concern themselves about how the
files are actually written, just that they are.

#### Files (fs)

Under the store abstraction, the file system is modelled so that better testing
can be provided. Mock files are used extensively through out the testing of
the application, where we can record and monitor the expectations of the
mocking. This provides greater confidence and reassurance when it comes to
manual testing.

#### API (query)

The API query package, is devised to set up a series of routes that can then be
farmed off to the controllers. The controllers are created at runtime as a
simple facade around logic. The dependencies for the controllers are passed into
the controllers to prevent passing of store and template references throughout
the code, which in turn makes it easier to reason about.

#### Templates

The templates are encoded into the binary itself, but can also be viewed in
`/views` directory. They're encoded into the binary to help with distribution
of the binary, without the need for passing the view folder around.

There is an option of using the files directly from the `/views` folder by
passing the `ui.local=true` argument to the CLI. With this option enabled it
becomes easier to change the views without having to rebuild the binary
every time.

### Tests

Most of the application is tested to some degree, either via built in stdlib
testing, quick check testing (fuzz testing), mock testing and then finally
manual testing.

Of course more could be done, but with in the timeframe I personally think it's
a promising start.

#### Backend CLI Tests

Because we're using `glide` - it offers us a quick helpful command to test
everything:

```
go test -v $(glide nv)
```
