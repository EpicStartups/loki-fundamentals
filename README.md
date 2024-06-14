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


## Watch all these videos in sequence
https://www.youtube.com/watch?v=TLnH7efQNd0&list=PLDGkOdUX1Ujr9QOsM--ogwJAYu6JD48W7
https://youtube.com/live/pfwDbZHVW1w?feature=share \
https://youtube.com/live/YcSs-jvI0xw?feature=share \
https://youtube.com/live/jckCXI87Osg?feature=share
https://youtube.com/live/PLgwz5mDZ_Y?feature=share (golang with slog)



sample slog template
```
package main

import (
	"context"
	"errors"
	"log/slog"
	"os"
	"time"
)

// CustomType represents a custom type that implements the LogValuer interface.
type CustomType struct {
	ID   int
	Name string
}

// LogValue implements the LogValuer interface for CustomType.
func (c CustomType) LogValue() slog.Value {
	return slog.GroupValue(
		slog.Int("id", c.ID),
		slog.String("name", c.Name),
	)
}

// Secret represents a sensitive value that should be redacted in logs.
type Secret string

// LogValue implements the LogValuer interface for Secret.
func (s Secret) LogValue() slog.Value {
	return slog.StringValue("[REDACTED]")
}

func main() {
	// Create a new JSON handler with custom options
	opts := &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelInfo,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.MessageKey {
				a.Key = "message"
			}
			return a
		},
	}
	jsonHandler := slog.NewJSONHandler(os.Stdout, opts)

	// Create a new logger with the JSON handler
	logger := slog.New(jsonHandler)

	// Create a logger with additional context
	userLogger := logger.With(
		slog.Int("userID", 123),
		slog.String("requestID", "abc123"),
		slog.String("path", "/api/users"),
	)

	// Log messages at different levels
	pass2 := Secret("thetestsecret")
	userLogger.Debug("Debug message", "key", "value")
	userLogger.Debug("Debug message 1", "key", "value")
	userLogger.Debug("Debug message 2", "key", "value")
	userLogger.Debug("Debug message 3", "key", "value")
	userLogger.Debug("Debug message 4", "key", "value")
	userLogger.Info("Info message", "key", "value", "key2", "value2", "just a key", pass2)
	userLogger.Warn("Warning message", "key", "value")
	userLogger.Error("Error message", "key", "value")

	// Log with a custom type that implements LogValuer
	customValue := CustomType{ID: 456, Name: "Custom Value"}
	userLogger.Info("Custom type", "custom", customValue)

	// Log with a sensitive value that should be redacted
	password := Secret("secret123")
	userLogger.Info("Sensitive data ---------------------", "password", password)

	// Log with a group
	userLogger.Info("Grouped message",
		slog.Group("group1",
			slog.Int("intKey", 123),
			slog.String("stringKey", "value"),
		),
		slog.Group("group2",
			slog.Bool("boolKey", true),
			slog.Duration("durationKey", time.Second),
		),
	)

	// Log with a context
	ctx := context.Background()
	userLogger.LogAttrs(ctx, slog.LevelInfo, "Context message",
		slog.String("contextKey", "contextValue"),
		slog.Time("timestamp", time.Now()),
	)

	// Use a LevelVar to dynamically control the logging level
	levelVar := new(slog.LevelVar)
	levelVar.Set(slog.LevelInfo)

	dynamicLogger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: levelVar,
	}))

	dynamicLogger.Info("Dynamic logging", "key", "value")
	levelVar.Set(slog.LevelDebug)
	dynamicLogger.Debug("Debug message after level change", "key", "value")

	// Log an error with additional context
	err := performOperation()
	if err != nil {
		userLogger.Error("Operation failed", err,
			slog.String("context", "additional context"),
			slog.Int("retryCount", 3),
			slog.Duration("retryDelay", 5*time.Second),
		)
	}
}

func performOperation() error {
	// Simulating an error condition
	return errors.New("operation failed")
}
```
