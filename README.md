# Harbor: Web UI

[![Build Status](http://github.dronehippie.de/api/badges/harborapp/harbor-ui/status.svg)](http://github.dronehippie.de/harborapp/harbor-ui)
[![Coverage Status](http://coverage.dronehippie.de/badges/harborapp/harbor-ui/coverage.svg)](http://coverage.dronehippie.de/harborapp/harbor-ui)
[![Join the chat at https://gitter.im/harborapp/harbor](https://badges.gitter.im/Join%20Chat.svg)](https://gitter.im/harborapp/harbor-api)
![Release Status](https://img.shields.io/badge/status-beta-yellow.svg?style=flat)

**This project is under heavy development, it's not in a working state yet!**


## Build

This project requires NodeJS to build the sources, the installation of NodeJS
won't be covered by those instructions. To build the sources just execute the
following command after NodeJS setup:

```
make
```


## Running

We are using a simple proxy written in Go that forwards the API requests to a
server instance running whereever you want. We are open for improvements to get
a better behavior.

```
go run server.go --scheme http --addr localhost:8080
```

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
