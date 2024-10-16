FROM python:3.12-slim

WORKDIR /app

COPY ./requirements /app/requirements

RUN pip install --no-cache-dir -r /app/requirements/dev.txt

COPY . .

EXPOSE 8081

CMD uvicorn user.app.main:app --port 8080 && celery -A user.app.tasks worker --loglevel=info
