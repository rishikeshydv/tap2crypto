# Tap2Crypto

Tap2Crypto is a cutting-edge project that combines the power of Natural Language Processing, Deep Learning, and blockchain technology to provide a secure and user-friendly crypto wallet with advanced investment insights and payment functionalities.

## Table of Contents

- [Introduction](#introduction)
- [Features](#features)
- [Technologies Used](#technologies-used)
- [Installation](#installation)
- [Usage](#usage)
- [Architecture](#architecture)
- [Contributing](#contributing)
- [License](#license)

## Introduction

Tap2Crypto is designed to facilitate seamless scan-to-pay and tap-to-pay transactions using cryptocurrencies. Additionally, it incorporates a sophisticated sentiment analysis model that examines stock market sentiment from Twitter and Yahoo News to identify highly volatile stocks and their associated call/put options ratio. Based on this analysis, Tap2Crypto can execute smart contracts to convert Ethereum to USD and perform call/put purchases, providing users with valuable investment insights and automated trading capabilities.

## Features

- **Scan-to-Pay and Tap-to-Pay**: Effortless cryptocurrency transactions using QR codes and NFC technology.
- **Sentiment Analysis**: Natural Language Processing and Deep Learning to analyze stock market sentiment from Twitter and Yahoo News.
- **Investment Insights**: Identifies highly volatile stocks and their call/put options ratio for informed decision-making.
- **Smart Contract Integration**: Converts Ethereum to USD and executes call/put purchases based on sentiment analysis recommendations.
- **User-Friendly Interface**: Intuitive and secure interface for managing crypto transactions and viewing investment insights.

## Technologies Used

- **Backend**:
  - Golang
  - Python
- **Frontend**:
  - ViteJS
- **Machine Learning**:
  - Natural Language Processing
  - Deep Learning
- **Blockchain**:
  - Ethereum
  - Smart Contracts
- **Cloud Services**:
  - AWS ECR (Elastic Container Registry)
  - AWS ECS (Elastic Container Service)

## Installation

### Prerequisites

- Go (Golang)
- Python
- Node.js
- Docker
- AWS CLI

### Steps

1. Clone the repository:
   ```sh
   git clone https://github.com/rishikeshydv/tap2crypto.git
   cd tap2crypto
   ```

2. Install dependencies:

   - Backend (Golang):
     ```sh
     cd backend
     cd golang
     go mod download
     ```

   - Frontend (ViteJS):
     ```sh
     cd frontend
     cd tap2crypto
     npm install
     ```

   - Machine Learning models (Python):
     ```sh
     cd backend
     cd golang
     cd ai
     pip install -r requirements.txt
     ```
## Usage

1. Access the Tap2Crypto via localhost.
2. Register an account and log in.
3. Use the scan-to-pay or tap-to-pay functionality for crypto transactions.
4. View investment insights and automated trading recommendations based on sentiment analysis.

## Architecture

The architecture of Tap2Crypto is designed to ensure scalability, security, and high performance. It includes the following components:

- **Frontend**: Developed with ViteJS for a fast and responsive user interface.
- **Backend**: Built with Golang to handle API requests, transaction processing, and smart contract interactions.
- **Machine Learning**: Python-based models for sentiment analysis and stock market prediction.
- **Blockchain**: Ethereum smart contracts for automated trading and currency conversion.
- **Cloud Services**: AWS ECR for container image storage and AWS ECS for container orchestration and deployment.

## Contributing

We welcome contributions to Tap2Crypto! If you would like to contribute, please follow these steps:

1. Fork the repository.
2. Create a new branch: `git checkout -b feature-branch`
3. Make your changes and commit them: `git commit -m 'Add new feature'`
4. Push to the branch: `git push origin feature-branch`
5. Submit a pull request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.


