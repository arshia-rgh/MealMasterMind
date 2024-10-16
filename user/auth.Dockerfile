FROM python:3.12-slim

WORKDIR /app

COPY requirements .

RUN pip install --no-cache-dir -r requirements/dev.txt

COPY . .

EXPOSE 8081

CMD ["uvicorn", "main:app", "--host", "0.0.0.0", "--port", "8081"]
