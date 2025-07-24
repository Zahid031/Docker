#!/bin/sh

echo "Waiting for database..."
while ! nc -z db 5432; do
  sleep 1
done

echo "Running migrations..."
python manage.py migrate

echo "Collecting static files..."
python manage.py collectstatic --noinput

echo "Starting Django server..."
python manage.py runserver 0.0.0.0:8000
