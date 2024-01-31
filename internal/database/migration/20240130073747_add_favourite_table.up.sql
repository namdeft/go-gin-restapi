CREATE TABLE favourite (
    user_id INT NOT NULL, dish_id INT NOT NULL, PRIMARY KEY (user_id, dish_id), FOREIGN KEY (user_id) REFERENCES user (ID) ON DELETE CASCADE, FOREIGN KEY (dish_id) REFERENCES dish (ID) ON DELETE CASCADE
);