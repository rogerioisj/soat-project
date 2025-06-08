DROP TABLE IF EXISTS orders_itens;

DROP TABLE IF EXISTS orders;

DROP TYPE IF EXISTS order_status;

ALTER TABLE public.users DROP CONSTRAINT IF EXISTS orders_user_id_fkey;