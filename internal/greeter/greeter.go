package greeter

// SayHello возвращает персонализированное приветствие
func SayHello(name string) string {
	return "Привет, " + name + "! Добро пожаловать в мир Go-программирования!"
}

// Внутренняя функция (не экспортируется)
func internalFunction() string {
	return "Этот текст недоступен вне пакета greeter"
}
