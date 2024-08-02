CREATE TABLE IF NOT EXISTS customers (
  id INT AUTO_INCREMENT,
  customer_key VARCHAR(255) NOT NULL,
  customer_zip_code_prefix INT NOT NULL,
  customer_city VARCHAR(255) NOT NULL,
  customer_state CHAR(2) NOT NULL,

  PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS products (
  id INT AUTO_INCREMENT,
  product_key VARCHAR(255) NOT NULL,
  product_category_name VARCHAR(255),
  product_name_length INT,
  product_description_length INT,
  product_photos_qty INT,
  product_weight_g INT,
  product_length_cm INT,
  product_height_cm INT,
  product_width_cm INT,

  PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS orders (
  id INT AUTO_INCREMENT,
  order_key VARCHAR(255) NOT NULL,
  customer_id INT NOT NULL,
  order_status VARCHAR(50),
  order_purchase_timestamp DATETIME,
  order_approved_at DATETIME,
  order_delivered_carrier_date DATETIME,
  order_delivered_customer_date DATETIME,
  order_estimated_delivery_date DATETIME,

  PRIMARY KEY (id),
  FOREIGN KEY (customer_id) REFERENCES customers(id)
);

CREATE TABLE IF NOT EXISTS orderItems (
  id INT AUTO_INCREMENT,
  order_id INT NOT NULL,
  product_id INT NOT NULL,
  order_item_num INT NOT NULL,
  seller_key VARCHAR(255),
  shipping_limit_date DATETIME,
  price DECIMAL(10,2),
  freight_value DECIMAL(10,2),

  PRIMARY KEY (id, order_id, order_item_num),
  FOREIGN KEY (order_id) REFERENCES orders(id),
  FOREIGN KEY (product_id) REFERENCES products(id)
);

CREATE TABLE IF NOT EXISTS sellers (
  id INT AUTO_INCREMENT,
  seller_key VARCHAR(255),
  seller_zip_code_prefix INT NOT NULL,
  seller_city VARCHAR(255) NOT NULL,
  seller_state CHAR(2) NOT NULL,

  PRIMARY KEY (id)
);
