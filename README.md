# stocksbeat

Welcome to stocksbeat.

Ensure that this folder is at the following location:
`${GOPATH}/src/github.com/tuxiedev/stocksbeat`

## Development

### Requirements

* [Golang](https://golang.org/dl/) 1.7

### Clone

To clone stocksbeat from the git repository, run the following commands:

```
mkdir -p ${GOPATH}/src/github.com/tuxiedev/stocksbeat
git clone https://github.com/tuxiedev/stocksbeat ${GOPATH}/src/github.com/tuxiedev/stocksbeat
```

### Init Project
To get running with stocksbeat and also install the
dependencies, run the following command:

```
make setup
```

It will create a clean git history for each major step. Note that you can always rewrite the history if you wish before pushing your changes.

### Build

To build the binary for stocksbeat run the command below. This will generate a binary
in the same directory with the name stocksbeat.

```
make
```


### Run

To run stocksbeat with debugging output enabled, run:

```
FINNHUB_TOKEN=<your_finnhub_token> ./stocksbeat -c stocksbeat.yml -e -d "*"
```


### Cleanup

To clean  stocksbeat source code, run the following command:

```
make fmt
```

To clean up the build directory and generated artifacts, run:

```
make clean
```


## Packaging

The beat frameworks provides tools to crosscompile and package your beat for different platforms. This requires [docker](https://www.docker.com/) and vendoring as described above. To build packages of your beat, run the following command:

```
make release
```

This will fetch and create all images required for the build process. The whole process to finish can take several minutes.
