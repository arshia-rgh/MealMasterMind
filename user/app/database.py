from sqlalchemy import create_engine
from sqlalchemy.orm import sessionmaker, declarative_base

from .config import DBConfig

DB_URL = f"mysql+pymysql://{DBConfig.DB_USER}:{DBConfig.DB_PASSWORD}@{DBConfig.DB_HOST}:{DBConfig.DB_PORT}/{DBConfig.DB_NAME}"

engine = create_engine(DB_URL)

SessionLocal = sessionmaker(autocommit=False, autoflush=False, bind=engine)
base = declarative_base()


def get_db():
    db = SessionLocal()

    try:
        yield db

    finally:
        db.close()
