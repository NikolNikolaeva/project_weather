# Weather project

## Description

The purpose of the project is to manage cities' data and fetching weather forecast using a third party.

## Features

Add, delete city
Retrieve weather forecasts for a registered city by hour, day, or week.

## Technologies

DB: PostgreSQL

## Generate repositories mocks:

    1. cd ./repositories
    2. `mockgen --build_flags=--mod=mod -package repositories -destination ./../mocks/mock_city_repo.go . CityDB`
    3. mockgen --build_flags=--mod=mod -package repositories -destination ./../mocks/mock_forecast_repo.go . ForecastDB

## Generate services mocks:

    1. cd ./services
    2. `mockgen --build_flags=--mod=mod -package repositories -destination ./../mocks/mock_city_repo.go . WeatherDataGetter`
