-- 创建数据库
CREATE DATABASE IF NOT EXISTS vital_fitness CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
USE vital_fitness;

-- 用户表
CREATE TABLE IF NOT EXISTS users (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,

    username VARCHAR(50) NOT NULL UNIQUE,
    email VARCHAR(100) NOT NULL UNIQUE,
    phone VARCHAR(20) UNIQUE,
    password VARCHAR(255) NOT NULL,
    avatar VARCHAR(255),
    nickname VARCHAR(50),
    gender TINYINT DEFAULT 0 COMMENT '0:未知 1:男 2:女',
    birthday DATE,
    height DECIMAL(5,2) COMMENT '身高(cm)',
    weight DECIMAL(5,2) COMMENT '体重(kg)',

    INDEX idx_username (username),
    INDEX idx_email (email),
    INDEX idx_phone (phone),
    INDEX idx_deleted_at (deleted_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- 健身动作表
CREATE TABLE IF NOT EXISTS exercises (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,

    name VARCHAR(100) NOT NULL,
    category VARCHAR(50),
    description TEXT,
    video_url VARCHAR(255),
    image_url VARCHAR(255),

    INDEX idx_name (name),
    INDEX idx_category (category),
    INDEX idx_deleted_at (deleted_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- 训练记录表
CREATE TABLE IF NOT EXISTS workouts (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,

    user_id BIGINT UNSIGNED NOT NULL,
    exercise_id BIGINT UNSIGNED NOT NULL,
    workout_date DATE NOT NULL,

    weight DECIMAL(10,2),
    sets INT NOT NULL DEFAULT 1,
    reps INT NOT NULL DEFAULT 1,
    rest_time INT COMMENT '休息时间(秒)',

    feeling TINYINT DEFAULT 3 COMMENT '训练感受(1-5分)',
    notes TEXT,
    image_url VARCHAR(255),

    INDEX idx_user_id (user_id),
    INDEX idx_exercise_id (exercise_id),
    INDEX idx_workout_date (workout_date),
    INDEX idx_deleted_at (deleted_at),

    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (exercise_id) REFERENCES exercises(id) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- 心情记录表
CREATE TABLE IF NOT EXISTS mood_records (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,

    user_id BIGINT UNSIGNED NOT NULL,
    record_date DATE NOT NULL,
    mood_score TINYINT NOT NULL COMMENT '心情分数(1-10)',
    mood_tags VARCHAR(255),
    description TEXT,

    INDEX idx_user_id (user_id),
    INDEX idx_record_date (record_date),
    INDEX idx_mood_score (mood_score),
    INDEX idx_deleted_at (deleted_at),

    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    CONSTRAINT chk_mood_score CHECK (mood_score >= 1 AND mood_score <= 10)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- 体重记录表
CREATE TABLE IF NOT EXISTS weight_records (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,

    user_id BIGINT UNSIGNED NOT NULL,
    record_date DATE NOT NULL,
    weight DECIMAL(5,2) NOT NULL COMMENT '体重(kg)',
    height DECIMAL(5,2) COMMENT '身高(cm)',
    bmi DECIMAL(5,2) COMMENT 'BMI值',

    INDEX idx_user_id (user_id),
    INDEX idx_record_date (record_date),
    INDEX idx_bmi (bmi),
    INDEX idx_deleted_at (deleted_at),

    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- 饮食记录表
CREATE TABLE IF NOT EXISTS diet_records (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,

    user_id BIGINT UNSIGNED NOT NULL,
    record_date DATE NOT NULL,
    meal_type VARCHAR(20) NOT NULL COMMENT 'breakfast,lunch,dinner,snack',
    food_items TEXT COMMENT '食物列表(JSON格式)',
    total_calories DECIMAL(10,2),
    protein DECIMAL(10,2) COMMENT '蛋白质(g)',
    carbs DECIMAL(10,2) COMMENT '碳水化合物(g)',
    fat DECIMAL(10,2) COMMENT '脂肪(g)',
    water_intake INT COMMENT '饮水量(ml)',
    notes TEXT,

    INDEX idx_user_id (user_id),
    INDEX idx_record_date (record_date),
    INDEX idx_meal_type (meal_type),
    INDEX idx_deleted_at (deleted_at),

    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- 插入一些基础动作数据
INSERT INTO exercises (name, category, description) VALUES
('杠铃卧推', '胸部', '经典的胸部训练动作，主要锻炼胸大肌'),
('哑铃飞鸟', '胸部', '孤立训练胸部的动作，增加胸部宽度'),
('俯卧撑', '胸部', '自重训练动作，方便实用'),
('引体向上', '背部', '经典的背部训练动作，主要锻炼背阔肌'),
('杠铃划船', '背部', '复合型背部训练动作'),
('哑铃划船', '背部', '单侧背部训练，纠正左右不平衡'),
('杠铃深蹲', '腿部', '最全面的下肢训练动作'),
('腿举', '腿部', '器械训练动作，安全性较高'),
('腿弯举', '腿部', '专门训练腘绳肌的动作'),
('杠铃推举', '肩部', '经典的肩部训练动作'),
('哑铃侧平举', '肩部', '孤立训练三角肌中束'),
('哑铃弯举', '手臂', '经典的肱二头肌训练动作'),
('杠铃弯举', '手臂', '大重量肱二头肌训练'),
('平板支撑', '核心', '静态核心稳定性训练');