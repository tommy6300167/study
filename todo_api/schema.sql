-- Todo API Database Schema

-- Create database
CREATE DATABASE IF NOT EXISTS todo_api;
USE todo_api;

-- Create todos table
CREATE TABLE todos (
    id INT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    completed BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    
    INDEX idx_completed (completed),
    INDEX idx_created_at (created_at)
);

-- Insert sample data
INSERT INTO todos (title, description, completed) VALUES
('Learn Go', 'Study Go programming language basics', false),
('Build REST API', 'Create a todo list REST API with Gin framework', false),
('Write Tests', 'Add unit tests for the API endpoints', false),
('Deploy Application', 'Deploy the application to production', false);