import json
import logging
import math
import os
import time
from typing import Tuple

import pika
from pika.adapters.blocking_connection import BlockingChannel, BlockingConnection

logging.basicConfig(level=logging.INFO)


async def publish_message(routing_key: str, data: dict):
    conn, ch = await connect()
    if not conn or not ch:
        return False
    try:
        ch.queue_declare(queue=routing_key)
        message = json.dumps(data).encode("utf-8")
        ch.basic_publish(
            exchange="",
            routing_key=routing_key,
            body=message,
            properties=pika.BasicProperties(content_type="application/json"),
        )
        logging.info(f"published message: {data} to queue: {routing_key}")
        return True
    except Exception as e:
        logging.error(f"Failed to publish message: {e}")
        return False

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
