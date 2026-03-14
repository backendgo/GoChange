Вот простой пример консольного криптообменника на Go. Эта программа запрашивает актуальный курс BTC/USD через публичное API (CoinGecko) и рассчитывает сумму обмена.
Что делает эта программа:

   1. Подключается к CoinGecko API.
   2. Получает текущую цену Bitcoin в долларах.
   3. Позволяет пользователю ввести сумму в USD и конвертировать её в BTC (или наоборот).

package main
import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)
// Структура для ответа от CoinGecko APItype PriceResponse struct {
	Bitcoin struct {
		USD float64 `json:"usd"`
	} `json:"bitcoin"`
}
func getBTCPrice() (float64, error) {
	resp, err := http.Get("https://api.coingecko.com")
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	var result PriceResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return 0, err
	}

	return result.Bitcoin.USD, nil
}
func main() {
	fmt.Println("=== Мини-Криптообменник ===")
	
	price, err := getBTCPrice()
	if err != nil {
		fmt.Println("Ошибка при получении курса:", err)
		return
	}

	fmt.Printf("Текущий курс BTC: $%.2f\n\n", price)

	fmt.Print("Введите сумму в USD для покупки BTC: ")
	var usdAmount float64
	_, err = fmt.Scanln(&usdAmount)
	if err != nil {
		fmt.Println("Ошибка ввода. Пожалуйста, введите число.")
		return
	}

	btcAmount := usdAmount / price
	fmt.Printf("За $%.2f вы получите %.8f BTC\n", usdAmount, btcAmount)
}

Как запустить:

   1. Установите Go.
   2. Создайте файл main.go и вставьте туда код выше.
   3. Откройте терминал и выполните:
   
   go run main.go
   
   
Основные нюансы:

* API: Мы используем CoinGecko, так как оно бесплатное и не требует регистрации ключа для простых запросов.
* Безопасность: Настоящие обменники требуют работы с блокчейном (например, через библиотеку ethclient для Ethereum) и строгой проверки транзакций.
Хотите добавить в программу поддержку других валют (например, Ethereum) или сделать её в виде веб-сервиса?

