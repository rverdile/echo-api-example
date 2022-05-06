CREATE TABLE IF NOT EXISTS users(
                                    id serial PRIMARY KEY,
                                    email VARCHAR (300) UNIQUE NOT NULL,
                                    password VARCHAR (50) NOT NULL,
                                    name VARCHAR (50) NOT NULL,
                                    created_at timestamp NOT NULL,
                                    updated_at timestamp NOT NULL,
                                    deleted_at timestamp
);