import json
import logging
import math
import os
import time
from typing import Tuple

import pika
from pika.adapters.blocking_connection import BlockingChannel, BlockingConnection
from send_mail import send_mail

logging.basicConfig(level=logging.INFO)


def callback(ch, method, properties, body):
    data = json.loads(body)
    send_mail(data)
    logging.info(f"Received message: {data}")


async def consume(routing_key: str):
    conn, ch = await connect()

    if not ch or not conn:
        return
    while True:
        try:
            ch.queue_declare(queue=routing_key)

            ch.basic_consume(queue=routing_key, on_message_callback=callback, auto_ack=True)
            logging.info(f"Started consuming from queue: {routing_key}")
            ch.start_consuming()

        except Exception as e:
            logging.error(f"Failed to consume message: {e}")
            return

        finally:
            if ch:
                ch.close()
            if conn:
                conn.close()


async def connect() -> Tuple[None, None] | Tuple[BlockingConnection, BlockingChannel]:
    counts = 0

    rabbitmq_url = f"amqp://{os.getenv('RABBITMQ_USERNAME')}:{os.getenv('RABBITMQ_PASSWORD')}@{os.getenv('RABBITMQ_HOST')}:{os.getenv('RABBITMQ_PORT')}/"

    while True:
        try:
            connection = pika.BlockingConnection(pika.URLParameters(rabbitmq_url))
            channel = connection.channel()
            logging.info("Connected to the rabbitMQ")
            break
        except Exception as e:
            logging.warning("rabbitmq not yet ready...")
            counts += 1

            if counts > 5:
                logging.error(f"Failed to connect to RabbitMQ: {e}")
                return None, None

            back_off = math.pow(counts, 2)
            logging.info(f"Backing off for {back_off} seconds")

            time.sleep(back_off)

    return connection, channel
