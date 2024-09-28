import os

from user.app.database import engine
from sqlalchemy import text


def apply_migration(file_path):
    with open(file_path, "r") as file:
        sql = file.read()

    with engine.connect() as connection:
        connection.execute(text(sql))


if __name__ == "__main__":
    migrations_dir = 'versions'
    for filename in sorted(os.listdir(migrations_dir)):
        if filename.endswith('.sql'):
            file_path = os.path.join(migrations_dir, filename)
            apply_migration(file_path)
