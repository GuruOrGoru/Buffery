# Programming Learning App

A cross-platform mobile application that teaches coding interactively through lessons, quizzes, and code execution, powered by a Go backend and React Native frontend.

## Features

- **Interactive Lessons**: Step-by-step coding tutorials with Markdown content
- **Quizzes**: Multiple-choice questions to test understanding
- **Code Execution**: Run code snippets with real-time feedback via sandbox
- **Progress Tracking**: Monitor learning progress with streaks and XP
- **Gamification**: Badges, leaderboards, and achievements
- **Offline Support**: Cache lessons and quizzes for offline access
- **AI Hints**: Optional AI-powered code explanations (future enhancement)

## Architecture

- **Backend**: Go with Gin framework, PostgreSQL database, JWT authentication
- **Mobile**: React Native with Expo, file-based routing, secure storage
- **Database**: PostgreSQL with migrations
- **Deployment**: Dockerized backend, Expo EAS for mobile builds

## Project Structure

```
programmingCenter/
├── backend/          # Go API server
│   ├── cmd/
│   ├── internal/
│   ├── migrations/
│   └── README.md
├── mobile/           # React Native app
│   ├── app/
│   ├── components/
│   └── README.md
└── README.md         # This file
```

## Getting Started

### Prerequisites

- Go 1.21+
- Node.js 18+
- PostgreSQL
- Expo CLI

### Setup

1. Clone the repository:
   ```bash
   git clone https://github.com/yourname/programmingCenter.git
   cd programmingCenter
   ```

2. Setup backend:
   ```bash
   cd backend
   go mod download
   cp .env.example .env  # Configure your environment variables
   # Run migrations and start server
   make run
   ```
   See [backend/README.md](backend/README.md) for detailed instructions.

3. Setup mobile app:
   ```bash
   cd ../mobile
   npm install
   npx expo start
   ```
   See [mobile/README.md](mobile/README.md) for detailed instructions.

## API Documentation

The backend provides RESTful APIs for authentication, lessons, quizzes, submissions, and progress tracking. Full API docs available at `/swagger` when running the backend.

## Deployment

- **Backend**: Deploy to Render, Fly.io, or Supabase Edge Functions
- **Mobile**: Build APKs/IPAs with Expo EAS, deploy to app stores

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make changes and add tests
4. Submit a pull request

## License

MIT License - see LICENSE file for details.

## Roadmap

- Multi-language support (Python, JavaScript)
- AI tutor chat
- Web dashboard
- Daily challenges

For development progress and daily logs, see the project roadmap in the repository.