### Genel Bakış

Bu uygulama, GORM, PostgreSQL ve Go Fiber kullanarak oluşturulmuş bir URL kısaltma uygulamasıdır. Kullanıcıların uzun URL'leri kısa ve öz URL'lere dönüştürmelerine olanak tanır.

### Özellikler

* URL kısaltma ve açma

### Kurulum

1. Docker'ı yükleyin
2. docker compose up -d diyerek databasei docker üzerinde ayağa kaldırın.
3. Uygulamayı klonlayın ve `go run main.go` komutunu kullanarak çalıştırın.

### Kullanım

* **URL kısaltmak için:**
    * Kısaltmak istediğiniz url i seçin.
    * Uzun URL'yi girin ve Post isteği gönderin.
    * Kısaltılmış kod ile get isteği sonucunda uzun url i elde edeceksiniz.

### Amaç - Kapsam

* Uygulama gelişim amaçlı geliştirilmiştir.



