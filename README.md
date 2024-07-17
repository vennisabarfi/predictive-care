# African Proverbs Weekly Email API

This project is an API that sends African proverbs to users weekly by email. It includes user authentication and provides endpoints for registering, logging in, and viewing proverbs. The proverbs are sent out via a scheduled cron job.

## Features

- **User Registration**: Allows users to register with a username, email, and password.
- **User Login**: Enables users to log in and receive a JWT token for authentication.
- **User Logout**: Provides functionality for users to log out by clearing the JWT token.
- **View Proverbs**: Allows users to view all proverbs or a specific proverb by ID.
- **Weekly Proverbs Email**: Sends a weekly email with a random proverb to all registered users.

## Project Structure

```
.
├── main.go
├── handlers
│   └── handlers.go
├── models
│   └── models.go
├── storage
│   └── storage.go
├── cron
│   └── cron.go
├── proverbs_only.json
├── proverbs_data.json
├── proverbs_cleaned.json
├── meanings_data.json
├── proverbs_and_meanings_data.csv
└── README.md
```

## Endpoints

- **`POST /register`**: Register a new user.
- **`POST /login`**: Login a user.
- **`GET /logout`**: Logout a user.
- **`GET /viewproverbs`**: View all proverbs.
- **`GET /viewproverbs/:id`**: View a specific proverb by ID.

## Setup

### Prerequisites

- Golang
- PostgreSQL
- Gin framework
- GORM
- JWT
- GoDotEnv

### Installation

1. Clone the repository:
    ```sh
    git clone https://github.com/yourusername/african_proverbs_api.git
    cd african_proverbs_api
    ```

2. Install dependencies:
    ```sh
    go get -u github.com/gin-gonic/gin
    go get -u github.com/golang-jwt/jwt
    go get -u github.com/joho/godotenv
    go get -u gorm.io/gorm
    go get -u gorm.io/driver/postgres
    ```

3. Set up your PostgreSQL database and configure environment variables in a `.env` file:
    ```
    DB_HOST=your_db_host
    DB_PORT=your_db_port
    DB_USER=your_db_user
    DB_PASS=your_db_password
    DB_NAME=your_db_name
    DB_SSLMODE=disable
    SECRET=your_jwt_secret
    SMTP_HOST=your_smtp_host
    SMTP_PORT=your_smtp_port
    ```

4. Run the application:
    ```sh
    go run main.go
    ```

## Data Cleaning and Integration

This project builds upon earlier work where I scraped African proverbs and cleaned the data using Python scripts. The cleaned data is then integrated into the Golang application.

### Python Scripts

#### `clean_proverbs.py`

This script reads raw proverbs data, processes each line, and writes cleaned data to new files.

```python
import json

# Read data from input file
with open('proverbs_data.json', 'r') as f:
    data = f.readlines()

# Process each line of data
formatted_data = []
for line in data:
    try:
        json_data = json.loads(line)
        formatted_data.append(json_data)
    except json.JSONDecodeError as e:
        print(f"Error decoding JSON: {e}")

proverbs = json_data["Proverbs"]

# Convert to list of proverbs
proverbs_list = [proverbs[key] for key in proverbs]

# Create new JSON object
proverbs_json = {
    "Proverbs": proverbs_list
}

# Write JSON data to files
with open('proverbs_only.json', 'w') as f:
    json.dump(proverbs_json, f, indent=4)

with open('proverbs_cleaned.json', 'w') as f:
    json.dump(formatted_data, f, indent=4)
```

#### `extract_meanings.py`

This script uses Pandas to load a CSV file containing proverbs and their meanings, and writes the data to JSON files.

```python
import pandas as pd
import os

# Load dataset containing proverbs and their meanings
file_path = os.getcwd()
dataset = pd.read_csv(f"{file_path}/proverbs_and_meanings_data.csv", encoding='cp1252')
dataset = dataset.rename(columns={'Meaning ': 'Meaning'})  # Rename column for consistency

# Extract meanings and proverbs from dataset
meanings = dataset["Meaning"]
proverbs = dataset["Proverbs"]

# Save to JSON files
dataset.to_json('proverbs_data.json', index=False)
meanings.to_json('meanings_data.json', index=False)
```

### Previous Work

This project is an extension of my earlier work where I webscraped African proverbs and cleaned the data. You can find the previous project [here]([https://github.com/venn](https://github.com/vennisabarfi/Proverb-Recommendation-Engine-Website/)).

### Next Steps
- Complete Frontend
- Integrate More Tests
- Dockerize
- Add Swagger Documentation for API
  
## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.


