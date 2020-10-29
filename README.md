# BSNHub Service Provider Examples

Example service provider daemons for BSNHub.

## Get Started

Every example is a complete project.

You can enter any example directory for the following instructions.

### Install

```bash
make install
```

### Configure

Configuration is required to start the service provider daemon.

The default config lies in `./config/config.yaml`. The config items can be modified by demand.

### key Management

The key needs to be provided to interact with BSNHub.

The `keys` command is intended for key management.

```bash
<service-provider> keys add [args]

<service-provider> keys show [args]

<service-provider> keys import [args]
```

### Auth

Interaction with BSNHub needs certain authorities.

Please make sure that the specified key has the appropriate permissions.

### Service Deployment

If the service has not been deployed yet, the deployment process can be performed by running the following command:

```bash
<service-provider> deploy
```

_Note:_ For deployment automation, the metadata of the service definition and binding needs to be placed into `metadata` directory.

### Start

Start the service provider daemon:

```bash
<service-provider> start [config-file]
```
