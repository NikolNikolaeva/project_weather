CREATE TABLE IF NOT EXISTS City
(
    id         VARCHAR(36) DEFAULT gen_random_uuid() PRIMARY KEY,
    name       VARCHAR(100) UNIQUE NOT NULL,
    country    VARCHAR(100) UNIQUE NOT NULL,
    latitude   DECIMAL(9, 6)       NOT NULL,
    longitude  DECIMAL(9, 6)       NOT NULL,
    created_at TIMESTAMP   DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP   DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS Forecast
(
    id            VARCHAR(36) DEFAULT gen_random_uuid() PRIMARY KEY,
    city_id       INT  NOT NULL,
    forecast_date DATE NOT NULL,
    temperature   DECIMAL(5, 2),
    condition     VARCHAR(100),
    created_at    TIMESTAMP   DEFAULT CURRENT_TIMESTAMP,
    updated_at    TIMESTAMP   DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (city_id) REFERENCES City (id) ON DELETE CASCADE
);