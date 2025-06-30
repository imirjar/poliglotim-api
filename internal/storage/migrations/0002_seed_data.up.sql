-- migrations/0002_seed_data.up.sql

-- Добавляем тестовый курс
INSERT INTO courses (id, name, description, updated, logo_path)
VALUES ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11', 'Основы Go', 'Курс по основам программирования на Go', NOW(), '/logos/go-basic.png');

-- Добавляем главы для тестового курса
INSERT INTO chapters (id, course_id, name, description, updated, position)
VALUES 
('b0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11', 'a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11', 'Введение в Go', 'Основные концепции языка', NOW(), 1),
('c0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11', 'a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11', 'Типы данных', 'Работа с типами данных в Go', NOW(), 2),
('d0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11', 'a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11', 'Горутины', 'Параллельное программирование', NOW(), 3);