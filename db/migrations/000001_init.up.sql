-- Создаем схему (namespace) если она не существует
CREATE SCHEMA IF NOT EXISTS schema_name;

-- Создаем таблицу 'orders' в схеме 'schema_name', если она не существует
CREATE TABLE IF NOT EXISTS schema_name.orders (
    id    BIGSERIAL NOT NULL,  -- Автоинкрементируемый первичный ключ
    name  TEXT      NOT NULL,  -- Текстовое поле (обязательное)
    price INTEGER   DEFAULT 0, -- Числовое поле с значением по умолчанию 0

    -- Объявляем первичный ключ
    CONSTRAINT orders_pk PRIMARY KEY (id)
);