INSERT INTO statuses (title) VALUES 
('Bronze'), 
('Silver'), 
('Gold'), 
('Platinum'), 
('Diamond');

INSERT INTO tasks (title, points) VALUES 
('Entered referral code', 10), 
('Subscribed to Telegram channel', 20), 
('Followed on Twitter', 15);

INSERT INTO referrers (code) VALUES 
('REF123'), 
('REF456'), 
('REF789');

INSERT INTO users (telegram_id, active, registered_at, user_name, status_id, last_visit, balance, isadmin, referrer_id) VALUES 
(1001, TRUE, '2025-01-01 10:00:00', 'Alice', 2, '2025-01-06 15:00:00', 80, FALSE, 1),  
(1002, TRUE, '2025-01-02 11:00:00', 'Bob', 3, '2025-01-06 15:30:00', 30, FALSE, 2),    
(1003, TRUE, '2025-01-03 12:00:00', 'John', 1, '2025-01-06 16:00:00', 50, FALSE, NULL), 
(1004, TRUE, '2025-01-04 13:00:00', 'David', 4, '2025-01-06 16:30:00', 70, FALSE, 1),  
(1005, TRUE, '2025-01-05 14:00:00', 'Robert', 5, '2025-01-06 17:00:00', 100, TRUE, 3);

INSERT INTO tasks_activity (user_id, task_id) VALUES 
(1, 1), 
(1, 2), 
(2, 1), 
(3, 3), 
(4, 2), 
(5, 1), 
(5, 3);
