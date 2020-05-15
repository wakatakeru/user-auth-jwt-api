-- Create Table
CREATE DATABASE userauth.users (
  id INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
  name VARCHAR(64) NOT NULL,
  display_name VARCHAR(64),
  email VARCHAR(128),
  password VARCHAR(512)
);

-- Grant Access
GRANT ALL ON *.* TO userauth;
