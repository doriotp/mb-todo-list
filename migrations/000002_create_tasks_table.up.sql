CREATE TABLE tasks (
    id SERIAL PRIMARY KEY,                
    user_id int,                 
    title VARCHAR(255) NOT NULL,          
    description VARCHAR(255) NOT NULL,            
    is_completed BOOLEAN DEFAULT false,       
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
