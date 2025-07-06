1. users
CREATE TABLE users (
    user_id VARCHAR(36) PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    mobile VARCHAR(15) NOT NULL UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    date_updated TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    last_login TIMESTAMP,
    is_verified BOOLEAN DEFAULT FALSE,
    status ENUM('active', 'suspended', 'deleted') DEFAULT 'active'
);

2. user_sessions
CREATE TABLE user_sessions (
    session_id VARCHAR(36) PRIMARY KEY,
    user_id VARCHAR(36) NOT NULL,
    token TEXT NOT NULL,
    device_info JSON,
    ip_address VARCHAR(45),
    expires_at TIMESTAMP NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(user_id)
);

3. questions
CREATE TABLE questions (
    questionid VARCHAR(36) PRIMARY KEY,
    user_id VARCHAR(36) NOT NULL,
    question_text TEXT NOT NULL,
    chatgpt_response JSON,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(user_id)
);

Authentication Tables
4. otp_verification
CREATE TABLE otp_verification (
    verification_id VARCHAR(36) PRIMARY KEY,
    mobile VARCHAR(15) NOT NULL,
    otp VARCHAR(6) NOT NULL,
    is_used BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    expires_at TIMESTAMP NOT NULL,
    attempts TINYINT DEFAULT 0
);

Payment & Services Tables

5. advisor_tiers
CREATE TABLE advisor_tiers (
    tier_id VARCHAR(36) PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    base_chat_price DECIMAL(10,2) NOT NULL,
    base_call_price DECIMAL(10,2) NOT NULL,
    description TEXT
);

6. user_custom_tiers
CREATE TABLE user_custom_tiers (
    custom_tier_id VARCHAR(36) PRIMARY KEY,
    user_id VARCHAR(36) NOT NULL,
    tier_id VARCHAR(36) NOT NULL,
    custom_chat_price DECIMAL(10,2),
    custom_call_price DECIMAL(10,2),
    valid_from TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    valid_until TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(user_id),
    FOREIGN KEY (tier_id) REFERENCES advisor_tiers(tier_id)
);

7. payments

CREATE TABLE payments (
    payment_id VARCHAR(36) PRIMARY KEY,
    user_id VARCHAR(36) NOT NULL,
    amount DECIMAL(10,2) NOT NULL,
    currency VARCHAR(3) DEFAULT 'INR',
    service_type ENUM('chat', 'call') NOT NULL,
    advisor_tier_id VARCHAR(36),
    payment_gateway VARCHAR(50),
    gateway_reference VARCHAR(255),
    status ENUM('pending', 'completed', 'failed', 'refunded') DEFAULT 'pending',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(user_id),
    FOREIGN KEY (advisor_tier_id) REFERENCES advisor_tiers(tier_id)
);

8. call_recordings
CREATE TABLE call_recordings (
    recording_id VARCHAR(36) PRIMARY KEY,
    user_id VARCHAR(36) NOT NULL,
    advisor_id VARCHAR(36) NOT NULL,
    channel_name VARCHAR(255) NOT NULL,
    duration_seconds INT,
    storage_url VARCHAR(512) NOT NULL,
    encryption_key_id VARCHAR(36),
    is_encrypted BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    deletion_reason VARCHAR(255),
    FOREIGN KEY (user_id) REFERENCES users(user_id)
);

9. encryption_keys
CREATE TABLE encryption_keys (
    key_id VARCHAR(36) PRIMARY KEY,
    key_data TEXT NOT NULL, /* Encrypted in database */
    algorithm VARCHAR(20) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    expires_at TIMESTAMP,
    is_active BOOLEAN DEFAULT TRUE
);