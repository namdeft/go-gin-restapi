CREATE TABLE category (
    ID INT NOT NULL AUTO_INCREMENT, name VARCHAR(45), PRIMARY KEY (ID)
);

CREATE TABLE ingredient (
    ID INT NOT NULL AUTO_INCREMENT, name VARCHAR(45), quantity INT, import_date DATE, expired_date DATE, couting_unit VARCHAR(45), PRIMARY KEY (ID)
);

CREATE TABLE dish (
    ID INT NOT NULL AUTO_INCREMENT, name VARCHAR(45) NOT NULL, price DOUBLE NOT NULL, status ENUM(
        'unavailable', 'available', 'deleted'
    ) DEFAULT 'unavailable', created_at DATETIME DEFAULT CURRENT_TIMESTAMP, updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP, deleted_at DATETIME DEFAULT NULL, PRIMARY KEY (ID)
);

CREATE TABLE origin (
    ID INT NOT NULL AUTO_INCREMENT, name VARCHAR(45), PRIMARY KEY (ID)
);

CREATE TABLE dish_ingredient (
    dish_id INT, ingredient_id INT, PRIMARY KEY (dish_id, ingredient_id), FOREIGN KEY (dish_id) REFERENCES dish (ID) ON DELETE CASCADE, FOREIGN KEY (ingredient_id) REFERENCES ingredient (ID) ON DELETE CASCADE
);

CREATE TABLE dish_origin (
    dish_id INT NOT NULL, origin_id INT NOT NULL, PRIMARY KEY (dish_id, origin_id), FOREIGN KEY (dish_id) REFERENCES dish (ID) ON DELETE CASCADE, FOREIGN KEY (origin_id) REFERENCES origin (ID) ON DELETE CASCADE
);

CREATE TABLE dish_category (
    dish_id INT NOT NULL, category_id INT NOT NULL, PRIMARY KEY (dish_id, category_id), FOREIGN KEY (dish_id) REFERENCES dish (ID) ON DELETE CASCADE, FOREIGN KEY (category_id) REFERENCES category (ID) ON DELETE CASCADE
);

CREATE TABLE user (
    ID int NOT NULL AUTO_INCREMENT, name varchar(255) NOT NULL, email varchar(255) NOT NULL, password varchar(255) NOT NULL, PRIMARY KEY (ID), UNIQUE KEY email (email)
)

CREATE TABLE favourite (
    user_id INT NOT NULL, dish_id INT NOT NULL, PRIMARY KEY (user_id, dish_id), FOREIGN KEY (user_id) REFERENCES user (ID) ON DELETE CASCADE, FOREIGN KEY (dish_id) REFERENCES dish (ID) ON DELETE CASCADE
);