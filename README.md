ws-healthcheck
==============

A simple executable to quickly make a websocket healthcheck.

[![Build Status](https://travis-ci.org/packetloop/ws-healthcheck.svg?branch=master)](https://travis-ci.org/packetloop/ws-healthcheck.svg?branch=master)

[Binary Releases](https://github.com/packetloop/ws-healthcheck/releases)

Usage:
-----

Standalone

```bash
HOST=localhost ORIGIN=localhost PORT=10000 ./ws-healthcheck-linux-amd64
```

With Consul/Registrator

```bash
_CHECK_INTERVAL: 3s
SERVICE_CHECK_SCRIPT: "HOST=$SERVICE_IP ORIGIN=localhost PORT=$SERVICE_PORT /usr/bin/ws-healthcheck-linux-amd64"
```