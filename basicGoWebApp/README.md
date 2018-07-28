# Basic Go Web Application

Contains sample Go Web application

- Container built locally and in a local docker directory
- Hosted on Kubernetes

# How to run

Ensure that there is a local registry running. The following codebase

To do an initial application startup

```bash
make version=v1 build
make init
```

To update the application

```bash
make version=v2 update
```

To cleanup the applicaton

```bash
make cleanup
```
