from fastapi import FastAPI

from user.app.routes import user_public

app = FastAPI()

app.include_router(user_public.router, prefix="/api")
