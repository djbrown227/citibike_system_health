# Citi Bike Station Data Processor

This Go repository provides tools to fetch, process, analyze, and detect anomalies in Citi Bike station data, with a focus on helping operators maintain a balanced and efficient bike-sharing system. The project addresses issues such as bike availability monitoring, proximity analysis, and real-time anomaly detection to optimize operational performance and enhance user experience.

## Problem Statement

Citi Bike’s Bike Angels program, introduced in 2016 by its parent company Lyft, was designed to encourage users to help balance the distribution of bikes across New York City. By moving bikes from overcrowded docking stations to those in need, participants could earn points that could be redeemed for Lyft credits, future bike rentals, or even cash transfers.

While this program has successfully mobilized a community of riders, an unintended consequence has emerged: some users have begun manipulating the system to profit significantly by intentionally shifting bikes between stations to earn points.

One such participant, Mark Epperson, along with 10 to 15 other frequent users, has found ways to exploit the system for financial gain, earning between $1,000 and $2,000 per month by moving bikes for several hours daily. Some riders have even reported making upwards of $7,000 to $8,000 per month. This practice, referred to as "station flipping," involves intentionally creating bike shortages to maximize point accumulation.

Although this behavior isn’t illegal, it conflicts with the program's intended purpose and has prompted concerns from both Lyft and critics, such as David Shmoys, a data science professor at Cornell University. Lyft recently addressed the issue by sending a letter to Bike Angels participants, urging them to refrain from this exploitative practice.

This project aims to address this issue by improving monitoring and detection capabilities for anomalies in bike availability, helping to identify and mitigate suspicious patterns that may indicate system manipulation.

## Features

### 1. Citi Bike API Data Fetcher

This module fetches both real-time dynamic data and static station information from the Citi Bike API.

- **Dynamic Station Data**: Retrieves real-time information about bike availability, electric bikes (e-bikes), scooters, docks, and station status (e.g., whether renting or returning bikes).

  Key Data Includes:
  - Available vehicle types (e.g., bikes, e-bikes, scooters)
  - Disabled bikes, docks, and scooters
  - Last report timestamp
  - Station installation and operational status

- **Static Station Information**: Fetches static details such as station name, geographic coordinates (latitude/longitude), station capacity, and rental URLs for mobile apps.

#### Key Functions:
- `FetchData()`: Retrieves dynamic station data, including real-time bike and dock availability.
- `FetchStationInfo()`: Fetches static station details like location, name, and total capacity.

### 2. Station Data Processing

This module processes and logs Citi Bike station data, offering real-time insights and historical tracking.

- **Logging Station Data**: Records real-time station availability data (bikes, e-bikes, scooters, docks) for historical tracking.

  Key Features:
  - `SetupLogFile`: Initializes a log file for storing data.
  - `LogStationData`: Logs real-time data for each station, capturing the number of available bikes, scooters, and dock status.

- **Station Data Analytics**:
  - `PrintStationDetails`: Outputs real-time station data, including location and availability, to the console for easy monitoring.
  - `PrintClosestStations`: Displays the 10 closest stations to a given location based on calculated distances.
  - `CalculatePercentFilled`: Computes the percentage of bikes available relative to station capacity.
  - `CalculatePercentEmpty`: Calculates the percentage of empty docks based on station availability.

- **Station Mapping**:
  - `CreateStationMap`: Builds a map of stations indexed by `StationID` for efficient lookup and data processing.

### 3. Distance Calculation and Anomaly Detection

This module provides geographic distance calculations between stations and detects anomalies in bike availability.

- **Distance Calculation**:
  - **Haversine Formula**: Uses the Haversine formula to calculate the distance between two geographic points based on their latitude and longitude.
  - `FindClosestStations`: Returns the 10 closest stations to a given station based on the calculated distance.

- **Anomaly Detection**:
  - **AnomalyDetection**: Monitors station data over time to detect unusual changes in bike availability (e.g., a significant drop in bikes available within a short time window).
  - `abs`: A helper function to calculate the absolute difference in available bikes between two consecutive records.

#### Use Cases:
- **Proximity Analysis**: Helps users identify the closest stations for convenience.
- **Real-Time Monitoring**: Detects suspicious patterns in bike availability, providing operational alerts for potential issues.

## How to Improve Anomaly Detection

The current anomaly detection system uses simple rule-based logic to detect significant changes in bike availability. This system could be enhanced by applying more sophisticated methods, such as:

1. **Statistical Anomaly Detection**: 
   - Use statistical models such as moving averages, Z-scores, or standard deviation thresholds to detect outliers.
   - Example: A station could be flagged if its bike availability deviates significantly from its historical average.

2. **Time Series Analysis**: 
   - Apply time series analysis techniques like ARIMA or Exponential Smoothing to predict expected availability, and flag deviations from predictions.
   - This would allow the system to learn patterns over time and identify anomalies based on these trends.

3. **Machine Learning Models**: 
   - Implement machine learning models such as Random Forests, Gradient Boosting, or Neural Networks to detect anomalies based on a combination of features (e.g., time of day, weather, and location).
   - These models could be trained on historical data to improve anomaly detection accuracy.

4. **Clustering Techniques**:
   - Use clustering techniques like DBSCAN or k-means to detect stations with behavior that deviates from their peers (e.g., stations with consistently lower availability than others in the same geographic area).
  
5. **Real-Time Alerting System**:
   - Add a real-time alert system that notifies operators when anomalies are detected via email, SMS, or a dashboard, allowing for quick intervention.

## Conclusion

This Go repository provides comprehensive tools for working with Citi Bike station data. From fetching real-time availability to detecting anomalies and calculating proximity, the system aims to enhance operational efficiency and provide actionable insights. Future improvements in anomaly detection will make the system even more robust and responsive to potential issues in the bike-sharing network.
