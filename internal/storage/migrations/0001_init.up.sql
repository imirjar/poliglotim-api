-- migrations/0001_init.up.sql

-- Создаем таблицу для курсов
CREATE TABLE IF NOT EXISTS courses (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL,
    description TEXT,
    updated TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    logo_path VARCHAR(512) -- Путь к файлу логотипа
);

-- Создаем таблицу для глав
CREATE TABLE IF NOT EXISTS chapters (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    course_id UUID NOT NULL REFERENCES courses(id) ON DELETE CASCADE,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    updated TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    position INT NOT NULL DEFAULT 0 -- Для упорядочивания глав внутри курса
);

-- Создаем индекс для ускорения поиска глав по курсу
CREATE INDEX IF NOT EXISTS idx_chapters_course_id ON chapters(course_id);