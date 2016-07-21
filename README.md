# Umschlag: Web UI

[![Build Status](http://github.dronehippie.de/api/badges/umschlag/umschlag-ui/status.svg)](http://github.dronehippie.de/umschlag/umschlag-ui)
[![Coverage Status](http://coverage.dronehippie.de/badges/umschlag/umschlag-ui/coverage.svg)](http://coverage.dronehippie.de/umschlag/umschlag-ui)
[![Join the chat at https://gitter.im/umschlag/umschlag](https://badges.gitter.im/Join%20Chat.svg)](https://gitter.im/umschlag/umschlag)
![Release Status](https://img.shields.io/badge/status-beta-yellow.svg?style=flat)

**This project is under heavy development, it's not in a working state yet!**

Where does this name come from or what does it mean? It's quite simple, it's one
german word for transshipment, I thought it's a good match as it is related to
containers and a harbor.


## Build

This project requires NodeJS to build the sources, the installation of NodeJS
won't be covered by those instructions. To build the sources just execute the
following command after NodeJS setup:

```
npm install
npm run build
```

If you also want to publish it as a single binary with our server based on Go
make sure you have a working Go environment, for further reference or a guide
take a look at the [install instructions](http://golang.org/doc/install.html).
As this project relies on vendoring of the dependencies and we are not
exporting `GO15VENDOREXPERIMENT=1` within our makefile you have to use a Go
version `>= 1.6`

```bash
go get -d github.com/umschlag/umschlag-ui
cd $GOPATH/src/github.com/umschlag/umschlag-ui
make generate build

bin/umschlag-ui -h
```

With the `make generate` command we are embedding all the static assets into the
binary so there is no need for any webserver or anything else beside launching
this binary.


## Development

To start developing on this UI you have to execute only a few commands. To setup
a NodeJS environment or even a Go environment is out of the scope of this
document. To start development just execute those commands:

```
npm install
npm run start -- --host localhost:8080
```

The development server proxies all requests to the define host. So in order to
properly work with it you need to start the API separately.

After launching this command on a terminal you can access the web interface at
[http://localhost:9000](http://localhost:9000)


## Contributing

Fork -> Patch -> Push -> Pull Request


## Authors

* [Thomas Boerger](https://github.com/tboerger)


## License

Apache-2.0


## Copyright

```
Copyright (c) 2016 Thomas Boerger <http://www.webhippie.de>
```
