# Performance considerations

Performance is often important when running GitHub action runners with garm. This document shows some ways to improve the creation time of a GitHub action runner.

## GARM specific performance considerations

### Bundle the GitHub action runner

When a new instance is created by garm, it usually downloads the latest available GitHub action runner binary, installs the requirements and starts it afterwards. This can be a time consuming task that quickly adds up when a lot of instances are created by garm throughout the day. Therefore it is recommended to include the GitHub action runner binary inside of the used image.

Example steps for setting a cached runner on a linux image in LXD:

```bash
# Create a temporary instance from your base image
lxc launch <BASE_IMAGE> temp

# Enter bash inside the container
lxc exec temp -- bash

# Get and install the runner
mkdir -p /home/runner/actions-runner
cd /home/runner/actions-runner
curl -O -L https://github.com/actions/runner/releases/download/v2.320.0/actions-runner-linux-x64-2.320.0.tar.gz
# Extract the installer
tar xzf ./actions-runner-linux-x64-2.320.0.tar.gz 

# Exit the container
exit

# Stop the instance and publish it as a new image
lxc stop temp
lxc publish temp --alias BASE_IMAGE-2.320.0

# Delete the temporary instance
lxc delete temp

# Update garm to use the new image
garm-cli pool update <POOL_ID> \
  --image=BASE_IMAGE-2.320.0
```

You can read more about cached runners in the [Using Cached Runners](https://github.com/cloudbase/garm/blob/main/doc/using_cached_runners.md) documentation.

### Disable updates

By default garm configures the `cloud-init` process of a new instance to update packages on startup. To prevent this from happening (and therefore reduce the time needed to start an instance) garm can be configured accordingly.

Example to disable this on LXD provider:

```bash
garm-cli pool update <POOL_ID> \
  --extra-specs='{"disable_updates": true}'
```

## LXD specific performance considerations

### Storage driver

LXD supports various [storage drivers](https://linuxcontainers.org/lxd/docs/latest/reference/storage_drivers/) out of the box. These storage drivers support different features which influence the creation time of a new instance. Most notably check if the driver supports `Optimized image storage` and `Optimized instance creation` as these have the biggest impact on instance creation time.

If you're not sure which storage driver is currently used, check your storages with `lxc storage list`.

### Use shiftfs/idmapped mounts

Whenever a new unprivileged instance is started on LXD, its filesystem gets remapped. This is a time consuming task which depends on the image size that's being used. For large images this can easily take over a minute to complete. There are two ways to get around this: `shiftfs` or `idmapped mounts`. While the latter is the preferred one, not all filesystems currently support it, so in most cases enabling `shiftfs` show a significant performance improvement.

Example on how to enable it on a snap installed LXD:

```bash
snap set lxd shiftfs.enable=true
systemctl reload snap.lxd.daemon
```

Some details and discussions around `shiftfs` can be found [here](https://discuss.linuxcontainers.org/t/trying-out-shiftfs/5155).

Note: When `shiftfs` is used, mounting between host and container might need some extra steps to be secure. See [here](https://discuss.linuxcontainers.org/t/share-folders-and-volumes-between-host-and-containers/7735) for details.