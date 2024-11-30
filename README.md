# Feedoku

Feedoku is a modern RSS Feed aggregator built with Next.js, Go, Postgres, Redis, and Kafka. It allows users to aggregate, manage, and read their favorite RSS feeds in one place.

## 🚀 Features

- Real-time feed updates using Kafka ( Not implemented yet )
- Fast content delivery with Redis caching
- Responsive web interface built with Next.js
- RESTful API powered by Go
- Feed management and organization
- Search functionality across all feeds
- Mobile-friendly interface

## 🛠️ Tech Stack

- **Frontend**: Next.js
- **Backend**: Go (Golang)
- **Cache**: Redis
- **Message Broker**: Apache Kafka
- **Database**: PostgreSQL

## 🔧 Installation

1. Clone the repository:
```bash
git clone https://github.com/codebyaadi/feedoku.git
cd feedoku
```

2. Set up the frontend:
```bash
cd www
npm install
cp .env.example .env
# Configure your environment variables
```

3. Set up the backend: (work in main directory)
```bash
cp .env.example .env
# Configure your environment variables
```

4. Start Redis:
```bash
redis-server
```

5. Start Kafka:
```bash
# Start Zookeeper first
bin/zookeeper-server-start.sh config/zookeeper.properties

# Start Kafka
bin/kafka-server-start.sh config/server.properties
```

## 🚀 Running the Application

1. Start the frontend development server:
```bash
cd frontend
npm run dev
```

2. Start the backend server: (work in main directory)
```bash
go run main.go
```

## 🤝 Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📝 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 👥 Authors

- Aditya Rajbhar - [codebyaadi](https://github.com/codebyaadi)

## 🙏 Acknowledgments

- Thanks to all contributors
- Inspired by various RSS readers and aggregators
- Built with amazing open source technologies