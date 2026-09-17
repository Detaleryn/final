# Планировщик задач
Планировщик позволяет добавлять задачи, получать список ближайших задач, редактировать и удалять их, а также отмечать задачи выполненными.
Задания повышенной трудности не выполнялись.
# Запуск локально
Для запуска необходим установленный Go.  
Запуск проекта: go run .
По умолчанию веб-сервер запускается на порту 7540.
После запуска планировщик доступен по адресу:
http://localhost:7540/
# Запуск тестов
Перед запуском тестов необходимо запустить веб-сервер.
## Полный набор тестов:
go test ./tests
## Отдельные тесты:
go test -run ^TestApp$ ./tests  
go test -run ^TestDB$ ./tests  
go test -run ^TestNextDate$ ./tests  
go test -run ^TestAddTask$ ./tests  
go test -run ^TestTasks$ ./tests  
go test -run ^TestTask$ ./tests  
go test -run ^TestEditTask$ ./tests  
go test -run ^TestDone$ ./tests  
go test -run ^TestDelTask$ ./tests  
# Настройки тестов
Настройки находятся в файле: tests/settings.go  
Используемые значения:  
var Port = 7540  
var DBFile = "../scheduler.db"  
var FullNextDate = false  
var Search = false  
var Token = ``  
FullNextDate = false оставляет отключёнными тесты правил повторения для задания с повышенной сложностью.  
Search = false соответствует отсутствию реализации поиска задач.