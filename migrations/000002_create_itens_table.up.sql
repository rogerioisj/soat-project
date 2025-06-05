CREATE TYPE item_type AS ENUM (
    'snack',
    'drink',
    'dessert',
    'accompaniment'
);

CREATE TABLE itens (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    price INTEGER NOT NULL,
    description TEXT,
    type item_type NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO itens (name, price, description, type)
VALUES
    ('Coxinha', 500, 'Coxinha de frango com massa crocante', 'snack'),
    ('Pão de Queijo', 300, 'Pão de queijo quentinho e saboroso', 'snack'),
    ('Pastel de Carne', 450, 'Pastel de carne moída com massa fina', 'snack'),
    ('Hamburgue de Patinho', 700, 'Hamburgue de patinho grelhado com queijo', 'snack'),
    ('Água com Gás', 100, 'Água com gás refrescante', 'drink'),
    ('Água Mineral', 150, 'Água mineral gelada', 'drink'),
    ('Refrigerante', 250, 'Refrigerante gelado', 'drink'),
    ('Suco Natural', 350, 'Suco natural de laranja', 'drink'),
    ('Bolo de Chocolate', 400, 'Bolo de chocolate com cobertura cremosa', 'dessert'),
    ('Sorvete', 450, 'Sorvete de baunilha com calda de chocolate', 'dessert'),
    ('Batata Frita', 200, 'Porção de batata frita crocante', 'accompaniment'),
    ('Salada Caesar', 600, 'Salada Caesar com frango grelhado e croutons', 'accompaniment'),
    ('Arroz Branco', 150, 'Arroz branco soltinho', 'accompaniment'),
    ('Feijão Tropeiro', 350, 'Feijão tropeiro com bacon e linguiça', 'accompaniment');