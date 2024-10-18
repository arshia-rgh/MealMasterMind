FROM python:3.12-slim

WORKDIR /app

RUN apt-get update && apt-get install -y \
    protobuf-compiler \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*

COPY ./requirements /app/requirements

RUN pip install --no-cache-dir -r /app/requirements/dev.txt

COPY . .

RUN python -m grpc_tools.protoc -I./mails --python_out=. --grpc_python_out=. ./mails/*.proto


CMD ["python", "main.py"]