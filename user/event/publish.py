import logging
import math
import os
import time

import pika

logging.basicConfig(level=logging.INFO)


def connect():
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
                return

            back_off = math.pow(counts, 2)
            logging.info(f"Backing off for {back_off} seconds")

            time.sleep(back_off)

    return connection, channel
