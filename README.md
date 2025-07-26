# BeBot - WhatsApp Order Assistant

A WhatsApp chatbot built with Go that helps users place orders and manage their profiles through WhatsApp messaging.

## Features

- **User Registration**: Create and manage user profiles with name, address, and store preferences
- **Order Management**: Place orders through WhatsApp messages
- **Image Generation**: Generate images using AI APIs
- **API Integration**: Connect with external APIs for jokes, deployment status, and chat responses
- **Multi-store Support**: Support for multiple store options (ABC, BKC, MLP stores)

## Prerequisites

- Go 1.22 or higher
- SQLite3 (for local database storage)
- WhatsApp account for bot authentication

## Installation

1. Clone the repository:
```bash
git clone <repository-url>
cd OrderAssist
```

2. Install dependencies:
```bash
go mod download
```

3. Build the application:
```bash
go build -o main .
```

## Usage

### Running the Application

1. Start the bot:
```bash
./main
```

2. Scan the QR code with your WhatsApp mobile app to authenticate the bot

3. Start chatting with the bot using these commands:

### Available Commands

#### User Registration
- Send `hello`, `hi`, `hey`, or `hello there` to start the registration process
- Follow the prompts to provide:
  - Name: `name-<your name>`
  - Address: `address-<your address>`
  - Store ID: `storeid-<1 or 2 or 3>`

#### Order Placement
- Place an order: `order:<order details>`

#### Image Generation
- Generate an image: `Image of <description>`

#### API Features
- Get deployment status: Send any message containing "deploy"
- Get a joke: Send any message containing "joke"
- Chat with AI: Send any message containing "plate"

## Environment Variables

- `GROQ_API_KEY`: API key for Groq AI chat integration

## Development

### Project Structure

```
OrderAssist/
├── commands/          # Command handlers for different features
├── handlers/          # WhatsApp event handlers
├── utils/            # Utility functions
├── whatsapp/         # WhatsApp client management
├── main.go           # Application entry point
├── go.mod            # Go module file
└── Dockerfile        # Docker configuration
```

### Building

```bash
# Build for current platform
make build

# Build for all platforms
make build-all

# Run with hot reload
make run
```

### Testing

```bash
# Run tests
make test

# Generate coverage report
make cover-html
```

## Docker

Build and run with Docker:

```bash
# Build the image
docker build -t bebot .

# Run the container
docker run -it bebot
```

## API Endpoints

The bot integrates with several external APIs:

- **User Management**: `https://oms-bebot-backend.onrender.com/api/user/`
- **Order Management**: `https://oms-bebot-backend.onrender.com/api/orders/`
- **Image Generation**: `https://openai-image-7la8.onrender.com/image/generate`
- **Should I Deploy**: `https://shouldideploy.today/api`
- **Joke API**: `https://v2.jokeapi.dev/joke/`
- **Groq AI Chat**: `https://api.groq.com/openai/v1/chat/completions`

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests if applicable
5. Submit a pull request

## License

This project is licensed under the MIT License.
