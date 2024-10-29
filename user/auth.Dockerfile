FROM python:3.12-slim

WORKDIR /app

COPY ./requirements /app/requirements

RUN pip install --no-cache-dir -r /app/requirements/dev.txt

RUN apt-get update && apt-get install -y \
    protobuf-compiler \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*

COPY . .

RUN python -m grpc_tools.protoc -I./grpc_ --python_out=. --grpc_python_out=. ./grpc_/*.proto


EXPOSE 8080

CMD ["uvicorn", "app.main:app", "--host", "auth-service", "--port", "8080"]
