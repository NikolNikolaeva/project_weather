CREATE TABLE IF NOT EXISTS City
(
    id         VARCHAR(36) DEFAULT gen_random_uuid() PRIMARY KEY,
    name       VARCHAR(100) NOT NULL,
    country    VARCHAR(100) NOT NULL,
    latitude   VARCHAR(100) NOT NULL,
    longitude  VARCHAR(100) NOT NULL,
    created_at TIMESTAMP   DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP   DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT unique_City UNIQUE (name, country)
);

CREATE TABLE IF NOT EXISTS Forecast
(
    id            VARCHAR(36) DEFAULT gen_random_uuid() PRIMARY KEY,
    city_id       varchar(36) NOT NULL,
    forecast_date DATE        NOT NULL,
    temperature   VARCHAR(100),
    condition     VARCHAR(100),
    created_at    TIMESTAMP   DEFAULT CURRENT_TIMESTAMP,
    updated_at    TIMESTAMP   DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (city_id) REFERENCES City (id) ON DELETE CASCADE,
    CONSTRAINT unique_Forecast UNIQUE (city_id, forecast_date)

);