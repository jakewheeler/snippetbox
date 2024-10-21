# Let's Go - SnippetBox

This repository contains the sample project code from the [Let's Go](https://lets-go.alexedwards.net/) book by Alex Edwards.

## Database

### Starting the database

```sh
docker compose up -d
```

### Connecting to the database

```sh
mysql -h 127.0.0.1 -P 3306 -D snippetbox -u devuser -p
```

Enter the password used in `docker-compose.yml`.

Note: If `devuser` does not have appropriate permissions to access `snippetbox` you can use the `root` user and password (also found in `docker-compose.yml`) to grant those permissions.
