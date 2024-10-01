from celery import Celery

from user.app import config, tasks

celery_app = Celery(
    "user_service", broker=config.CELERY_BROKER, backend=config.CELERY_BACKEND, broker_connection_retry_on_startup=True
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
