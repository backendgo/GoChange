package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

// Структура для ответа от CoinGecko API
type PriceResponse struct {
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
