-- Active: 1718919020656@@127.0.0.1@5432@master@public
INSERT INTO users (id, username, email, password_hash, full_name, user_type, address, phone_number, bio, specialties, years_of_experience, is_verified)
VALUES 
    (gen_random_uuid(), 'john_doe', 'john.doe@example.com', 'hashed_password1', 'John Doe', 'customer', '123 Main St, City', '+1234567890', 'Bio for John Doe', '{"Cooking"}', 5, true),
    (gen_random_uuid(), 'jane_smith', 'jane.smith@example.com', 'hashed_password2', 'Jane Smith', 'chef', '456 Oak St, City', '+1987654321', 'Bio for Jane Smith', '{"Baking", "Grilling"}', 7, true),
    (gen_random_uuid(), 'chef_tony', 'tony@example.com', 'hashed_password3', 'Tony Stark', 'chef', '789 Elm St, City', '+1122334455', 'Bio for Tony Stark', '{"Italian Cuisine"}', 10, true),
    (gen_random_uuid(), 'foodie123', 'foodie123@example.com', 'hashed_password4', 'Sarah Lee', 'customer', '321 Maple St, City', '+1223344556', 'Bio for Sarah Lee', '{"Tasting"}', 2, false),
    (gen_random_uuid(), 'homecook_mary', 'mary@example.com', 'hashed_password5', 'Mary Johnson', 'chef', '654 Pine St, City', '+1445566778', 'Bio for Mary Johnson', '{"Home Cooking"}', 8, true),
    (gen_random_uuid(), 'baker_bob', 'bob@example.com', 'hashed_password6', 'Bob Baker', 'chef', '987 Birch St, City', '+1667788990', 'Bio for Bob Baker', '{"Baking"}', 6, true),
    (gen_random_uuid(), 'alice_wonder', 'alice@example.com', 'hashed_password7', 'Alice Wonderland', 'customer', '159 Cedar St, City', '+1889900112', 'Bio for Alice Wonderland', '{"Reviewing"}', 1, false),
    (gen_random_uuid(), 'gordon_ramsay', 'gordon@example.com', 'hashed_password8', 'Gordon Ramsay', 'chef', '753 Spruce St, City', '+1990011223', 'Bio for Gordon Ramsay', '{"Gourmet Cooking"}', 15, true),
    (gen_random_uuid(), 'customer_kate', 'kate@example.com', 'hashed_password9', 'Kate Brown', 'customer', '258 Willow St, City', '+1011121314', 'Bio for Kate Brown', '{"Tasting"}', 3, false),
    (gen_random_uuid(), 'chef_mike', 'mike@example.com', 'hashed_password10', 'Mike Tyson', 'chef', '369 Hickory St, City', '+2021222324', 'Bio for Mike Tyson', '{"Grilling"}', 12, true);
