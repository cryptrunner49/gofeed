# API Technical Design Document

## Introduction

The Game API serves as a optional interface for interacting with and modifying elements within the game. This technical design document provides a comprehensive overview of the architecture, components, APIs, and data flow associated with the Game API. It aims to offer a clear understanding of how the API functions, its key features, and how it integrates with the game's ecosystem.

## Architecture Overview

![architecture overview](architecture_overview.drawio.png)

### API

The **API** serves as an optional interface for players to interact with the system. It is designed to be consumed directly by using tools like curl, postman or custom scripts without the need for a user interface or frontend application. The API will handle requests, process logic, and communicate with the **API Database** to enable a cheat or with the **Game** to make some change directly into the game.

### API Database

The **API Database** is the persistent storage layer that the API interacts with directly. It is specifically designed to store cheat-related data, including the current state of each cheat (enabled or disabled). This database support the API's functionality.

### Game

The **Game** is the core application where gameplay occurs. It includes all game logic, mechanics, and real-time interactions. The game communicates with the **Game Database** to save and retrieve player progress, ensuring a seamless experience across sessions.

### Game Database

The **Game Database** is the persistent storage layer dedicated to the game. It stores all player-specific data, such as saved progress, achievements, and in-game states.

## API Specification

### /player (Modify the player)

#### 1. **Endpoint URL**

- **URL:** <http://localhost:8844/player>
- **HTTP Method:** PUT

#### 2. **Payload**

- **Headers:**

  ```json
  {
    "Content-Type": "application/json",
    "Authorization": "Bearer YOUR_ACCESS_TOKEN"
  }
  ```

- **Body:**

  ```json
  {
    "class": "value",
  }
  ```

#### 3. **JSON Response**

- **Success Response (200 OK):**

  ```json
  {
    "status": "success",
    "data": {
      "class": "value",
    },
    "message": "The player class was successfully updated."
  }
  ```

- **Error Response (4xx/5xx):**

  ```json
  {
    "status": "error",
    "message": "Detailed error message",
    "code": "ERROR_CODE"
  }
  ```

#### 4. **Example Request**

- **cURL Example:**

  ```sh
  curl -X PUT "http://localhost:8844/player" \
       -H "Content-Type: application/json" \
       -H "Authorization: Bearer YOUR_ACCESS_TOKEN" \
       -d '{
             "class": "archer",
           }'
  ```