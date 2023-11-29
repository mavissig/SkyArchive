FROM python:3.11.1-alpine

WORKDIR /server

ENV PYTHONDONTWHRITEBYTECODE=1 \
    PYTHONUNBUFFERED=1

RUN apk update \
    && apk add postgresql-dev gcc python3-dev musl-dev

RUN pip install --upgrade pip
COPY ../../../Downloads/FlyCloud-connect/server/requirements.txt .
RUN pip install -r requirements.txt

COPY ../../../Downloads/FlyCloud-connect/server/entrypoint.sh .
RUN chmod +x ./entrypoint.sh

EXPOSE 8000

COPY ../../../Downloads/FlyCloud-connect/server .


# ENTRYPOINT ["chmod", "+x", "/app/entrypoint.sh"]