import asyncio
from contextlib import asynccontextmanager

from app.config import origins
from app.routes import user_protected, user_public
from event.publish import publish_message
from fastapi import FastAPI, Request
from fastapi.middleware.cors import CORSMiddleware

from .grpc_server import serve

app = FastAPI()

app.add_middleware(
    CORSMiddleware,
    allow_origins=origins,
    allow_credentials=True,
    allow_methods=["GET", "POST", "PUT", "DELETE", "OPTIONS"],
    allow_headers=["Origin", "Content-Length", "Content-Type", "Authorization", "Accept"],
    max_age=12 * 60,
)


@app.middleware("http")
async def log_requests(request: Request, call_next):
    # Log the request
    await publish_message("logs", {"name": "auth", "level": "info", "data": f"Request: {request.method} {request.url}"})

    response = await call_next(request)

    # Log the response
    await publish_message("logs", {"name": "auth", "level": "info", "data": f"Response status: {response.status_code}"})

    return response


@asynccontextmanager
async def lifespan(app: FastAPI):
    loop = asyncio.get_event_loop()
    grpc_server_task = loop.run_in_executor(None, serve)
    print("Starting grpc server")
    yield
    grpc_server_task.cancel()


app.add_event_handler("startup", lambda: asyncio.create_task(lifespan(app).__aenter__()))
app.add_event_handler("shutdown", lambda: asyncio.create_task(lifespan(app).__aexit__(None, None, None)))

app.include_router(user_public.router, prefix="/api")
app.include_router(user_protected.router, prefix="/api")
