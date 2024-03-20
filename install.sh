#!/bin/bash

# Installer script for 'yourtool'

# Define the installation directory and the binary name
INSTALL_DIR=/usr/local/bin
BINARY_NAME=yourtool
BINARY_PATH=./$BINARY_NAME # Assuming the binary is in the current directory
GO_VERSION=1.18.1 # Specify the version of Go to install if not installed

# Function to check if Go is installed
check_go() {
    if ! command -v go &> /dev/null
    then
        echo "Go could not be found, installing GoLang $GO_VERSION..."
        install_go
    else
        echo "Go is already installed."
    fi
}

# Function to install Go
install_go() {
    # Download GoLang
    wget https://golang.org/dl/go$GO_VERSION.linux-amd64.tar.gz -O /tmp/go$GO_VERSION.linux-amd64.tar.gz

    # Extract and install
    sudo tar -C /usr/local -xzf /tmp/go$GO_VERSION.linux-amd64.tar.gz

    # Set Go environment variables (you may want to add this to ~/.profile or ~/.bashrc)
    export PATH=$PATH:/usr/local/go/bin
    echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.profile

    # Verify installation
    go version
}

# Check if the user has sudo privileges
if [ "$(id -u)" != "0" ]; then
    echo "This script requires superuser access."
    echo "Please run this script with 'sudo' or as root."
    exit 1
fi

# Check for Go installation
check_go

# Check if the binary file exists
if [ ! -f "$BINARY_PATH" ]; then
    echo "Error: The binary '$BINARY_PATH' does not exist!"
    exit 1
fi

# Copy the binary to the installation directory
echo "Installing ${BINARY_NAME} to ${INSTALL_DIR}"
sudo cp $BINARY_PATH $INSTALL_DIR/$BINARY_NAME

# Set the binary to be executable
sudo chmod +x $INSTALL_DIR/$BINARY_NAME

# Optional: Check if installation was successful
if [ -f "$INSTALL_DIR/$BINARY_NAME" ]; then
    echo "Installation completed successfully."
    echo "You can now use '${BINARY_NAME}' from anywhere on your system."
else
    echo "Installation failed. Please check the permissions and try again."
    exit 1
fi

# End of the script

