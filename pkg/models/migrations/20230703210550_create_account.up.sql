CREATE TABLE accounts (
  id BIGINT PRIMARY KEY AUTO_INCREMENT,
  created_at DATETIME NOT NULL DEFAULT current_timestamp,
  updated_at DATETIME NOT NULL DEFAULT current_timestamp,
  type TINYINT NOT NULL,
  status TINYINT NOT NULL,
  first_name VARCHAR(255),
  last_name VARCHAR(255),
  email VARCHAR(255) NOT NULL UNIQUE,
  mobile_number VARCHAR(255),
  password VARCHAR(255)
);
