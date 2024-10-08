# NRL Tipping Application

The NRL Tipping Application is a project designed to create a user-friendly tipping platform for the Australian NRL competitions, including NRL, NRL Women's Championship, State of Origin, and State of Origin Women's. The application allows users to view fixtures, team details, and match outcomes and place tips on various games. The backend is built with Go, and the frontend uses Vue.js, TailwindCSS, and Pinia.

## Project Goals

- **Develop a Full-Stack Application**: Create a robust full-stack application to support NRL tipping, with a backend API built in Go and a frontend developed in Vue.js.
- **Real-Time Data Integration**: Integrate with NRL's API to fetch real-time data about fixtures, odds, and results.
- **Scalable Architecture**: Use Docker for containerization and Kubernetes for potential scaling and deployment in a clustered environment.
- **Type-Safe Database Interactions**: Use sqlc for generating type-safe Go code from SQL queries, ensuring efficient and error-free database operations.

## Local Development

Ensure you have [Docker] installed as the best way to mimic the live environment is via containerisation. First it would be best to setup the database, which can be down by pulling it up first and using [migrate]. The following is a single script to set up and run:

```bash
# Build the whole project
docker compose build

# Setup the database and migrate to the current version
docker compose up -d db
migrate -source 'file:./backend/internal/db/migration' -database 'postgres://postgres:password@localhost:5432/nrl_tipping?sslmode=disable' up

# Run the whole project
docker compose up
```
This will be a blocking session unless you add the `-d` flag to daemonise it. Some form of testing stage will be added at some point soon to ensure everything works accordingly.

### Adding a New Database Change

If you want to add a new table or modify existing tables, you will need to create a new database migration. For example, if you want to add a new field to the teams table called city, you can do the following:

```bash
# Generate new migration files
migrate create -ext sql -dir backend/internal/db/migration -seq add_team_city
```
Then, modify the generated migration files:

```sql
-- file: backend/db/internal/migration/000003_add_team_city.up.sql
ALTER TABLE teams
ADD COLUMN city VARCHAR(255);

-- file: backend/db/internal/migration/000003_add_team_city.down.sql
ALTER TABLE teams
DROP COLUMN city;
```

After modifying the migration files, run the following command to apply the migration:

```bash
migrate -source 'file:./backend/internal/db/migration' -database 'postgres://postgres:password@localhost:5432/nrl_tipping?sslmode=disable' up
```

### Adding a New SQL Query

This project uses [sqlc] to generate type-safe Go code from SQL queries, avoiding the use of traditional ORMs and maintaining better control over database interactions. To add a new query, follow these steps:

1. **Create a New SQL Query**: Add a new SQL query to one of the `.sql` files located in `backend/internal/db/query`. For example, to get a fixture by its ID:

```sql
-- file: backend/internal/db/query/fixtures.sql

-- name: GetFixtureByID :one
SELECT *
FROM fixtures
WHERE id = $1;
```

2. **Generate Go Code**: Run `sqlc generate` in the backend directory of the project to generate Go code for the new query:

```bash
cd backend
sqlc generate
```

This command will generate new Go functions based on your SQL queries in the `backend/internal/db` directory. Make sure to run `sqlc generate` every time you modify the `.sql` files.

## API Endpoints

The NRL Tipping Application backend provides several API endpoints for interacting with competitions, fixtures, and match details.

- **Get Competitions**
    - **URL**: `GET /api/v1/competitions`
    - **Description**: Retrieves a list of all available competitions.
    - **Response**: JSON array of competitions.

- **Get All Fixtures**
    - **URL**: `GET /api/v1/fixtures`
    - **Description**: Retrieves a list of all fixtures.
    - **Response**: JSON array of fixtures.

- **Get Fixtures by Competition ID**
    - **URL**: `GET /api/v1/fixtures/{competition_id}`
    - **Description**: Retrieves fixtures for a specific competition.
    - **Parameters**:
        - `competition_id` *(required)*: The ID of the competition.
    - **Response**: JSON array of fixtures for the specified competition.

- **Get Match Details**
    - **URL**: `GET /api/v1/fixtures/{competition_id}/{match_id}`
    - **Description**: Retrieves details for a specific match within a competition.
    - **Parameters**:
        - `competition_id` *(required)*: The ID of the competition.
        - `match_id` *(required)*: The ID of the match.
    - **Response**: JSON object with match details.

Here are some example commands using curl to interact with the API.

```bash
# Get Competitions
curl -X GET http://localhost:8080/api/v1/competitions

# Get All Fixtures
curl -X GET http://localhost:8080/api/v1/fixtures

# Get Fixtures by Competition ID
#   - 111 NRL
#   - 161 NRLW
#   - 116 State of Origin
#   - 156 Womens State of Origin 
curl -X GET "http://localhost:8080/api/v1/fixtures/111"

# Get Match Details
curl -X GET "http://localhost:8080/api/v1/fixtures/111/20241112610"
```

## Contributing

To contribute to this project, please fork the repository and create a pull request with your changes. Ensure that all new code follows the project’s coding standards and is well-documented.

[Docker]: https://www.docker.com/get-started/
[migrate]: https://github.com/golang-migrate/migrate
[sqlc]: https://docs.sqlc.dev/en/stable/overview/install.html