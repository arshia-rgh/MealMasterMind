import json
import logging
import os

from rabbit_helper.rabbit import Rabbit
from send_mail import send_mail


async def consume(routing_key):
    rabbitmq_url = f"amqp://{os.getenv('RABBITMQ_USERNAME')}:{os.getenv('RABBITMQ_PASSWORD')}@{os.getenv('RABBITMQ_HOST')}:{os.getenv('RABBITMQ_PORT')}/"
    rabbit = Rabbit(rabbitmq_url)

    await rabbit.consume(routing_key, callback)


def callback(ch, method, properties, body):
    data = json.loads(body)
    send_mail(data)
    logging.info(f"Received message: {data}")
