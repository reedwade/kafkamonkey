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

```bash
$ export BROKER=host:port
$ kafkamonkey topics
```

You may want to set your kafka with topic auto-create enabled.

```bash
$ kafkamonkey produce
```

# Produce a lot of messages

Kafkamonkey likes to produce messages. It uses a worker pool
to send batches of any size of message you'd like.
You can set the message size, batch count and worker pool
size with command line args.

The default topic is `monkey`. You can include `NOW` in the topic to
get a timestamp (will be the same for all in a run).
You can use `BATCHID` for the batch number if you want a distinct topic
for each message batch. This is helpful for replicating the case where
sometime people like to have way too many topics in their kafka.

```bash
$ kafkamonkey produce \
    --topic NOW-monkey \
    --batches 100 \
    --message-length 1000 \
    --messages-per-batch 1000
```

```bash
$ kafkamonkey produce \
    --topic NOW-BATCHID \
    --batches 10000
```

# Configuration

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
