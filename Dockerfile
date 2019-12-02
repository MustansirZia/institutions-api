ARG EXECUTABLE_DIR=/insitutions
ARG EXECUTABLE_NAME=insitutions-api
ARG EXECUTABLE_PATH=${EXECUTABLE_DIR}/${EXECUTABLE_NAME}

FROM golang:1.12 as builder

ARG EXECUTABLE_DIR
ARG EXECUTABLE_NAME
ARG EXECUTABLE_PATH

WORKDIR ${EXECUTABLE_DIR}

# Copy go mod and sum files.
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed.
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container.
COPY . .

# Build the executable.
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ${EXECUTABLE_NAME} .

# Use a tiny base image for running our application.
FROM alpine:latest 

ARG EXECUTABLE_DIR
ARG EXECUTABLE_PATH
ARG EXECUTABLE_NAME
ENV EXECUTABLE_NAME=${EXECUTABLE_NAME}

ENV PORT=5000

WORKDIR ${EXECUTABLE_DIR}

RUN apk --no-cache add ca-certificates

# Copy the Pre-built executable file from the previous stage.
COPY --from=builder ${EXECUTABLE_PATH} ${EXECUTABLE_PATH}

# Copy json files inside this image.
COPY data data

# Run our application.
CMD "./${EXECUTABLE_NAME}"
