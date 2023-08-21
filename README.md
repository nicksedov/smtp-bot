sbconn-bot
========================

Параметры командной строки 
------------------------------------------------------------------------------------
Имя           | Описание                                | Значение по умолчанию
------------------------------------------------------------------------------------
name          | Отображаемое имя сервера                | sbconn-bot
listen        | Адрес SMTP-сервера                      | :smtp
msglimit      | Макс. размер входящего сообщения, байт  | 2097152 ( = 2 Мб)
timeout.read  | Таймаут чтения, сек.                    | 5
timeout.write | Таймаут записи, сек.                    | 5
config        | файл конфигурации (YAML)                | sbconn-settings.yaml
bot.token     | Токен Telegram-бота                     |
openai.token  | API-токен OpenAI                        |
------------------------------------------------------------------------------------

Пример вызова командной строки:
/opt/sbconn-bot/sbconn-bot -listen=:<port> -config=<path/to/settings_yaml> -bot.token=<telegram_token> -openai.token=<openai_api_token>

Contribution
============
SMTP Based on original repo from @alash3al
Thanks to @aranajuan


