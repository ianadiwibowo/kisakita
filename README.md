# kisakita

_"Ada sakit dalam Kisakita"_ (_"There was pain in our story"_)

Kisakita is an (hypotethical) app for story writing collaboration. You write story together with your friends, or even total strangers, in a turn-based manner. This repository is a monolith backend written in Go, with these software engineering principles in mind:
* Domain-Driven Design (DDD)
* Test-Driven Development (TDD)
* Uncle Bob's Clean Code and Clean Architecture
* SOLID principle
* Gang of Four design patterns

> "Kisakita" is a wordplay in Bahasa Indonesia, roughly translated to "our story".

## 1. Setup

### 1.1 Prerequisites

* Go v14.1
* Ruby v2.7.0

### 1.1 Install External Dependencies

```bash
make dep
bundle install
```

### 1.2 Run Unit Tests

To run complete suite of unit tests:

```bash
make test
```

To see unit tests coverage:

```bash
make cov
```

### 1.3 Compile

```bash
make build
```

### 1.4 Deploy Database

```bash
bundle exec rails db:migrate
```

## 2. Contributing

### 2.x Using Mockery

```bash
cd domain/storywriting
mockery -name=StoryRepository
# => Generating mock for: StoryRepository in file: mocks/StoryRepository.go
```
