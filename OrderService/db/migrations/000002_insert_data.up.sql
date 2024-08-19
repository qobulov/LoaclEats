-- Insert data into kitchens table
INSERT INTO kitchens (owner_id, name, description, cuisine_type, address, phone_number, rating, total_orders)
VALUES 
  (gen_random_uuid(), 'Italian Bistro', 'Authentic Italian Cuisine', 'Italian', '123 Pasta St.', '123-456-7890', 4.5, 100),
  (gen_random_uuid(), 'Sushi Place', 'Fresh Sushi Daily', 'Japanese', '456 Sushi Blvd.', '987-654-3210', 4.8, 150),
  (gen_random_uuid(), 'Burger Joint', 'Best Burgers in Town', 'American', '789 Burger Ave.', '555-555-5555', 4.2, 200),
  (gen_random_uuid(), 'Taco House', 'Tasty Tacos', 'Mexican', '321 Taco Ln.', '111-222-3333', 4.6, 250),
  (gen_random_uuid(), 'Curry Spot', 'Spicy Indian Food', 'Indian', '654 Curry Rd.', '444-666-7777', 4.7, 300),
  (gen_random_uuid(), 'Vegan Delights', 'Healthy and Delicious', 'Vegan', '987 Vegan Way', '888-999-0000', 4.3, 50),
  (gen_random_uuid(), 'BBQ Pit', 'Slow Cooked BBQ', 'BBQ', '111 BBQ St.', '222-333-4444', 4.9, 180),
  (gen_random_uuid(), 'Pizza Place', 'Hot and Fresh Pizza', 'Italian', '222 Pizza Pl.', '333-444-5555', 4.4, 220),
  (gen_random_uuid(), 'Dim Sum Corner', 'Delicious Dim Sum', 'Chinese', '333 Dim Sum Ct.', '666-777-8888', 4.8, 270),
  (gen_random_uuid(), 'French Bakery', 'Freshly Baked Goods', 'French', '444 Bakery Blvd.', '999-000-1111', 4.9, 320);

-- Insert data into dishes table
INSERT INTO dishes (kitchen_id, name, description, price, category, ingredients, allergens, nutrition_info, dietary_info, available)
VALUES
  ((SELECT id FROM kitchens ORDER BY RANDOM() LIMIT 1), 'Margherita Pizza', 'Classic pizza with tomatoes and mozzarella', 10.99, 'Main Course', ARRAY['tomatoes', 'mozzarella', 'basil'], ARRAY['dairy'], '{"calories": 250, "protein": 12}', ARRAY['vegetarian'], true),
  ((SELECT id FROM kitchens ORDER BY RANDOM() LIMIT 1), 'Spicy Tuna Roll', 'Fresh tuna with a kick', 12.99, 'Sushi', ARRAY['tuna', 'spicy mayo'], ARRAY['fish'], '{"calories": 200, "protein": 18}', ARRAY['3 times a week'], true),
  ((SELECT id FROM kitchens ORDER BY RANDOM() LIMIT 1), 'Cheeseburger', 'Juicy burger with cheese', 8.99, 'Main Course', ARRAY['beef', 'cheese', 'lettuce'], ARRAY['dairy'], '{"calories": 350, "protein": 20}', ARRAY['3 times a week'], true),
  ((SELECT id FROM kitchens ORDER BY RANDOM() LIMIT 1), 'Beef Taco', 'Savory beef taco', 3.99, 'Appetizer', ARRAY['beef', 'taco shell'], ARRAY['3 times a week'], '{"calories": 150, "protein": 10}', ARRAY['3 times a week'], true),
  ((SELECT id FROM kitchens ORDER BY RANDOM() LIMIT 1), 'Chicken Curry', 'Spicy chicken curry', 14.99, 'Main Course', ARRAY['chicken', 'curry sauce'], ARRAY['3 times a week'], '{"calories": 400, "protein": 25}', ARRAY['3 times a week'], true),
  ((SELECT id FROM kitchens ORDER BY RANDOM() LIMIT 1), 'Vegan Salad', 'Fresh and healthy salad', 7.99, 'Salad', ARRAY['lettuce', 'tomatoes', 'cucumbers'], ARRAY['3 times a week'], '{"calories": 120, "protein": 5}', ARRAY['vegan'], true),
  ((SELECT id FROM kitchens ORDER BY RANDOM() LIMIT 1), 'BBQ Ribs', 'Tender BBQ ribs', 16.99, 'Main Course', ARRAY['ribs', 'bbq sauce'], ARRAY['3 times a week'], '{"calories": 600, "protein": 35}', ARRAY['3 times a week'], true),
  ((SELECT id FROM kitchens ORDER BY RANDOM() LIMIT 1), 'Pepperoni Pizza', 'Pizza with pepperoni', 11.99, 'Main Course', ARRAY['pepperoni', 'mozzarella', 'tomato sauce'], ARRAY['dairy'], '{"calories": 300, "protein": 15}', ARRAY['3 times a week'], true),
  ((SELECT id FROM kitchens ORDER BY RANDOM() LIMIT 1), 'Pork Dumplings', 'Delicious pork dumplings', 8.99, 'Appetizer', ARRAY['pork', 'dumpling wrapper'], ARRAY['3 times a week'], '{"calories": 180, "protein": 12}', ARRAY['3 times a week'], true),
  ((SELECT id FROM kitchens ORDER BY RANDOM() LIMIT 1), 'Croissant', 'Buttery croissant', 2.99, 'Dessert', ARRAY['flour', 'butter'], ARRAY['dairy'], '{"calories": 220, "protein": 4}', ARRAY['3 times a week'], true);

-- Insert data into orders table
INSERT INTO orders (user_id, kitchen_id, items, total_amount, status, delivery_address, delivery_time)
VALUES
  (gen_random_uuid(), (SELECT id FROM kitchens ORDER BY RANDOM() LIMIT 1), '[{"dish_id": "dish1", "quantity": 2}, {"dish_id": "dish2", "quantity": 1}]'::jsonb, 24.97, 'pending', '123 Delivery St.', CURRENT_TIMESTAMP + INTERVAL '1 hour'),
  (gen_random_uuid(), (SELECT id FROM kitchens ORDER BY RANDOM() LIMIT 1), '[{"dish_id": "dish3", "quantity": 1}, {"dish_id": "dish4", "quantity": 3}]'::jsonb, 20.96, 'completed', '456 Delivery Ave.', CURRENT_TIMESTAMP + INTERVAL '2 hours'),
  (gen_random_uuid(), (SELECT id FROM kitchens ORDER BY RANDOM() LIMIT 1), '[{"dish_id": "dish5", "quantity": 2}]'::jsonb, 29.98, 'completed', '789 Delivery Rd.', CURRENT_TIMESTAMP + INTERVAL '3 hours'),
  (gen_random_uuid(), (SELECT id FROM kitchens ORDER BY RANDOM() LIMIT 1), '[{"dish_id": "dish6", "quantity": 4}]'::jsonb, 31.96, 'pending', '101 Delivery Blvd.', CURRENT_TIMESTAMP + INTERVAL '4 hours'),
  (gen_random_uuid(), (SELECT id FROM kitchens ORDER BY RANDOM() LIMIT 1), '[{"dish_id": "dish7", "quantity": 1}]'::jsonb, 16.99, 'completed', '202 Delivery St.', CURRENT_TIMESTAMP + INTERVAL '5 hours'),
  (gen_random_uuid(), (SELECT id FROM kitchens ORDER BY RANDOM() LIMIT 1), '[{"dish_id": "dish8", "quantity": 3}]'::jsonb, 35.97, 'pending', '303 Delivery Ave.', CURRENT_TIMESTAMP + INTERVAL '6 hours'),
  (gen_random_uuid(), (SELECT id FROM kitchens ORDER BY RANDOM() LIMIT 1), '[{"dish_id": "dish9", "quantity": 2}, {"dish_id": "dish10", "quantity": 2}]'::jsonb, 21.98, 'completed', '404 Delivery Rd.', CURRENT_TIMESTAMP + INTERVAL '7 hours'),
  (gen_random_uuid(), (SELECT id FROM kitchens ORDER BY RANDOM() LIMIT 1), '[{"dish_id": "dish1", "quantity": 5}]'::jsonb, 54.95, 'pending', '505 Delivery Blvd.', CURRENT_TIMESTAMP + INTERVAL '8 hours'),
  (gen_random_uuid(), (SELECT id FROM kitchens ORDER BY RANDOM() LIMIT 1), '[{"dish_id": "dish2", "quantity": 1}]'::jsonb, 12.99, 'completed', '606 Delivery St.', CURRENT_TIMESTAMP + INTERVAL '9 hours'),
  (gen_random_uuid(), (SELECT id FROM kitchens ORDER BY RANDOM() LIMIT 1), '[{"dish_id": "dish3", "quantity": 2}]'::jsonb, 17.98, 'pending', '707 Delivery Ave.', CURRENT_TIMESTAMP + INTERVAL '10 hours');

-- Insert data into reviews table
INSERT INTO reviews (order_id, user_id, kitchen_id, rating, comment)
VALUES
  ((SELECT id FROM orders ORDER BY RANDOM() LIMIT 1), gen_random_uuid(), (SELECT id FROM kitchens ORDER BY RANDOM() LIMIT 1), 4.5, 'Great food!'),
  ((SELECT id FROM orders ORDER BY RANDOM() LIMIT 1), gen_random_uuid(), (SELECT id FROM kitchens ORDER BY RANDOM() LIMIT 1), 4.0, 'Very tasty!'),
  ((SELECT id FROM orders ORDER BY RANDOM() LIMIT 1), gen_random_uuid(), (SELECT id FROM kitchens ORDER BY RANDOM() LIMIT 1), 3.5, 'Good, but could be better.'),
  ((SELECT id FROM orders ORDER BY RANDOM() LIMIT 1), gen_random_uuid(), (SELECT id FROM kitchens ORDER BY RANDOM() LIMIT 1), 5.0, 'Excellent!'),
  ((SELECT id FROM orders ORDER BY RANDOM() LIMIT 1), gen_random_uuid(), (SELECT id FROM kitchens ORDER BY RANDOM() LIMIT 1), 2.0, 'Not as expected.'),
  ((SELECT id FROM orders ORDER BY RANDOM() LIMIT 1), gen_random_uuid(), (SELECT id FROM kitchens ORDER BY RANDOM() LIMIT 1), 4.8, 'Delicious and fresh.'),
  ((SELECT id FROM orders ORDER BY RANDOM() LIMIT 1), gen_random_uuid(), (SELECT id FROM kitchens ORDER BY RANDOM() LIMIT 1), 3.8, 'Good quality.'),
  ((SELECT id FROM orders ORDER BY RANDOM() LIMIT 1), gen_random_uuid(), (SELECT id FROM kitchens ORDER BY RANDOM() LIMIT 1), 4.9, 'Amazing flavors.'),
  ((SELECT id FROM orders ORDER BY RANDOM() LIMIT 1), gen_random_uuid(), (SELECT id FROM kitchens ORDER BY RANDOM() LIMIT 1), 1.5, 'Not good.'),
  ((SELECT id FROM orders ORDER BY RANDOM() LIMIT 1), gen_random_uuid(), (SELECT id FROM kitchens ORDER BY RANDOM() LIMIT 1), 4.7, 'Very good.');

-- Insert data into payments table
INSERT INTO payments (order_id, amount, status, payment_method, transaction_id, card_number, expiry_date, cvv)
VALUES
  ((SELECT id FROM orders ORDER BY RANDOM() LIMIT 1), 24.97, 'completed', 'credit_card', 'txn_123456', '4111111111111111', '12/25', '123'),
  ((SELECT id FROM orders ORDER BY RANDOM() LIMIT 1), 20.96, 'completed', 'credit_card', 'txn_234567', '4111111111111112', '11/24', '124'),
  ((SELECT id FROM orders ORDER BY RANDOM() LIMIT 1), 29.98, 'pending', 'debit_card', 'txn_345678', '4111111111111113', '10/23', '125'),
  ((SELECT id FROM orders ORDER BY RANDOM() LIMIT 1), 31.96, 'completed', 'paypal', 'txn_456789', '4111111111111114', '09/22', '126'),
  ((SELECT id FROM orders ORDER BY RANDOM() LIMIT 1), 16.99, 'completed', 'credit_card', 'txn_567890', '4111111111111115', '08/21', '127'),
  ((SELECT id FROM orders ORDER BY RANDOM() LIMIT 1), 35.97, 'completed', 'credit_card', 'txn_678901', '4111111111111116', '07/20', '128'),
  ((SELECT id FROM orders ORDER BY RANDOM() LIMIT 1), 21.98, 'pending', 'debit_card', 'txn_789012', '4111111111111117', '06/19', '129'),
  ((SELECT id FROM orders ORDER BY RANDOM() LIMIT 1), 54.95, 'completed', 'paypal', 'txn_890123', '4111111111111118', '05/18', '130'),
  ((SELECT id FROM orders ORDER BY RANDOM() LIMIT 1), 12.99, 'completed', 'credit_card', 'txn_901234', '4111111111111119', '04/17', '131'),
  ((SELECT id FROM orders ORDER BY RANDOM() LIMIT 1), 17.98, 'pending', 'debit_card', 'txn_012345', '4111111111111120', '03/16', '132');

-- Insert data into working_hours table
INSERT INTO working_hours (kitchen_id, day_of_week, open_time, close_time)
VALUES
  ((SELECT id FROM kitchens ORDER BY RANDOM() LIMIT 1), 1, '08:00', '20:00'),
  ((SELECT id FROM kitchens ORDER BY RANDOM() LIMIT 1), 2, '08:00', '20:00'),
  ((SELECT id FROM kitchens ORDER BY RANDOM() LIMIT 1), 3, '08:00', '20:00'),
  ((SELECT id FROM kitchens ORDER BY RANDOM() LIMIT 1), 4, '08:00', '20:00'),
  ((SELECT id FROM kitchens ORDER BY RANDOM() LIMIT 1), 5, '08:00', '22:00'),
  ((SELECT id FROM kitchens ORDER BY RANDOM() LIMIT 1), 6, '09:00', '22:00'),
  ((SELECT id FROM kitchens ORDER BY RANDOM() LIMIT 1), 7, '09:00', '18:00'),
  ((SELECT id FROM kitchens ORDER BY RANDOM() LIMIT 1), 1, '08:00', '18:00'),
  ((SELECT id FROM kitchens ORDER BY RANDOM() LIMIT 1), 2, '08:00', '20:00'),
  ((SELECT id FROM kitchens ORDER BY RANDOM() LIMIT 1), 3, '08:00', '20:00');
