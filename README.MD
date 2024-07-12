# API Documentation

## Health Check

- **Endpoint:** `/health`
- **Method:** `GET`
- **Description:** Checks the health of the server.
- **Response:** A string "OK" if the server is running properly.

## User Struct Explanation

The User struct represents a user with various attributes:

- **ID**: A unique identifier for the user, represented as a string. It uses the `uuid_digit` format.
- **Name**: The name of the user, represented as a string. It uses the `name` format.
- **Email**: The user's email address, represented as a string. It uses the `email` format.
- **Phone**: The user's phone number, represented as a string. It uses the `phone_number` format.
- **CreditCardNumber**: The user's credit card number, represented as a string. It uses the `cc_number` format.
- **Avatar**: The user's avatar, represented as a string. It does not have a specific format.
- **JoinedDate**: The date the user joined, represented as an integer in Unix time format. It uses the `unix_time` format.
- **Age**: The user's age, represented as an integer. It can be one of the following values: 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, or 33.

## Get Users

- **Endpoint:** `/users`
- **Method:** `GET`
- **Description:** Retrieves a list of users. Supports pagination through query parameters.
- **Query Parameters:**
  - `page`: The page number (optional).
  - `pageSize`: The number of users per page (optional).
- **Response:** A JSON object containing an array of users and pagination details.

## Upsert User

- **Endpoint:** `/users`
- **Method:** `POST`
- **Description:** Inserts a new user or updates an existing user. If the user ID is provided, the existing user is updated. If the user ID is not provided, a new user is created.
- **Request Body:** A JSON object representing a user.
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