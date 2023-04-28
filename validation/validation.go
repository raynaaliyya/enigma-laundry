package validation

import (
	"fmt"
	"os"
	"strconv"

	"github.com/raynaaliyya/enigma-laundry/entity"
)

func ValidateService(service entity.Service) error {
	minPrice, err := strconv.Atoi(os.Getenv("MIN_PRICE"))
	if err != nil {
		return fmt.Errorf("failed to read MIN_PRICE env variable: %v", err)
	}

	minLenServiceName, err := strconv.Atoi(os.Getenv("MIN_LEN_SERVICE_NAME"))
	if err != nil {
		return fmt.Errorf("failed to read MIN_LEN_SERVICE_NAME env variable: %v", err)
	}

	if service.Harga < minPrice {
		return fmt.Errorf("harga pelayanan tidak boleh kurang dari %d", minPrice)
	}

	if len(service.Pelayanan) < minLenServiceName {
		return fmt.Errorf("panjang nama produk layanan tidak boleh kurang dari %d kata", minLenServiceName)
	}
	return nil
}

func ValidateTransaction(service entity.Service) error {
	minPrice, err := strconv.Atoi(os.Getenv("MIN_PRICE"))
	if err != nil {
		return fmt.Errorf("failed to read MIN_PRICE env variable: %v", err)
	}

	serviceCuciSetrika := os.Getenv("CUCI_SETRIKA")

	if service.Pelayanan == serviceCuciSetrika && service.Harga < minPrice {
		return fmt.Errorf("pelayanan %s must at least cost %d", serviceCuciSetrika, minPrice)
	}
	return nil
}
