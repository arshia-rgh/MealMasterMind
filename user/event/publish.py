import os

from rabbit_helper.rabbit import Rabbit


async def publish_message(routing_key: str, data: dict):
    rabbitmq_url = f"amqp://{os.getenv('RABBITMQ_USERNAME')}:{os.getenv('RABBITMQ_PASSWORD')}@{os.getenv('RABBITMQ_HOST')}:{os.getenv('RABBITMQ_PORT')}/"
    rabbit = Rabbit(rabbitmq_url)

    try:
        await rabbit.publish(routing_key, data)
    finally:
        rabbit.close()
