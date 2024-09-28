from fastapi import FastAPI

from user.app.routes import user_public, user_protected

app = FastAPI()

app.include_router(user_public.router, prefix="/api")
app.include_router(user_protected.router, prefix="/api")
