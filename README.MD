# API Documentation

## Health Check

- **Endpoint:** `/health`
- **Method:** `GET`
- **Description:** Checks the health of the server.
- **Response:** A string "OK" if the server is running properly.

## User Model

- **ID:** A unique identifier for the user. It is a string that follows the UUID format.
- **Name:** The name of the user. It is a string.
- **Email:** The email address of the user. It is a string that follows the email format.
- **Phone:** The phone number of the user. It is a string.
- **CreditCardNumber:** The credit card number of the user. It is a string.

## Get Users

- **Endpoint:** `/users`
- **Method:** `GET`
- **Description:** Retrieves a list of users. Supports pagination through query parameters.
- **Query Parameters:**
    - `page`: The page number (optional).
    - `pageSize`: The number of users per page (optional).
- **Response:** A JSON array of users.

## Upsert User

- **Endpoint:** `/users`
- **Method:** `POST`
- **Description:** Inserts a new user or updates an existing user.
- **Request Body:** A JSON object representing a user. A new user is created if its ID is not provided, otherwise an existing user is updated.
- **Response:** A JSON object representing the created or updated user.

## Delete User

- **Endpoint:** `/users/:id`
- **Method:** `DELETE`
- **Description:** Deletes a user by ID.
- **Path Parameters:**
    - `id`: The ID of the user to delete.
- **Response:** A string "User deleted" if the user was deleted, or "User not found" if no user with the given ID exists.

## Reset User Data

- **Endpoint:** `/users/reset`
- **Method:** `POST`
- **Description:** Resets the user data by generating a new set of mock users.
- **Query Parameters:**
    - `num`: The number of mock users to generate (optional, default is 10).
- **Response:** A string "Data reset".

