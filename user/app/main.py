from fastapi import FastAPI

from user.app.routes import user_protected, user_public

app = FastAPI()

app.include_router(user_public.router, prefix="/api")
app.include_router(user_protected.router, prefix="/api")
