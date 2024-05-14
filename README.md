![release](https://img.shields.io/github/v/release/project-n-oss/interchange)

# Interchange

Interchange is a [sidecar](https://learn.microsoft.com/en-us/azure/architecture/patterns/sidecar) proxy process that help redirect or reduce chances of overwrites in native crunch.

## Getting started

- go1.21

## Running Interchange

### Local

You will need to create a config.yml file in the root of the project. You can use the following template:

```yaml
App:
  CloudPlatform: AWS
```

You can then run Interchange directly from the command line:

```bash
go run main.go serve
```

This will run Interchange localy on your machine on `localhost:7075`.

run the following command to learn more about the options:

```bash
go run main.go serve --help
```

## Using Interchange

### Docker

You can pull the docker image from the [containers page](https://github.com/project-n-oss/interchange/pkgs/container/interchange)

You can then run the docker image with the following command:

```bash
docker run -p 7075:7075 --env INTERCHANGE_APP_CLOUDPLATFORM=AWS <interchange-image> serve 
```


### Pre Built binaries

Interchange binaries are hosted and released from GitHub. Please check our [releases page](./releases).
To download any release of our linux amd64 binary run:

```bash
wget https://github.com/project-n-oss/interchange/releases/download/${release}/interchange-linux-amd64.tar.gz
```

You can then run the binary directly:

```bash
INTERCHANGE_APP_CLOUDPLATFORM=AWS ./interchange serve
```

## Contributing

### Versioning

This repository uses [release-please](https://github.com/google-github-actions/release-please-action) to create and manage release.

### Commits

We follow [conventional commits](https://www.conventionalcommits.org/en/v1.0.0/) for our commits and PR titles. This allows us to use release-please to manage our releases.

The most important prefixes you should have in mind are:

- fix: which represents bug fixes, and correlates to a SemVer patch.
- feat: which represents a new feature, and correlates to a SemVer minor.
- feat!:, or fix!:, refactor!:, etc., which represent a breaking change (indicated by the !) and will result in a SemVer major.
