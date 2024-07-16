-- Active: 1718919020656@@127.0.0.1@5432@userservice

INSERT INTO users (username, email, password_hash, full_name, user_type, address, phone_number, bio, specialties, years_of_experience, is_verified)
VALUES 
    ('john_doe', 'john.doe@example.com', 'hashed_password1', 'John Doe', 'customer', '123 Main St, City', '+1234567890', 'Bio for John Doe', '{"Cooking"}', 5, true),
    ('jane_smith', 'jane.smith@example.com', 'hashed_password2', 'Jane Smith', 'chef', '456 Oak St, City', '+1987654321', 'Bio for Jane Smith', '{"Baking", "Grilling"}', 7, true),
    ('chef_tony', 'tony@example.com', 'hashed_password3', 'Tony Stark', 'chef', '789 Elm St, City', '+1122334455', 'Bio for Tony Stark', '{"Italian Cuisine"}', 10, true),
    ('foodie123', 'foodie123@example.com', 'hashed_password4', 'Sarah Lee', 'customer', '321 Maple St, City', '+1223344556', 'Bio for Sarah Lee', '{"Tasting"}', 2, false),
    ('homecook_mary', 'mary@example.com', 'hashed_password5', 'Mary Johnson', 'chef', '654 Pine St, City', '+1445566778', 'Bio for Mary Johnson', '{"Home Cooking"}', 8, true),
    ('baker_bob', 'bob@example.com', 'hashed_password6', 'Bob Baker', 'chef', '987 Birch St, City', '+1667788990', 'Bio for Bob Baker', '{"Baking"}', 6, true),
    ('alice_wonder', 'alice@example.com', 'hashed_password7', 'Alice Wonderland', 'customer', '159 Cedar St, City', '+1889900112', 'Bio for Alice Wonderland', '{"Reviewing"}', 1, false),
    ('gordon_ramsay', 'gordon@example.com', 'hashed_password8', 'Gordon Ramsay', 'chef', '753 Spruce St, City', '+1990011223', 'Bio for Gordon Ramsay', '{"Gourmet Cooking"}', 15, true),
    ('customer_kate', 'kate@example.com', 'hashed_password9', 'Kate Brown', 'customer', '258 Willow St, City', '+1011121314', 'Bio for Kate Brown', '{"Tasting"}', 3, false),
    ('chef_mike', 'mike@example.com', 'hashed_password10', 'Mike Tyson', 'chef', '369 Hickory St, City', '+2021222324', 'Bio for Mike Tyson', '{"Grilling"}', 12, true);


INSERT INTO kitchens_profile (name, description, cuisine_type, address, phone_number, rating, total_orders)
VALUES 
  ( 'Italian Bistro', 'Authentic Italian Cuisine', 'Italian', '123 Pasta St.', '123-456-7890', 4.5, 100),
  ( 'Sushi Place', 'Fresh Sushi Daily', 'Japanese', '456 Sushi Blvd.', '987-654-3210', 4.8, 150),
  ( 'Burger Joint', 'Best Burgers in Town', 'American', '789 Burger Ave.', '555-555-5555', 4.2, 200),
  ( 'Taco House', 'Tasty Tacos', 'Mexican', '321 Taco Ln.', '111-222-3333', 4.6, 250),
  ( 'Curry Spot', 'Spicy Indian Food', 'Indian', '654 Curry Rd.', '444-666-7777', 4.7, 300),
  ( 'Vegan Delights', 'Healthy and Delicious', 'Vegan', '987 Vegan Way', '888-999-0000', 4.3, 50),
  ( 'BBQ Pit', 'Slow Cooked BBQ', 'BBQ', '111 BBQ St.', '222-333-4444', 4.9, 180),
  ( 'Pizza Place', 'Hot and Fresh Pizza', 'Italian', '222 Pizza Pl.', '333-444-5555', 4.4, 220),
  ( 'Dim Sum Corner', 'Delicious Dim Sum', 'Chinese', '333 Dim Sum Ct.', '666-777-8888', 4.8, 270),
  ( 'French Bakery', 'Freshly Baked Goods', 'French', '444 Bakery Blvd.', '999-000-1111', 4.9, 320);