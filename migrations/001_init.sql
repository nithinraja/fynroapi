CREATE DATABASE IF NOT EXISTS fynroapi;

USE fynroapi;

CREATE TABLE `questions` (
  `id` int(11) NOT NULL AUTO_INCREMENT PRIMARY KEY,
  `questionid` varchar(36) NOT NULL,
  `question` text NOT NULL,
  `username` varchar(255) DEFAULT NULL,
  `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  KEY `idx_questionid` (`questionid`)  -- Optional index for faster lookups
);




CREATE TABLE IF NOT EXISTS responses (
    id INT AUTO_INCREMENT PRIMARY KEY,
    questionid VARCHAR(36) NOT NULL,
    response TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

