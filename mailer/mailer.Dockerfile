FROM python:3.12-slim

WORKDIR /app

COPY ./requirements /app/requirements

RUN pip install --no-cache-dir -r /app/requirements/dev.txt

COPY . .


CMD ["python", "main.py"]