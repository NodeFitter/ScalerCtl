<h1 align="center"> ScalerCtl </h1>

<p align="center"> CLI tool to interact with NodeFitter </p>

<div align="center">
  <img alt="Static Badge" src="https://img.shields.io/badge/-1.26.5-blue?style=flat-square&logo=go&logoSize=auto&labelColor=FFF&color=00ADD8">
</div>

## About This Project

NodeFitter is a simple VM autoscaler created for the "Fog and Cloud Computing" course at <a href="https://www.unitn.it/it">University of Trento</a>, Italy.

The goal was to create a VM autoscaler capable to automatically deploy VMs and make them join a Kubernetes cluster for it to be able to schedule pods.

## Installation Instructions

For the purpose of this project the CLI is meant to be run inside a docker container alongside the
autoscaler. From the `Submission` repository, run:

```sh
cd nodefitter
docker compose up
```

To run this project by itself, install the Go language package on your device (`golag-go` on Debian-based
distributions), then run: 

```sh
go install
```

Ensure the installation path (`~/go/bin` by default) is included in your `$PATH`

## Usage

For the docker container, run:

```sh
docker exec <container> /app/scalerctl <subcommand> <args>
```

The available subcommands are:

- `get`: returns the current _local_ list of VMs. The resource usage gets updated at every scheduling
cycle, otherwise it will appear as "Not Available".
- `set`: overwrite the current threshold of free resources below which a new VM will be scheduled. The initial
values come from the `schedulerConfig.yaml` file required by the autoscaler.
    - `set cpu <percentage>` to set the free % of CPU usage (1–100).
    - `set ram <Mb>` to set the free memory in Mb.
- `stop`: pause the scheduling process.
- `start`: resume the scheduling process.
