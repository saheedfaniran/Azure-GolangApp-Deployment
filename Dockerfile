# Start from the official Go Docker image (Defines a base for the image)
FROM golang:1.16 AS build

# Set the working directory inside the container
WORKDIR /app

# Copy the Go modules files
COPY go.mod go.sum ./

# Download the Go modules (Executes and commit the go application commands) 
RUN go mod download

# Copy the application source code
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -o app

# Start a new stage with a minimal base image
FROM scratch

# Copy the application binary from the previous build stage
COPY --from=build /app/app /app
#Expose port 1337
EXPOSE 1337

#Declare variable
ENV CANDIDATE_FIRSTNAME=Saheed

# Set the entrypoint to run the application binary
ENTRYPOINT ["/app"]


