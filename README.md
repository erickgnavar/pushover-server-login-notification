# Login notifier

Sent a push notification when someone login to your remote server.

## Configuration

Setup these environment variables to run the program

```shell
PUSHOVER_APP_TOKEN=""
PUSHOVER_USER_KEY=""
```

These variables will be used when the program is compiled into a unique binary. The values will be in the compiled binary.

## Compile for the current arch

```shell
make build
```

## Compile for linux amd64

```shell
make build_linux
```

## Compile for linux arm, raspberry

```shell
make build_arm
```

## Copy the binary to the remote host

```shell
scp main remote_host:/home/$USER/
```

## Setup the binary to run at login

Add `$HOME/main` at the end of `~/.profile` or `.bash_profile`.
