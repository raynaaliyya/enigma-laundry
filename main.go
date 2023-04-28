package main

import (
	"fmt"
	"log"
	"os"

	"github.com/raynaaliyya/enigma-laundry/config"
	"github.com/raynaaliyya/enigma-laundry/entity"
	"github.com/raynaaliyya/enigma-laundry/repository"
	"github.com/raynaaliyya/enigma-laundry/usecase"
	"github.com/raynaaliyya/enigma-laundry/validation"

	_ "github.com/lib/pq"
)

func main() {
	os.Setenv("MIN_PRICE", "0")
	os.Setenv("MIN_LEN_SERVICE_NAME", "3")

	db := config.NewDatabase()

	// dependencies
	serviceRepo := repository.NewServiceRepo(db)
	serviceUsecase := usecase.NewServiceUsecase(serviceRepo)

	service := entity.Service{ID: "S003", Pelayanan: "Laundry Bedcover", Harga: 50000}
	if err := validation.ValidateService(service); err != nil {
		log.Println(err)
		return
	}

	// tambahkan data mahasiswa
	fmt.Println("==> INSERT PELAYANAN <==")
	serviceUsecase.Insert(service)

	// // tampilkan data mahasiswa
	// fmt.Println("==> Find All <==")

	// // tampilkan berdasarkan ID
	// fmt.Println("==> Find By Id<==")

	// // update berdasarkan ID
	// fmt.Println("==> Edit By Id <==")

	// // delete berdasarkan ID
	// fmt.Println("==> Delete By Id <==")
}
