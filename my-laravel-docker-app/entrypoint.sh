#!/bin/bash
set -e

cleanup() {
    echo "Shutting down Laravel..."
    kill -TERM $LARAVEL_PID 2>/dev/null || true
    wait $LARAVEL_PID 2>/dev/null || true
    exit 0
}

trap cleanup SIGTERM SIGINT
php artisan serve --host=0.0.0.0 --port=8000 &
LARAVEL_PID=$!

wait $LARAVEL_PID