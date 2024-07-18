# Event Based Digital Signage System
## Project Overview

The project aims to display trending events on digital signage by fetching data from Twitter.
This repository facilitates backend of the system, which is implemented in Go, and it interacts with Firestore to store user credentials.

## Technologies Used

- **Backend**: Go, Firestore
- **Hosting**: Google Cloud Platform (GCP)


### Setup
1. **Clone the Repository:**

   ```bash
   git clone https://github.com/mayani-ak/signage-system.git
   cd signage-system
   ```

2. **Set Up Firestore:**
    - Create a project in the [Firebase Console](https://console.firebase.google.com/).
    - Create a Firestore database and add a collection 'users'
    - Download the service account key JSON file and place it in the backend directory.

3. **Install Go Dependencies:**

   ```bash
   go mod tidy
   ```

4. **Set Environment Variables:**

   ```env
   GOOGLE_CLOUD_PROJECT=<gcp-project-id>
   JWT_SECRET_KEY=<your-jwt-secret-key>
   WEATHER_API_KEY=<for-weatherapi.com>
   TWITTER_BEARER_TOKEN=<twitter-bearer-token>
   ```

5. **Run the Backend Server:**

   ```bash
   go run main.go
   ```

## Running the Application

1. **Start the server**:
    ```sh
    go run main.go
    ```

2. The server will start on port `8080` by default. You can access it at `http://localhost:8080`.

## API Endpoints

### Public Routes

- `POST /signup`
   - Endpoint for user signup.
   - **Request Body**:
     ```json
     {
       "username": "your-username",
       "password": "your-password"
     }
     ```
   - **Response**:
     ```json
     {
       "message": "Signup successful",
       "token": "your-jwt-token"
     }
     ```

- `POST /login`
   - Endpoint for user login.
   - **Request Body**:
     ```json
     {
       "username": "your-username",
       "password": "your-password"
     }
     ```
   - **Response**:
     ```json
     {
       "message": "Login successful",
       "token": "your-jwt-token"
     }
     ```

### Protected Routes

- `GET /restricted/content`
   - Endpoint to fetch content.
   - **Headers**:
      - `Authorization: Bearer your-jwt-token`
   - **Response**:
     ```json
     {
       "content": "Your content here"
     }
     ```

- `POST /restricted/update`
   - Endpoint to update content.
   - **Headers**:
      - `Authorization: Bearer your-jwt-token`
   - **Request Body**:
     ```json
     {
       "content": "New content"
     }
     ```
   - **Response**:
     ```json
     {
       "id": "uuid",
       "created_time": "",
       "location": "",
       "weather": "",
       "social_media": "" 
     }
     ```

## Current Implementation

1. Backend APIs to Register a new user and Login
2. External API calls to fetch trending tweets (for Seattle location) using a Twitter API, and current weather using weather.com API

## Next Implementation
1. Frontend using Vue.js
2. Allow user to specify location through UI to fetch trending tweets

## Some Possible Enhancements For A Real System
1. Implement OAuth for more secure user authentication and authorization, 
allowing users to log in using their Google, Twitter, or Facebook accounts.
2. Event Sources
   1. Multiple Social Media Integrations: Extend the system to fetch events from other social media platforms like Instagram and Facebook.
   2. Custom Event Feeds: Allow users to integrate custom event feeds via RSS or APIs.
3. UI/UX
   1. Add support for interactive digital signage
   2. Dynamic content and elements to fit in different size of displays
4. Admin Dashboard
   1. Allow Admin user to manage user accounts and display settings.
   2. Allow scheduling events for specified date-time
5. Load balancing and caching
6. Deployment on Kubernetes
7. CI/CD pipeline
