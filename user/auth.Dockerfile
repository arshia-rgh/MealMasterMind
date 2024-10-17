FROM python:3.12-slim

WORKDIR /app

COPY ./requirements /app/requirements

RUN pip install --no-cache-dir -r /app/requirements/dev.txt

COPY . .

EXPOSE 8080

CMD ["uvicorn", "app.main:app", "--host", "auth-service", "--port", "8080"]
