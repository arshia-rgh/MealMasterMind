import os

from sqlalchemy import text

from user.app.db.database import engine


def apply_migration(file_path):
    with open(file_path, "r") as file:
        sql = file.read()

    # Split the SQL file content into individual SQL statements
    sql_statements = sql.split(";")

    with engine.connect() as connection:
        for statement in sql_statements:
            if statement.strip():  # Skip empty statements
                connection.execute(text(statement))


if __name__ == "__main__":
    migrations_dir = "../../../migrations"
    for filename in sorted(os.listdir(migrations_dir)):
        if filename.endswith(".sql"):
            file_path = os.path.join(migrations_dir, filename)
            apply_migration(file_path)
