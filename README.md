# Kafka and Vector Integration Using Docker Compose and Golang

This repository demonstrates how to set up Kafka and Vector for log management, with a Golang producer program sending messages to Kafka and Vector consuming those messages and outputting them to a file and the console.

## Prerequisites

To run and use this project, ensure that you have the following installed:

- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)
- [Golang](https://golang.org/)

## Project Structure

```text
.
├── docker-compose.yml          # Docker Compose file for services
├── vector_config.toml          # Configuration file for Vector
├── main.go                     # Golang program to produce messages to Kafka
├── vector_output/              # Directory containing Vector output log files
└── README.md                   # Project documentation
```

## Services

### 1. Kafka
Kafka is used as a message broker. The topic used in this project is `example-topic`. The Kafka setup uses **Bitnami Kafka** and **Zookeeper** Docker images.

### 2. Vector
Vector collects messages from Kafka (`example-topic`) and writes logs into a file as well as the console. The configuration for Vector is stored in `vector_config.toml`.

### 3. Golang Producer
The Golang script (`main.go`) acts as a producer to send messages to the `example-topic` on Kafka.

## How It Works

1. **Kafka Producer**:
   - The Golang `main.go` file creates and sends messages to Kafka topic `example-topic` every second.
   - The producer connects to Kafka exposed port `9092`.

2. **Vector Consumer**:
   - Vector consumes messages from the Kafka topic `example-topic` using the configuration defined in `vector_config.toml`.
   - Consumed logs are written to:
     - A file: `vector_output/output.log`. Create the vector_output dir before running the setup.
     - The console

3. **Integration**:
   - Components are connected via Docker Compose which defines how Kafka, Vector, and Golang work together.

## Setup and Usage

### Step 1: Clone the Repository
Clone this repository to your local machine:
```bash
git clone <your-repo-url>
cd <repository-folder>
```

### Step 2: Start the Docker Services
Use `docker-compose` to spin up Kafka, Zookeeper, and Vector:
```bash
docker-compose up
```

### Step 3: Run the Golang Kafka Producer
Run the Golang producer program to send messages to Kafka:
```bash
go run main.go
```

### Step 4: Check the Vector Output
1. **File Output**: Open the generated log file in the `vector_output/` directory:
   ```bash
   cat vector_output/output.log
   ```
2. **Console output**: check vector container log to see the output.


## Configuration Details

### Docker Compose
The `docker-compose.yml` defines the following services:
- **Kafka**: Configured with topic `example-topic`, available on port `9092`.
- **Vector**: Reads logs from Kafka and outputs them to `vector_output/output.log`.

### Vector Configuration
The `vector_config.toml` file specifies:
- Kafka topic: `example-topic`
- Output file: `/var/lib/vector/output.log`
- Console logging

### Golang Producer
The producer connects to Kafka on `localhost:9092` using the [Sarama library](https://github.com/IBM/sarama).
It produces messages to the `example-topic` topic every second.

## References

- [Kafka Documentation](https://kafka.apache.org/documentation/)
- [Vector Documentation](https://vector.dev/docs/)
- [Sarama Library (Golang)](https://github.com/IBM/sarama)

## License

This project is licensed under the [MIT License](LICENSE).