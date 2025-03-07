package main

func main() {

    tickers := map[string]string{
        "GOOG": "Google Inc",
        "MSFT": "Microsoft",
        "FB":   "FaceBook",
        "AMZN": "Amazon",
    }

    // map 키 체크
    val, exists := tickers["MSFT"]

    if !exists {
        println("No MSFT ticker")
    }

}

