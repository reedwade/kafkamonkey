# kafkamonkey

An experiment for Kafka harsh use.

```
$ kafkamonkey
This is a tool for manipulating a kafka service -- mainly for creating stressful situations.

Usage:
  kafkamonkey [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  ping        test kafka broker connectivity
  produce     generate a lot of messages and send them to your kafka
  topics      list the topics and in sync replica situation by partition

Flags:
      --broker string              broker host:port
      --ca-file string             broker ca file (default "ca.pem")
  -h, --help                       help for kafkamonkey
      --no-tls                     don't make a TLS connection
      --service-cert-file string   broker cert file (default "service.cert")
      --service-key-file string    broker key file (default "service.key")

Use "kafkamonkey [command] --help" for more information about a command.
```


## Quick Start

### Install

```bash
$ go install github.com/reedwade/kafkamonkey@latest
```

The `kafkamonkey` executable will appear in `~/go/bin/`

### Run

```bash
$ kafkamonkey --no-tls --broker host:port ping
```
or you have `ca.pem`, `service.cert`, `service.key` files

```bash
$ kafkamonkey --broker host:port ping
```

## Configuration

Broker and tls config can be set by envars, command line args
or a config file located in the current directory or the user
home directory.

```bash
$ cat <<EOF >.kafkamonkey.env
BROKER=host:port

# CA_FILE=ca.pem
# SERVICE_CERT_FILE=service.cert
# SERVICE_KEY_FILE=service.key
EOF
```
