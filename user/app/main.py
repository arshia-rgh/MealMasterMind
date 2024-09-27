from fastapi import FastAPI

from user.app.routes.user import router

app = FastAPI()

app.include_router(router, prefix="/api")
