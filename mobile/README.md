# Mobile App

React Native application for the Programming Learning App, providing an interactive mobile experience for coding education.

## Tech Stack

- **Framework**: React Native with Expo
- **Navigation**: React Navigation (file-based routing)
- **State Management**: React hooks + Context
- **Storage**: AsyncStorage for offline data, SecureStore for tokens
- **UI**: Custom themed components
- **Markdown**: react-native-markdown-display
- **Code Editor**: TextInput with optional syntax highlighting
- **Build**: Expo EAS

## Setup

### Prerequisites

- Node.js 18+
- Expo CLI (`npm install -g @expo/cli`)
- iOS Simulator (macOS) or Android Emulator

### Installation

1. Navigate to mobile directory:
   ```bash
   cd mobile
   ```

2. Install dependencies:
   ```bash
   npm install
   ```

3. Start development server:
   ```bash
   npx expo start
   ```

4. Run on device/emulator:
   - Press `i` for iOS simulator
   - Press `a` for Android emulator
   - Scan QR code with Expo Go app

## App Structure

```
mobile/
├── app/              # File-based routing screens
│   ├── (tabs)/       # Tab navigation
│   ├── modal.tsx     # Modal screens
│   └── _layout.tsx   # Root layout
├── components/       # Reusable UI components
│   ├── ui/           # Base UI components
│   └── ...           # Feature components
├── constants/        # Theme and config
├── hooks/            # Custom React hooks
└── assets/           # Images and icons
```

## Key Screens

- **Auth**: Login/Register with JWT storage
- **Home**: Lesson list with progress overview
- **Lesson**: Markdown content viewer
- **Quiz**: Interactive MCQ interface
- **Code Editor**: Text input for code execution
- **Profile**: User progress, badges, leaderboard

## Features

- **Offline Mode**: Cache lessons and quizzes
- **Progress Sync**: Automatic sync when online
- **Gamification**: XP, streaks, achievements
- **Code Execution**: Submit code to backend sandbox
- **AI Hints**: Optional AI assistance (future)

## Building for Production

1. Configure EAS:
   ```bash
   npx eas build:configure
   ```

2. Build APK/IPA:
   ```bash
   npx eas build --platform android
   npx eas build --platform ios
   ```

3. Submit to stores:
   ```bash
   npx eas submit --platform android
   ```

## Environment

Create `.env` file for API base URL and other configs.

## Testing

Run tests:
```bash
npm test
```

## Development Notes

- Use `SecureStore` for sensitive data
- Implement proper error handling and loading states
- Follow React Native performance best practices
- Test on multiple devices/screen sizes

See root [README.md](../README.md) for project overview and backend setup.