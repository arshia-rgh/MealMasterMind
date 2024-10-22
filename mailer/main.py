import asyncio
from typing import Final

from event.consume import consume

routing_keys: Final = ["send-mail"]


async def main():
    for key in routing_keys:
        await consume(key)


if __name__ == "__main__":
    asyncio.run(main())
