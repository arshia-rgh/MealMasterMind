from app.config import origins
from app.routes import user_protected, user_public
from fastapi import FastAPI
from fastapi.middleware.cors import CORSMiddleware

app = FastAPI()

app.add_middleware(
    CORSMiddleware,
    allow_origins=origins,
    allow_credentials=True,
    allow_methods=["GET", "POST", "PUT", "DELETE", "OPTIONS"],
    allow_headers=["Origin", "Content-Length", "Content-Type", "Authorization", "Accept"],
    max_age=12 * 60,
)

app.include_router(user_public.router, prefix="/api")
app.include_router(user_protected.router, prefix="/api")
