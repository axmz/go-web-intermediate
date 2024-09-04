FROM mcr.microsoft.com/devcontainers/go:1.22

RUN go install github.com/air-verse/air@latest
RUN go install github.com/gobuffalo/pop/v6/soda@latest
# Change ownership of the /go directory to vscode user and set permissions
RUN chown -R vscode:vscode /go && chmod -R u+rwx /go
RUN git config --global --add safe.directory /go/src/app

# Set the working directory
WORKDIR /go/src/app
# RUN go mod 

EXPOSE 4000, 4001, 5000

# Keep the container running
CMD ["tail", "-f", "/dev/null"]
