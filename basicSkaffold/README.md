# Using Skaffold

Skaffold is an interesting piece of tool that allows for faster iteration of development for applications that are meant to be deployed to Kubernetes.

Refer to the following link for more details:
https://github.com/GoogleContainerTools/skaffold

## Running the project

The following line would do the following:

- Do a multi stage build of a golang web application
- Create the deployment and service resources on a local kubernetes cluster
- Streams logs from the running instance
- Watch the files specified in the `skaffold.yaml` file. Essentially, the files within the Docker context as well as the kubernetes resource configuration files. On change, it would rebuild and run the above steps (except with cache if possible)

```bash
skaffold dev
```

This is on the assumption that the user has already installed both docker and kubernetes already
