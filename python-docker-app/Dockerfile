FROM python:3.12-alpine

WORKDIR /code

RUN apk add --no-cache gcc musl-dev libffi-dev

COPY requirements.txt ./
RUN python -m pip install --upgrade pip && \
    python -m pip install -r requirements.txt

COPY . .

EXPOSE 8000

CMD ["python", "manage.py", "runserver", "0.0.0.0:8000"]

