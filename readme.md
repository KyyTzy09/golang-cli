# KYY Agent CLI

KYY Agent is a powerful Command Line Interface (CLI) application built with **Golang**. It's designed to assist with various development tasks, acting as an intelligent agent right from your terminal.

## Features

- **Ask**: Get quick answers to your questions.
- **Chat**: Engage in interactive conversations with the AI.
- **Config**: Manage and update application settings.
- **Scan**: Analyze files and project structures.

## Project Structure

The project is organized into two main top-level directories:

- `cmd/`: Contains the main entry points for the CLI commands.
- `internal/`: Houses internal packages and logic not meant for external consumption.

## Installation

```bash
# Example: How to install your CLI tool
go install github.com/KyyTzy09/golang-cli@latest
```

## Usage

```bash
# Example: Ask a question
kyy ask "How do I compile a Go program?"

# Example: Start a chat session
kyy chat

# Example: Configure settings
kyy config --gemini YOUR_API_KEY

# Example: Scan a file or directory
kyy scan ./internal/cli/

# Example: Use tools like scan or create file
kyy tools "Create test.txt in root project"

```