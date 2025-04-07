# Ethereum Balance to USD Converter

Этот проект позволяет получить баланс Ethereum (ETH) и Wrapped Ethereum (WETH) на заданном Ethereum-адресе и пересчитать их в эквивалент в USD с использованием данных с Chainlink Price Feed, все через on-chain запросы. Это приложение использует Infura для подключения к сети Ethereum и взаимодействует с контрактами Chainlink для получения цены ETH в USD.

## Описание

- Программа получает Ethereum-адрес из командной строки.
- Подключается к сети Ethereum через Infura для получения баланса ETH.
- Получает баланс WETH с помощью стандартного контракта ERC20.
- Использует Chainlink Price Feed для получения актуальной цены ETH в USD.
- Пересчитывает балансы ETH и WETH в USD и выводит результаты.

## Структура проекта

Проект разделен на несколько пакетов для улучшения структуры и читабельности:

- `cmd/internal/main.go`: Основной исполнимый файл.
- `pkg/ethclient`: Пакет для взаимодействия с Ethereum, получения баланса и цены.
- `cmd/config`: Пакет для загрузки конфигурации и работы с переменными окружения.

## Установка

1. Склонируйте репозиторий:

    ```bash
    git clone https://github.com/yourusername/ethereum-balance-to-usd.git
    cd ethereum-balance-to-usd
    ```

2. Установите зависимости:

    ```bash
    go mod tidy
    ```

3. Создайте файл `.env` в корневой директории проекта и добавьте в него следующие переменные:

    ```env
    INFURA_PROJECT_ID=your_infura_project_id
    WETH_CONTRACT_ADDRESS=0xC02aaA39b223FE8D0A0e5C4f27eAD9083C756Cc2
    CHAINLINK_ETH_USD_ADDRESS=0x5f4ec3df9cbd43714fe2740f5e3616155c5b8419
    ```

    **Замените** `your_infura_project_id` на ваш реальный Infura Project ID.

## Запуск

Чтобы запустить приложение, используйте команду:

```bash
go run cmd/ethereum-balance-to-usd/main.go <ethereum-address>
