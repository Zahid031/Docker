# Stage 1: Frontend asset compilation
FROM node:20 AS frontend

WORKDIR /app
COPY package.json package-lock.json vite.config.ts tsconfig.json ./
COPY resources/js ./resources/js
RUN npm install
RUN npm run build

# Stage 2: Backend setup
FROM php:8.3-fpm AS backend

# Install system dependencies and PHP extensions
RUN apt-get update && apt-get install -y \
    git \
    unzip \
    libzip-dev \
    libpng-dev \
    libjpeg-dev \
    libfreetype6-dev \
    libonig-dev \
    libxml2-dev \
    && docker-php-ext-install pdo_mysql zip gd mbstring exif pcntl bcmath opcache

# Install Composer
COPY --from=composer/composer:latest-bin /composer /usr/bin/composer

WORKDIR /app

# Copy composer files and install dependencies
COPY composer.json composer.lock ./
RUN composer install --no-dev --no-interaction --no-plugins --no-scripts --prefer-dist

# Copy application files
COPY . .

# Copy compiled frontend assets from the frontend stage
COPY --from=frontend /app/public/build ./public/build

# Set permissions
RUN chown -R www-data:www-data /app/storage /app/bootstrap/cache
RUN chmod -R 775 /app/storage /app/bootstrap/cache

# Generate application key
RUN php artisan key:generate

# Expose port and start PHP-FPM
EXPOSE 9000
CMD ["php-fpm"]

