-- Insert Users
INSERT INTO "User" (username, password, created_at) VALUES
('admin1', '123456', NOW()),
('admin2', '123456', NOW()),
('chef1', '123456', NOW()),
('chef2', '123456', NOW()),
('customer1', '123456', NOW()),
('customer2', '123456', NOW()),
('customer3', '123456', NOW()),
('customer4', '123456', NOW()),
('customer5', '123456', NOW()),
('customer6', '123456', NOW()),
('customer7', '123456', NOW());

-- Insert Admins
INSERT INTO "Admin" (user_id, name, created_at) VALUES
(1, 'Admin One', NOW()),
(2, 'Admin Two', NOW());

-- Insert Chefs
INSERT INTO "Chef" (user_id, name, created_at) VALUES
(3, 'Chef One', NOW()),
(4, 'Chef Two', NOW());

-- Insert Customers
INSERT INTO "Customer" (user_id, name, created_at) VALUES
(5, 'Customer One', NOW()),
(6, 'Customer Two', NOW()),
(7, 'Customer Three', NOW()),
(8, 'Customer Four', NOW()),
(9, 'Customer Five', NOW()),
(10, 'Customer Six', NOW()),
(11, 'Customer Seven', NOW());

-- Insert Menu Items
INSERT INTO "Item" (name, price, created_at) VALUES
('Nasi Goreng', 30000, NOW()),
('Mie Goreng', 25000, NOW()),
('Sate Ayam', 40000, NOW()),
('Ayam Goreng', 35000, NOW()),
('Es Teh Manis', 5000, NOW());

-- Insert Orders with Customer Relationships
INSERT INTO "Order" (customer_id, total_price, discount, rating, created_at) VALUES
(7, 75000, 10000, 3, NOW()),
(2, 120000, 5000, 5, NOW()),
(1, 50000, 0, 4, NOW());

-- Insert Order Items with Order and Item Relationships
INSERT INTO "OrderItem" (order_id, item_id, quantity, status, created_at) VALUES
(1, 5, 1, 'ordered', NOW()), -- Es Teh Manis
(2, 2, 2, 'ordered', NOW()), -- Mie Goreng
(2, 5, 1, 'ordered', NOW()), -- Es Teh Manis
(2, 4, 1, 'ordered', NOW()), -- Ayam Goreng
(3, 3, 1, 'ordered', NOW()); -- Sate Ayam
