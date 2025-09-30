package main

import "fmt"

type Loggers interface {
	Log() string
}

type OldLogger struct {
	message string
}

func (l *OldLogger) Log() string { //старый логгер удовлетворяет интерфейсу Loggers
	return l.message
}

type NewLogger struct {
}

func (l *NewLogger) PrintMessage(message string) string { //новый логгер не удовлетворяет интерфейсу Loggers
	return message
}

type LoggerAdapter struct { //вписываем в адаптер дополнительное поле - message, помимо указателя на newlogger
	logger  *NewLogger
	message string
}

func NewLoggerAdapter(message string) Loggers { //конструктор адаптера нового логгера, удовлетворяющего интерфейсу Loggers
	return &LoggerAdapter{
		logger:  &NewLogger{},
		message: message}
}

func (l *LoggerAdapter) Log() string { //вписываем в уже знакомую форму Log() функцию нового логгера
	return l.logger.PrintMessage(l.message)
}

func main() {
	old := &OldLogger{message: "old logger"}               //старый логгер
	newloggerAdapter := NewLoggerAdapter("logger adapter") //новый логгер

	loggers := []Loggers{old, newloggerAdapter}
	for _, v := range loggers {
		fmt.Println(v.Log()) //используйм Log() через адаптер для нового логгера
	}

}
