from celery import Celery

import config

celery_app = Celery(
    "user_service",
    broker=config.CELERY_BROKER,
    backend=config.CELERY_BACKEND,

)

celery_app.conf.update(
    result_expires=config.CELERY_RESULT_EXPIRE,
    task_serializer=config.CELERY_TASK_SERIALIZER,
    accept_content=config.CELERY_ACCEPT_CONTENT,
    result_serializer=config.CELERY_RESULT_SERIALIZER,
    timezone=config.CELERY_TIMEZONE,
    enable_utc=config.CELERY_ENABLE_UTC,
)

celery_app.autodiscover_tasks()