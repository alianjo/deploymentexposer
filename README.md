# Kubernetes Controller README
=====================================

Table of Contents
-----------------

- [Kubernetes Controller README](#kubernetes-controller-readme)
  - [Table of Contents](#table-of-contents)
  - [Overview](#overview)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
    - [Option 2: Using YAML manifests](#option-2-using-yaml-manifests)
  - [Usage](#usage)
    - [Running Locally](#running-locally)
    - [Creating a Custom Resource](#creating-a-custom-resource)
    - [Interacting with the Custom Resource](#interacting-with-the-custom-resource)
  - [TODO](#todo)

## Overview

 The Kubernetes controller is designed to automate the creation of Kubernetes ingress and service resources for each Deployment created in the cluster. Its primary purpose is to simplify the process of exposing Deployments to external traffic, thereby streamlining the deployment and management of applications in a Kubernetes environment.

## Prerequisites

* Kubernetes cluster (version v1.28.4 or higher)

## Installation

### Option 2: Using YAML manifests

1. Clone the repository: `git clone <repo-url>`
2. Apply the manifests: `kubectl apply -f <manifests-path>`

## Usage
### Running Locally

1. Clone the repository: `git clone <repo-url>`
2. Build the controller: `go build -o controller main.go`
3. Run the controller: `./controller -kubeconfig ~/.kube/config`
4. Verify the controller is running: `kubectl get deployments -n <namespace>`

### Creating a Custom Resource

1. Create a YAML file with the custom resource definition: `apiVersion: <group>.<version> kind: <kind> metadata: name: <name> spec: ...`
2. Apply the YAML file: `kubectl apply -f <yaml-file>`

### Interacting with the Custom Resource

## TODO 
- [ ] Complete README
- [ ] Add manifests related to the controller and examples
