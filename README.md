# Modern-Go-App-Structure

Это учебное приложение на Go демонстрирует современную структуру проекта и работу с User-Agent. Приложение запускает веб-сервер, который показывает информацию о клиенте (тип устройства, архитектуру процессора, ОС, браузер и т.д.).

---

## 🚀 Что делает приложение?

После запуска сервера (по умолчанию `localhost:8080`) вы увидите:

- **Тип устройства**: Десктоп, мобильное устройство или планшет.
- **Операционная система**: Например, Windows, macOS, Android или iOS.
- **Браузер и его версия**: Chrome, Safari, Firefox и т.д.
- **Количество ядер процессора** (на стороне сервера).
- **Архитектура процессора** (на стороне сервера).

---

## 🛠️ Технологии

- **Go**: Язык программирования.
- **github.com/mssola/useragent**: Библиотека для анализа User-Agent.
- **Makefile**: Автоматизация процессов сборки/тестирования.

---

## 📦 Установка и запуск

### 1. Клонирование репозитория

```bash
git clone https://github.com/Folombas/modern-go-app-structure.git
cd modern-go-app-structure
```
