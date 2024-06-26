![release](https://img.shields.io/github/v/release/project-n-oss/sidekick-router)

# sidekick-router

sidekick-router is a [sidecar](https://learn.microsoft.com/en-us/azure/architecture/patterns/sidecar) proxy process that help redirect or reduce chances of overwrites in native crunch.

## Getting started

- go1.21

## Running sidekick-router

### Local

You will need to create a config.yml file in the root of the project. You can use the following template:

```yaml
App:
  CloudPlatform: AWS
```

You can then run sidekick-router directly from the command line:

```bash
go run main.go serve
```

This will run sidekick-router localy on your machine on `localhost:7075`.

run the following command to learn more about the options:

```bash
go run main.go serve --help
```

## Using sidekick-router

### Docker

You can pull the docker image from the [containers page](https://github.com/project-n-oss/sidekick-router/pkgs/container/sidekick-router)

You can then run the docker image with the following command:

```bash
docker run -p 7075:7075 --env SIDEKICKROUTER_APP_CLOUDPLATFORM=AWS <sidekick-router-image> serve 
```


### Pre Built binaries

sidekick-router binaries are hosted and released from GitHub. Please check our [releases page](https://github.com/project-n-oss/sidekick-router/releases).
To download any release of our linux amd64 binary run:

```bash
wget https://github.com/project-n-oss/sidekick-router/releases/download/${release}/sidekick-router-linux-amd64.tar.gz
```

You can then run the binary directly:

```bash
SIDEKICKROUTER_APP_CLOUDPLATFORM=AWS ./sidekick-router serve
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
