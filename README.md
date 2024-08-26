
# go-streamer

![go-streamer Logo](./assets/logo.png)

go-streamer is a simple and opinionated tool to stream logs via kafka

## Features

- **Send Messages**
- **Receive Messages**
- **Flexible Configuration**: Configure environment variables and config files.

## Installation

To install KafkaStreamer, use `go get`:

```sh
go get -u github.com/sivaramsajeev/log_streamer/cmd/go-streamer
```


## Usage

go-streamer provides two primary commands: `send` and `receive`. The `--help` flag is available for detailed usage.


Example:

```sh
export CONFIG_PROPERTIES_FILE=/path/to/client.properties
export LOG_FILE_PATH=/path/to/logfile
export CONFIG_TOPIC_NAME=<name_of_the_topic>

go-streamer send

go-streamer receive 

```



