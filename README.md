# epicshelter

[![Built with Devbox](https://www.jetify.com/img/devbox/shield_moon.svg)](https://www.jetify.com/devbox/docs/contributor-quickstart/)

> "I swear I never did that when I went about finding a name for my backup system. I swear that I just rolled the bones and came up with EPICSHELTER."
>
> -- <cite>Edward Snowden</cite>

## Intro

epicshelter is a cloud-agnostic infrastructure architecture that enables scalable and resilient application deployment across any cloud provider. Built on GitOps principles with Kubernetes, ArgoCD, and Crossplane, it provides a unified development experience from local clusters to production environments while maintaining complete portability and avoiding vendor lock-in.

## Getting started

To get started, the following dependencies need to be installed:

- [AWS CLI](https://docs.aws.amazon.com/cli/latest/userguide/getting-started-install.html)
- [Devbox](https://www.jetify.com/docs/devbox/)
- [Docker](https://docs.docker.com/manuals/)

If you'd rather not use Devbox, you can manually install the dependencies listed in the `devbox.json` file under the "packages" section.

## Documentation

Detailed documentation is available on the following components:

- [Infrastructure Architecture](docs/infra-architecture.md)

## Contributing

Feel free to open any issues on pull requests!

