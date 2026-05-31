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
go install github.com/your-username/kyy-agent-cli@latest
```

## Usage

```bash
# Example: Ask a question
kyy-agent ask "How do I compile a Go program?"

# Example: Start a chat session
kyy-agent chat

# Example: Configure settings
kyy-agent config set openai_api_key YOUR_API_KEY

# Example: Scan a file or directory
kyy-agent scan ./internal/cli/
```
