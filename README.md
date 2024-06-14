# Carnivorous Green House

![Carnivorous Green House](./static/index_image.png)

## Description

This project is a simulation of a carnivorous plant greenhouse. The greenhouse is populated with a variety of carnivorous plant. This project is used within the Grafana Loki fundermentals course to showcase log monitoring and alerting.

## Features

- [x] Plant simulation
- [x] User authentication
- [x] Bug toggle switch (cause the server to randomly throw errors)
- [ ] Historic plant data
- [ ] Microservices architecture mode

## Installation - Local

This project requires `python 3.10` to run locally. To install the project, run the following commands:

1. Clone the repository
```bash
git clone https://github.com/Jayclifford345/carnivorous_green_house.git
```

2. Install the project dependencies
```bash
pip install -r requirements.txt
```

3. Run the project
```bash
python app.py
```

4. Open a web browser and navigate to `http://localhost:5000`



## Installation - Docker

This project can also be run using Docker. To run the project using Docker, run the following commands:
1. Clone the repository
```bash
git clone https://github.com/Jayclifford345/carnivorous_green_house.git
```

2. Run using Docker Compose
```bash
docker-compose up -d
```

3. Open a web browser and navigate to `http://localhost:5005`


## Contributing

If you would like to contribute to this project, please follow these steps:

1. Fork the repository.
2. Create a new branch for your feature or bug fix.
3. Make your changes and commit them.
4. Push your changes to your forked repository.
5. Submit a pull request to the original repository.

Please ensure that your code follows the project's coding conventions and includes appropriate tests. Thank you for your contribution!


## What is Loki Overview
https://www.youtube.com/watch?v=TLnH7efQNd0&list=PLDGkOdUX1Ujr9QOsM--ogwJAYu6JD48W7

## Reference Videos
https://youtube.com/live/PLgwz5mDZ_Y?feature=share (golang with slog)
https://youtube.com/live/pfwDbZHVW1w?feature=share
https://youtube.com/live/YcSs-jvI0xw?feature=share
https://youtube.com/live/HMyRPltA_dc?feature=share (not working for now)
https://youtube.com/live/jckCXI87Osg?feature=share